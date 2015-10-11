package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"gopkg.in/gcfg.v1"
)

type Main struct {
	Db_Type            string
	Db_Name            string
	Db_Host            string
	Db_User            string
	Db_Pass            string
	Db_Table           string
	Db_Sqlite_Filename string
}

type IPTAConfig struct {
	Main
}

func read_config(filename string) IPTAConfig {
	cfgStr, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var cfg IPTAConfig

	err = gcfg.ReadStringInto(&cfg, string(cfgStr[:]))
	if err != nil {
		log.Fatal("Failed to parse gcfg data: %s", err)
		os.Exit(1)
	}
	fmt.Println(cfg.Main.Db_Type)

	return cfg
}

func main() {
	var syslogflag, followflag string
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	cfg := read_config(usr.HomeDir + "/.ipta")

	flag.StringVar(&syslogflag, "import", "", "To import syslog you need to specify a filename")
	flag.StringVar(&syslogflag, "i", "", "To import syslog you need to specify a filename")
	flag.StringVar(&followflag, "follow", "", "To follow syslog you need to specify a filename")
	flag.StringVar(&followflag, "f", "", "To follow syslog you need to specify a filename")

	longCreateF := flag.Bool("create-table", false, "")
	createF := flag.Bool("ct", false, "")
	longDeleteF := flag.Bool("delete-db", false, "")
	deleteF := flag.Bool("dt", false, "")
	longClear := flag.Bool("clear", false, "")
	clearF := flag.Bool("c", false, "")

	flag.Parse()
	if syslogflag != "" {
		fmt.Printf("Arg = %s\n", syslogflag)
		import_syslog(cfg, syslogflag)
	}

	if followflag != "" {
		fmt.Printf("Arg = %s\n", followflag)
		follow_syslog(cfg, followflag)
	}

	if *longCreateF || *createF {
		fmt.Printf("Arg = %s\n", syslogflag)

		create_table(cfg)
	}

	if *longDeleteF || *deleteF {
		fmt.Printf("Arg = %s\n", syslogflag)

		delete_table(cfg)
	}

	if *longClear || *clearF {
		clear_db(cfg)
	}
}
