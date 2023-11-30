package rkbackend	

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/aiteung/atdb"

)

func InsertOneDoc(db *mongo.Database, col string, docs interface{}) (insertedID primitive.ObjectID, err error) {
	cols := db.Collection(col)
	result, err := cols.InsertOne(context.Background(), docs)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return
}

func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	cols := db.Collection(col)
	filter := bson.M{}
	cursor, err := cols.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error GetAllDocs in colection", col, ":", err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		fmt.Println(err)
	}
	return docs
}

func UpdateOneDoc(db *mongo.Database, col string, filter, update interface{}) (err error) {
    cols := db.Collection(col)
    result, err := cols.UpdateOne(context.Background(), filter, bson.M{"$set": update})
    if err != nil {
        fmt.Printf("UpdateOneDoc: %v\n", err)
    }
    if result.ModifiedCount == 0 {
        err = errors.New("no data has been changed with the specified filter")
        return err
    }
    return
}


func DeleteOneDoc(db *mongo.Database, col string, filter bson.M) (err error) {
    cols := db.Collection(col)
    result, err := cols.DeleteOne(context.Background(), filter)
    if err != nil {
        fmt.Printf("DeleteOneDoc: %v\n", err)
    }
    if result.DeletedCount == 0 {
        err = fmt.Errorf("no data has been deleted with the specified filter")
        return err
    }
    return
}


//Admin
func InsertProduk(db *mongo.Database, col string, produkdata Produk) (insertedID primitive.ObjectID, err error) {
	insertedID, err = InsertOneDoc(db, col, produkdata)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
	}
	return insertedID, err
}

func GetAllDataProduk(db *mongo.Database, col string) (produklist [] Produk) {
	cols := db.Collection(col)
	filter := bson.M{}
	cursor, err := cols.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error GetAllDocs in colection", col, ":", err)
	}
	err = cursor.All(context.TODO(), &produklist)
	if err != nil {
		fmt.Println(err)
	}
	return produklist
}

func InsertUser(db *mongo.Database, collection string, userdata Admin) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "Username : " + userdata.Username + "\nPassword : " + userdata.Password
}

func DeleteProduk(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
	cols := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := cols.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}
	if result.DeletedCount == 0 {
		err = fmt.Errorf("Data tidak berhasil dihapus")
		return false, err
	}
	return true, nil
}




