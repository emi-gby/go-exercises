package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"sort"
)

//calculate the word frequency in a sentence
func WordFrequency(text string) map[string] int {
	frequencyMap := make(map[string]int)
	words := strings.FieldsSeq(text)

	for word := range words {
		frequencyMap[word] += 1
	}
	return frequencyMap
}

//map struct to sort by value logic
type mapStruct struct{
	Key string
	Value int
}

//logic for sorting maps by value
func SortMapByValue(initMap map[string]int) []mapStruct{

	var mapSlice []mapStruct
	for k, v:= range initMap {
		mapSlice = append(mapSlice, mapStruct{k, v})
	}

	sort.Slice(mapSlice, func(i, j int) bool{
		if mapSlice[i].Value == mapSlice[j].Value {
			//sort alphabetically if value is equal
			return mapSlice[i].Key < mapSlice[j].Key
		}
		return mapSlice[i].Value > mapSlice[j].Value
	})
	
	return mapSlice
}

//return the top 3 words more frenquent
func Top3Words(freqSlice []mapStruct) {

	if len(freqSlice) >= 3 {
		for i:=range 3 {
			fmt.Println(i+1,"° :" , freqSlice[i].Key, "(",freqSlice[i].Value, ")")
		}
	}else{
		for i:=range len(freqSlice) {
			fmt.Println(i+1,"° :" , freqSlice[i].Key, "(",freqSlice[i].Value, ")")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')

	frequencyMap := WordFrequency(strings.ToLower(sentence))

	orderdWords := SortMapByValue(frequencyMap)

	/*for k, v := range frequencyMap {
		fmt.Println(k, ": ", v)
	}*/
	
	Top3Words(orderdWords)
}