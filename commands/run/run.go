package run

import (
	"flag"
	"fmt"
	"time"

	"github.com/Wjinlei/golib"
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
	now := time.Now()
	public.Logfile = golib.FormatNowTime("2006-01-02") + ".log"

	conn, err := public.GetConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	tables, _, callback, err := conn.QueryRows("SHOW TABLES")
	if err != nil {
		return err
	}

	for tables.Next() {
		conn.FindScript(callback(0))
	}

	golib.FileWrite(public.Logfile, fmt.Sprintf("[%s] 扫描耗时: %v\n", golib.GetNowTime(), time.Since(now)), golib.FileAppend)
	golib.FileWrite(public.Logfile, "------------------------------------------------\n", golib.FileAppend)
	return nil
}
