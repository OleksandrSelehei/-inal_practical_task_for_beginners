package comments

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Coments struct {
	PostId      int    `json:"postId"`
	ComentsId   int    `json:"id"`
	ComentsName string `json:"name"`
	UserEmail   string `json:"email"`
	ComentsBody string `json:"body"`
}

// запит до АПІ по ID посту для отримання всіх коментарів які налужать цьому посту
func RequestsComents(PostId int, wg *sync.WaitGroup, ch chan Coments) {
	defer wg.Done()
	urlApi := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", PostId)

	response, err := http.Get(urlApi)

	if err != nil {
		fmt.Printf("Erorr %s", err)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	var coments []Coments
	err = json.Unmarshal(data, &coments)
	if err != nil {
		fmt.Printf("Erorr %s", err)
		return
	}

	for _, comment := range coments {
		ch <- comment
	}
}
