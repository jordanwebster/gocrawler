package main

import (
    "fmt"
	"net/http"
	"net/url"
	"os"

    "golang.org/x/net/html"
)

func main() {
    s := os.Args[1]
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    results := crawl(u)
    for _, result := range results {
        fmt.Println(result.String())
    }
}


func crawl(u *url.URL) []url.URL {
    results := make([]url.URL, 10)
    response, err := http.Get(u.String())
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    var doc, err2 = html.Parse(response.Body)
    if err2 != nil {
        panic(err2)
    }
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    new_url, new_url_err := url.Parse(attr.Val)
                    if new_url_err != nil {
                        panic(new_url_err)
                    }
                    results = append(results, *new_url)
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    return results
}
