package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

type UnDoneCommand struct {
	Meta
}

func (c *UnDoneCommand) Run(args []string) int {
	if len(args) != 1 {
		return 1
	}
	id := args[0]
	fmt.Printf("Task %s is undone\n", id)

	db, err := gorm.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var todo Todo
	db.Find(&todo, "id = ?", id)
	db.Model(&todo).Update("IsDone", 0)
	if err != nil {
		log.Fatal(err)
	}
	return 0
}

func (c *UnDoneCommand) Synopsis() string {
	return ""
}

func (c *UnDoneCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
