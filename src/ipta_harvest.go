package main

import (
	"flag"
	"fmt"
)

func main() {
	var syslogflag string

	flag.StringVar(&syslogflag, "import", "", "To import syslog you need to specify a filename")
	flag.StringVar(&syslogflag, "i", "", "To import syslog you need to specify a filename")

	flag.Parse()
	if syslogflag != "" {
		fmt.Printf("Arg = %s\n", syslogflag)
		import_syslog(syslogflag)
	}
}
