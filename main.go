package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	jsonMap := make(map[string]interface{})

	rand.Seed(time.Now().UnixNano())
	for i := 1; ; i++ {
		// resp, err := http.Get("https://voteone.cqcb.com/vote/2021project/q1/2021redstory/controller.php?enews=detail&id=39")
		resp, err := http.Get("https://voteone.cqcb.com/vote/2021project/q1/2021redstory/controller.php?enews=detail&id=100")
		if err != nil {
			_ = fmt.Sprintf("err = %v\n", err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				continue
			}

			_ = json.Unmarshal(body, &jsonMap)
			if jsonMap["code"] != 200 {
				_ = fmt.Sprintf("resp = %s\n", string(body))
				_ = resp.Body.Close()
				break
			}

			_ = resp.Body.Close()
		}

		n := rand.Intn(10)
		time.Sleep(time.Duration(n) * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	for {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go worker(&wg)
		}

		wg.Wait()
	}
}
