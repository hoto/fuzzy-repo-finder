package traverse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	root := "/home/andrzej.rehmann/projects"
	//root := "/home/andrzej.rehmann/go/src"
	matchingPaths := FindDir(root, ".git")
	for i, e := range matchingPaths {
		fmt.Printf("%d %s\n", i, e)
	}
}

func FindDir(root string, needle string) []string {
	matchingPaths := make([]string, 0)
	scan(root, needle, &matchingPaths)
	return matchingPaths
}

func scan(dir string, needle string, matchingPaths *[]string) {
	haystack, err := ioutil.ReadDir(dir)
	check(err)
	if containsNeedle(haystack, needle) {
		*matchingPaths = append(*matchingPaths, dir)
		return
	}
	for _, file := range haystack {
		if file.IsDir() {
			dir := fmt.Sprintf("%s/%s", dir, file.Name())
			scan(dir, needle, matchingPaths)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func containsNeedle(files []os.FileInfo, needle string) bool {
	for _, file := range files {
		if file.Name() == needle {
			return true
		}
	}
	return false
}
