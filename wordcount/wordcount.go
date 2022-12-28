package wordcount

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func Countwords(filename string) (freqmap map[string]int) {

	bytes, err := os.ReadFile("./wordcount/files/" + filename)
	if err != nil {
		log.Println(err)
	}

	content := string(bytes)
	freqmap = make(map[string]int)

	// Split the content string into words
	words := strings.Split(content, " ")

	for _, word := range words {
		if _, ok := freqmap[word]; !ok {
			freqmap[word] = 1
		} else {
			freqmap[word] += 1
		}
	}
	return freqmap
}

func Reducewords(freqmaps, combined map[string]int) map[string]int {
	fmt.Println("check1")
	for key, val := range freqmaps {

		if _, ok := combined[key]; !ok {
			combined[key] = val
		} else {
			combined[key] += val
		}

	}

	return combined
}

//func MapReducefunc(filenames []string) map[string]int {
//
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//
//	mapchan := make(chan map[string]int, 100)
//	//reducechan := make(chan map[string]int, 100)
//
//	wg.Add(len(filenames))
//
//	for _, filename := range filenames {
//		go func(name string) {
//			wordcount := Countwords(name)
//			mapchan <- wordcount
//			defer wg.Done()
//
//		}(filename)
//
//		go func() {
//			wg.Wait()
//			close(mapchan)
//		}()
//	}
//	combined := make(map[string]int)
//	for wordcount := range mapchan {
//		wg.Add(1)
//		go func(counts map[string]int) {
//			mu.Lock()
//			combined = Reducewords(counts, combined)
//			mu.Unlock()
//			wg.Done()
//		}(wordcount)
//
//	}
//	wg.Wait()
//
//	return combined
//
//}

//func

func MapReducefunc(filenames []string) map[string]int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	mapchan := make(chan map[string]int, 100)

	wg.Add(len(filenames))

	for _, filename := range filenames {
		go func(name string) {
			wordcount := Countwords(name)
			mapchan <- wordcount
			wg.Done()
		}(filename)
	}

	go func() {
		wg.Wait()
		close(mapchan)
	}()

	combined := make(map[string]int)
	for wordcount := range mapchan {
		mu.Lock()
		combined = Reducewords(wordcount, combined)
		mu.Unlock()
	}

	return combined
}
