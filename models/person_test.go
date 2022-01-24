package models_test

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	// db "github.com/ggrrrr/bui_lib/db/mgo"
	"github.com/ggrrrr/bui_api_people/models"
	db "github.com/ggrrrr/bui_lib/db/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// github.com/ggrrrr/bui_lib/db/mgo
)

const (
	NS   string = "ns"
	PIN1 string = "PIN1"
)

var (
	PHONES1 = map[string]string{"label1": "val1"}
	LABEL1  = []string{"l1", "l2", "l1=1", "l1=2"}
)

func OK(t *testing.T, str string, v ...interface{}) {
	a := fmt.Sprintf(str, v...)
	t.Logf("OK WITH %v", a)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

const shortForm = "2006-01-02"

func TestPasswd1(t *testing.T) {
	ctx := context.Background()

	err := db.Configure(ctx)
	if err != nil {
		// mgo.Configure(ctx)
		t.Fatal(err)
	}

	_, err = models.List(ctx, bson.D{})
	if err != nil {
		t.Fatalf("F list All %v", err)
	}
	// if P2.Id != UUID1 {
	// t.Logf("F cant find %v", )

	LAEL2 := RandomString(7)
	DOB, _ := time.Parse(shortForm, "2020-09-29")
	P0 := &models.Person{FullName: "Pesho peshov", Labels: []string{"tours:mtb", "tours:winter", LAEL2}, DateOfBirth: primitive.NewDateTimeFromTime(DOB)}

	err = P0.Insert(ctx)
	if err != nil {
		t.Errorf("F person insert")
	}

	P1, err := models.GetById(ctx, P0.Id)
	if err != nil {
		t.Errorf("F person insert")
	}
	if !reflect.DeepEqual(P0, P1) {
		t.Errorf("F P0 %v", P0)
		t.Errorf("F P1 %v", P1)
		t.Errorf("F P0 != P2")
	}
	// t.Logf("found by id: %v", P1)

	// filter := bson.M{"name": "name1"}
	// labelFilter := bson.M{"label1": "val1"}
	// filter := bson.D{{"$match", bson.D{{"name", "name1"}}}}
	// filter := bson.M{

	// 	"labels": bson.D{{"$elemMatch", labelFilter}},
	// }
	// filter := bson.D{{"labels", bson.M{"label1": "val1"}}}
	// filter := bson.M{
	// 	"labels": bson.M{"keys": bson.M{"$all": bson.A{"label1"}}},
	// }

	// list1, err := models.ListAll(ctx)
	// if err != nil {
	// 	t.Fatalf("F cant find %v", err)
	// }
	// // if P2.Id != UUID1 {
	// t.Logf("F cant find %v", list1)

	// return
	filter := bson.D{{"labels", bson.D{{"$all", bson.A{LAEL2}}}}}
	// filter := bson.D{{"labels.label1", bson.D{{"$exists", true}}}}
	// filter := bson.D{{"labels", bson.D{{"$exists", true}}}}

	list, err := models.List(ctx, filter)
	if err != nil {
		t.Fatalf("F cant find %v", err)
	}
	// if P2.Id != UUID1 {
	t.Logf("F cant find %v, 0: %+v", len(*list), (*list)[0].Id)
	// }
	// return
	// record := bson.M{"name": "newName", "$push": bson.M{"labels": bson.A{"new:shit"}}}
	update := bson.M{
		"$addToSet": bson.M{"labels": "new:shit"},
	}
	err = models.Update(ctx, (*list)[0].Id, update)
	if err != nil {
		t.Fatalf("F update error %v", err)
	}

	update = bson.M{
		"$addToSet": bson.M{"labels": "new:shit"},
	}
	err = models.Update(ctx, (*list)[0].Id, update)
	if err != nil {
		t.Fatalf("F update error %v", err)
	}
	update = bson.M{
		"$addToSet": bson.M{"labels": "new:shit"},
	}
	err = models.Update(ctx, (*list)[0].Id, update)
	if err != nil {
		t.Fatalf("F update error %v", err)
	}
	update = bson.M{
		"$addToSet": bson.M{"labels": "new:shit"},
	}
	err = models.Update(ctx, (*list)[0].Id, update)
	if err != nil {
		t.Fatalf("F update error %v", err)
	}
	// return
	// record1 := bson.M{"labels.label2": "shit2"}
	update1 := bson.M{
		"$pull": bson.M{"labels": "label3:va1"},
	}
	err = models.Update(ctx, (*list)[0].Id, update1)
	if err != nil {
		t.Fatalf("F update1 error %v", err)
	}
	update1 = bson.M{
		"$pull": bson.M{"labels": "new:shit"},
	}
	err = models.Update(ctx, (*list)[0].Id, update1)
	if err != nil {
		t.Fatalf("F update1 error %v", err)
	}

}
