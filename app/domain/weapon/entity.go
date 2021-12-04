package weapon

type Weapon struct {
	ID          int64      `bson:"_id"`
	Description string     `bson:"description"`
	Type        TypeWeapon `bson:"type"`
}

type TypeWeapon struct {
	Category    string `bson:"category"`
	Description string `bson:"description"`
}
