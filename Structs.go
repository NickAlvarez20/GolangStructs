package main

//Import fmt for formatting I/O Operations
import (
	"errors"
	"fmt"
)

const maxCookies = 2 //Define max cookies here

//Cookie represents the structure of a cookie with its brand, size, and quantity
type Cookie struct{
	//Brands holds the name of manufacturer
	brand string

	//Size describes the physical size of the cookie, e.g., "small"
	size string

	//Count indicates the number of cookies in this instance
	count int
}

//Cookiebox represents a box that holds cookies
type CookieBox struct{
	//Cookies is a slice containing Cookie structs, representing the cookies in the box
	Cookies []Cookie
}

//AddCookie appends the given cookie to the Cookiebox
func (c *CookieBox) AddCookie(cookie Cookie) error{
	//Error Check Edge Case
	if len(c.Cookies) >= maxCookies {
		return errors.New("cookie box is full")
	}
	c.Cookies = append(c.Cookies, cookie)
	return nil
}

//Display cookies displayes all the cookies in the cookiebox
func (c *CookieBox) DisplayCookies(){
	for _, cookie := range c.Cookies{
		fmt.Printf("Brand: %s, Size: %s, Count: %d\n",cookie.brand,cookie.size,cookie.count)
	}
}

//Test Methods and Instantiate a Cookiebox
func main(){
	//Instantiate a cookie box with 3 cookies
	ChristmasCookies := CookieBox{[]Cookie{
		{brand: "Hersheys", size: "Large", count: 12},
		{brand: "DoubleChocy", size: "Extra Large", count: 24},
	}}
	//Add cookies to the box
	ChristmasCookies.AddCookie(Cookie{brand: "Nestle",size: "Medium",count: 12})
	//Format a print line in between
	fmt.Println()
	//Add Cookie
	ChristmasCookies.AddCookie(Cookie{brand: "GrokChristmasSnackos",size: "Small",count: 1200})
	
	//View cookies in the box
	ChristmasCookies.DisplayCookies()
}


//go run Structs.go