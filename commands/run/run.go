package run

import (
	"flag"
	"fmt"

	"github.com/Wjinlei/hwsmysqlclear/commands/public"
	"github.com/genshen/cmds"
	_ "github.com/go-sql-driver/mysql"
)

var runCommand = &cmds.Command{
	Name:        "run",
	Summary:     "running",
	Description: "running",
	CustomFlags: false,
	HasOptions:  true,
	FlagSet:     &flag.FlagSet{},
	Runner:      nil,
}

/* Option */
type Option struct {
	user     string
	password string
	dbname   string
}

var opt Option

func init() {
	opt = Option{}
	runCommand.Runner = &run{}
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	runCommand.FlagSet = fs
	runCommand.FlagSet.StringVar(&opt.user, "u", "root", `username`)
	runCommand.FlagSet.StringVar(&opt.password, "p", "", `password`)
	runCommand.FlagSet.StringVar(&opt.dbname, "db", "", `database name`)
	runCommand.FlagSet.Usage = runCommand.Usage // use default usage provided by cmds.Command.
	cmds.AllCommands = append(cmds.AllCommands, runCommand)
}

type run struct{}

func (v *run) PreRun() error {
	return public.Connect(opt.user, opt.password, opt.dbname)
}

func (v *run) Run() error {
	conn, err := public.GetConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	rows, columns, callback, err := conn.QueryRows("SHOW TABLES")
	if err != nil {
		return err
	}

	for rows.Next() {
		for i := range columns {
			fmt.Println(columns[i], ": ", callback(i))
		}
		fmt.Println("--------------------------")
	}
	return nil
}
