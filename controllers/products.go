package controllers

import (
	"alura-go-store/models"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	err := temp.ExecuteTemplate(w, "Index", products)
	if err != nil {
		return
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "New", nil)
	if err != nil {
		panic(err.Error())
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")
	err := models.AddProduct(name, description, price, quantity)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var id, _ = strconv.Atoi(r.URL.Query().Get("id"))
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var id, _ = strconv.Atoi(r.URL.Query().Get("id"))
	p, err := models.FindById(id)
	if err != nil {
		panic(err.Error())
	}
	err = temp.ExecuteTemplate(w, "Edit", p)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")
	_, err := models.UpdateProduct(id, name, description, price, quantity)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/", 301)
}
