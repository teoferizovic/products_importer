package procesor

import (
	"io/ioutil"
	"log"
	"products_importer/model"
)

func ReadFile(conf model.Config) []string {

	var filesSlice []string

	files, err := ioutil.ReadDir(conf.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//fmt.Println(f.Name())
		filesSlice = append(filesSlice, f.Name())
	}

	return filesSlice


}
