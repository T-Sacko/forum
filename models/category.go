package models

import "fmt"

func SaveCategory(category string) {
	_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", category)
	if err != nil {
		fmt.Println("Error inserting category: ",err)
		return
	}
	fmt.Println("Successfully inserted category!!!!!")
}

// type Category struct {
//     ID   int    `json:"id"`
//     Name string `json:"name"`
// }

// func (c *Category) Save() error {
//     // save the category to the database
// }

// func GetCategoryByID(id int) (*Category, error) {
//     // query the database for a category with the given ID
// }

// func GetCategories() ([]*Category, error) {
//     // query the database for all categories
// }
