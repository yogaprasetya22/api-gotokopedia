package db

type Category struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

var CategoryData = `[
  {
    "name": "Pc Gaming",
    "slug": "pc-gaming",
    "description": "Pc Gaming Category"
  },
  {
    "name": "Otomotif",
    "slug": "otomotif",
    "description": "Otomotif Category"
  },
  {
    "name": "Handphone",
    "slug": "handphone",
    "description": "Handphone Category"
  },
  {
    "name": "Casing Hp",
    "slug": "casing-hp",
    "description": "Casing Hp Category"
  },
  {
    "name": "Dekorasi Kamar",
    "slug": "dekorasi-kamar",
    "description": "Dekorasi Kamar Category"
  }
]`