package rkbackend

import (
	"context"
	"errors"
	"fmt"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// Admin
func InsertProduk(db *mongo.Database, col string, produkdata Produk) (insertedID primitive.ObjectID, err error) {
	insertedID, err = InsertOneDoc(db, col, produkdata)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
	}
	return insertedID, err
}

func InsertPesanReview(db *mongo.Database, col string, pesandata Pesan) (insertedID primitive.ObjectID, err error) {
	objectid := primitive.NewObjectID()
	data := bson.M{
		"_id":    objectid,
		"nama":   pesandata.Nama,
		"subjek": pesandata.Subjek,
		"pesan":  pesandata.Pesan,
	}
	insertedID, err = InsertOneDoc(db, col, data)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
	}
	return insertedID, err
}

func GetAllDataProduk(db *mongo.Database, col string) (produklist []Produk) {
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

func GetAllDataReview(db *mongo.Database, col string) (pesanlist []Pesan) {
	cols := db.Collection(col)
	filter := bson.M{}
	cursor, err := cols.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error GetAllDocs in colection", col, ":", err)
	}
	err = cursor.All(context.TODO(), &pesanlist)
	if err != nil {
		fmt.Println(err)
	}
	return pesanlist
}

func InsertUser(db *mongo.Database, collection string, userdata Admin) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "Username : " + userdata.Username + "\nPassword : " + userdata.Password
}

func UpdateProduk(db *mongo.Database, col string, produk Produk) (produks Produk, status bool, err error) {
	cols := db.Collection(col)
	filter := bson.M{"_id": produk.ID}
	update := bson.M{
		"$set": bson.M{
			"nama":      produk.Nama,
			"harga":     produk.Harga,
			"deskripsi": produk.Deskripsi,
			"stok":      produk.Stok,
			"image":     produk.Image,
		},
	}

	result, err := cols.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return produks, false, err
	}
	if result.ModifiedCount == 0 && result.UpsertedCount == 0 {
		err = fmt.Errorf("Data tidak berhasil diupdate")
		return produks, false, err
	}

	err = cols.FindOne(context.Background(), filter).Decode(&produks)
	if err != nil {
		return produks, false, err
	}

	return produks, true, nil
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

func GetProdukFromID(db *mongo.Database, col string, _id primitive.ObjectID) (*Produk, error) {
	cols := db.Collection(col)
	filter := bson.M{"_id": _id}

	produklist := new(Produk)

	err := cols.FindOne(context.Background(), filter).Decode(produklist)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no data found for ID %s", _id.Hex())
		}
		return nil, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
	}

	return produklist, nil
}

func GetAllDataKeluhan(db *mongo.Database, col string) (keluhanlist []Keluhan) {
	cols := db.Collection(col)
	filter := bson.M{}
	cursor, err := cols.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error GetAllDocs in colection", col, ":", err)
	}
	err = cursor.All(context.TODO(), &keluhanlist)
	if err != nil {
		fmt.Println(err)
	}
	return keluhanlist
}

func GetAllDataTransaksi(db *mongo.Database, col string) (transaksilist []Transaksi) {
	cols := db.Collection(col)
	filter := bson.M{}
	cursor, err := cols.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error GetAllDocs in colection", col, ":", err)
	}
	err = cursor.All(context.TODO(), &transaksilist)
	if err != nil {
		fmt.Println(err)
	}
	return transaksilist
}

// func TransaksiCleared(db *mongo.Database, col string, done TransaksiClear) (bool, error) {
// 	cols := db.Collection(col)

// 	transaksi, err := GetTransaksiFromID(db, "transaksi", done.Transaksi.ID)
// 	if err != nil {
// 		fmt.Println("Error GetTodoFromID in colection", col, ":", err)
// 		return false, err
// 	}

// 	fmt.Println("todo: ", transaksi)

// 	insert := bson.D{
// 		{Key: "isdone", Value: true},
// 		{Key: "transaksi", Value: bson.D{
// 			{Key: "_id", Value: transaksi.ID},
// 			{Key: "transaksi_number", Value: transaksi.Transaksi_number},
// 			{Key: "status", Value: transaksi.Status},
// 			{Key: "user_phone", Value: transaksi.User_phone},
// 			{Key: "formatted_time", Value: transaksi.Formatted_time},
// 			{Key: "Buktitf", Value: transaksi.Buktitf},
// 		}},
// 	}

// 	fmt.Println("insert: ", insert)

// 	_, err = cols.InsertOne(context.Background(), insert)
// 	if err != nil {
// 		return false, err
// 	}

// 	status, err := DeleteProduk(db, "transaksi", transaksi.ID)
// 	if err != nil {
// 		fmt.Println("Error DeleteTodo in colection", col, ":", err)
// 		return false, err
// 	}

// 	if !status {
// 		fmt.Println("Data tidak berhasil dihapus")
// 		return false, err
// 	}

// 	return true, nil
// }

// func GetTransaksiFromID(db *mongo.Database, col string, _id primitive.ObjectID) (*Transaksi, error) {
// 	cols := db.Collection(col)
// 	filter := bson.M{"_id": _id}

// 	transaksilist := new(Transaksi)

// 	err := cols.FindOne(context.Background(), filter).Decode(transaksilist)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, fmt.Errorf("no data found for ID %s", _id.Hex())
// 		}
// 		return nil, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
// 	}

// 	return transaksilist, nil
// }
// func DeleteTransaksi(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
// 	cols := db.Collection(col)
// 	filter := bson.M{"_id": _id}
// 	result, err := cols.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return false, err
// 	}
// 	if result.DeletedCount == 0 {
// 		err = fmt.Errorf("Data tidak berhasil dihapus")
// 		return false, err
// 	}
// 	return true, nil
// }

// func GetTransaksiDone(db *mongo.Database, col string) (todo []TransaksiClear, err error) {
// 	cols := db.Collection(col)
// 	filter := bson.M{"isdone": true}

// 	cur, err := cols.Find(context.Background(), filter)
// 	if err != nil {
// 		fmt.Println("Error GetTodoDone in colection", col, ":", err)
// 		return todo, err
// 	}

// 	err = cur.All(context.Background(), &todo)
// 	if err != nil {
// 		fmt.Println("Error reading documents:", err)
// 		return todo, err
// 	}

// 	return todo, nil
// }
