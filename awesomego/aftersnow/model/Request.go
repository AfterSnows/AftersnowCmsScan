package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Url     string
	Body    []byte
	Headers http.Header
	Cert    []byte
}

func Request(url string) *Response {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return &Response{
		Url:     url,
		Body:    body,
		Headers: resp.Header,
		Cert:    Getcert(resp),
	}
}

func Getcert(resp *http.Response) []byte {
	var certs []byte
	if resp.TLS != nil {
		cert := resp.TLS.PeerCertificates[0]
		var str string
		if js, err := json.Marshal(cert); err == nil {
			certs = js
		}
		str = string(certs) + cert.Issuer.String() + cert.Subject.String()
		certs = []byte(str)
	}
	return certs
}
