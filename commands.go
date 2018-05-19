package main

import (
	"github.com/mitchellh/cli"
	"github.com/tos-miyake/todo/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Meta: *meta,
			}, nil
		},
		"add": func() (cli.Command, error) {
			return &command.AddCommand{
				Meta: *meta,
			}, nil
		},
		"done": func() (cli.Command, error) {
			return &command.DoneCommand{
				Meta: *meta,
			}, nil
		},
		"undone": func() (cli.Command, error) {
			return &command.UnDoneCommand{
				Meta: *meta,
			}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
