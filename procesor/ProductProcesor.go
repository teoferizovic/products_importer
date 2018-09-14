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

	var products []model.Product
	mapProducts := make(map[string]int)


	rows, err := db.Query("SELECT p.id,p.external_name FROM products as p WHERE p.external_name IS NOT NULL")

	if err != nil {
		return nil, err
	}

	var id int
	var externalName sql.NullString

	for rows.Next() {

		err = rows.Scan(&id,&externalName)

		if err != nil {
			return nil,err
		}

		products = append(products, model.Product{ID:id,ExternalName:externalName.String,})
	}

	for _, prod := range products {
		mapProducts[prod.ExternalName] = prod.ID
	}

	return mapProducts,nil
}

//https://stackoverflow.com/questions/548541/insert-ignore-vs-insert-on-duplicate-key-update