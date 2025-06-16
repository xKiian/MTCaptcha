package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"mtcaptcha/internal/mtcaptcha"
	"net/http"
	"strings"
)

func main() {
	solver, _ := mtcaptcha.New("MTPublic-KzqLY1cKH", "2captcha.com", "")
	
	token, err := solver.Solve()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("[+] Token:" + token[:100])
	check(token)
}

func check(token string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	
	body := strings.NewReader(fmt.Sprintf(`{"token":"%s"}`, token))
	
	req, err := http.NewRequest("POST", "https://2captcha.com/api/v1/captcha-demo/mtcaptcha/verify", body)
	if err != nil {
		log.Fatal(err)
	}
	
	req.Header.Set("content-type", "application/json")
	
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s\n", bodyText)
}
