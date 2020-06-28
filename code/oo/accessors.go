package main

import "fmt"

type Document struct {
	title string
	pages []string
}

func main() {
	doc := NewDocument("A Work of Stunning Meaninglessness",
		[]string{"Page 1", "Page 2, no better than the first"})
	fmt.Println(doc.Title())
	doc.SetTitle("A Work of Stunning Genius")
	fmt.Println(doc.Title())
}

func NewDocument(title string, pages []string) Document {
	return Document{title: title, pages: pages}
}

func (d Document) Title() string         { return d.title }
func (d *Document) SetTitle(title string) { d.title = title }
