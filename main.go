package main

import (
	"crypto/tls"
	//"context"
	"io/ioutil"
	"log"
	"net/http"

	//"github.com/minio/minio-go"
)

func main() {
	//accessKey := "NLYUPR3DZFAJ62SPHTPQ"
	//secKey := "p5xvehVAaR+Nmxkc5hxMvgHYJCmukN3eKpk+quwnZPA"
	//endpoint := "ams3.digitaloceanspaces.com"
	//spaceName := "qaqtus.images" // Space names must be globally unique
	//ssl := true
	//
	//// Initiate a client using DigitalOcean Spaces.
	//client, err := minio.New(endpoint, accessKey, secKey, ssl)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get("https://qaqtus.images.ams3.digitaloceanspaces.com/a.png")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}