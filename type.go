package rkbackend

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

type Credential struct {
	Status  bool     `json:"status" bson:"status"`
	Token   string   `json:"token,omitempty" bson:"token,omitempty"`
	Message string   `json:"message,omitempty" bson:"message,omitempty"`
	Data    []Produk `bson:"data,omitempty" json:"data,omitempty"`
}

type Produk struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Harga     string             `bson:"harga,omitempty" json:"harga,omitempty"`
	Deskripsi string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Stok      string             `bson:"stok,omitempty" json:"stok,omitempty"`
}

type Transaksi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaProduk  string             `bson:"namaproduk,omitempty" json:"namaproduk,omitempty"`
	Harga       string             `bson:"harga,omitempty" json:"harga,omitempty"`
	NamaPembeli string             `bson:"namapembeli,omitempty" json:"namapembeli,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Alamat      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NoHP        string             `bson:"nohp,omitempty" json:"nohp,omitempty"`
	Quantity    string             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Total       string             `bson:"total,omitempty" json:"total,omitempty"`
}
