package cuttergo

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadRunes(filename string) (map[string]float64, [][]rune, error) {
	dict := make(map[string]float64)
	total := 0.0
	arrays := make([][]rune, 0)

	f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	if err != nil {
		return dict, nil, err
	}

	r := bufio.NewReader(f)
	for {
		l, _, err := r.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		segs := strings.Split(string(l), "\t")
		if len(segs) < 2 {
			continue
		}
		val, err := strconv.ParseFloat(segs[1], 64)
		if err != nil {
			continue
		}
		dict[segs[0]] = val
		total += val
		arrays = append(arrays, []rune(segs[0]))
	}

	for k, v := range dict {
		dict[k] = math.Log(v) - math.Log(total)
	}

	return dict, arrays, nil
}
