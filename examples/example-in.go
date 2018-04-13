package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
	url := "https://weebs.cf/random/hug"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", html)
}