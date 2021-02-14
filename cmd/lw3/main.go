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
	buckwheats := make(chan parsing.Buckwheat, 300)
	go consume(buckwheats)
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

func consume(buckwheats <-chan parsing.Buckwheat) {
	for buckwheat := range buckwheats {
		log.Info(buckwheat)
		time.Sleep(time.Second)
	}
}

func produce(parser parsing.Parser, buckwheats chan<- parsing.Buckwheat) {
	parse(parser, buckwheats)
	for range time.Tick(time.Minute * 2) {
		parse(parser, buckwheats)
	}
}

func parse(parser parsing.Parser, buckwheats chan<- parsing.Buckwheat) {
	products, err := parser.ParseBuckwheats()
	if err != nil {
		log.WithField("parser", parser.Name()).Error(err)
	} else {
		for _, product := range products {
			buckwheats <- product
		}
	}
}
