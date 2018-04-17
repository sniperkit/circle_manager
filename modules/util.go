package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type strCase bool

const (
	lower strCase = false
	upper strCase = true
)

var commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UI", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
var commonInitialismsReplacer *strings.Replacer

func init() {
	var commonInitialismsForReplacer []string
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
}

// Gorm
func toDBName(name string) string {
	if name == "" {
		return ""
	}

	var (
		value                        = commonInitialismsReplacer.Replace(name)
		buf                          = bytes.NewBufferString("")
		lastCase, currCase, nextCase strCase
	)

	for i, v := range value[:len(value)-1] {
		nextCase = strCase(value[i+1] >= 'A' && value[i+1] <= 'Z')
		if i > 0 {
			if currCase == upper {
				if lastCase == upper && nextCase == upper {
					buf.WriteRune(v)
				} else {
					if value[i-1] != '_' && value[i+1] != '_' {
						buf.WriteRune('_')
					}
					buf.WriteRune(v)
				}
			} else {
				buf.WriteRune(v)
				if i == len(value)-2 && nextCase == upper {
					buf.WriteRune('_')
				}
			}
		} else {
			currCase = upper
			buf.WriteRune(v)
		}
		lastCase = currCase
		currCase = nextCase
	}

	buf.WriteByte(value[len(value)-1])

	return strings.ToLower(buf.String())
}

func convInterface(raw interface{}) string {
	switch raw := raw.(type) {
	case time.Time:
		if raw.Hour() == 0 {
			return fmt.Sprintf("%d.%d.%d", raw.Year(), raw.Month(), raw.Day())
		}
		return fmt.Sprintf("%d.%d.%d %d시", raw.Year(), raw.Month(), raw.Day(), raw.Hour())
	case uint:
		return fmt.Sprintf("%d", raw)
	case string:
		return fmt.Sprintf("%s", raw)
	case []byte:
		return fmt.Sprintf("%s", string(raw))
	case int:
		return fmt.Sprintf("%d", raw)
	}

	fmt.Printf("Unknown type:%s", raw)

	return ""
}

func req(method string, url string, body interface{}, headers map[string]string, model interface{}) (int, error) {
	var r *http.Request
	if body == nil {
		r, _ = http.NewRequest(method, url, nil)
	} else {
		bodyBytes, _ := json.Marshal(body)
		r, _ = http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
		r.Header.Add("Content-Type", "application/json")
	}
	if headers != nil {
		for key, value := range headers {
			r.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	w, err := client.Do(r)
	if err != nil {
		return 0, err
	}
	if w.StatusCode >= 200 && w.StatusCode <= 400 {
		if model != nil {
			resBody, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(resBody, model); err != nil {
				return -1, err
			}
			//(w.StatusCode, ". BODY : ", string(resBody))
		}
	}

	return w.StatusCode, nil
}