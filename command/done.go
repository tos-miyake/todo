package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

type DoneCommand struct {
	Meta
}

func (c *DoneCommand) Run(args []string) int {
	if len(args) != 1 {
		return 1
	}
	id := args[0]
	fmt.Printf("Task %s is done\n", id)

	db, err := gorm.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var todo Todo
	db.Find(&todo, "id = ?", id)
	db.Model(&todo).Update("IsDone", 1)
	if err != nil {
		log.Fatal(err)
	}
	return 0
}

func (c *DoneCommand) Synopsis() string {
	return ""
}

func (c *DoneCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
