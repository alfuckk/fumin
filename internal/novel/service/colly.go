package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alfuckk/fumin/pkg/novel"
)

func fetchNovel() *[]novel.Novel {
	file, err := os.Open("config/novel/sy.json")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var data []novel.Novel
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &data
}

// func fetchNovelCategory(book *biquge.Biquge) (result biquge.Categories) {
// 	var categories []biquge.Category
// 	err := json.Unmarshal([]byte(book.ExploreURL), &categories)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON:", err)
// 		return
// 	}

// 	for k, v := range categories {
// 		v.URL = book.BookSourceURL[:len(book.BookSourceURL)-1] + v.URL
// 		categories[k] = v
// 	}

// 	return categories
// }
