package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	PostId int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// запит до АПІ для отримання постів по ID користувача
func RequestsPost(UserID int) ([]Post, error) {
	var urlApi string = fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", UserID)
	response, err := http.Get(urlApi)

	if err != nil {
		fmt.Printf("Erorr %s", err)
		return nil, fmt.Errorf("Error connecting to the database: %v", err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %v", err)
	}

	var posts []Post
	err = json.Unmarshal(data, &posts)

	if err != nil {
		fmt.Printf("Erorr %s", err)
		return nil, fmt.Errorf("Error connecting to the database: %v", err)
	}
	return posts, nil
}
