package rkbackend

type Admin struct {
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Produk struct {
	Nama string `bson:"nama,omitempty" json:"nama,omitempty"`
	Harga string `bson:"harga,omitempty" json:"harga,omitempty"`
	Deskripsi string `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
}