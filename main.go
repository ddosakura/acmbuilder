package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func builder(i int, name string) {
	fmt.Printf("生成　%d　中。。。\n", i)
	cmd := exec.Command(name, strconv.Itoa(i))
	stdout, err := os.OpenFile(strconv.Itoa(i)+".in", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	defer stdout.Close()
	cmd.Stdout = stdout
	cmd.Run()
	fmt.Printf("生成完成\n")
}

func buildans(i int, name string) {
	fmt.Printf("计算　%d　答案中。。。\n", i)
	cmd := exec.Command(name)
	stdin, e := os.OpenFile(strconv.Itoa(i)+".in", os.O_RDONLY, 0644)
	stdout, err := os.OpenFile(strconv.Itoa(i)+".out", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil && e != nil {
		log.Fatalln(err, e)
	}
	defer stdin.Close()
	defer stdout.Close()
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Run()
	fmt.Printf("计算完成\n")
}

func running(i int, name string) {
	fmt.Printf("答题　%d　中。。。\n", i)
	cmd := exec.Command(name)
	stdin, e := os.OpenFile(strconv.Itoa(i)+".in", os.O_RDONLY, 0644)
	stdout, err := os.OpenFile(strconv.Itoa(i)+".my", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil && e != nil {
		log.Fatalln(err, e)
	}
	defer stdin.Close()
	defer stdout.Close()
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	start := time.Now()
	cmd.Run()
	cost := time.Since(start)
	fmt.Printf("答题完成　cost=[%s]\n", cost)
}

func cmp(i int) {
	fmt.Printf("比较　%d　中。。。\n", i)
	cmd := exec.Command("cmp", strconv.Itoa(i)+".out", strconv.Itoa(i)+".my")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Printf("比较完成\n")
}

func main() {
	args := os.Args
	n, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < n; i++ {
		builder(i, args[2])
		buildans(i, args[3])
		running(i, args[4])
		cmp(i)
	}
}
