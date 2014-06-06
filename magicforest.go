package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Forest struct {
	goats  int
	wolves int
	lions  int
}

func (f Forest) wolfDevoursGoat() *Forest {
	if f.goats > 0 && f.wolves > 0 {
		return &Forest{f.goats - 1, f.wolves - 1, f.lions + 1}
	}
	return nil
}

func (f Forest) lionDevoursGoat() *Forest {
	if f.goats > 0 && f.lions > 0 {
		return &Forest{f.goats - 1, f.wolves + 1, f.lions - 1}
	}
	return nil
}

func (f Forest) lionDevoursWolf() *Forest {
	if f.lions > 0 && f.wolves > 0 {
		return &Forest{f.goats + 1, f.wolves - 1, f.lions - 1}
	}
	return nil
}

func (f Forest) meal() []*Forest {
	a := []*Forest{f.wolfDevoursGoat(), f.lionDevoursGoat(), f.lionDevoursWolf()}
	return a
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 3 {
		goats, err := strconv.Atoi(args[0])
		if err != nil {
			usage()
		}
		wolves, err := strconv.Atoi(args[1])
		if err != nil {
			usage()
		}
		lions, err := strconv.Atoi(args[2])
		if err != nil {
			usage()
		}
		f := Forest{goats, wolves, lions}
		fmt.Println(f)

	} else {
		usage()
	}
}

func usage() {
	fmt.Println("USAGE: magicforest <goats> <wolves> <lions>")
	os.Exit(0)
}
