package images_test

import (
	"log"
	"os"

	images "github.com/fundipper/goldmark-images"
	"github.com/yuin/goldmark"
)

var source = []byte(`![alt](image.png "title")`)

func Example() {
	md := goldmark.New(
		goldmark.WithExtensions(
			images.NewExtender(func(src string) (string, map[string]string) {
				return "loading.gif",
					map[string]string{
						"class":    "lazy",
						"data-src": src,
					}
			}),
		),
	)
	if err := md.Convert(source, os.Stdout); err != nil {
		log.Fatal(err)
	}

	// Output:
	// <p>
	// <img src="loading.gif" alt="alt" title="title" class="lazy" data-src="image.png">
	// </p>
}
