package models

type Item struct {
	ID          int
	CategoryID  int
	Title       string
	Description string
	ImageURL    string
	Price       string
}

type Category struct {
	ID    int
	URL   string
	Title string
	items []Item
}

var categories []Category = []Category{
	Category{
		ID:    1,
		URL:   "#drinks",
		Title: "Drinks",
	},
	Category{
		ID:    2,
		URL:   "#lunch",
		Title: "Lunch",
	},
	Category{
		ID:    3,
		URL:   "#dinner",
		Title: "Dinner",
	},
}

var items []Item = []Item{
	Item{
		ID:          1,
		CategoryID:  2,
		ImageURL:    "images/img-01.jpg",
		Title:       "Special Drinks 1",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$7.79",
	},
	Item{
		ID:          2,
		CategoryID:  2,
		ImageURL:    "images/img-02.jpg",
		Title:       "Special Drinks 2",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$9.79",
	},
	Item{
		ID:          3,
		CategoryID:  2,
		ImageURL:    "images/img-03.jpg",
		Title:       "Special Drinks 3",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$10.79",
	},
	Item{
		ID:          4,
		CategoryID:  3,
		ImageURL:    "images/img-04.jpg",
		Title:       "Special Lunch 1",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$15.79",
	},
	Item{
		ID:          5,
		CategoryID:  3,
		ImageURL:    "images/img-05.jpg",
		Title:       "Special Lunch 2",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$18.79",
	},
	Item{
		ID:          6,
		CategoryID:  3,
		ImageURL:    "images/img-06.jpg",
		Title:       "Special Lunch 3",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$20.79",
	},
	Item{
		ID:          7,
		CategoryID:  4,
		ImageURL:    "images/img-07.jpg",
		Title:       "Special Dinner 1",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$25.79",
	},
	Item{
		ID:          8,
		CategoryID:  4,
		ImageURL:    "images/img-08.jpg",
		Title:       "Special Dinner 2",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$22.79",
	},
	Item{
		ID:          9,
		CategoryID:  4,
		ImageURL:    "images/img-09.jpg",
		Title:       "Special Dinner 3",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$24.79",
	},
}

func GetCategories() []Category {
	for i := 1; i <= len(categories); i++ {
		categories[i-1].items = GetItemsforCategories(i)
	}
	return categories
}

func GetItemsforCategories(CategoryID int) []Item {
	var result []Item
	for _, i := range items {
		if i.CategoryID == CategoryID+1 {
			result = append(result, i)
		}
	}
	return result
}
