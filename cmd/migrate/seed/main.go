package main

import (
	"log"

	"github.com/yogaprasetya22/api-gotokopedia/internal/db"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost/gotokopedia?sslmode=disable")

	log.Println("Database address:", addr)

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}

// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"math/rand"
// 	"strconv"
// 	"strings"

// 	"github.com/yogaprasetya22/api-gotokopedia/internal/db"
// 	"github.com/yogaprasetya22/api-gotokopedia/internal/db/dummy"
// 	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
// 	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
// )

// func main() {
// 	// mau melihat log generated products
// 	addr := env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost/gotokopedia?sslmode=disable")

// 	log.Println("Database address:", addr)

// 	conn, err := db.New(addr, 3, 3, "15m")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer conn.Close()

// 	store := store.NewStorage(conn)

// 	// Generate products
// 	products := generateProducts(store)
// 	for _, product := range products {
// 		log.Printf("Sold: %d", product.Sold)
// 		log.Printf("Discount product: %+v", product.Discount)
// 		log.Printf("Discount product: %+v", product.DiscountPrice)
// 		log.Printf("Price: %+v", product.Price)
// 	}

// }

// func generateProducts(storage store.Storage) []*store.Product {
// 	var allProductDummy []dummy.Product

// 	// Gabungkan semua data dummy
// 	dummyDataList := []string{
// 		dummy.CasingData,
// 		dummy.PcGamingData,
// 		dummy.OtomotifData,
// 		dummy.HandphoneData,
// 		dummy.DekorasiKamarData,
// 	}

// 	for _, data := range dummyDataList {
// 		var productDummy []dummy.Product
// 		if err := json.Unmarshal([]byte(data), &productDummy); err != nil {
// 			log.Fatalf("Error unmarshalling products: %v", err)
// 		}
// 		allProductDummy = append(allProductDummy, productDummy...)
// 	}

// 	products := make([]*store.Product, len(allProductDummy))

// 	for i, product := range allProductDummy {
// 		if product.ProductName == "" {
// 			log.Printf("Skipping nil product at index %d", i)
// 			continue
// 		}


// 		// Konversi rating menjadi float64
// 		parseFloat := func(value string) float64 {
// 			if value == "null" || value == "" {
// 				return 0
// 			}
// 			cleaned := strings.ReplaceAll(value, ".", "")
// 			cleaned = strings.ReplaceAll(cleaned, ",", "")
// 			parsed, _ := strconv.ParseFloat(cleaned, 64)
// 			return parsed
// 		}

// 		// Konversi diskon menjadi float64 "10%"
// 		parseDiscount := func(value string) float64 {
// 			if value == "null" || value == "" {
// 				return 0
// 			}
// 			cleaned := strings.ReplaceAll(value, "%", "")
// 			cleaned = strings.ReplaceAll(cleaned, ",", "")
// 			parsed, _ := strconv.ParseFloat(cleaned, 64)
// 			return parsed
// 		}

// 		// konversi harga diskon menjadi float64 "Rp 1.000.000"
// 		parcePrice := func(value string) float64 {
// 			if value == "" || value == "null" {
// 				return 0
// 			}
// 			cleaned := strings.ReplaceAll(value, ".", "")
// 			cleaned = strings.ReplaceAll(cleaned, "Rp", "")
// 			cleaned = strings.ReplaceAll(cleaned, " ", "")
// 			parsed, err := strconv.ParseFloat(cleaned, 64)
// 			if err != nil {
// 				log.Printf("Error parsing discount price: %v", err)
// 				return 0
// 			}
// 			return parsed
// 		}

// 		// Konversi stok menjadi int
// 		parseStock := func(value string) int {
// 			if value == "Habis" {
// 				return 100
// 			}
// 			cleaned := strings.ReplaceAll(value, ".", "")
// 			cleaned = strings.ReplaceAll(cleaned, ",", "")
// 			cleaned = strings.ReplaceAll(cleaned, " ", "")
// 			parsed, _ := strconv.Atoi(cleaned)
// 			return parsed
// 		}

// 		// Konversi sold menjadi int
// 		parseSold := func(value string) int {
// 			if value == "null" || value == "" {
// 				return 0
// 			}
// 			cleaned := strings.ReplaceAll(value, "Terjual ", "")
// 			cleaned = strings.ReplaceAll(cleaned, "+", "")
// 			cleaned = strings.ReplaceAll(cleaned, " ", "")
// 			baseSold, _ := strconv.Atoi(cleaned)
// 			if strings.Contains(value, "+") {
// 				return baseSold + rand.Intn(100) // Tambahkan nilai acak antara 0 dan 99
// 			}
// 			return baseSold
// 		}

// 		// Dapatkan kategori berdasarkan slug
// 		category, err := storage.Categoris.GetBySlug(context.Background(), product.Category)
// 		if err != nil {
// 			log.Fatalf("Error getting category by slug: %v", err)
// 		}

// 		products[i] = &store.Product{
// 			Name:          product.ProductName,
// 			Slug:          product.Slug,
// 			Description:   product.Description,
// 			Country:       product.Country,
// 			Price:         parcePrice(product.Price),
// 			DiscountPrice: parcePrice(product.DiscountPrice),
// 			Discount:      parseDiscount(product.Discount),
// 			Rating:        parseFloat(product.Rating),
// 			Estimation:    product.Estimation,
// 			Stock:         parseStock(product.Stock),
// 			Sold:          parseSold(product.Sold),
// 			IsForSale:     product.Discount != "" && product.Discount != "null",
// 			IsApproved:    true,
// 			ImageUrls:     product.ImageURL,
// 			Category:      *category,
// 		}
// 	}

// 	return products
// }
