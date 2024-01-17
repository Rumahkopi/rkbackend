package rkbackend

import (
	"encoding/json"
	"net/http"
	"os"

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
	} else if produkdata.Nama == "" || produkdata.Harga == "" || produkdata.Deskripsi == "" || produkdata.Stok == "" || produkdata.Image == "" {
		resp.Message = "Data Tidak Boleh Kosong"
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

func InsertDataPesan(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	pesandata := new(Pesan)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&pesandata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		insertedID, err := InsertPesanReview(conn, "review", *pesandata)
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

func GetDataReview(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	data := GetAllDataReview(mconn, collectionname)
	return GCFReturnStruct(data)
}

func UpdateDataProduk(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	produkdata := new(Produk)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)

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
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		produks, status, err := UpdateProduk(conn, "produk", *produkdata)
		if err != nil || !status {
			resp.Message = "Gagal update data : " + err.Error()
		} else {
			resp.Message = "Berhasil Update data dengan ID: " + produks.ID.Hex()
		}
	}
	return GCFReturnStruct(resp)
}

func DeleteDataProduk(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	produkdata := new(Produk)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)

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

	// err = json.NewDecoder(r.Body).Decode(&produkdata)
	// if err != nil {
	// 	resp.Message = "error parsing application/json: " + err.Error()
	// } else {

	resp.Status = true
	status, err := DeleteProduk(conn, "produk", produkdata.ID)
	if err != nil || !status {
		resp.Message = "Gagal Delete data : " + err.Error()
	} else {
		resp.Message = "Berhasil Delete Data Produk"
	}

	return GCFReturnStruct(resp)
}

func GetOneDataProduk(MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	resp := new(Credential)
	produkdata := new(Produk)
	resp.Status = false
	err := json.NewDecoder(r.Body).Decode(&produkdata)

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

	// Menggunakan fungsi GetProdukFromID untuk mendapatkan data produk berdasarkan ID
	produkdata, err = GetProdukFromID(mconn, collectionname, ID)
	if err != nil {
		resp.Message = err.Error()
		return GCFReturnStruct(resp)
	}

	resp.Status = true
	resp.Message = "Get Data Berhasil"
	resp.Data = []Produk{*produkdata}

	return GCFReturnStruct(resp)
}
