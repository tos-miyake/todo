package command

import (
	"flag"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {

	var all bool
	flags := flag.NewFlagSet("delete", flag.ContinueOnError)
	flags.BoolVar(&all, "all", false, "select all mode")
	flags.BoolVar(&all, "a", false, "select all mode")

	db, err := gorm.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if all {
		db.Delete(Todo{})
	} else {
		var id string
		if len(args) == 1 {
			id = args[0]
		}

		var todo Todo
		db.Find(&todo, "id = ?", id)
		db.Delete(&todo)
		if err != nil {
			log.Fatal(err)
		}
	}
	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return ""
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
