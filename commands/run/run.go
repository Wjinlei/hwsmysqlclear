package run

import (
	"flag"
	"fmt"
	"strings"
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
	include  string // Include tables, comma separated
	exclude  string // Exclude tables, comma separated
	ntime    int64  // How many seconds between scans
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
	runCommand.FlagSet.Int64Var(&opt.ntime, "t", 10, `How many seconds between scans`)
	runCommand.FlagSet.StringVar(&opt.include, "include", "", `Include tables, comma separated`)
	runCommand.FlagSet.StringVar(&opt.exclude, "exclude", "", `Exclude tables, comma separated`)
	runCommand.FlagSet.Usage = runCommand.Usage // use default usage provided by cmds.Command.
	cmds.AllCommands = append(cmds.AllCommands, runCommand)
}

type run struct{}

func (v *run) PreRun() error {
	return public.Connect(opt.user, opt.password, opt.dbname)
}

func (v *run) Run() error {
	if opt.ntime < 10 {
		opt.ntime = 10
	}

	for {
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
			table := callback(0)
			if contains(strings.Split(opt.exclude, ","), table) {
				continue
			}

			if opt.include == "" {
				conn.FindScript(table)
				continue
			}

			if contains(strings.Split(opt.include, ","), table) {
				conn.FindScript(table)
			}
		}

		if opt.exclude != "" {
			golib.FileWrite(
				public.Logfile,
				fmt.Sprintf("排除表: %v\n", opt.exclude),
				golib.FileAppend)
		}
		if opt.include != "" {
			golib.FileWrite(
				public.Logfile,
				fmt.Sprintf("指定表: %v\n", opt.include),
				golib.FileAppend)
		}
		golib.FileWrite(public.Logfile, "------------------------------------------------\n", golib.FileAppend)
		golib.FileWrite(
			public.Logfile,
			fmt.Sprintf("[%s] 扫描耗时: %v\n",
				golib.GetNowTime(),
				time.Since(now)),
			golib.FileAppend)
		golib.FileWrite(public.Logfile, "------------------------------------------------\n", golib.FileAppend)
		golib.FileWrite(public.Logfile, "\n", golib.FileAppend)

		time.Sleep(time.Duration(time.Second * time.Duration(opt.ntime)))
	}
}

func contains(keys []string, key string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}
