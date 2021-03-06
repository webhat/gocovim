/*
		Copyright (C) 2019  Daniël W. Crompton

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
		along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var files map[string]map[int]*Status = make(map[string]map[int]*Status)

func verbose() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

func main() {
	verbose()
	var filename string
	if len(os.Args) < 2 {
		filename = "cover.out"
	} else {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cover, err := parseCoverLine(scanner.Text())
		if cover.filename == "mode" {
			continue
		}
		if err != nil {
			log.Println(err)
		}
		log.Println(cover)

		err = addToCoverageMap(cover)
		if err != nil {
			log.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	createCoverageFile()
}

type Status struct {
	status   string // hit, minn, ignored
	line     int
	filename string
}

func addToCoverageMap(cover Cover) error {
	filename := strings.Split(cover.filename, "/")
	cover.filename = filename[len(filename)-1]

	if files[cover.filename] == nil {
		files[cover.filename] = make(map[int]*Status)
	}
	for i := cover.lineFrom; i <= cover.lineTo; i++ {
		files[cover.filename][i] = &Status{}
		files[cover.filename][i].filename = filename[len(filename)-1]
		files[cover.filename][i].line = i
		if cover.numberOfStatements < 1 {
			files[cover.filename][i].status = "ignored"
		} else if cover.count > 0 {
			files[cover.filename][i].status = "hit"
		} else {
			files[cover.filename][i].status = "miss"
		}
		log.Println(files[cover.filename][i].status, "line:", i)
	}
	return nil
}

func createCoverageFile() {
	out, err := os.Create(".cadre/coverage.vim")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer out.Close()

	fmt.Fprintf(out, ":hi  HitSign     ctermfg=6      cterm=bold   gui=bold    guifg=Green\n")
	fmt.Fprintf(out, ":hi  MissSign    ctermfg=Red    cterm=bold   gui=bold    guifg=Red\n")
	fmt.Fprintf(out, ":hi  IgnoredSign ctermfg=6      cterm=bold   gui=bold    guifg=Grey\n")
	fmt.Fprintf(out, ":sign  define  hit      linehl=HitLine      texthl=HitSign      text=✔\n")
	fmt.Fprintf(out, ":sign  define  miss     linehl=MissLine     texthl=MissSign     text=✘\n")
	fmt.Fprintf(out, ":sign  define  ignored  linehl=IgnoredLine  texthl=IgnoredSign  text=◌\n")

	for f := range files {
		fmt.Fprintf(out, "if bufwinnr(\"%s\") > 0\n", f)
		for l := range files[f] {
			fmt.Fprintf(out, ":sign place 1 line=%d name=%s file=%s\n", files[f][l].line, files[f][l].status, files[f][l].filename)
		}
		fmt.Fprintf(out, "endif\n")
	}
}

func parseCoverLine(line string) (item Cover, err error) {
	item = Cover{}
	log.Println(line)
	line = strings.Replace(line, ":", " ", 1)
	// bitbucket.org/specialbrands/master-control-unit/controller/config.go:122.13,151.2 20 1
	cnt, err := fmt.Sscanf(line, "%s %d.%d,%d.%d %d %d", &item.filename, &item.lineFrom, &item.columnFrom, &item.lineTo, &item.columnTo, &item.numberOfStatements, &item.count)
	log.Printf(" Count: %d\n", cnt)
	log.Println(item)
	return
}

type Cover struct {
	filename                  string
	lineFrom, lineTo          int
	columnFrom, columnTo      int
	numberOfStatements, count int
}

/*
	:sign place 1 line=37 name=hit file=main.go
	:sign place 1 line=38 name=miss file=main.go
	:sign place 1 line=39 name=ignored file=main.go
*/
