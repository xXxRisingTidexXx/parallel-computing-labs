package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/parallel-computing-labs/internal/pp"
	"io"
	"io/ioutil"
	"net"
	"time"
)

const (
	trials   = 1000
	protocol = "tcp"
	address  = "127.0.0.1:8585"
)

func main() {
	data, err := ioutil.ReadFile("source.jpg")
	if err != nil {
		log.Fatal(err)
	}
	var channelLatency int64
	for i := 0; i < trials; i++ {
		payloads := make(chan pp.Payload, 1)
		go writeToChannel(data, payloads)
		payload := <-payloads
		channelLatency += time.Now().Sub(payload.Time).Nanoseconds()
	}
	var pipeLatency int64
	for i := 0; i < trials; i++ {
		reader, writer := io.Pipe()
		starts := make(chan time.Time, 1)
		go writeToPipe(data, writer, starts)
		if _, err := ioutil.ReadAll(reader); err != nil {
			log.Fatal(err)
		}
		pipeLatency += time.Now().Sub(<-starts).Nanoseconds()
		if err := reader.Close(); err != nil {
			log.Fatal(err)
		}
	}
	var socketLatency int64
	listener, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	quits, starts := make(chan struct{}, 1), make(chan time.Time, 1)
	go writeToSocket(data, listener, quits, starts)
	for i := 0; i < 1; i++ {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			log.Fatal(err)
		}
		socketLatency += time.Now().Sub(<-starts).Nanoseconds()
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}
	quits <- struct{}{}
	log.Infof("Channel: %d ns", channelLatency/trials)
	log.Infof("Pipe: %d ns", pipeLatency/trials)
	log.Infof("Socket: %d ns", socketLatency/trials)
}

func writeToChannel(data []byte, payloads chan<- pp.Payload) {
	payloads <- pp.Payload{Bytes: data, Time: time.Now()}
}

func writeToPipe(
	data []byte,
	writer io.WriteCloser,
	starts chan<- time.Time,
) {
	start := time.Now()
	if _, err := writer.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}
	starts <- start
}

func writeToSocket(
	data []byte,
	listener net.Listener,
	quits <-chan struct{},
	starts chan<- time.Time,
) {
	for ok := true; ok; {
		select {
		case <-quits:
			ok = false
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			start := time.Now()
			if _, err := conn.Write(data); err != nil {
				log.Fatal(err)
			}
			if err := conn.Close(); err != nil {
				log.Fatal(err)
			}
			starts <- start
		}
	}
}
