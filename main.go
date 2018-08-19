package main



import ("fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)


func task() {
	fmt.Println("I am runnning task.")

	// Open up our database connection.
	db, err := sql.Open("mysql", "username:root@tcp(127.0.0.1:3306)/test_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func main() {
	/*s := gocron.NewScheduler()
	s.Every(2).Seconds().Do(task)
	<- s.Start()*/
	task();
}


//https://github.com/jasonlvhit/gocron
