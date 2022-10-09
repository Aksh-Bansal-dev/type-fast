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

func main() {
	flag.Parse()
	para := getRandPara(*length)
	fmt.Println(para)
	startTime := time.Now()
	reader := bufio.NewReader(os.Stdin)
	userStr, _, _ := reader.ReadLine()
	if string(userStr) == para {
		fmt.Println("succcess")
	} else {
		fmt.Println("false")
	}
	fmt.Printf("Speed: %v wps", (*length)/int(time.Now().Sub(startTime).Minutes()))
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
