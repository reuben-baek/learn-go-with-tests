package logs

import (
	"log"
	"os"
	"testing"
	"time"
)

const logFileName = "test.log"

var logFile = func() *os.File {
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", logFileName, err)
	}
	return file
}()

var logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)

var worker = 100
var writeCountPerWorker = 1000

func TestLog(t *testing.T) {
	defer logFile.Close()

	logger.Printf("Hello logger ~~")
	ch := make(chan bool)

	for i := 0; i < worker; i++ {
		go writeLogs(ch, i)
	}

	for i := 0; i < worker; i++ {
		<-ch
	}
	logger.Printf("Done ~~")
}

func writeLogs(ch chan bool, workerId int) {
	for i := 0; i < writeCountPerWorker; i++ {
		logger.Printf("worker[%d] - iteration %d", workerId, i)
		time.Sleep(10 * time.Millisecond)
	}
	ch <- true
}
