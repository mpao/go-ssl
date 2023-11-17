package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var client = func() *http.Client {
	cert := "cert/localCA.crt"
	b, err := os.ReadFile(cert)
	ca, _ := x509.SystemCertPool()
	if ok := ca.AppendCertsFromPEM(b); !ok {
		log.Fatal(err)
	}
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{RootCAs: ca},
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 0,
		},
		Timeout: 20 * time.Second,
	}
	return c
}()

var url = "https://localhost:8080"

func main() {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	msg, _ := io.ReadAll(resp.Body)
	log.Println(string(msg))
}
