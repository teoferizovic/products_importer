package procesor

import (
	"database/sql"
	"fmt"
	"products_importer/model"
)


func InsertCategories(db *sql.DB) error {

	fmt.Println("aaaa")

	//var vs []model.Category

	vs := []model.Category{{"Category111","hh"},{"Category2112","gaga"}}

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
	/*v := model.Category{"Category11",""}

	//stmt, err := db.Prepare("INSERT INTO categories VALUES(?,?)")
	stmt, err := db.Prepare("INSERT categories SET name=?,description=?")

	if err != nil {
		fmt.Println("error",err)
		return err
	}

	_, err = stmt.Exec(v.Name,v.Description)

	if err != nil {
		fmt.Println("error",err)
		return err
	}*/
	//_, err = stmt.Exec(v)

	return nil
}

func AllCategories() error {
	return nil
}

//https://stackoverflow.com/questions/548541/insert-ignore-vs-insert-on-duplicate-key-update