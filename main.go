package main

import (
	"fmt"
	"mtcaptcha/internal/mtcaptcha"
)

func main() {
	solver, _ := mtcaptcha.New("MTPublic-KzqLY1cKH", "2captcha.com", "")
	res, err := solver.GetChallenge()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	res2, err := solver.GetImage()
	if err != nil {
		panic(err)
	}
	fmt.Println(res2)
	
	fmt.Print(mtcaptcha.SolveFoldChallenge("VSUFhjii7scj36xjdJJgYVCyK7pJ_s0tUX8cjI68iDb5izqgVyHXZDwcJJcDNHSb", 30, 995))
}
