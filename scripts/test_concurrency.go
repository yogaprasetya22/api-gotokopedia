package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type UpdateProductPayload struct {
	Name          *string   `json:"name" validate:"omitempty"`
	Slug          *string   `json:"slug" validate:"omitempty,max=100"`
	Description   *string   `json:"description" validate:"omitempty,max=100"`
	Price         *float64  `json:"price" validate:"omitempty"`
	DiscountPrice *float64  `json:"discount_price" validate:"omitempty"`
	Discount      *float64  `json:"discount" validate:"omitempty"`
	Rating        *float64  `json:"rating" validate:"omitempty"`
	Estimation    *string   `json:"estimation" validate:"omitempty"`
	Stock         *int      `json:"stock" validate:"omitempty"`
	Sold          *int      `json:"sold" validate:"omitempty"`
	IsForSale     *bool     `json:"is_for_sale" validate:"omitempty"`
	IsApproved    *bool     `json:"is_approved" validate:"omitempty"`
	ImageUrls     *[]string `json:"image_urls" validate:"omitempty"`
	Version       *int      `json:"version"`
}

func updateProduct(productID int, p UpdateProductPayload, wg *sync.WaitGroup) {
	defer wg.Done()

	// Construct the URL for the update endpoint
	url := fmt.Sprintf("http://localhost:8080/api/v1/catalogue/%d", productID)

	// Create the JSON payload
	b, _ := json.Marshal(p)

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers as needed, for example:
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Update response status:", resp.Status)
}

func main() {
	var wg sync.WaitGroup

	// Assuming the post ID to update is 1
	productID := 440

	// Simulate User A and User B updating the same post concurrently
	wg.Add(3)
	name := "Product A"
	slug := "product-a"
	description := "Product A description"

	go updateProduct(productID, UpdateProductPayload{Name: &name}, &wg)
	go updateProduct(productID, UpdateProductPayload{Slug: &slug}, &wg)
	go updateProduct(productID, UpdateProductPayload{Description: &description}, &wg)

	wg.Wait()
}
