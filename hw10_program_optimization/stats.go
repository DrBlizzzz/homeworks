package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"
	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	var parser fastjson.Parser
	statistics := make(DomainStat)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		v, err := parser.Parse(scanner.Text())
		if err != nil {
			return nil, err
		}
		email := v.Get("Email").String()
		if strings.Contains(email, domain) {
			key := strings.ToLower(
				strings.SplitN(
					email[1:len(email)-1], "@", 2,
				)[1],
			)
			_, ok := statistics[key]
			if ok {
				statistics[key]++
			} else {
				statistics[key] = 1
			}
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
	return statistics, nil
}
