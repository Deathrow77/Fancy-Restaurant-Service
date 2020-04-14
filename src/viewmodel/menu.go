package viewmodel

type Menu struct {
	Title       string
	Description string
	Active      string
	Types       []Categories
}

type Categories struct {
	URL      string
	Title    string
	Contents []Items
}

type Items struct {
	ImageURL    string
	Title       string
	Description string
	Price       string
}

func CreateMenu() Menu {

	result := Menu{Title: "Special Menu",
		Description: "Our Delicacies from around the world",
		Active:      "menu",
	}

	AllClass := Categories{URL: "",
		Title: "All",
	}

	DrinksClass := Categories{
		URL:   "",
		Title: "Drinks",
	}

	LunchClass := Categories{
		URL:   "",
		Title: "Lunch",
	}

	DinnerClass := Categories{
		URL:   "",
		Title: "Dinner",
	}

	itemDrink1 := Items{
		ImageURL:    "images/img-01.jpg",
		Title:       "Special Drinks 1",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$7.79",
	}
	itemDrink2 := Items{
		ImageURL:    "images/img-02.jpg",
		Title:       "Special Drinks 2",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$9.79",
	}
	itemDrink3 := Items{
		ImageURL:    "images/img-03.jpg",
		Title:       "Special Drinks 3",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$10.79",
	}
	itemLunch1 := Items{
		ImageURL:    "images/img-04.jpg",
		Title:       "Special Lunch 1",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$15.79",
	}
	itemLunch2 := Items{
		ImageURL:    "images/img-05.jpg",
		Title:       "Special Lunch 2",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$18.79",
	}
	itemLunch3 := Items{
		ImageURL:    "images/img-06.jpg",
		Title:       "Special Lunch 3",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$20.79",
	}
	itemDinner1 := Items{
		ImageURL:    "images/img-07.jpg",
		Title:       "Special Dinner 1",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$25.79",
	}
	itemDinner2 := Items{
		ImageURL:    "images/img-08.jpg",
		Title:       "Special Dinner 2",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$22.79",
	}
	itemDinner3 := Items{
		ImageURL:    "images/img-09.jpg",
		Title:       "Special Dinner 3",
		Description: "Sed id magna vitae eros sagittis euismod.",
		Price:       "$24.79",
	}
	AllClass.Contents = []Items{itemDrink1, itemDrink2, itemDrink3, itemLunch1, itemLunch2, itemLunch3, itemDinner1, itemDinner2, itemDinner3}
	DrinksClass.Contents = []Items{itemDrink1, itemDrink2, itemDrink3}
	LunchClass.Contents = []Items{itemLunch1, itemLunch2, itemLunch3}
	DinnerClass.Contents = []Items{itemDinner1, itemDinner2, itemDinner3}

	result.Types = []Categories{AllClass, DrinksClass, LunchClass, DinnerClass}
	return result
}
