package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var sem = make(chan struct{}, 1)

func read() {
	reader := bufio.NewReader(os.Stdin)
	m1, _ := reader.ReadString('\n')
	<-sem // Attendre le sémaphore
	for i := 0; i < 10; i++ {
		fmt.Fprint(os.Stderr, ".")
		time.Sleep(1 * time.Second)
	}
	process(m1)
	sem <- struct{}{} // Rendre le sémaphore
}

func process(m1 string) {
	fmt.Fprintf(os.Stderr, "<PID=%d> Réception de %s", os.Getpid(), m1)
	os.Stderr.Sync()
}

func write() {
	m2 := "Hello SR05\n"
	for i := 0; i < 10; i++ {
		fmt.Fprint(os.Stderr, ".")
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(os.Stdout, m2)
}

func main() {
	sem <- struct{}{} // Initialiser le sémaphore

	go func() {
		for {
			read()
		}
	}()

	for {
		<-sem // Attendre le sémaphore
		write()
		sem <- struct{}{} // Rendre le sémaphore
		time.Sleep(10 * time.Second)
	}
}
