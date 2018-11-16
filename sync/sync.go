package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getSite(url string, list *[]string, index int) {
	defer wg.Done()
	response, _ := http.Get(url)
	if response.StatusCode == http.StatusOK {
		*list = append(*list, url)
	}
}

func handlerSites(arr []string) []string {
	wg.Add(len(arr))
	var list = make([]string, 0)
	for index, url := range arr {
		go getSite(url, &list, index)
	}
	wg.Wait()
	return list
}

func main() {
	// var chanel = make(chan []string)
	var list = []string{
		"https://github.com/werickvieira",
		"https://github.com/ayrtonteshima",
		"https://github.com/wallacebenevides",
		"https://github.com/paulirish",
		"https://github.com/addyosmani",
		"https://github.com/yyx990803",
	}

	var newList = handlerSites(list)
	fmt.Println("NEWLIST", newList)

}
