package cuttergo

import (
	goahocorasick "github.com/anknown/ahocorasick"
	"regexp"
)

type Cutter struct {
	m   *goahocorasick.Machine
	dic map[string]float64
}

// Init initializes the read, filename is the path to the dictionary file
// The dictionary file should be in the format of: word\tweight\tfeature
func (c *Cutter) Init(filename string) error {
	dic, arr, err := ReadRunes(filename)
	if err != nil {
		return err
	}

	c.m = new(goahocorasick.Machine)
	if err := c.m.Build(arr); err != nil {
		return err
	}
	c.dic = dic

	return nil
}

// Tokenize tokenizes the content
func (c *Cutter) tokenize(content string) []string {
	if c.m == nil {
		err := c.Init("dict.txt")
		if err != nil {
			return []string{}
		}
	}
	// inf is the least score of the dictionary
	inf := -1e10
	flo := make([]float64, len([]rune(content))+1)
	flo[0] = 0
	for i := 1; i <= len([]rune(content)); i++ {
		flo[i] = inf
	}
	routes := make([]int, len([]rune(content))+1)
	for i := 0; i <= len([]rune(content)); i++ {
		routes[i] = i
	}
	bound := -1
	tokens := make([]string, 0)

	// MultiPatternSearch returns the positions of the words in the content
	// process every result of the MultiPatternSearch
	// compare the score and get the segmentation
	for _, item := range c.m.MultiPatternSearch([]rune(content), false) {
		bound = item.Pos + len(item.Word) - 1
		scores := bound - len(item.Word) + 1
		bound += 1
		if flo[scores] == inf {
			last := scores
			for ; last > 0; last-- {
				if flo[last] != inf {
					break
				}
			}

			flo[scores] = flo[last] - 10.0
			routes[scores] = last
		}
		s := flo[scores] + c.dic[string(item.Word)]
		if s > flo[bound] {
			flo[bound] = s
			routes[bound] = scores
		}
	}

	if bound < 0 {
		tokens = append(tokens, content)
		return tokens
	}

	if bound < len([]rune(content)) {
		tokens = append(tokens, string([]rune(content)[bound:]))
		content = content[:bound]
	}

	for len([]rune(content)) > 0 {
		b := routes[bound]
		tokens = append(tokens, string([]rune(content)[b:bound]))
		content = string([]rune(content)[:b])
		bound = b
	}
	tmp := make([]string, 0, len(tokens))
	for i := len(tokens) - 1; i >= 0; i-- {
		tmp = append(tmp, tokens[i])
	}
	return tmp
}

// Cut cuts the content into words
func (c *Cutter) Cut(content string) ([]string, error) {
	results := make([]string, 0)
	regText := regexp.MustCompile("[\u4e00-\u9fa5]+")

	regSkip := regexp.MustCompile("([a-zA-Z0-9]+(?:\\.\\d+)?%?)")

	blo := proc(content, regText)

	for _, item := range blo {
		if item == "" {
			continue
		}
		if regText.MatchString(item) {
			results = append(
				results,
				c.tokenize(item)...,
			)
		} else {
			tmp := proc(item, regSkip)
			for _, it := range tmp {
				if it != "" {
					results = append(results, it)
				}
			}
		}
	}
	return results, nil
}

// proc is a helper function to process the content
func proc(content string, text *regexp.Regexp) []string {
	nonChi := text.Split(content, -1)
	segChi := text.FindAllStringSubmatchIndex(content, -1)
	words := make([]string, 0)
	for i, item := range segChi {
		words = append(words, nonChi[i])
		words = append(words, content[item[0]:item[1]])
		if i == len(segChi)-1 {
			words = append(words, nonChi[i+1])
		}
	}
	if segChi == nil {
		words = append(words, content)
	}
	return words
}
