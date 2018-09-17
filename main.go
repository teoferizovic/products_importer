package main



import (
	"database/sql"
	"fmt"
	"products_importer/procesor"
	"github.com/BurntSushi/toml"
	"products_importer/model"
	_ "github.com/go-sql-driver/mysql"
)

var conf model.Config

func task() error {


	db, err := sql.Open("mysql", conf.Username+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.DB)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	extProducts,err := procesor.ReadFile(db,conf)
	if err != nil {
		panic(err.Error())
	}

	if len(extProducts) == 0 {
		return nil
	}

	err = procesor.InsertCategories(db,extProducts)
	if err != nil {
		panic(err.Error())
	}

	categories,err:=procesor.AllCategories(db)
	if err != nil {
		panic(err.Error())
	}

	err = procesor.InsertProducts(db,extProducts,categories)
	if err != nil {
		panic(err.Error())
	}

	products,err := procesor.AllProducts(db)
	if err != nil {
		panic(err.Error())
	}

	err = procesor.InsertProductImages(db,products,extProducts,conf)
	if err != nil {
		panic(err.Error())
	}

	return nil

}

func init(){

	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", conf)
}

func main() {
	/*s := gocron.NewScheduler()
	s.Every(2).Seconds().Do(task)
	<- s.Start()*/
	task();
}


//https://github.com/jasonlvhit/gocron
