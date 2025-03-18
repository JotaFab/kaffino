package server

import (
	"encoding/json"
	"kaffino/internal/database"
	"log"
	"net/http"
)

func (s *Server) createProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	product := &database.Product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Create the product
	if err := s.db.CreateProduct(r.Context(), product); err != nil {
		log.Printf("Failed to create product: %v", err)
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	// Marshal the response
	jsonResp, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from the URL
	
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Get the product from the database
	product, err := s.db.GetProduct(r.Context(), id)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		http.Error(w, "Failed to get product", http.StatusInternalServerError)
		return
	}

	// Marshal the response
	jsonResp, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) listProductsHandler(w http.ResponseWriter, r *http.Request) {
	// List all products
	products, err := s.db.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to list products", http.StatusInternalServerError)
		return
	}

	// Marshal the response
	jsonResp, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	product := &database.Product{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Update the product
	if err := s.db.UpdateProduct(r.Context(), product); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	// Marshal the response
	jsonResp, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// Get the product ID from the URL
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Delete the product
	if err := s.db.DeleteProduct(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusNoContent)
}
