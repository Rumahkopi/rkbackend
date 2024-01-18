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
	ID     primitive.ObjectID 		   `bson:"_id,omitempty" json:"_id,omitempty"`
	User_phone   		string             `bson:"user_phone,omitempty" json:"user_phone,omitempty"`
	FormattedTime 		string             `bson:"formattedtime,omitempty" json:"formatedtime,omitempty"`
	Content  			string             `bson:"content,omitempty" json:"content,omitempty"`
	Complaint_number	int				   `bson:"Complaint_number,omitempty" json:"Complaint_number,omitempty"`	
}

type Transaksi struct {
	ID     primitive.ObjectID 		   `bson:"_id,omitempty" json:"_id,omitempty"`
	User_phone   		string             `bson:"user_phone,omitempty" json:"user_phone,omitempty"`
	Formatted_time 		string             `bson:"formatted_time,omitempty" json:"formated_time,omitempty"`
	transaksi_number	int				   `bson:"transaksi_number,omitempty" json:"transaksi_number,omitempty"`	
	payment_proof		string			   `bson:"payment_proof,omitempty" json:"payment_proof,omitempty"`	
	buktitf				string			   `bson:"buktitf,omitempty" json:"content,omitempty"`
}