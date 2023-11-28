go_crawler

Fast golang crawler

## Quick start

``` Console
$ go rum main.go -h
  -same_host
    	Crawl only the seed domain (default true)
  -seed string
    	Enter seed url (default "https://www.google.com/")
$ got rum main.go -seed "https://www.google.com/" -same_host=true
2023/11/28 14:33:20 https://www.google.com/
2023/11/28 14:33:21 https://www.google.com/imghp?hl=en&tab=wi
2023/11/28 14:33:21 https://www.google.com/preferences%3Fhl=en
2023/11/28 14:33:21 https://www.google.com/advanced_search%3Fhl=en-IN&authuser=0
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=hi&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAU
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=bn&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAY
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=te&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAc
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=mr&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAg
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=ta&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAk
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=gu&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAo
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=kn&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAs
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=ml&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCAw
2023/11/28 14:33:21 https://www.google.com/setprefs?sig=0_JOpMk_oJn2ku-iDU6MAnxCe6yL8%3D&hl=pa&source=homepage&sa=X&ved=0ahUKEwib97iFq-aCAxXuQjABHaJ1DJUQ2ZgBCA0
2023/11/28 14:33:21 https://www.google.com/intl/en/ads/
2023/11/28 14:33:21 https://www.google.com/intl/en/about.html
...
```

``` Golang 
func main() {
	var seed string = "https://www.google.com/"
	var same_host bool = true

	seedUrl, err := url.Parse(seed)
    if err != nil {
		log.Fatalf("--> Invalid seed URL: %s\n", seed)
	}
    c := NewCrawler(seedUrl, same_host)
    c.Crawl()
}
```
