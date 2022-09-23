package service

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"server/internal/config"
	"strings"
	"time"
)

type RequestConfig struct {
	TryTimes int
	Duration time.Duration
}

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{
		TryTimes: 3,
		Duration: 10,
	}
}

type HttpProxyHandler struct {
}

func (h HttpProxyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request: %s %s %s %s\n", req.Method, req.Host, req.RemoteAddr, req.URL)
	log.Infof("Received request: %s %s %s %s\n", req.Method, req.Host, req.RemoteAddr, req.URL)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("body = NULL:", err.Error())
	}

	reqConfig := NewRequestConfig()
	res, err := reqConfig.requestProxy(req.Context(), body, w, req)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			log.Infof("request context canceled, err, %v\n", err)
		} else {
			log.Infof("request context failed, err, %v\n", err)
		}
		return
	}

	defer res.Body.Close()
	for k, v := range res.Header {
		w.Header().Set(k, v[0])
	}
	io.Copy(w, res.Body)
}

func (reqConfig *RequestConfig) requestProxy(ctx context.Context, body []byte, w http.ResponseWriter, req *http.Request) (*http.Response, error) {
	endpoint := config.GetTarget()
	reqURL := "http://" + endpoint.EndPoint + req.URL.String()

	proxy_req, err := http.NewRequest(req.Method, reqURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	for k, v := range req.Header {
		proxy_req.Header.Set(k, v[0])
	}

	return reqConfig.Retry(ctx, proxy_req, w)
}

func (reqConfig *RequestConfig) Retry(ctx context.Context, proxy_req *http.Request, w http.ResponseWriter) (*http.Response, error) {
	times := 0
	var res *http.Response

	for {

		cli := &http.Client{}
		var err error
		res, err = cli.Do(proxy_req)

		if err == nil && res.StatusCode == http.StatusOK { // response succeed
			break
		}

		times++
		if times > reqConfig.TryTimes {
			fmt.Printf("request %v times over the max\n", proxy_req.URL.String())
			return nil, fmt.Errorf("client retry sent times over the max %d\n", reqConfig.TryTimes)
		}

		fmt.Printf("client try to request %v for the %v times\n", proxy_req.URL.String(), times)

		time.Sleep(time.Second * reqConfig.Duration)
		select {
		case <-ctx.Done():
			return nil, context.Canceled
		default:
		}
	}

	return res, nil
}
