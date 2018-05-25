package modules

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	beegoContext "github.com/astaxie/beego/context"
	"github.com/olimgroup/makezip/envs"
	"github.com/olimgroup/makezip/utils"
	"github.com/sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v5"
	elogrus "gopkg.in/sohlich/elogrus.v2"
)

const (
	HTTP_LOG_NAME = "http"
)

var (
	esClient *elastic.Client
)

// AccessLogRecord struct for holding access log data.
type AccessLogRecord struct {
	RemoteAddr     string    `json:"remote_addr"`
	RequestTime    time.Time `json:"request_time"`
	RequestMethod  string    `json:"request_method"`
	Request        string    `json:"request"`
	ServerProtocol string    `json:"server_protocol"`
	Host           string    `json:"host"`
	Status         int       `json:"status"`
	BodyBytesSent  int64     `json:"body_bytes_sent"`
	ElapsedTime    float64   `json:"elapsed_time"`
	HTTPReferrer   string    `json:"http_referrer"`
	HTTPUserAgent  string    `json:"http_user_agent"`
	RemoteUser     string    `json:"remote_user"`
}

func InitLogsToES(logLevel logrus.Level, esURL string, httpLog bool, excludURLs ...string) {
	initLogs(logLevel, esURL, httpLog, excludURLs...)
}

func InitLogs(logLevel logrus.Level) {
	initLogs(logLevel, "", false)
}

func initLogs(logLevel logrus.Level, esURL string, httpLog bool, excludURLs ...string) {
	logrus.SetLevel(logLevel)

	if esURL != "" {
		var err error
		if esClient, err = elastic.NewClient(elastic.SetURL(esURL), elastic.SetSniff(false)); err != nil {
			panic(err)
		}

		go runESFlush()
	}

	if httpLog {
		initHTTPLog(excludURLs...)
	}
}

func NewLCircleESLogger(logLevel logrus.Level, esURL string, esIndex string) (*logrus.Logger, error) {
	logger := logrus.New()
	if esURL != "" && esIndex != "" {
		hostname := utils.GetHostname()
		client, err := elastic.NewClient(elastic.SetURL(esURL), elastic.SetSniff(false))
		if err != nil {
			return nil, err
		}
		hook, err := elogrus.NewElasticHook(client, hostname, logLevel, esIndex)
		if err != nil {
			return nil, err
		}
		logger.AddHook(hook)
	}
	return logger, nil
}

func genIndex(typeName string) string {
	return fmt.Sprintf("%s-%s", envs.EsIndexPrefix, typeName)
}

func initHTTPLog(excludURLs ...string) {
	beego.InsertFilter("*", beego.BeforeRouter, func(context *beegoContext.Context) {
		context.Input.SetData("startTime", time.Now())
	}, false)

	beego.InsertFilter("*", beego.FinishRouter, func(context *beegoContext.Context) {
		r := context.Request

		if matched, _ := regexp.MatchString("^/system/status$", r.RequestURI); matched {
			return
		}
		for _, excludURLs := range excludURLs {
			if matched, _ := regexp.MatchString(excludURLs, r.RequestURI); matched {
				return
			}
		}

		startTime := context.Input.GetData("startTime").(time.Time)
		statusCode := context.ResponseWriter.Status
		if statusCode == 0 {
			statusCode = 200
		}
		timeDur := time.Since(startTime)

		put(genIndex(HTTP_LOG_NAME), "accesslog", &AccessLogRecord{
			RemoteAddr:     context.Input.IP(),
			RequestTime:    startTime,
			RequestMethod:  r.Method,
			Request:        fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto),
			ServerProtocol: r.Proto,
			Host:           r.Host,
			Status:         statusCode,
			ElapsedTime:    timeDur.Seconds(),
			HTTPReferrer:   r.Header.Get("Referer"),
			HTTPUserAgent:  r.Header.Get("User-Agent"),
			RemoteUser:     r.Header.Get("Remote-User"),
			BodyBytesSent:  int64(len(context.Input.RequestBody)),
		})
	}, false)
}

func put(index string, esType string, data interface{}) error {
	if esClient != nil {
		logrus.
			WithField("index", index).
			WithField("esType", esType).
			Debug("Send event log to elasticsearch")
		_, err := esClient.Index().Index(index).Type(esType).BodyJson(data).Do(context.Background())
		return err
	}
	return nil
}

func runESFlush() {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if _, err := esClient.Flush().Do(context.Background()); err != nil {
				logrus.WithError(err).Error("runESFlush")
			}
		}
	}
}
