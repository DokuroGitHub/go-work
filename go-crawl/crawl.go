package crawl

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/TwiN/go-color"
)

func sendValue(c chan<- string, wg *sync.WaitGroup, nURLs int) {
	defer wg.Done()
	for i := 1; i <= nURLs; i++ {
		item := color.InBlack("[") + color.InBlue(fmt.Sprintf("URL%03d", +i)) + color.InBlack("]")
		c <- item
		fmt.Println(color.Autof(item) + ": enqueued")
	}
	close(c)
	fmt.Println(color.InGreen("Channel has closed"))
}

func receiveValue(c <-chan string, wg *sync.WaitGroup, nWorkers, sPerURL int) {
	defer wg.Done()
	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		worker := color.InBlack("[") + color.InPurple(fmt.Sprintf("Worker%02d", +i)) + color.InBlack("]")
		go func() {
			defer wg.Done()
			for item := range c {
				fmt.Println(color.Autof(item) + ":" + color.Autof(worker) + color.InCyan(": running"))
				time.Sleep(time.Second * time.Duration(sPerURL))
				fmt.Println(color.Autof(item) + color.InGreen(": finished"))
			}
			fmt.Println(color.Autof(worker) + color.InGreen(": finished"))
		}()
	}
}

func Crawl(numberOfURLs, numberOfWorkers, secondsPerURL, channelCap int) error {
	if numberOfURLs < 0 {
		return errors.New("numberOfURLs can not be positive (>=0)")
	}
	if numberOfWorkers <= 0 {
		return errors.New("numberOfWorkers must be positive (>0)")
	}
	if secondsPerURL <= 0 {
		return errors.New("secondsPerURL must be positive (>0)")
	}
	if channelCap <= 0 {
		return errors.New("channelCap must be positive (>0)")
	}

	// WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(2) // .Done() in sendValue() + receiveValue()

	// channel
	c := make(chan string, channelCap)

	// sender
	go sendValue(c, wg, numberOfURLs)

	// receiver
	go receiveValue(c, wg, numberOfWorkers, secondsPerURL)

	// wait
	wg.Wait()
	fmt.Print(color.InGreen("Crawl ended\n"))

	return nil
}
