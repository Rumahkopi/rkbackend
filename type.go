package rkbackend

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
	Role     string `bson:"role,omitempty" json:"role,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Produk struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Harga     string             `bson:"harga,omitempty" json:"harga,omitempty"`
	Deskripsi string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Stok      string             `bson:"stok,omitempty" json:"stok,omitempty"`
}

type ResponseProduk struct {
	Status  bool     `bson:"status" json:"status"`
	Message string   `bson:"message,omitempty" json:"message,omitempty"`
	Data    []Produk `bson:"data,omitempty" json:"data,omitempty"`
}

type Transaksi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Harga    string             `bson:"harga,omitempty" json:"harga,omitempty"`
	Alamat   string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NoHP     string             `bson:"nohp,omitempty" json:"nohp,omitempty"`
	Quantity string             `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
