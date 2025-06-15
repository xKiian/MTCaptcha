package main

import "mtcaptcha/internal/mtcaptcha"

func main() {
	solver := mtcaptcha.New("MTPublic-KzqLY1cKH", "2captcha.com")
	solver.GetChallenge()
	
	mtcaptcha.GetPulseData()
}
