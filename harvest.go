package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type FWItem struct {
	action string
	inIf   string
	outIf  string
	mac    string
	src    string
	dst    string
	proto  string
	spt    string
	dpt    string
	date   time.Time
}

var months = map[string]time.Month{
	"Jan": time.January,
	"Feb": time.February,
	"Mar": time.March,
	"Apr": time.April,
	"May": time.May,
	"Jun": time.June,
	"Jul": time.July,
	"Aug": time.August,
	"Sep": time.September,
	"Oct": time.October,
	"Nov": time.November,
	"Dec": time.December,
}

func get_tokendata(line, sep string) string {
	return line[strings.LastIndex(line, sep)+1 : len(line)]
}

func get_date(line string, reDate *regexp.Regexp) time.Time {
	year := time.Now().Year() //very evil for unittesters to do this
	date := reDate.FindStringSubmatch(line)

	day, _ := strconv.Atoi(date[2])
	hour, _ := strconv.Atoi(date[3])
	minute, _ := strconv.Atoi(date[4])
	second, _ := strconv.Atoi(date[5])

	retVal := time.Date(year, months[date[1]], day, hour, minute, second, 0, time.UTC)

	return retVal
}

func parse_line(line string) (bool, FWItem) {
	var item FWItem
	valid := false
	reIPT := regexp.MustCompile(`IPT:\s(\w+)`)
	reDate := regexp.MustCompile(`^(\w+)\s+(\d+)\s(\d\d):(\d\d):(\d\d)`)
	start := reIPT.FindStringSubmatch(line)
	if len(start) > 0 {
		valid = true //maybe bad asumptation but will do for now.
		item.action = start[1]
		tokens := strings.Split(line, " ")
		for index := range tokens {
			if strings.Contains(tokens[index], "IN=") {
				item.inIf = get_tokendata(tokens[index], "=")
			}
			if strings.Contains(tokens[index], "OUT=") {
				item.outIf = get_tokendata(tokens[index], "=")
			}
			if strings.Contains(tokens[index], "MAC=") {
				item.mac = get_tokendata(tokens[index], "=")
			}
			if strings.Contains(tokens[index], "SRC=") {
				item.src = get_tokendata(tokens[index], "=")
			}
			if strings.Contains(tokens[index], "DST=") {
				item.dst = get_tokendata(tokens[index], "=")
			}
			if strings.Contains(tokens[index], "PROTO=") {
				item.proto = get_tokendata(tokens[index], "=")
			}
			if strings.Contains(tokens[index], "SPT=") {
				item.spt = get_tokendata(tokens[index], "=")
			}
			if strings.Contains(tokens[index], "DPT=") {
				item.dpt = get_tokendata(tokens[index], "=")
			}
			item.date = get_date(line, reDate)
		}
	}
	return valid, item
}

func import_syslog(cfg IPTAConfig, filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	db := open_db(cfg)
	defer db.Close()
	//re := regexp.MustCompile(`IPT:\s(\w+)\sIN=(\w+)\sOUT=(\w*)\sMAC=([0-9a-f:]+)\sSRC=([0-9.]+)\sDST=([0-9.]+)\s.*PROTO=(\w+)\sSPT=([0-9]+)\sDPT=([0-9]+)`)

	//var items [1024]FWItem
	//index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		valid, item := parse_line(line)

		if valid {
			/*	items[index]=item
				++index
				//fmt.Println(item)
				if index > len(items){
					index = 0
					add_items(cfg, db, items)
				}
			*/
			add_item(cfg, db, item)
		}
	}

	file.Close()
}

func follow_syslog(cfg IPTAConfig, filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := open_db(cfg)
	defer db.Close()

	reader := bufio.NewReader(file)
	fileinfo, _ := file.Stat()

	for {
		reader.Peek(1)
		bytes_left := reader.Buffered()
		offset_eof, _ := file.Seek(0, os.SEEK_END)
		if bytes_left > 0 {
			line, err := reader.ReadString('\n')
			if err == io.EOF || err == nil {
				fmt.Printf("--> line = %s\t | bytes_left= %d offset_eof = %d size = %d\n", string(line[:len(line)-1]), bytes_left, offset_eof, fileinfo.Size())
			}
		} else if fileinfo.Size() > offset_eof {
			file.Close()
			file, _ = os.Open(filename)
			file.Seek(offset_eof, os.SEEK_SET)

			reader = bufio.NewReader(file)
		}
		checker, _ := os.Open(filename)
		fileinfo, _ = checker.Stat()
		checker.Close()

		if bytes_left < 10 {
			time.Sleep(500 * time.Millisecond)
		}
	}
	/*
		if valid {
			add_item(cfg, db, item)
		}
	*/

	file.Close()
}
