package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"gitlab.com/clitool/sitemap/fetch"
)

var usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s %s\nFlags:\n", os.Args[0], "<hostname>")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	// argsWithProg := os.Args
	wordPtr := flag.String("url", "", "top level sitemap url")
	// depthPtr := flag.Int("depth", 1, "crawl depth")
	// boolPtr := flag.Bool("fork", false, "a bool")

	flag.Parse()

	// // fmt.Println(argsWithProg)
	// fmt.Println("url:", *wordPtr)
	// fmt.Println("depth:", *depthPtr)
	// // fmt.Println("fork:", *boolPtr)
	// // fmt.Println("args:", flag.Args())

	args := flag.Args()
	if *wordPtr == "" && len(args) != 1 || *wordPtr != "" && len(args) >= 1 {
		exitOnError()
	}

	var sitemapURL string
	if *wordPtr != "" {
		if !isURL(*wordPtr) {
			exitOnError()
		}
		sitemapURL = *wordPtr
	} else {
		hostname := args[0]
		sitemapURL = findSitemapURL(hostname)
	}

	// fmt.Println("sitemapURL:", sitemapURL)
	fetch.SitemapItems(sitemapURL)
}

func findSitemapURL(hostname string) string {
	smURL := hostname
	if !isURL(smURL) {
		if !strings.HasPrefix(smURL, "http") {
			// TBD: Add prefix "www." to hostname ????
			smURL = "http://" + hostname
		} else {
			exitOnError()
		}
	}
	if !strings.HasSuffix(smURL, "/sitemap.xml") && !strings.HasSuffix(smURL, "/sitemap.txt") {
		// tbd: Check robots.txt ("Sitemap:") ????
		if !strings.HasSuffix(smURL, "/") {
			// tbd: check if it already includes sitemap file name at the end???
			smURL += "/sitemap.xml"
		} else {
			smURL += "sitemap.xml"
		}
	}
	return smURL
}

func isURL(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	return true
}

func exitOnError() {
	flag.Usage()
	os.Exit(1)
}
