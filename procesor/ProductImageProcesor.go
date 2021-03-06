package procesor

import (
	"database/sql"
	"fmt"
	"products_importer/model"
)

func InsertProductImages(db *sql.DB,products map[string]int,extProducts []model.ExtProduct,conf model.Config) error {

	vs := []model.ProductImage{}

	for _, extProd := range extProducts {

		productId := products[extProd.ProductName]
		var imgName string

		if len(extProd.Image) == 0 {
			imgName = "default.jpg"
		} else {
			imgName = extProd.Image
		}

		path := conf.ImagePath+imgName

		item := model.ProductImage{Path:path,ProductId:productId}
		vs = append(vs,item)

	}

	sqlStr := "INSERT INTO product_images(path,product_id) VALUES "

	vals := []interface{}{}

	for _, row := range vs {
		sqlStr += "(?, ?),"
		vals = append(vals, row.Path,row.ProductId)
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
