package viewmodel

import (
	"models"
)

type Menu struct {
	Title       string
	Description string
	Types       []Category
}

type Category struct {
	ID    int
	URL   string
	Title string
	items []Item
}

type Item struct {
	ID          int
	CategoryID  int
	Title       string
	Description string
	ImageURL    string
	Price       string
}

func CreateMenu(categories []models.Category) Menu {
	result := Menu{
		Title:       "Our Special Menu",
		Description: "A List of all our Special Dishes",
	}
	result.Types = make([]Category, len(categories))
	for i := 1; i <= len(categories); i++ {
		vm := categorytoVM(categories[i-1])
		vm.items = AddItems(i)
		result.Types[i-1] = vm
	}
	return result
}

func categorytoVM(c models.Category) Category {
	return Category{
		URL:   c.URL,
		Title: c.Title,
		ID:    c.ID,
	}
}

//
// func AddtoAll(c Category) {
// 	items := models.GetItemsforAll()
// 	for _, i := range items {
// 		c.items = append(c.items, ItemtoVM(i))
// 	}
// }
//
func AddItems(c int) []Item {
	items := models.GetItemsforCategories(c)
	var result []Item
	for _, i := range items {
		if i.CategoryID == c {
			result = append(result, ItemtoVM(i))
		}
	}
	return result
}

func ItemstoVM(a []models.Item) []Item {
	var contents []Item
	for _, i := range a {
		contents = append(contents, ItemtoVM(i))
	}
	return contents
}

func ItemtoVM(i models.Item) Item {
	return Item{
		ID:          i.ID,
		CategoryID:  i.CategoryID,
		Title:       i.Title,
		Description: i.Description,
		ImageURL:    i.ImageURL,
		Price:       i.Price,
	}
}
