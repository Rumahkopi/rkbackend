package rkbackend

import (
	"encoding/json"
	"net/http"
	"os"
	"fmt"
	"strconv"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var dataadmin Admin
	err := json.NewDecoder(r.Body).Decode(&dataadmin)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, dataadmin) {
			Response.Status = true
			tokenstring, err := watoken.Encode(dataadmin.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func InsertData(db *mongo.Database, collection string, userdata Admin) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "Username : " + userdata.Username + "\nPassword : " + userdata.Password
}

func InsertProduk(Mongostring, dbname, colname string, r *http.Request) string {
	req := new(Credential)
	conn := SetConnection(Mongostring, dbname)
	resp := new(Produk)
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
			req.Status = strconv.Itoa(http.StatusNotFound)
			req.Message = "error parsing application/json: " + err.Error()
	} else {
			req.Status = strconv.Itoa(http.StatusOK)
			Ins := InsertDataProduk(conn, colname,
				resp.Nama,
				resp.Harga)
			req.Message = fmt.Sprintf("%v:%v", "Berhasil Input data", Ins)
	}
	
	return GCFReturnStruct(req)
}