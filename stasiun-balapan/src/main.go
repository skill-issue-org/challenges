package main

/**
 * Skill Issue Weekly Challenges
 *
 * Title: Stasiun Balapan
 * Points: 6
 * Author: dwisiswant0
**/

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"

	"math/rand"
	"os/signal"

	"skill-issue.org/stasiun-balapan/pkg/consts"
	"skill-issue.org/stasiun-balapan/pkg/errors"
	"skill-issue.org/stasiun-balapan/pkg/vars"
)

var (
	attempt, i = 0, 0

	mu sync.Mutex
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	sigs := make(chan os.Signal, 1)

	fmt.Println(consts.Title)
	fmt.Printf("%s: ", consts.Subtitle)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cobaLagi()
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if i%2 != 0 {
			wg.Add(1)
			go masuk()
		}

		i++
	}

	wg.Wait()
	getFlag()
}

func masuk() {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	attempt++
}

func getFlag() {
	println()

	if attempt >= rand.Intn(consts.N-50)+50 {
		fmt.Printf("%s: %s\n", consts.Result, vars.Flag)
	} else {
		cobaLagi()
	}
}

func cobaLagi() {
	panic(errors.ErrCobaLagi)
}
