package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	data, err := ioutil.ReadFile("logs/source.log")
	if err != nil {
		log.Fatalf("main: failed to read the log, %v", err)
	}
	go produce(data)
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-signals
	log.Infof("main: first signal received, %s", s)
	s = <-signals
	log.Infof("main: second signal received, %s", s)
}

func produce(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	for {
		rand.Shuffle(
			len(lines),
			func(i, j int) {
				lines[i], lines[j] = lines[j], lines[i]
			},
		)
		newLines := make([][]byte, rand.Intn(10)+1)
		for i := range newLines {
			newLines[i] = actualize(lines[i])
		}
		if err := write(append(bytes.Join(newLines, []byte{'\n'}), '\n')); err != nil {
			log.Errorf("main: failed to write the log, %v", err)
		}
		time.Sleep(time.Duration(rand.Intn(1201)+100) * time.Millisecond)
	}
}

func actualize(line []byte) []byte {
	leftIndex := bytes.Index(line, []byte{'['})
	rightIndex := bytes.Index(line, []byte{']'})
	if leftIndex == -1 || rightIndex == -1 || rightIndex <= leftIndex {
		return line
	}
	return append(
		append(line[:leftIndex+1], time.Now().Format("02/Jan/2006:15:04:05 -0700")...),
		line[rightIndex:]...,
	)
}

func write(bytes []byte) error {
	file, err := os.OpenFile("logs/target.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	if _, err = file.Write(bytes); err != nil {
		_ = file.Close()
		return err
	}
	return file.Close()
}
