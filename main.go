package main




import (
	"fmt"
	"github.com/adrake333/gatorcli/internal/config"
	"log"
	"os"
)




func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	appState := state{
		cfg: cfg,
	}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("invalid command")
		os.Exit(1)
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]
	cmd := command{
		name: commandName,
		args: commandArgs,
	}
	err = cmds.run(&appState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
