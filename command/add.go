package command

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type AddCommand struct {
	Meta
}

func (c *AddCommand) Run(args []string) int {
	title := strings.Join(args, "")
	fmt.Println("added new task :", title)

	db, err := gorm.Open("sqlite3", dbPath())
	if err != nil {
		return 1
	}
	defer db.Close()

	db.Create(&Todo{Title: title, IsDone: 0})

	if err != nil {
		return 1
	}
	return 0
}

func (c *AddCommand) Synopsis() string {
	return ""
}

func (c *AddCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
