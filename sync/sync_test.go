package sync

import (
	"testing"
	"github.com/robfig/cron"
	
	conf "github.com/irisnet/irishub-sync/conf/server"
	
	"sync"
	"github.com/irisnet/irishub-sync/module/logger"
	"time"
)

func TestStart(t *testing.T) {
	var (
		limitChan    chan int
		unBufferChan chan int
	)
	limitChan = make(chan int, 3)
	unBufferChan = make(chan int)
	goroutineNum := 5
	activeGoroutineNum := goroutineNum
	for i := 1; i <= goroutineNum; i++ {
		limitChan <- i
		go func(goroutineNum int, ch chan int) {
			logger.Info.Println("release limitChan")
			<- limitChan
			defer func() {
				logger.Info.Printf("%v goroutine send data to channel\n",
					goroutineNum)
				ch <- goroutineNum
			}()
			
		}(i, nil)
	}
	
	for
	{
		select {
		case <-unBufferChan:
			activeGoroutineNum = activeGoroutineNum - 1
			logger.Info.Printf("active goroutine num is %v", activeGoroutineNum)
			if activeGoroutineNum == 0 {
				logger.Info.Println("All goroutine complete")
				break
			}
		}
	}
	
	
}

func Test_startCron(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(2)
	
	spec := conf.SyncCron
	c := cron.New()
	c.AddFunc(spec, func() {
		mutex.Lock()
		var sleepSecond time.Duration
		sleepSecond = 3
		time.Sleep(time.Second * sleepSecond)
		logger.Info.Printf("awake up after %v second\n", sleepSecond)
		mutex.Unlock()
	})
	go c.Start()
	
	wg.Wait()
}


