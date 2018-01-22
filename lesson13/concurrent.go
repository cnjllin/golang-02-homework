package main

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"
)

func worker(ch chan string) {
	for url := range ch {
		_ = url
	}
}

func main() {
	urls := []string{} // 1000个url
	// 以并发为10来抓取url
	wg := new(sync.WaitGroup)
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ch)
		}()
	}
	for _, url := range urls {
		ch <- url
	}
	close(ch)
	wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	dialer := new(net.Dialer)
	conn, err := dialer.DialContext(ctx, "tcp", "")
	deadline, _ := ctx.Deadline()
	conn.SetReadDeadline(deadline)

	t := NewTaskRunner(10)
forloop:
	for _, url := range urls {
		url := url
		select {
		case <-ctx.Done():
			break forloop
		default:
		}
		t.Put(func() {
			req, _ := http.NewRequest("GET", url, nil)
			req = req.WithContext(ctx)
			http.DefaultClient.Do(req)
		})
	}
	t.Wait()
}

func main1() {
	urls := []string{} // 1000个url
	token := make(chan bool, 10)
	for _, url := range urls {
		token <- true
		go func(url string) {
			_ = url
			<-token
		}(url)
	}
}

type TaskRunner struct {
	token chan bool
	wg    sync.WaitGroup
}

func NewTaskRunner(concurrent int) *TaskRunner {
	return &TaskRunner{
		token: make(chan bool, concurrent),
	}
}

func (t *TaskRunner) Put(task func()) {
	t.token <- true
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		task()
		<-t.token
	}()
}

func (t *TaskRunner) Wait() {
	t.wg.Wait()
}
