package rkbackend

import (
	"encoding/json"
	"net/http"
	"os"


	"github.com/whatsauth/watoken"
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

func GetAllData(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	data := GetAllDataProduk(mconn, collectionname)
	return GCFReturnStruct(data)
}