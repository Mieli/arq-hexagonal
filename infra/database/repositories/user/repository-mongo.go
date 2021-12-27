package user

import (
	"context"
	"fmt"
	"time"

	pkguser "delegacia.com.br/app/domain/user"
	pkgutils "delegacia.com.br/app/domain/utils"
	"delegacia.com.br/infra/dl"
	pkginfrautils "delegacia.com.br/infra/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	dataBase          mongo.Client
	collection        mongo.Collection
	contextRepository context.Context
}

func NewUserRepository(db mongo.Client) pkguser.UserRepository {
	collection := db.Database(dl.GetEnv("MONGO_DATABASE")).Collection("user")

	return &repository{
		dataBase:          db,
		collection:        *collection,
		contextRepository: context.TODO(),
	}
}

func (r *repository) Save(user pkguser.User) (*pkguser.User, error) {
	id := user.ID.Hex()
	if id != "000000000000000000000000" {
		newUser, err := update(r.contextRepository, r.collection, id, user)
		if err != nil {
			return nil, err
		}
		return newUser, nil

	} else {
		oid, err := insert(r.contextRepository, r.collection, &user)
		id, _ := primitive.ObjectIDFromHex(*oid)
		user.ID = id
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
}

func insert(ctx context.Context, collection mongo.Collection, user *pkguser.User) (*string, error) {
	pkginfrautils.SetEventRecord(&user.EventRecord, pkgutils.CREATED_AT, true)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		oid := id.Hex()
		return &oid, nil
	}
	return nil, err
}

func update(ctx context.Context, collection mongo.Collection, id string, newUser pkguser.User) (*pkguser.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	filterFind := bson.M{
		"_id":                objId,
		"eventRecord.active": true,
	}

	oldUser := pkguser.User{}
	options := &options.FindOneOptions{
		AllowPartialResults: &options.DefaultOrdered,
	}
	err = collection.FindOne(ctx, filterFind, options).Decode(&oldUser)
	if err != nil {
		return nil, err
	}

	newUser.EventRecord = pkginfrautils.MergeEventRecord(newUser.EventRecord, oldUser.EventRecord)
	today := time.Now()
	newUser.EventRecord.UpdatedAt = &today

	filter := bson.M{
		"_id":                objId,
		"eventRecord.active": true,
	}
	filterUpdate := bson.M{}
	filterUpdate["$set"] = &newUser

	_, err = collection.UpdateOne(ctx, filter, filterUpdate)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (r *repository) FindByEmailAndPassord(email, password string) (*pkguser.User, error) {

	filter := bson.M{
		"email":              email,
		"password":           password,
		"eventRecord.active": true,
	}

	user := pkguser.User{}
	options := &options.FindOneOptions{
		AllowPartialResults: &options.DefaultOrdered,
	}
	err := r.collection.FindOne(r.contextRepository, filter, options).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindAll() ([]*pkguser.User, error) {

	filter := bson.M{
		"eventRecord.active": true,
	}
	cursor, err := r.collection.Find(r.contextRepository, filter)
	if err != nil {
		return nil, err
	}
	users := make([]*pkguser.User, 0)
	if err = cursor.All(r.contextRepository, &users); err != nil {
		return nil, err
	}
	cursor.Close(r.contextRepository)

	return users, nil
}

func (r *repository) FindById(id string) (*pkguser.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":                objId,
		"eventRecord.active": true,
	}
	user := pkguser.User{}
	options := &options.FindOneOptions{
		AllowPartialResults: &options.DefaultOrdered,
	}
	err = r.collection.FindOne(r.contextRepository, filter, options).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Remove(id string) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	pkginfrautils.SetEventRecord(&user.EventRecord, pkgutils.DELETED_AT, false)

	objId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objId,
	}

	filterUpdate := bson.M{}
	filterUpdate["$set"] = &user

	result := r.collection.FindOneAndUpdate(r.contextRepository, filter, filterUpdate)

	if result != nil && result.Err() != nil {
		return fmt.Errorf(result.Err().Error())
	}
	return nil
}
