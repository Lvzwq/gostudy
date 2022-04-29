package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished\n", id)
}

func main() {
	var s sync.WaitGroup

	for i := 1; i < 5; i++ {
		s.Add(1)
		j := i
		go func() {
			defer s.Done()
			worker(j)
		}()
	}
	s.Wait()

	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				b := make([]byte, 65536)

				resp.Body.Read(b)
				fmt.Println(url, string(b))
				defer resp.Body.Close()
			}

			return err

		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}

}
