package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var length = flag.Int("n", 100, "Number of words")
var timeout = flag.Int("t", 3600, "Timeout time (in seconds)")

func main() {
	flag.Parse()
	para := getRandPara(*length)
	fmt.Println(para)

	reader := bufio.NewReader(os.Stdin)
	startTime := time.Now()
	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	go func() {
		<-timer.C
		fmt.Println("\n Time over!")
		os.Exit(1)
	}()
	userPara, _, _ := reader.ReadLine()

	numCorrWords := countCorrWord(para, string(userPara))
	fmt.Printf("Speed: %v wps\n", int(float64(numCorrWords)/time.Now().Sub(startTime).Minutes()))
	fmt.Printf("Accuracy: %v%%\n", int(float64(numCorrWords)/float64(*length)*100))
}

func getRandPara(n int) string {
	rand.Seed(time.Now().Unix())
	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(wordlist[rand.Intn(len(wordlist))])
		if i != n-1 {
			str.WriteByte(' ')
		}
	}
	return str.String()
}

func countCorrWord(para string, userPara string) int {
	words := strings.Split(para, " ")
	userWords := strings.Split(userPara, " ")
	count := len(userWords)
	for i := range userWords {
		if userWords[i] != words[i] {
			count--
		}
	}
	return count
}
