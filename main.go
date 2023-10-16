package main

import (
	"fmt"
	workdb "practiceDB/db"
	comments "practiceDB/requests/coment"
	posts "practiceDB/requests/post"
	"sync"
)

func main() {
	// дані для підключення до БД та підключення
	meta := workdb.DataConfig{UserName: "root", Password: "Anim_2016_Go", DbName: "blog"}
	db, err := meta.Conection()

	if err != nil {
		fmt.Printf("Erorr %s", err)
		return
	}

	// запт до АПІ(отримання постів юзера ID n)
	posts, err := posts.RequestsPost(7)
	if err != nil {
		fmt.Printf("Erorr %s", err)
		return
	}
	// створення каналу
	ch := make(chan comments.Coments)
	//створеня групи горутин
	var wg sync.WaitGroup
	// функция яка чикає на завершення всих потоків та закріває канал
	go func() {
		wg.Wait()
		close(ch)
	}()
	//запис всіх отриманих постів в БД та запит до АПІ по ID посту(для отримання коментарів)
	for _, i := range posts {
		wg.Add(1)
		go workdb.AddPost(i, &wg, db)
		go comments.RequestsComents(i.PostId, &wg, ch)
	}
	//запис до БД всіх отриманих коментарів
	for item := range ch {
		wg.Add(1)
		go workdb.AddComment(item, &wg, db)
	}
	db.Close()
}
