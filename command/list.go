package command

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/olekukonko/tablewriter"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	var all bool
	flags := flag.NewFlagSet("list", flag.ContinueOnError)
	flags.BoolVar(&all, "all", false, "select all mode")
	flags.BoolVar(&all, "a", false, "select all mode")
	if err := flags.Parse(args); err != nil {
		return 1
	}
	db, err := gorm.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var todos []Todo
	if all {
		db.Table("todos").Find(&todos)
	} else {
		db.Where("is_done = 0").Find(&todos)
	}

	data := [][]string{}

	for _, todo := range todos {
		data =
			append(data, []string{strconv.FormatUint(uint64(todo.ID), 10), todo.Title, doneLabel(todo.IsDone), todo.CreatedAt.Format("2006-01-02 15:04:05")})
	}

	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No", "Title", "Status", "Created"})
		table.SetBorder(true)
		table.AppendBulk(data)
		table.Render()
	}

	return 0
}

func (c *ListCommand) Synopsis() string {
	return ""
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

func doneLabel(isDone int) string {
	if isDone == 0 {
		return "-"
	}
	return "Done"
}
