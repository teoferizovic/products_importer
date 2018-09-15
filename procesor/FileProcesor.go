package procesor

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"products_importer/helper"
	"products_importer/model"
	"time"
)

var externalProducts []model.ExtProduct
var uploadfiles []model.UploadFile

func ReadFile(db *sql.DB,conf model.Config) ([]model.ExtProduct,error) {

	currentFiles :=[]string{}

	storedFiles,err := AllFiles(db)

	if err != nil {
		return nil,err
	}

	files, err := ioutil.ReadDir(conf.FilePath)

	if err != nil {
		return nil, err
	}

	for _, f := range files {
		currentFiles = append(currentFiles,f.Name())
	}


	newFiles := helper.DifferenceSlices(currentFiles,storedFiles)

	if len(newFiles) == 0 {
		return externalProducts, nil
	}

	err = InsertFiles(db,newFiles)

	if err != nil{
		return nil, err
	}

	t := time.Now()

	for _, f := range newFiles {

		csvFile, _ := os.Open(conf.FilePath+f)
		reader := csv.NewReader(bufio.NewReader(csvFile))

		lines, err := reader.ReadAll()

		if err != nil {
			fmt.Println("error reading all lines: %v", err)
			return nil, err
		}

		prod := make([]string, len(lines)-1)
		cat := make([]string, len(lines)-1)
		stat := make([]string, len(lines)-1)
		pric := make([]string, len(lines)-1)
		img := make([]string, len(lines)-1)

		for i, line := range lines {

			// skip header line
			if i == 0 {
				continue
			}

			prod[i-1] = line[0]
			cat[i-1] = line[1]
			stat[i-1] = line[2]
			pric[i-1] = line[3]
			img[i-1] = line[4]

		}

		for i, _ := range prod {
			extProduct := model.ExtProduct{prod[i]+"##"+t.Format("20060102150405") ,cat[i],1,2,img[i]}
			externalProducts = append(externalProducts, extProduct)
		}

	}
	//fmt.Println(externalProducts)
	return externalProducts, nil


}

func AllFiles(db *sql.DB) ([]string,error) {

	var uploadFiles []model.UploadFile
	files := []string{}

	rows, err := db.Query("SELECT f.id,f.name FROM upload_files as f")

	if err != nil {
		return nil, err
	}

	var id int
	var name string

	for rows.Next() {

		err = rows.Scan(&id,&name)

		if err != nil {
			return nil,err
		}

		uploadFiles = append(uploadFiles, model.UploadFile{ID:id,Name:name})
	}

	for _, file := range uploadFiles {
		files = append(files,file.Name)
	}

	return files,nil

}


func InsertFiles(db *sql.DB,newFiles []string) error {

	vs := []model.UploadFile{}

	for _, file := range newFiles {


		item := model.UploadFile{Name:file}
		vs = append(vs,item)

	}

	sqlStr := "INSERT INTO upload_files(name) VALUES "

	vals := []interface{}{}

	for _, row := range vs {
		sqlStr += "(?),"
		vals = append(vals, row.Name)
	}

	//trim the last ,
	sqlStr = sqlStr[0:len(sqlStr)-1]

	//prepare the statement
	stmt, _ := db.Prepare(sqlStr)

	//format all vals at once
	_,err := stmt.Exec(vals...)

	if err != nil {
		return err
	}

	return nil
}
