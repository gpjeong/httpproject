package rest

import (
	"bytes"
	"crypto/tls"
	"httpproject/util/logger"
	"io"
	"net/http"
	"time"
)

var client *http.Client

type Uri struct {
	Id   string `uri:"id"`
	Name string `uri:"name"`
}

// NewClientInit Rest API(HTTP1.*) 연결 공통 로직 - 커넥션 정보 설정
func NewClientInit() {
	var defaultClient *http.Client

	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		logger.Log.Error().Msgf("defaultRoundTripper not an *http.Transport")
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.IdleConnTimeout = 90 * time.Second
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	defaultTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	defaultClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: &defaultTransport,
	}

	client = defaultClient
}

func ExecuteService(uri string, method string, reqBody *bytes.Buffer) (respBody []byte, statusCode int, err error) {

	var req *http.Request
	req, err = http.NewRequest(method, uri, reqBody)
	if err != nil {
		logger.Log.Error().Msgf("HTTP Client Error: %s", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Error().Msgf("Http Client Do Error: %s", err.Error())
		return nil, http.StatusRequestTimeout, err
	}

	if resp.StatusCode == http.StatusForbidden {
		logger.Log.Error().Msgf("Http Client Response StatusCode Error: %s", resp.StatusCode)
		return nil, resp.StatusCode, nil
	}

	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error().Msgf("ReadIo Error: %s", err.Error())
	}

	return respBody, resp.StatusCode, nil
}
