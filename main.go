package main

import (
	"bufio"
	"bytes"
	"fmt"
	"kd.kz/hashmap"
	"log"
	"os"
	"sort"
)

var numberOfFrequentWords = 20

func main() {
	hashMap := hashmap.NewHashMap()

	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Bytes()
		word = bytes.ToLower(word)
		word = bytes.TrimSpace(word)
		word = bytes.Trim(word, `,.!@#$%^&*()_+=-:;'"0123456789`)

		if len(word) == 0 {
			continue
		}

		if !hashMap.Contains(word) {
			hashMap.Set(word, 1)
		} else {
			val := hashMap.Get(word)
			val++
			hashMap.Set(word, val)
		}
	}

	var intArr []int
	for k := range hashMap.Table {
		if hashMap.Table[k] != nil {
			intArr = append(intArr, hashMap.Table[k].Value)
		}
	}

	sort.Ints(intArr)

	if numberOfFrequentWords > len(intArr) {
		numberOfFrequentWords = len(intArr) - 1
	}

	arr := intArr[len(intArr)-numberOfFrequentWords:]
	checkMap := hashmap.NewHashMap()

	for i := len(arr) - 1; i >= 0; i-- {
		for k := range hashMap.Table {
			if hashMap.Table[k] != nil {
				if hashMap.Table[k].Value == arr[i] {
					if checkMap.Contains(hashMap.Table[k].Key) {
						continue
					}
					fmt.Printf(" %v %v\n", arr[i], string(hashMap.Table[k].Key))
					checkMap.Set(hashMap.Table[k].Key, 1)
				}
			}
		}
	}
}
