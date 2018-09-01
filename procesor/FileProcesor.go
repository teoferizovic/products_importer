package procesor

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"products_importer/model"
)

var externalProducts []model.ExtProduct

func ReadFile(conf model.Config) ([]model.ExtProduct,error) {

	files, err := ioutil.ReadDir(conf.FilePath)

	if err != nil {
		return nil, err
	}

	for _, f := range files {

		csvFile, _ := os.Open(conf.FilePath+f.Name())
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
			extProduct := model.ExtProduct{prod[i] ,cat[i],1,2,img[i]}
			externalProducts = append(externalProducts, extProduct)
		}

	}

	return externalProducts, nil


}


