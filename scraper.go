package main
	import "github.com/gocolly/colly"
	import "fmt"

func main(){
 c := colly.NewCollector(colly.AllowedDomains("www.google.com"))

 c.OnHTML("a[href]", func(h *colly.HTMLElement) {
 fmt.Println(h.Text)
 })

c.Visit("www.google.com/")


}
