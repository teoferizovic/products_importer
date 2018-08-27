package main



import ("fmt"
	"products_importer/procesor"
	"github.com/BurntSushi/toml"
	"products_importer/model"
)

var conf model.Config

func task() {

	fmt.Println("I am runnning task.")
	files := procesor.ReadFile(conf)
	fmt.Println(files)
	//fmt.Println(model.Product{"edo",33,10,66});

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
