package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/mailru/easyjson"
)

type Email string

type DomainStat map[string]int

var emailPool = sync.Pool{
	New: func() interface{} {
		var email Email = ""
		return &email
	},
}

func GetDomainStat(r io.Reader, domain string) (emails DomainStat, err error) {
	emails = make(DomainStat, 1)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		email := emailPool.Get().(*Email)
		if err = easyjson.Unmarshal([]byte(sc.Text()), email); err != nil {
			return nil, fmt.Errorf("get users error: %w", err)
		}

		str := string(*email)
		if strings.Contains(str, domain) {
			emails[strings.ToLower(strings.SplitN(str, "@", 2)[1])]++
		}

		*email = ""
		emailPool.Put(email)
	}

	return
}
