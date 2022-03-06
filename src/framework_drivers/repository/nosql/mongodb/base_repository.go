package repository_mongodb

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	username = "username"
	password = "password"
	host     = "host"
	database = "database"
	uri      = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", username, password, host, database)
	client   *mongo.Client
)

func getCLient() *mongo.Client {
	url := os.Getenv("MONGO_URL")
	if client == nil {
		//client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
		//client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	}
	return client
}

func FindByID(id *string, dest interface{}) (interface{}, error) {
	method := reflect.ValueOf(dest).MethodByName("CollectionName")
	var collection string
	if method.IsValid() {
		res := method.Call([]reflect.Value{})
		ret := res[0].Interface()
		collection = ret.(string)
	}
	if len(collection) <= 0 {
		return dest, errors.New("table not found")
	}
	coll := getCLient().Database(database).Collection(collection)
	objID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}
	opts := options.FindOne()
	result := coll.FindOne(context.TODO(), bson.D{{"_id", objID}}, opts)
	err = result.Decode(dest)
	return dest, err
}

func Update(dest interface{}) error {
	method := reflect.ValueOf(dest).MethodByName("CollectionName")
	var collection string
	if method.IsValid() {
		res := method.Call([]reflect.Value{})
		ret := res[0].Interface()
		collection = ret.(string)
	}
	if len(collection) <= 0 {
		return errors.New("collection not found")
	}
	coll := getCLient().Database(database).Collection(collection)
	updated := time.Now().UTC()
	setValueToFieldByName(dest, "updated", &updated)

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", getValueFieldByName(dest, "_id")}}
	result, err := coll.UpdateOne(context.TODO(), filter, getObjToUpdate(dest), opts)
	if err != nil {
		return err
	}
	if oid, ok := result.UpsertedID.(primitive.ObjectID); ok {
		setValueToFieldByName(dest, "_id", oid)
	}
	return err
}

func Insert(dest interface{}) (interface{}, error) {
	method := reflect.ValueOf(dest).MethodByName("CollectionName")
	var collection string
	if method.IsValid() {
		res := method.Call([]reflect.Value{})
		ret := res[0].Interface()
		collection = ret.(string)
	}
	if len(collection) <= 0 {
		return dest, errors.New("table not found")
	}
	coll := getCLient().Database(database).Collection(collection)
	inserted := time.Now().UTC()
	setValueToFieldByName(dest, "inserted", &inserted)

	opts := options.InsertOne()
	result, err := coll.InsertOne(context.TODO(), &dest, opts)
	if err != nil {
		return nil, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		setValueToFieldByName(dest, "_id", oid)
	}
	return dest, nil
}

func getObjToUpdate(dest interface{}) bson.M {
	var result bson.M = bson.M{}
	val := reflect.Indirect(reflect.ValueOf(dest))
	for i := 0; i < val.Type().NumField(); i++ {
		fieldName, success := val.Type().Field(i).Tag.Lookup("bson")
		if success {
			fieldName = strings.Split(fieldName, ",")[0]
			if fieldName != "inserted" && fieldName != "u_inserted" {
				result[fieldName] = val.Field(i).Interface()
			}
		}
	}
	return bson.M{"$set": result}
}

func setValueToFieldByName(dest interface{}, field_name string, value interface{}) {
	val := reflect.Indirect(reflect.ValueOf(dest))
	for i := 0; i < val.Type().NumField(); i++ {
		fieldName, success := val.Type().Field(i).Tag.Lookup("bson")
		if success {
			fieldName = strings.Split(fieldName, ",")[0]
			if fieldName == field_name {
				if val.Field(i).Type() == reflect.TypeOf((*primitive.ObjectID)(nil)) {
					v := reflect.New(reflect.TypeOf(value)).Elem()
					v.Set(reflect.ValueOf(value))
					val.Field(i).Set(v.Addr())
				} else {
					val.Field(i).Set(reflect.ValueOf(value))
				}
			}
		}
	}
}

func getValueFieldByName(dest interface{}, field_name string) interface{} {
	var fieldValue interface{}
	val := reflect.Indirect(reflect.ValueOf(dest))
	for i := 0; i < val.Type().NumField(); i++ {
		fieldName, success := val.Type().Field(i).Tag.Lookup("bson")
		if success {
			fieldName = strings.Split(fieldName, ",")[0]
			if fieldName == field_name {
				fieldValue = val.Field(i).Interface()
			}
		}
	}
	return fieldValue
}
