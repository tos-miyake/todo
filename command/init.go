package command

import (
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

type InitCommand struct {
	Meta
}

type Todo struct {
	gorm.Model
	Title  string
	IsDone int
}

func (c *InitCommand) Run(args []string) int {

	isDatabaseExists := exists(dbPath())

	if isDatabaseExists {
		println("Datab ase is already existed")
		return 1
	}

	db, err := gorm.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if db.HasTable(&Todo{}) {
		db.CreateTable(&Todo{})
	}
	db.AutoMigrate(&Todo{})

	return 0
}

func (c *InitCommand) Synopsis() string {
	return ""
}

func (c *InitCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
