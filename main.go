package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	findAndClick()
	findAndClick100Times()
}

func findAndClick() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://player.antyradio.pl/Turbo-Top")

	el := page.MustElementX("/html/body/main/div[2]/div[2]/div[1]/div/div[2]/div[1]/div[2]/div/div[18]/div/div[4]/div/a")
	fmt.Println(el)
	fmt.Println(el.Describe(0, false))
	el.MustClick()

	el2 := page.MustElementX("/html/body/main/div[2]/div[2]/div[1]/div/div[2]/div[1]/div[2]/div/div[18]/div/div[4]/div/a")
	fmt.Println(el2)
	fmt.Println(el2.Describe(0, false))
}

func findAndClick100Times() {
	for i := 0; i < 100; i++ {
		findAndClick()
	}
}
