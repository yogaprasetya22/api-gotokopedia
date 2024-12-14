package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/yogaprasetya22/api-gotokopedia/internal/db/dummy"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
	"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
	"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
	"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
	"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
	"walter", "xenia", "yasmin", "zoe",
}

var comments = []string{
	"Produk bagus, sesuai deskripsi.",
	"Pengiriman cepat, barang aman.",
	"Penjual ramah dan responsif.",
	"Harga terjangkau, kualitas oke.",
	"Barang diterima dalam kondisi baik.",
	"Terima kasih, sangat memuaskan.",
	"Produk asli, tidak mengecewakan.",
	"Rekomendasi, penjual terpercaya.",
	"Barang sesuai gambar, mantap.",
	"Pengemasan rapi, aman sampai tujuan.",
	"Produk berkualitas, sesuai harapan.",
	"Respon cepat, pengiriman kilat.",
	"Barang bagus, harga bersaing.",
	"Penjual sangat membantu, terima kasih.",
	"Produk sesuai ekspektasi, puas.",
	"Barang diterima dengan baik, terima kasih.",
	"Pengiriman cepat, barang sesuai.",
	"Penjual responsif, barang bagus.",
	"Produk original, sangat puas.",
	"Barang sesuai deskripsi, pengiriman cepat.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	/// Users
	users := generateUsers(10)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	/// Categories
	categories := generateCategories()
	for _, category := range categories {
		if err := store.Categoris.Create(ctx, category); err != nil {
			log.Println("Error creating category:", err)
			return
		}
	}

	/// Products dan Toko
	tokos := generateToko()
	for _, toko := range tokos {
		if err := store.Tokos.Create(ctx, toko); err != nil {
			log.Println("Error creating toko:", err)
			return
		}
	}
	products := generateProducts(store, tokos)
	for _, product := range products {
		if err := store.Products.Create(ctx, product); err != nil {
			log.Println("Error creating product:", err)
			return
		}
	}

	comments := generateComments(1000, users, products)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding completed.")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
		if err := users[i].Password.Set("asdasdasd"); err != nil {
			log.Println("Error setting password:", err)
			return nil
		}
	}

	return users
}

func generateCategories() []*store.Category {
	var categoryDummy []Category

	if err := json.Unmarshal([]byte(CategoryData), &categoryDummy); err != nil {
		log.Fatalf("Error unmarshalling categories: %v", err)
	}

	categories := make([]*store.Category, len(categoryDummy))

	for i, category := range categoryDummy {
		categories[i] = &store.Category{
			Name:        category.Name,
			Slug:        category.Slug,
			Description: category.Description,
		}
	}

	return categories
}

func generateToko() []*store.Toko {
	var allProductDummy []dummy.Product

	// Gabungkan semua data dummy
	dummyDataList := []string{
		dummy.CasingData,
		dummy.PcGamingData,
		dummy.OtomotifData,
		dummy.HandphoneData,
		dummy.DekorasiKamarData,
	}

	for _, data := range dummyDataList {
		var productDummy []dummy.Product
		if err := json.Unmarshal([]byte(data), &productDummy); err != nil {
			log.Fatalf("Error unmarshalling products: %v", err)
		}
		log.Printf("Unmarshalled %d products", len(productDummy))
		allProductDummy = append(allProductDummy, productDummy...)
	}

	toko := make([]*store.Toko, len(allProductDummy))

	for i, product := range allProductDummy {
		if product.ProductName == "" {
			log.Printf("Skipping nil product at index %d", i)
			continue
		}

		// Properti toko
		toko[i] = &store.Toko{
			Slug:         product.Toko.Username,
			Name:         product.Toko.TokoName,
			ImageProfile: product.Toko.ImageToko,
			Country:      product.Country,
			UserID:       rand.Int63n(10) + 1,
		}

	}
	return toko
}

func generateProducts(storage store.Storage, tokos []*store.Toko) []*store.Product {
	var allProductDummy []dummy.Product

	// Gabungkan semua data dummy
	dummyDataList := []string{
		dummy.CasingData,
		dummy.PcGamingData,
		dummy.OtomotifData,
		dummy.HandphoneData,
		dummy.DekorasiKamarData,
	}

	for _, data := range dummyDataList {
		var productDummy []dummy.Product
		if err := json.Unmarshal([]byte(data), &productDummy); err != nil {
			log.Fatalf("Error unmarshalling products: %v", err)
		}
		allProductDummy = append(allProductDummy, productDummy...)
	}

	products := make([]*store.Product, len(allProductDummy))

	for i, product := range allProductDummy {
		if product.ProductName == "" {
			log.Printf("Skipping nil product at index %d", i)
			continue
		}

		// Konversi harga menjadi float64 tanpa karakter non-digit
		parsePrice := func(value string) float64 {
			cleaned := strings.ReplaceAll(value, ".", "")
			cleaned = strings.ReplaceAll(cleaned, ",", "")
			cleaned = strings.ReplaceAll(cleaned, "Rp", "")
			cleaned = strings.ReplaceAll(cleaned, " ", "")
			parsed, _ := strconv.ParseFloat(cleaned, 64)
			return parsed
		}

		// Konversi rating menjadi float64
		parseFloat := func(value string) float64 {
			if value == "null" || value == "" {
				return 0
			}
			cleaned := strings.ReplaceAll(value, ".", "")
			cleaned = strings.ReplaceAll(cleaned, ",", "")
			parsed, _ := strconv.ParseFloat(cleaned, 64)
			return parsed
		}

		// Konversi stok menjadi int
		parseStock := func(value string) int {
			if value == "Habis" {
				return 100
			}
			cleaned := strings.ReplaceAll(value, ".", "")
			cleaned = strings.ReplaceAll(cleaned, ",", "")
			cleaned = strings.ReplaceAll(cleaned, " ", "")
			parsed, _ := strconv.Atoi(cleaned)
			return parsed
		}

		// Konversi sold menjadi int
		parseInt := func(value string) int {
			if value == "null" || value == "" {
				return 0
			}
			cleaned := strings.ReplaceAll(value, ".", "")
			cleaned = strings.ReplaceAll(cleaned, ",", "")
			parsed, _ := strconv.Atoi(cleaned)
			return parsed
		}

		// Dapatkan kategori berdasarkan slug
		category, err := storage.Categoris.GetBySlug(context.Background(), product.Category)
		if err != nil {
			log.Fatalf("Error getting category by slug: %v", err)
		}

		// Dapatkan toko berdasarkan slug
		var toko *store.Toko
		for _, t := range tokos {
			if t.Slug == product.Toko.Username {
				toko = t
				break
			}
		}

		if toko == nil {
			log.Fatalf("Toko with slug %s not found", product.Toko.Username)
		}

		products[i] = &store.Product{
			Name:          product.ProductName,
			Slug:          product.Slug,
			Description:   product.Description,
			Price:         parsePrice(product.Price),
			DiscountPrice: parseFloat(product.DiscountPrice),
			Discount:      parseFloat(product.Discount),
			Rating:        parseFloat(product.Rating),
			Estimation:    product.Estimation,
			Stock:         parseStock(product.Stock),
			Sold:          parseInt(product.Sold),
			IsForSale:     product.Discount != "" && product.Discount != "null",
			IsApproved:    true,
			ImageUrls:     product.ImageURL,
			Category:      *category,
			Toko:          *toko,
		}
	}

	return products
}

func generateComments(num int, users []*store.User, products []*store.Product) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			ProductID: products[rand.Intn(len(products))].ID,
			UserID:    users[rand.Intn(len(users))].ID,
			Content:   comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
