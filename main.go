package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const url string = "https://api.pwnedpasswords.com/range/"

func main() {

	pass := flag.String("p", "foo", "Password under test")
	flag.Parse()

	sha1Sum := sha1.Sum([]byte(*pass))
	sha1SumHexString := hex.EncodeToString(sha1Sum[:])

	req, err := http.NewRequest("GET", url + sha1SumHexString[:5], nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var found string
	for _, value := range strings.Split(string(body), "\n"){
		if strings.Contains(value, strings.ToUpper(sha1SumHexString[5:])) {
			found = value
			break
		}
	}

	if found != "" {
		fmt.Printf("Your password with hash %s was leaked %s times!\n",
			strings.TrimSpace(strings.Split(found,":")[0]),
			strings.TrimSpace(strings.Split(found,":")[1]))
	} else {
		fmt.Println("You're safe!")
	}
}