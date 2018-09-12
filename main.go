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


	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	extProducts,err := procesor.ReadFile(conf)

	if err != nil {
		panic(err.Error())
	}

	err = procesor.InsertCategories(db,extProducts)

	if err != nil {
		panic(err.Error())
	}

	categories,err:=procesor.AllCategories(db)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(categories)

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
