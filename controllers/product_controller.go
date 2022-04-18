package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"latihan/models"

	"github.com/go-martini/martini"
)

//GetAllProducts
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * from products"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		SendErrorResponse(400, "Error Query", w)
		return
	}

	var product models.Product
	var products []models.Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Fatal(err.Error())
			SendErrorResponse(400, "Error Scan", w)
			return
		} else {
			products = append(products, product)
		}
	}

	var response models.ProductsResponse
	response.Response.Status = 200
	response.Response.Message = "Success"
	response.Data = products

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Insert Product
func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		SendErrorResponse(400, "Parse Form Failed", w)
		return
	}

	var product models.Product
	product.Name = r.Form.Get("name")
	product.Price, _ = strconv.Atoi(r.Form.Get("price"))

	res, errQuery := db.Exec("INSERT INTO products(name, price) values (?,?)",
		product.Name,
		product.Price,
	)

	var response models.ProductResponse
	if errQuery != nil {
		SendErrorResponse(400, "Error Query", w)
		return
	}

	lastId, _ := res.LastInsertId()
	product.ID = int(lastId)

	response.Response.Status = 200
	response.Response.Message = "Insert Success"
	response.Data = product
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Delete Product
func DeleteProduct(w http.ResponseWriter, r *http.Request, args martini.Params) {
	db := connect()
	defer db.Close()

	productId := args["id"]

	var response models.GeneralResponse
	_, errQuery := db.Exec("DELETE FROM products WHERE id=?",
		productId,
	)

	if errQuery != nil {
		SendErrorResponse(400, "Error Query", w)
		return
	}

	response.Status = 200
	response.Message = "Delete Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Update Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		SendErrorResponse(400, "Parse Form Failed", w)
		return
	}

	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		SendErrorResponse(400, "Id Not Found", w)
		return
	}
	name := r.Form.Get("name")
	price, _ := strconv.Atoi(r.Form.Get("price"))

	_, errQuery := db.Exec("UPDATE products SET name=?, price=? WHERE id = ?",
		name,
		price,
		id,
	)

	var response models.GeneralResponse
	if errQuery != nil {
		SendErrorResponse(400, "Error Query", w)
		return
	}

	response.Status = 200
	response.Message = "Update Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Login
func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		SendErrorResponse(400, "Parse Form Failed", w)
		return
	}

	email := r.Form.Get("email")
	pass := r.Form.Get("pass")

	h := sha1.New()
	h.Write([]byte(pass))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	query := "SELECT id, name, age, address, email, userType FROM users WHERE email = '" + email + "' AND pass = '" + sha1_hash + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		SendErrorResponse(400, "Error Query", w)
		return
	}

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.UserType); err != nil {
			log.Fatal(err.Error())
			SendErrorResponse(400, "Error Scan", w)
			return
		}
	}

	if user.ID == 0 {
		log.Println(err)
		SendErrorResponse(400, "Email/Password Salah", w)
		return
	}

	generateToken(w, user.ID, user.Name, user.UserType)
	var response models.UserResponse
	response.Response.Status = 200
	response.Response.Message = "Login Berhasil, Selamat Datang User"
	response.Data = user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	var response models.GeneralResponse
	response.Status = 200
	response.Message = "Logout Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//send Error Message
func SendErrorResponse(status int, message string, w http.ResponseWriter) {
	var response models.GeneralResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//send Error Message
func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response models.GeneralResponse
	response.Status = 401
	response.Message = "Unauthorized"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendGeneralResponse(w http.ResponseWriter, s int, m string) {
	var response models.GeneralResponse
	response.Status = s
	response.Message = m
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
