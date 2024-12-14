package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type UpdateProductPayload struct {
	Name          *string  `json:"name"`
	Slug          *string  `json:"slug"`
	Description   *string  `json:"description,omitempty"`
	Price         *float64 `json:"price"`
	DiscountPrice *float64 `json:"discount_price"`
	Discount      *float64 `json:"discount"`
	Rating        *float64 `json:"rating"`
	Estimation    *string  `json:"estimation"`
	Stock         *int     `json:"stock"`
	Sold          *int     `json:"sold"`
	IsForSale     *bool    `json:"is_for_sale"`
	IsApproved    *bool    `json:"is_approved"`
	ImageUrls     []string `json:"image_urls"`
}

func updatePost(productID int, p UpdateProductPayload, wg *sync.WaitGroup) {
	defer wg.Done()

	// Construct the URL for the update endpoint
	url := fmt.Sprintf("http://localhost:8080/api/v1/product/%d", productID)

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
	productID := 3

	// Simulate User A and User B updating the same post concurrently
	wg.Add(3)
	name := "Product A"
	slug := "product-a"
	description := "Product A description"

	go updatePost(productID, UpdateProductPayload{Name: &name}, &wg)
	go updatePost(productID, UpdateProductPayload{Slug: &slug}, &wg)
	go updatePost(productID, UpdateProductPayload{Description: &description}, &wg)

	wg.Wait()
}
