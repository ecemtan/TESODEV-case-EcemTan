package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tesodev-product-api/config"
	"tesodev-product-api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var productCollection = config.GetCollection("product_database", "products") // get the "products" collection

// CreateProduct handles creating a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = primitive.NewObjectID()

	_, err := productCollection.InsertOne(context.Background(), product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting product: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProductByID handles retrieving a product by its ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = productCollection.FindOne(context.Background(), models.Product{ID: objectID}).Decode(&product)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// UpdateProduct handles updating an existing product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedProduct models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = productCollection.UpdateOne(
		context.Background(),
		models.Product{ID: objectID},
		map[string]interface{}{
			"$set": updatedProduct,
		},
	)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	updatedProduct.ID = objectID
	json.NewEncoder(w).Encode(updatedProduct)
}

// DeleteProduct handles deleting a product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	_, err = productCollection.DeleteOne(context.Background(), models.Product{ID: objectID})
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
