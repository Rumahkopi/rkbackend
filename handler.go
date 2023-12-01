package rkbackend

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func InsertDataProduk(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	produkdata := new(Produk)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&produkdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		insertedID, err := InsertProduk(conn, "produk", *produkdata)
		if err != nil {
			resp.Message = "Gagal memasukkan data ke database: " + err.Error()
		} else {
			resp.Message = "Berhasil Input data dengan ID: " + insertedID.Hex()
		}
	}
	return GCFReturnStruct(resp)
}

func InsertDataTransaksi(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	transaksidata := new(Transaksi)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&transaksidata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		insertedID, err := InsertTransaksi(conn, "transaksi", *transaksidata)
		if err != nil {
			resp.Message = "Gagal memasukkan data ke database: " + err.Error()
		} else {
			resp.Message = "Berhasil Input data dengan ID: " + insertedID.Hex()
		}
	}
	return GCFReturnStruct(resp)
}

func GetAllData(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	data := GetAllDataProduk(mconn, collectionname)
	return GCFReturnStruct(data)
}

func GetDataTransaksi(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	data := GetAllDataTransaksi(mconn, collectionname)
	return GCFReturnStruct(data)
}

func UpdateDataProduk(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	resp := new(ResponseProduk)
	produkdata := new(Produk)
	resp.Status = false

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		resp.Message = "error parsing application/json1:"
		return GCFReturnStruct(resp)
	}

	_, err := watoken.Decode(os.Getenv(PASETOPUBLICKEY), token)

	id := r.URL.Query().Get("_id")
	if id == "" {
		resp.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(resp)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		resp.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(resp)
	}

	produkdata.ID = ID

	err = json.NewDecoder(r.Body).Decode(&produkdata)
	if err != nil {
		resp.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(resp)
	}

	produk, _, err := UpdateProduk(mconn, "produk", *produkdata)
	if err != nil {
		resp.Message = err.Error()
		return GCFReturnStruct(resp)
	}

	resp.Status = true
	resp.Message = "Update todo success"
	resp.Data = []Produk{produk}

	return GCFReturnStruct(resp)
}

func DeleteDataProduk(PASETOPUBLICKEY, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	resp := new(ResponseProduk)
	produkdata := new(Produk)
	resp.Status = false

	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		resp.Message = "error parsing application/json1:"
		return GCFReturnStruct(resp)
	}

	id := r.URL.Query().Get("_id")
	if id == "" {
		resp.Message = "Missing '_id' parameter in the URL"
		return GCFReturnStruct(resp)
	}

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		resp.Message = "Invalid '_id' parameter in the URL"
		return GCFReturnStruct(resp)
	}

	produkdata.ID = ID

	err = json.NewDecoder(r.Body).Decode(&produkdata)
	if err != nil {
		resp.Message = "error parsing application/json3: " + err.Error()
		return GCFReturnStruct(resp)
	}

	_, err = DeleteProduk(mconn, collectionname, produkdata.ID)
	if err != nil {
		resp.Message = err.Error()
		return GCFReturnStruct(resp)
	}

	resp.Status = true
	resp.Message = "Delete todo success"

	return GCFReturnStruct(resp)
}
