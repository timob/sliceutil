sliceutil
==========
Go packge to provide slice utility functions

Example
-------
```Go
package main 

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/bradfitz/slice"
	"github.com/timob/sliceutil"
	"os"
)

func main() {
	key := flag.Int("k", 1, "key to sort lines on")
	flag.Parse()
	
	var keyedLines []*struct{key, line string}
	lineScanner := bufio.NewScanner(os.Stdin)	
	for lineScanner.Scan() {
		i := sliceutil.Append(&keyedLines)
		keyedLines[i].line = lineScanner.Text()
		
		wordScanner := bufio.NewScanner(bytes.NewBuffer(lineScanner.Bytes()))
		wordScanner.Split(bufio.ScanWords)
		for n := 1; wordScanner.Scan(); n++ {
			if n == *key {
				keyedLines[i].key = wordScanner.Text()
				break
			}
		}		
	}

	slice.Sort(keyedLines, func(i, j int) bool { return keyedLines[i].key < keyedLines[j].key})	

	for _, v := range keyedLines {
		fmt.Println(v.line)
	}
}```

