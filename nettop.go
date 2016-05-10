package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"bufio"
	"os"
	"strconv"
	"strings"
)

var T = flag.Float64("t", 2, "update time(s)")
var C = flag.Uint("c", 0, "count (0 == unlimit)")
var Inter = flag.String("i", "*", "interface")

var verbosity = flag.Int("v", 2, "verbosity")

type NetStat struct {
	Dev  []string
	Stat map[string]*DevStat
}

type DevStat struct {
	Name string
	Rx   uint64
	Tx   uint64
}

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	var ret []string

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}
	return ret, nil
}

func getInfo() (ret NetStat) {

	lines, _ := ReadLines("/proc/net/dev")

	ret.Dev = make([]string, 0)
	ret.Stat = make(map[string]*DevStat)

	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.Fields(strings.TrimSpace(fields[1]))

		// Vlogln(5, key, value)

		if *Inter != "*" && *Inter != key {
			continue
		}

		c := new(DevStat)
		// c := DevStat{}
		c.Name = key
		r, err := strconv.ParseInt(value[0], 10, 64)
		if err != nil {
			Vlogln(4, key, "Rx", value[0], err)
			break
		}
		c.Rx = uint64(r)

		t, err := strconv.ParseInt(value[8], 10, 64)
		if err != nil {
			Vlogln(4, key, "Tx", value[8], err)
			break
		}
		c.Tx = uint64(t)

		ret.Dev = append(ret.Dev, key)
		ret.Stat[key] = c
	}

	return
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	flag.Parse()

	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)

	var stat0 NetStat
	var stat1 NetStat
	var delta NetStat
	delta.Dev = make([]string, 0)
	delta.Stat = make(map[string]*DevStat)

	i := *C
	if i > 0 {
		i += 1
	}

	if *T < 0.01 {
		*T = 0.01
	}

	start := time.Now()
	elapsed := time.Since(start)
	if *Inter == "*" {
		fmt.Printf("\033[0;0H                                                        \r")
	}
	fmt.Printf("iface\tRx\tTx\n")
	for {

		elapsed = time.Since(start)
		stat1 = getInfo()
		// Vlogln(5, stat0)
		for _, value := range stat1.Dev {
			t0, ok := stat0.Stat[value]
			// fmt.Println("k:", key, " v:", value, ok)
			if ok {
				dev, ok := delta.Stat[value]
				if !ok {
					delta.Stat[value] = new(DevStat)
					dev = delta.Stat[value]
					delta.Dev = append(delta.Dev, value)
				}
				t1 := stat1.Stat[value]
				dev.Rx = t1.Rx - t0.Rx
				dev.Tx = t1.Tx - t0.Tx
			}
		}
		stat0 = stat1

		multi := len(delta.Dev)
		if multi > 1 {
			for i := 0; i < multi; i++ {
				fmt.Printf("\033[%d;0H                                                        \r", i)
			}
			fmt.Printf("\033[0;0Hiface\tRx\tTx\n")
			fmt.Printf("\033[2;0H")
		}
		for _, iface := range delta.Dev {
			stat := delta.Stat[iface]
			if multi > 1 {
				fmt.Printf("%v\t%v\t%v\n", iface, Vsize(stat.Rx, *T), Vsize(stat.Tx, *T))
			} else {
				fmt.Printf("\r%v\t%v\t%v", iface, Vsize(stat.Rx, *T), Vsize(stat.Tx, *T))
			}
		}
		// elapsed := time.Since(start)
		Vlogf(5, "[delta] %s", elapsed)
		start = time.Now()

		i -= 1
		if i == 0 {
			break
		}

		time.Sleep(time.Duration(*T*1000) * time.Millisecond)

	}
}

func Vsize(bytes uint64, delta float64) (ret string) {
	var tmp float64 = float64(bytes) / delta
	var s string = " "

	bytes = uint64(tmp)

	switch {
	case bytes < uint64(2<<9):

	case bytes < uint64(2<<19):
		tmp = tmp / float64(2<<9)
		s = "K"

	case bytes < uint64(2<<29):
		tmp = tmp / float64(2<<19)
		s = "M"

	case bytes < uint64(2<<39):
		tmp = tmp / float64(2<<29)
		s = "G"

	case bytes < uint64(2<<49):
		tmp = tmp / float64(2<<39)
		s = "T"

	}
	ret = fmt.Sprintf("%06.2f %sB/s", tmp, s)
	return
}

func Vlogf(level int, format string, v ...interface{}) {
	if level <= *verbosity {
		log.Printf(format, v...)
	}
}
func Vlog(level int, v ...interface{}) {
	if level <= *verbosity {
		log.Print(v...)
	}
}
func Vlogln(level int, v ...interface{}) {
	if level <= *verbosity {
		log.Println(v...)
	}
}
