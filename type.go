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
	Image     string             `bson:"image,omitempty" json:"image,omitempty"`
}

type Pesan struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama   string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Subjek string             `bson:"subjek,omitempty" json:"subjek,omitempty"`
	Pesan  string             `bson:"pesan,omitempty" json:"pesan,omitempty"`
}

type Keluhan struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Complaint_number int                `bson:"complaint_number,omitempty" json:"complaint_number,omitempty"`
	User_phone       string             `bson:"user_phone,omitempty" json:"user_phone,omitempty"`
	FormattedTime    string             `bson:"formattedtime,omitempty" json:"formatedtime,omitempty"`
	Content          string             `bson:"content,omitempty" json:"content,omitempty"`
}

type Transaksi struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Transaksi_number int                `bson:"transaksi_number,omitempty" json:"transaksi_number,omitempty"`
	Status           string             `bson:"status,omitempty" json:"status,omitempty"`
	User_phone       string             `bson:"user_phone,omitempty" json:"user_phone,omitempty"`
	Formatted_time   string             `bson:"formatted_time,omitempty" json:"formatted_time,omitempty"`
	Buktitf          string             `bson:"buktitf,omitempty" json:"buktitf,omitempty"`
}

type TransaksiClear struct {
	IsDone    bool      `bson:"isdone,omitempty" json:"isdone,omitempty"`
	Transaksi Transaksi `bson:"transaksi,omitempty" json:"transaksi,omitempty"`
}

type TransaksiClearResponse struct {
	Status  bool             `bson:"status" json:"status"`
	Message string           `bson:"message,omitempty" json:"message,omitempty"`
	Data    []TransaksiClear `bson:"data,omitempty" json:"data,omitempty"`
}
