package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func AnalyzeText(text string) {
    str := strings.ToLower(text)
    words := strings.FieldsFunc(str, func(r rune) bool {
        if unicode.IsSpace(r) {
            return true
        }
        if unicode.IsPunct(r) && r != '-' && r != '\'' {
            return true
        }
        return false
    })

    numberOfWords := len(words)                 //count number of words

    wordMap := make(map[string]int)             //make map of words
    for _, word := range words {
        wordMap[word]++
    }

    cntUnique := len(wordMap)                              //count unique words

    topMostFreq := getTopWords(wordMap, 5)                 //using func to form slice of top words rated from 5 to 1 

    fmt.Printf("Количество слов: %d\n", numberOfWords)
    fmt.Printf("Количество уникальных слов: %d\n", cntUnique)
    fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", topMostFreq[0], wordMap[topMostFreq[0]])
    fmt.Printf("Топ-5 самых часто встречающихся слов:\n")
    for _, word := range topMostFreq {
        fmt.Printf("\"%s\": %d раз\n", word, wordMap[word])
    }
}

func getTopWords(wordMap map[string]int, n int) []string {

    valueTop := make([]int, 0)                      // saving values of map in slice []int
    for _, val := range wordMap {
        valueTop = append(valueTop, val)
    }

    sort.Slice(valueTop, func(i, j int) bool {      //sorting this slice []int of map values from 5 to 1 
        return valueTop[i] > valueTop[j]
    })

    wordsTop := make([]string, 0, n)               //making slice of words corresponding to order of slice []int   
    for i:= 0; i < n; i++ {                         
        for key, val := range wordMap {
            if val == valueTop[i] {
                wordsTop = append(wordsTop, key)
                continue
            }
        }
    }

    return wordsTop 
}

