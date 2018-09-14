package procesor

import (
	"database/sql"
	"fmt"
	"products_importer/model"
	"strings"
)


func InsertProducts(db *sql.DB,extProducts []model.ExtProduct,categories map[string]int) error {

	vs := []model.Product{}

	for _, extProd := range extProducts {

		prodName := strings.Split(extProd.ProductName, "##")
		categoryId := categories[extProd.ProductCategory]

		item := model.Product{Name:prodName[0],Status:extProd.Status,CategoryId:categoryId,Price:extProd.Price,ExternalName:extProd.ProductName}
		vs = append(vs,item)

	}

	sqlStr := "INSERT INTO products(name,status,category_id,price,external_name) VALUES "

	vals := []interface{}{}

	for _, row := range vs {
		sqlStr += "(?, ?, ?, ?, ?),"
		vals = append(vals, row.Name, row.Status, row.CategoryId,row.Price,row.ExternalName)
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

func AllProducts(db *sql.DB) (map[string]int,error) {

	var categories []model.Category
	mapCategories := make(map[string]int)


	rows, err := db.Query("SELECT c.id,c.name,c.description FROM categories as c")

	if err != nil {
		return nil, err
	}

	var id int
	var name,description sql.NullString

	for rows.Next() {

		err = rows.Scan(&id,&name, &description)

		if err != nil {
			return nil,err
		}

		categories = append(categories, model.Category{ID:id,Name:name.String,Description:description.String})
	}

	for _, cat := range categories {
		mapCategories[cat.Name] = cat.ID
	}

	return mapCategories,nil
}

//https://stackoverflow.com/questions/548541/insert-ignore-vs-insert-on-duplicate-key-update