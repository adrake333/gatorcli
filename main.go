package main




import (
	"fmt"
	"github.com/adrake333/gatorcli/internal/config"
)




func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.SetUser("adam")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", cfg)
}
