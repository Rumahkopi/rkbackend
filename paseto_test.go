package rkbackend

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PASETO
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("admin", privateKey)
	fmt.Println(hasil, err)
}

// Hash Pass
func TestGeneratePasswordHash(t *testing.T) {
	password := "adminkopi"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var userdata Admin
	userdata.Username = "admin"
	userdata.Password = "adminkopi"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[Admin](mconn, "admin", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var userdata Admin
	userdata.Username = "admin"
	userdata.Password = "adminkopi"

	anu := IsPasswordValid(mconn, "admin", userdata)
	fmt.Println(anu)
}

// func TestInsertUser(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "proyek3")
// 	var userdata Admin
// 	userdata.Username = "admin"
// 	userdata.Role = "Admin"
// 	userdata.Password = "adminkopi"

// 	nama := InsertUser(mconn, "admin", userdata)
// 	fmt.Println(nama)
// }

// func TestGetKeluhan(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "proyek3")
// 	anu := GetAllDataKeluhan(mconn, "keluhan")
// 	fmt.Println(anu)
// }

// func TestGet(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "proyek3")
// 	anu := GetAllDataTransaksi(mconn, "transaksi")
// 	fmt.Println(anu)
// }

func TestTransaksiClear(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	var data TransaksiClear

	id := "65c0e451abb612354c96f789"

	ID, err := primitive.ObjectIDFromHex(id)
	data.Transaksi.ID = ID
	if err != nil {
		fmt.Printf("Data tidak berhasil di selesaikan")
	} else {

		status, err := TransaksiCleared(mconn, "transaksidone", data)
		fmt.Println("Status", status)
		if err != nil {
			t.Errorf("Error cleared todo with id: %v", err)
			return
		} else {
			fmt.Printf("Data berhasil di selesaikan untuk: %s\n", ID)
		}
		fmt.Println(data)
	}
}

func TestGetTodoDone(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "proyek3")
	anu, err := GetTransaksiDone(mconn, "transaksidone")
	if err != nil {
		t.Errorf("Error getting todo: %v", err)
		return
	}
	fmt.Println(anu)
}
