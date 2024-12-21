//Import package
package main

//Import fmt for formatting I/O Operations
import "fmt"

//Define a Cookie Struct
type Cookie struct{
	name string
	count int
	size string
}

//Define a CookieBox Struct | Holds Cookie Boxes | Slice of Cookie for Dynamic Sizing and Data Manipulation
type CookieBox struct{
	Cookies []Cookie
}

//create a function that adds a new cookie to the cookie box
func (c *CookieBox) AddCookie(cookie Cookie){
	c.Cookies = append(c.Cookies, cookie)
}

//Implement ViewCookies Function
func (c CookieBox) ViewCookies(){
	for _, cookie := range c.Cookies{
		fmt.Printf("Cookie Name: %s, Cookie Count: %d, Cookie Size: %s\n",cookie.name, cookie.count, cookie.size)
	}
}

//Test in main func 
func main(){
	//Instantiate a Cookie Box
	GrokCookieBox := CookieBox{
		Cookies: []Cookie{
			{name: "Chocy Chip", size: "Large", count: 12},
			{name: "White Chocy", size: "Medium", count: 12},
		},
	}
	//View Cookies
	GrokCookieBox.ViewCookies()
	//Add break in line here for formatting
	fmt.Println()
	//Add Cookies
	GrokCookieBox.AddCookie(Cookie{name: "White Chocy Chip", size: "Medium", count: 12})
	//View Cookies
	GrokCookieBox.ViewCookies()
}





//go run Structs.go