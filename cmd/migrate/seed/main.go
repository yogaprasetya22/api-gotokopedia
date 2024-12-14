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
// 	addr := env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost/gotokopedia?sslmode=disable")

// 	log.Println("Database address:", addr)

// 	conn, err := db.New(addr, 3, 3, "15m")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer conn.Close()

// 	storage := store.NewStorage(conn)

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
// 		log.Printf("Unmarshalled %d products", len(productDummy))
// 		allProductDummy = append(allProductDummy, productDummy...)
// 	}

// 	products := make([]*store.Product, len(allProductDummy))

// 	for i, product := range allProductDummy {
// 		if product.ProductName == "" {
// 			log.Printf("Skipping nil product at index %d", i)
// 			continue
// 		}

// 		// Konversi harga menjadi float64 tanpa karakter non-digit
// 		parsePrice := func(value string) float64 {
// 			cleaned := strings.ReplaceAll(value, ".", "")
// 			cleaned = strings.ReplaceAll(cleaned, ",", "")
// 			cleaned = strings.ReplaceAll(cleaned, "Rp", "")
// 			cleaned = strings.ReplaceAll(cleaned, " ", "")
// 			parsed, _ := strconv.ParseFloat(cleaned, 64)
// 			return parsed
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
// 		parseInt := func(value string) int {
// 			if value == "null" || value == "" {
// 				return 0
// 			}
// 			cleaned := strings.ReplaceAll(value, ".", "")
// 			cleaned = strings.ReplaceAll(cleaned, ",", "")
// 			parsed, _ := strconv.Atoi(cleaned)
// 			return parsed
// 		}

// 		// Properti toko
// 		toko := store.Toko{
// 			Slug:         product.Toko.Username,
// 			Name:         product.Toko.TokoName,
// 			ImageProfile: product.Toko.ImageToko,
// 			Country:      product.Country,
// 			UserID:       rand.Int63n(10) + 1,
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
// 			Price:         parsePrice(product.Price),
// 			DiscountPrice: parseFloat(product.DiscountPrice),
// 			Discount:      parseFloat(product.Discount),
// 			Rating:        parseFloat(product.Rating),
// 			Estimation:    product.Estimation,
// 			Stock:         parseStock(product.Stock),
// 			Sold:          parseInt(product.Sold),
// 			IsForSale:     product.Discount != "" && product.Discount != "null",
// 			IsApproved:    true,
// 			ImageUrls:     product.ImageURL,
// 			Category:      *category,
// 			Toko:          toko,
// 		}
// 	}

// 	log.Println("Total products:", len(products))
// 	log.Println("Total Tokos:", len(allProductDummy))
// }
