package models

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ggrrrr/bui_lib/db/mgo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
CREATE TABLE person (
	namespace text,
	id uuid,
	name text,
	email text,
	first_name text,
	last_name text,
	phones frozen<map<text,text>>,``
	attr frozen<map<text,text>>,
	labels frozen<map<text,text>>,
	created_time timestamp,
	PRIMARY KEY (namespace, id)
)
*/

type Person struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Namespace   string
	PIN         string             `json:"pin"`
	Email       string             `json:"email"`
	Name        string             `json:"name"`
	FullName    string             `bson:"full_name" json:"full_name"`
	DateOfBirth primitive.DateTime `bson:"dob" json:"dob"`
	Gender      string             `bson:"gender" json:"gender"`
	Phones      map[string]string  `json:"phones"`
	Labels      []string           `json:"labels"`
	Attr        map[string]string  `json:"attr"`
	Age         string             `bson:"-" json:"age"`
	CreatedTime time.Time
}

var EMPY_UUID = uuid.UUID{}
var COLLECTION_NAME string = "people"

// func CurrentNameSpace(ctx context.Context) string {
// 	return fmt.Sprint(ctx.Value(api.ContextPerson))
// }

func calcAge(dob time.Time) string {
	if dob.Year() > 1000 {
		return fmt.Sprintf("%d", time.Now().Year()-dob.Year())

	}
	return ""
}

func (p *Person) Insert(ctx context.Context) error {
	if len(p.FullName) > 0 {
		if p.Name == "" {
			names := strings.Split(p.FullName, " ")
			if len(names) > 1 {
				p.Name = names[0]
			}
		}
	}
	p.Id = primitive.NewObjectID()
	col := mgo.Database.Collection(COLLECTION_NAME)
	res, err := col.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	log.Printf("%v %v", res, p)
	// mgo.Database.Collection()
	return nil
}

func GetById(ctx context.Context, id primitive.ObjectID) (*Person, error) {
	col := mgo.Database.Collection(COLLECTION_NAME)
	var out Person
	res := col.FindOne(ctx, bson.M{"_id": id})

	if res.Err() != nil {
		return nil, res.Err()
	}
	err := res.Decode(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func List(ctx context.Context, filter bson.D) (*[]Person, error) {

	var out = make([]Person, 0)
	col := mgo.Database.Collection(COLLECTION_NAME)
	cur, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(context.Background()) {
		var result Person
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result.Age = calcAge(result.DateOfBirth.Time())
		out = append(out, result)
	}

	return &out, nil
}

func Update(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	log.Printf("Update[%v]: %v", id, update)
	col := mgo.Database.Collection(COLLECTION_NAME)
	res, err := col.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}
	log.Printf("udpate: %v", res)
	return nil

}
