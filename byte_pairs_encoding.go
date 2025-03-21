package main

import (
	"errors"
	"fmt"
)

var FinalText = string("")

type LookupTable map[string]string

type StringIntPair struct {
	Text    string
	Count   int
	Indexes []int
}

func the_most_frequent_pair(text string) (StringIntPair, error) {
	if len(text) < 2 {
		return StringIntPair{}, errors.New("No more pairs")
	}

	var mostRepeated = StringIntPair{}
	var localLoockupCount = make(map[string]int)
	var localLoockupIdx = make(map[string][]int)

	for index := range len(text) - 1 {
		pairText := text[index : index+2]
		localLoockupCount[pairText] += 1
		if len(localLoockupIdx[pairText]) == 0 || localLoockupIdx[pairText][len(localLoockupIdx[pairText])-1] != index-1 {
			localLoockupIdx[pairText] = append(localLoockupIdx[pairText], index)
		}
	}

	for key, count := range localLoockupCount {
		if count > mostRepeated.Count {
			mostRepeated.Text = key
			mostRepeated.Count = count
			mostRepeated.Indexes = localLoockupIdx[key]
		}
	}

	if mostRepeated.Count <= 1 {
		return StringIntPair{}, errors.New("No more pairs")
	}

	return mostRepeated, nil
}

func main() {

	pair, error := the_most_frequent_pair("aaabdaaabac")

	fmt.Print(pair)
	if error != nil {

	}
}
