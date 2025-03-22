package main

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
)

var FinalText = string("")

type LookupTable map[byte]string

type LookupTableLevels []LookupTable

type StringIntPair struct {
	Text    string
	Count   int
	Indexes []int
}

type KeyTable struct {
	Values string
	Pos    int
}

var Table = KeyTable{"ABCDEFGHIJKLMNOPQRSTUVXZ", 0}

func main() {
	text := "aaabdaaabac"
	FinalText = text
	for {
		var err error
		if FinalText, err = Proccess(FinalText); err != nil {
			break
		}
	}
	fmt.Println(FinalText)
}

func Proccess(text string) (string, error) {
	lookupTableLevels := LookupTableLevels{}
	lookupTableLevels = append(lookupTableLevels, LookupTable{})
	pos := 0
	pair, error := MostFrequentPair(text)
	if error != nil {
		slog.Error(error.Error())
		return text, error
	}
	key := GetNextKeyTolookupTable(pair.Text)
	textTrasnformed := TansformText(pair, string(key), text)

	if key == 0 {
		lookupTableLevels = append(lookupTableLevels, LookupTable{})
		pos = 0
	}

	lookupTableLevels[pos][key] = pair.Text

	pos += 1

	return textTrasnformed, nil
}

func GetNextKeyTolookupTable(pair string) byte {
	if Table.Pos == len(Table.Values) {
		return 0
	}
	next := Table.Values[Table.Pos]
	Table.Pos += 1
	return next
}

func MostFrequentPair(text string) (StringIntPair, error) {
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

func TansformText(pair StringIntPair, key, text string) string {
	var buffer bytes.Buffer
	start := 0
	for _, idx := range pair.Indexes {
		buffer.WriteString(text[start:idx] + key)
		start = idx + 2
	}
	buffer.WriteString(text[start:])
	return buffer.String()
}
