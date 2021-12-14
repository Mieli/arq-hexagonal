package weapon

import (
	"context"
	"fmt"
	"time"

	pkgutils "delegacia.com.br/app/domain/utils"
	pkgvictim "delegacia.com.br/app/domain/victim"
	"delegacia.com.br/infra/dl"
	pkginfrautils "delegacia.com.br/infra/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	dataBase           mongo.Client
	collection         mongo.Collection
	contextoRepository context.Context
}

func NewVictimRepository(db mongo.Client) pkgvictim.VictimRepository {
	collection := db.Database(dl.GetEnv("MONGO_DATABASE")).Collection("victim")

	return &repository{
		dataBase:           db,
		collection:         *collection,
		contextoRepository: context.TODO(),
	}
}

func (r *repository) Save(victim pkgvictim.Victim) (*pkgvictim.Victim, error) {
	id := victim.ID.Hex()
	if id != "" {
		victim, err := update(r.contextoRepository, r.collection, id, victim)
		if err != nil {
			return nil, err
		}
		return victim, nil
	} else {
		oid, err := insert(r.contextoRepository, r.collection, &victim)
		id, _ := primitive.ObjectIDFromHex(*oid)
		victim.ID = id
		if err != nil {
			return nil, err
		}
		return &victim, nil
	}

}

func insert(ctx context.Context, collection mongo.Collection, victim *pkgvictim.Victim) (*string, error) {
	pkginfrautils.SetEventRecord(&victim.EventRecord, pkgutils.CREATED_AT, true)
	result, err := collection.InsertOne(ctx, victim)
	if err != nil {
		return nil, err
	}
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		oid := id.Hex()
		return &oid, nil
	}
	return nil, err
}

func update(ctx context.Context, collection mongo.Collection, id string, newVictim pkgvictim.Victim) (*pkgvictim.Victim, error) {

	objId, err := primitive.ObjectIDFromHex(id)
	filterFind := bson.M{
		"_id":                objId,
		"eventRecord.active": true,
	}

	oldVictim := pkgvictim.Victim{}
	options := &options.FindOneOptions{
		AllowPartialResults: &options.DefaultOrdered,
	}
	err = collection.FindOne(ctx, filterFind, options).Decode(&oldVictim)
	if err != nil {
		return nil, err
	}

	newVictim.EventRecord = pkginfrautils.MergeEventRecord(newVictim.EventRecord, oldVictim.EventRecord)
	today := time.Now()
	newVictim.EventRecord.UpdatedAt = &today

	filter := bson.M{
		"_id":                objId,
		"eventRecord.active": true,
	}

	filterUpdate := bson.M{}
	filterUpdate["$set"] = &newVictim

	result := collection.FindOneAndUpdate(ctx, filter, filterUpdate)

	if result != nil && result.Err() != nil {
		return nil, fmt.Errorf(result.Err().Error())
	}
	return &newVictim, nil
}

func (r *repository) FindById(id string) (*pkgvictim.Victim, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":                objId,
		"eventRecord.active": true,
	}

	victim := pkgvictim.Victim{}
	options := &options.FindOneOptions{
		AllowPartialResults: &options.DefaultOrdered,
	}
	err = r.collection.FindOne(r.contextoRepository, filter, options).Decode(&victim)
	if err != nil {
		return nil, err
	}
	return &victim, nil
}

func (r *repository) FindAll() ([]*pkgvictim.Victim, error) {
	filter := bson.M{
		"eventRecord.active": true,
	}
	cursor, err := r.collection.Find(r.contextoRepository, filter)
	if err != nil {
		return nil, err
	}
	victims := make([]*pkgvictim.Victim, 0)
	if err = cursor.All(r.contextoRepository, &victims); err != nil {
		return nil, err
	}
	cursor.Close(r.contextoRepository)

	return victims, nil
}

func (r *repository) Remove(id string) error {
	victim, err := r.FindById(id)
	if err != nil {
		return err
	}

	pkginfrautils.SetEventRecord(&victim.EventRecord, pkgutils.DELETED_AT, false)

	objId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": objId,
	}

	filterUpdate := bson.M{}
	filterUpdate["$set"] = &victim

	result := r.collection.FindOneAndUpdate(r.contextoRepository, filter, filterUpdate)

	if result != nil && result.Err() != nil {
		return fmt.Errorf(result.Err().Error())
	}
	return nil
}
