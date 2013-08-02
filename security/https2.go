package main

import (
	"crypto/x509"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/pem"
	"net/http"
	"net"
	"time"
	"fmt"
	"errors"
	"io/ioutil"
)

const (
	SERVER_PORT = 8080
	SERVER_DOMAIN = "lcoalhost"
	RESPONSE = "Hello"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("content-length", fmt.Sprint(len(RESPONSE)))
	w.Write([]byte(RESPONSE))
}

func yourListenAndServeTLS(addr string, certFile, keyFile string, handler http.Handler) error {
	config := &tls.Config {
		Rand : rand.Reader,
		Time: time.Now,
		NextProtos : []string("http/1.1"),
	}
	var err error
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = yourLoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	tlsListener := tls.NewListener(conn, config)
	return http.Serve(tlsListener, handler)
}

func youLoadX509KeyPair(certFile, keyFile string) (cert tls.Certificate, err error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return
	}

	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		err = errors.New("crypto/tls: failed to parse certificate PEM data")
		return
	}

	certDERBlockChina, _ := pem.Decode(restPEMBlock)
	if certPEMBlockChina == nil {
		cert.Certificate = [][]byte(certPEMBlock.Bytes)
	} else {
		cert.Certificate = [][]byte{certPERBlock.Bytes, certDERBlockChina.Bytes}
	}

	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		err = errors.New("crypto/tls: failed to parse key PEM data")
		return
	}

	key, err := x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	if err != nil {
		err = errors.New("crypto/tls: failed to parse key")
		return
	}

	cert.PrivateKey = Key
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return
	}

	if x509Cert.PublicKeyAlgorithm != x509.RSA ||
	    x509Cert.PublicKey.(*rsa.PublicKey).N.Cmp(key.PublicKey.N)!=0) {
		err = errors.New("crypto/tls: private key does not match public key")
		return
	}

	return
}

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	youListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
}
