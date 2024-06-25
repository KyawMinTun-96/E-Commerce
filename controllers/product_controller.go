package controllers

import (
	"ecommerce-app/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	models.DB.Find(&products)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

type NewProduct struct {
	Item         string  `json:"item"`
	GenderType   string  `json:"gender_type"`
	Price        float64 `json:"price"`
	Qty          int     `json:"qty"`
	CategoryName string  `json:"category_name"`
	SizeValue    string  `json:"size_value"`
	StyleName    string  `json:"style_name"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct NewProduct

	json.NewDecoder(r.Body).Decode(&newProduct)

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		size := models.Size{
			Sizevalue: newProduct.SizeValue,
		}
		if err := tx.FirstOrCreate(&size, models.Size{Sizevalue: newProduct.SizeValue}).Error; err != nil {
			return err
		}

		gender := models.Gender{
			GenderType: newProduct.GenderType,
		}
		if err := tx.FirstOrCreate(&gender, models.Gender{GenderType: newProduct.GenderType}).Error; err != nil {
			return err
		}

		style := models.Style{
			StyleName: newProduct.StyleName,
		}
		if err := tx.FirstOrCreate(&style, models.Style{StyleName: newProduct.StyleName}).Error; err != nil {
			return err
		}

		qty := models.Qty{
			Quantity: newProduct.Qty,
		}
		if err := tx.FirstOrCreate(&qty, models.Qty{Quantity: newProduct.Qty}).Error; err != nil {
			return err
		}

		product := models.Product{
			Item: newProduct.Item,
		}
		if err := tx.FirstOrCreate(&product, models.Product{Item: newProduct.Item}).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
	// json.NewEncoder(w).Encode(price)
	// json.NewEncoder(w).Encode(catagory)

}

// package controllers

// import (
// 	"ecommerce-app/models"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"gorm.io/gorm"
// )

// func GetProducts(w http.ResponseWriter, r *http.Request) {
// 	var products []models.Product
// 	db := models.DB

// 	// Filters
// 	gender := r.URL.Query().Get("gender")
// 	style := r.URL.Query().Get("style")
// 	size := r.URL.Query().Get("size")

// 	if gender != "" {
// 		db = db.Where("gender = ?", gender)
// 	}
// 	if style != "" {
// 		db = db.Where("style = ?", style)
// 	}
// 	if size != "" {
// 		db = db.Where("size = ?", size)
// 	}

// 	// Pagination
// 	perPageStr := r.URL.Query().Get("per_page")
// 	pageIndexStr := r.URL.Query().Get("page")

// 	perPage := 10 // Default value
// 	if perPageStr != "" {
// 		perPage, _ = strconv.Atoi(perPageStr)
// 	}
// 	pageIndex := 1 // Default value
// 	if pageIndexStr != "" {
// 		pageIndex, _ = strconv.Atoi(pageIndexStr)
// 	}

// 	offset := (pageIndex - 1) * perPage
// 	result := db.Limit(perPage).Offset(offset).Find(&products)
// 	if result.Error != nil {
// 		if result.Error == gorm.ErrRecordNotFound {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}
// 		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(products)
// }
