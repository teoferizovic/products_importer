package procesor

import (
	"database/sql"
	"fmt"
	"products_importer/model"
)


func InsertCategories(db *sql.DB) error {

	fmt.Println("aaaa")

	//var vs []model.Category

	vs := []model.Category{{Name:"Category22",Description:"kk"},{Name:"Category45",Description:"wewewe"}}

	//fmt.Println(vs)

	sqlStr := "INSERT IGNORE INTO categories(name, description) VALUES "

	vals := []interface{}{}

	for _, row := range vs {
		sqlStr += "(?, ?),"
		vals = append(vals, row.Name, row.Description)
	}

	//trim the last ,
	sqlStr = sqlStr[0:len(sqlStr)-1]

	//prepare the statement
	stmt, _ := db.Prepare(sqlStr)

	//format all vals at once
	_,err := stmt.Exec(vals...)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func AllCategories(db *sql.DB) (map[string]int,error) {

	var categories []model.Category
	mapCategories := make(map[string]int)


	rows, err := db.Query("SELECT c.id,c.name,c.description FROM categories as c")

	if err != nil {
		return nil, err
	}

	var id int
	var name,description string

	for rows.Next() {

		err = rows.Scan(&id,&name, &description)

		if err != nil {
			return nil,err
		}

		categories = append(categories, model.Category{ID:id,Name:name,Description:description})
	}

	for _, cat := range categories {
		mapCategories[cat.Name] = cat.ID
	}

	return mapCategories,nil
}

//https://stackoverflow.com/questions/548541/insert-ignore-vs-insert-on-duplicate-key-update