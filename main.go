package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/2asm/bloom-filter/bloomfilter"
	"golang.org/x/net/html"
)

type Crawler struct {
	mu    sync.Mutex
	vis   map[string]bool
	bf    *bloomfilter.BloomFilter
	seed  *url.URL
	same  bool
	sites chan *url.URL
}

// TODO: handle initiliazation with non buffered channel
func NewCrawler(seed *url.URL, same bool) *Crawler {
	c := &Crawler{
		seed:  seed,
		same:  same,
		sites: make(chan *url.URL, 1),
		vis:   make(map[string]bool),
        bf: bloomfilter.NewBloomFilter(0.01,1000),
	}
	//c.vis[seed.String()] = true
    c.bf.Add(seed.String())
	c.sites <- seed
	return c
}

func (c *Crawler) Crawl() {
	for u := range c.sites {
		log.Printf("%s\n", u.String())
		go c.handle1(u)
	}
}

func (c *Crawler) handle1(URL *url.URL) {
	// blocking
	resp, err := http.Get(URL.String())
	if err != nil {
		log.Printf("--> Invalid URL: %s\n", URL)
		return
	}
	z := html.NewTokenizer(resp.Body)
	for {
		token_type := z.Next()
		switch {
		case token_type == html.ErrorToken:
			return
		case token_type == html.StartTagToken:
			token := z.Token()
			if token.Data != "a" {
				continue
			}
			for _, a := range token.Attr {
				if a.Key != "href" {
					continue
				}
				var u *url.URL
				if strings.HasPrefix(a.Val, "https://") || strings.HasPrefix(a.Val, "http://") {
					parsedUrl, err := url.Parse(a.Val)
					if err != nil {
						break
					}
					u = parsedUrl
				} else {
					u = &url.URL{Scheme: URL.Scheme, Host: URL.Host, Path: a.Val}
				}

				// don't break, return or error between mutext Lock and Unlock
				c.mu.Lock()
				if !c.same || (c.same && u.Host == c.seed.Host) {
                    /*
					if !c.vis[u.String()] {
						c.vis[u.String()] = true
						c.sites <- u
					}
                    */
					if !c.bf.Contains(u.String()) {
                        c.bf.Add(u.String())
						c.sites <- u
					}
				}
				c.mu.Unlock()
				break
			}
		}
	}
}

func main() {

	var seed string = "https://www.google.com/"
	var same_host bool = true
	flag.StringVar(&seed, "seed", seed, "Enter seed url")
	flag.BoolVar(&same_host, "same_host", same_host, "Crawl only the seed domain")
	flag.Parse()

	seedUrl, err := url.Parse(seed)
	if err != nil {
		log.Fatalf("--> Invalid seed URL: %s\n", seed)
	}
	c := NewCrawler(seedUrl, same_host)
	c.Crawl()

}
