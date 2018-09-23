package fetch

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const sitemapURL = "https://www.opentable.com/sitemap.xml"

type location struct {
	Loc string `xml:"loc"`
}

func (l location) String() string {
	return l.Loc
}

type sitemapIndex struct {
	Sitemap []location `xml:"sitemap"`
}

func Fetch1() {
	fmt.Println("Hi World!!!")

	res, err := http.Get(sitemapURL)
	if err != nil {
		fmt.Println("http.Get Error", err)
		return
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll Error", err)
		return
	}
	defer res.Body.Close()
	// body := string(bytes)
	// fmt.Println("body = ", body)

	var si sitemapIndex
	xml.Unmarshal(bytes, &si)

	sm := si.Sitemap
	// fmt.Println(sm)
	for k, v := range sm {
		// fmt.Println(k, v)
		fmt.Printf("*** %d: %s\n", k, v)

		url := v.String()
		res, err := http.Get(url)
		if err != nil {
			continue
		}

		if strings.HasSuffix(url, ".txt") {
			scanner := bufio.NewScanner(res.Body)
			defer res.Body.Close()
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

		} else if strings.HasSuffix(url, ".xml") {
			// bytes, err := ioutil.ReadAll(res.Body)
			// if err != nil {
			// 	continue
			// }
			// defer res.Body.Close()
			// body := string(bytes)
			// fmt.Println("XML Body = ", body)
		} else {
			// Ignore for now.
		}

	}

}
