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

type IPTAConfig struct {
	Main struct {
		Db_Type  string
		Db_Host  string
		Db_User  string
		Db_Pass  string
		Db_Table string
	}
}

func read_config(filename string) {
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
}

func main() {
	var syslogflag string
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	read_config(usr.HomeDir + "/.ipta")

	flag.StringVar(&syslogflag, "import", "", "To import syslog you need to specify a filename")
	flag.StringVar(&syslogflag, "i", "", "To import syslog you need to specify a filename")

	flag.Parse()
	if syslogflag != "" {
		fmt.Printf("Arg = %s\n", syslogflag)
		import_syslog(syslogflag)
	}
}
