package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/parallel-computing-labs/internal/parsing"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	buckwheats := make(chan parsing.Buckwheat, 50)
	for i := 0; i < 5; i++ {
		go consume(buckwheats, i)
	}
	parsers := []parsing.Parser{
		parsing.NewAuchanParser(),
		parsing.MewAquamarketParser(),
		parsing.NewFozzyParser(),
	}
	for _, parser := range parsers {
		go produce(parser, buckwheats)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
}

func consume(buckwheats <-chan parsing.Buckwheat, i int) {
	for buckwheat := range buckwheats {
		log.WithFields(log.Fields{"consumer": i, "buffer": len(buckwheats)}).Info(buckwheat.URL)
		time.Sleep(time.Millisecond * 200)
	}
}

func produce(parser parsing.Parser, buckwheats chan<- parsing.Buckwheat) {
	parse(parser, buckwheats)
	for range time.Tick(time.Second * 30) {
		parse(parser, buckwheats)
	}
}

func parse(parser parsing.Parser, buckwheats chan<- parsing.Buckwheat) {
	products, err := parser.ParseBuckwheats()
	if err != nil {
		log.WithField("producer", parser.Name()).Error(err)
	} else {
		for _, product := range products {
			buckwheats <- product
		}
	}
}
