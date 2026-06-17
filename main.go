package main




import (
	"database/sql"
	"fmt"
	"github.com/adrake333/gatorcli/internal/config"
	"github.com/adrake333/gatorcli/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)




func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	appState := state{
		db: dbQueries,
		cfg: &cfg,
	}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("invalid command")
		os.Exit(1)
	}
	cmds.register("register", handlerRegister)
	if len(os.Args) < 2 {
		fmt.Println("invalid command")
		os.Exit(1)
	}
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)
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
