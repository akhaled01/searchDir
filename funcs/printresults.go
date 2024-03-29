package funcs

import (
	"fmt"
	"os"
)

func PrintNormal(entry os.FileInfo) {
	if entry.IsDir() {
		fmt.Println(BlueANSI + BoldANSI + entry.Name() + ResetANSI)
	} else {
		fmt.Println(entry.Name() + " ")
	}
}

func PrintRes(mainfs string) {
	width, _ := getTerminalWidth()
	if width == 0 {
		width = 59
	}
	grouplen, susInfolen := 0, 0
	stat, err := os.Stat(mainfs)
	if err != nil {
		fmt.Println(RedANSI+BoldANSI+"[printresults.go] getting stat,", err)
		os.Exit(1)
	}

	if !stat.IsDir() {
		if !LongFormat {
			fmt.Println(mainfs)
		} else {
			if err != nil {
				fmt.Println(RedANSI+BoldANSI+"[printresults.go] error printing res,", err)
			}
			LFD(mainfs, grouplen, susInfolen)
		}
		fmt.Println()
		return
	}
	entries, dirs := SearchDir(mainfs)
	if NotFlags >= 2 && FileNum >= 1 {
		fmt.Println(mainfs)
	}
	if LongFormat {
		total := calculateTotal(entries, mainfs)
		fmt.Println("total", total)
	}
	Sort(entries)
	if Timesort && ReverseOrder {
		sortByReverseTime(mainfs, entries)
	} else if Timesort {
		sortByTime(mainfs, entries)
	} else if ReverseOrder {
		reverseSortAlphabet(entries)
	}

	// Print the entries in a list (normal one + if the files can be print in one line)
	for _, entry := range entries {
		if !LongFormat {
			PrintNormal(entry)
		} else {
			_, err := os.Stat(mainfs + "/" + entry.Name())
			if err != nil {
				fmt.Println(RedANSI+BoldANSI+"[printresults.go] error printing res,", err)
			}
			LFD(mainfs+"/"+entry.Name(), grouplen, susInfolen)
		}
	}
	if RecursiveSearch {
		for _, subFS := range dirs {
			fmt.Println(mainfs + "/" + subFS.Name())
			PrintRes(mainfs + "/" + subFS.Name())
		}
	}
	fmt.Println()
}
