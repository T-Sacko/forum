package models

import "fmt"



func GetCategoriesID(categories []string) []int {
	Ids := []int{}

	for i, v := range categories{
		var id int
		err := db.QueryRow("SELECT id FROM categories WHERE name = ?",v).Scan(&id)
		if err!= nil {
			fmt.Println("cant get categories, Error: ", err)
		}
		fmt.Printf("inserted category %v\n",i)
		Ids = append(Ids, id)
	}
	fmt.Println(Ids)
	return Ids
}

func LinkPostCategories(postId int, categoryIds []int){
	for i:=0;i<len(categoryIds);i++{
		_, err := db.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postId,categoryIds[i])
		if err!= nil{
			fmt.Println("Error inserting into post_categories table: ", err)
		}
		fmt.Printf("inserted post_category %v\n",categoryIds[i])
	}
}




// func SaveCategory(category string) {
// 	//db.Exec("SELECT FROM CATEGOR")
// 	_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", category)
// 	if err != nil {
// 		fmt.Println("Error inserting category: ",err)
// 		return
// 	}
// 	fmt.Println("Successfully inserted category!!!!!")
// }

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
