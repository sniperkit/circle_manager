package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/kardianos/osext"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Digest(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func Req(method string, url string, body interface{}, headers map[string]string, model interface{}) (int, error) {
	var r *http.Request
	logrus.Info(url, "호출합니다.")
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
		logrus.WithError(err).Error("요청 실패", err)
		return 0, err
	}
	if w.StatusCode >= 200 && w.StatusCode <= 400 {
		if model != nil {
			resBody, _ := ioutil.ReadAll(w.Body)
			if err := json.Unmarshal(resBody, model); err != nil {
				return -1, err
			}
			//logrus.Debug(w.StatusCode, ". BODY : ", string(resBody))
		}
	}

	logrus.Info("요청 성공. 상태 : ", method, url, w.StatusCode)
	return w.StatusCode, nil
}

func GetHostname() string {
	if hostname, err := os.Hostname(); err == nil {
		return hostname
	}
	return "unknown"
}

func GetAppPath() string {
	apppath, _ := osext.ExecutableFolder()
	return apppath
}

func MakeFirstLowerCase(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}

func IsExistsTag(reqTags string, notiTypeTags string) bool {
	mapTag := map[string]bool{}
	for _, tag := range strings.Split(reqTags, ",") {
		mapTag[tag] = true
	}

	mapNotiTypeTags := map[string]bool{}
	for _, notiTypeTag := range strings.Split(notiTypeTags, ",") {
		mapNotiTypeTags[notiTypeTag] = true
	}

	for _, tag := range strings.Split(reqTags, ",") {
		if _, ok := mapNotiTypeTags[tag]; !ok {
			return false
		}
	}
	return true
}

func MakeUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
