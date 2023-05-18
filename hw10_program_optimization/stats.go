package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	var mu sync.RWMutex
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(r)
	pipe := make(chan string)
	go func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			pipe <- scanner.Text()
		}
		close(pipe)
	}(scanner)
	domainStat := make(DomainStat)
	template := fmt.Sprintf("@[\\w]+\\.%s", domain)
	re, _ := regexp.Compile(template)
	for chanValue := range pipe {
		wg.Add(1)
		go func(chanValue string, wg *sync.WaitGroup) {
			key := strings.ToLower(
				re.FindString(chanValue),
			)
			if key != "" {
				mu.Lock()
				domainStat[key[1:]]++
				mu.Unlock()
			}
			wg.Done()
		}(chanValue, &wg)
	}
	wg.Wait()
	return domainStat, nil
}
