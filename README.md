# cuttergo

A fast **Chinese** word cut tool using golang based AC machine algorithm.

### Installation

go get github.com/Magic-Fy/cuttergo

### Example
```golang
package main

import (
	"fmt"
	"github.com/Magic-Fy/cuttergo"
	"strings"
)

func main() {
	wordCutter := cuttergo.Cutter{}
	segs, err := wordCutter.Cut("北京清华大学研究生院")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("'%s'\n", strings.Join(segs, "', '"))
}

//stdout
//
```
### Reference
https://github.com/liwenju0/cutword

