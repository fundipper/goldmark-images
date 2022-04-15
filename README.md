# goldmark-images

[Goldmark](https://github.com/yuin/goldmark) image replacer extension, base on [mdigger/goldmark-images](https://github.com/mdigger/goldmark-images). 

support image lazyload and more customize attribute.

## code

```
var source = []byte(`![alt](image.png "title")`)

func Example() {
	md := goldmark.New(
		goldmark.WithExtensions(
			images.NewExtender(
				"/loading.gif", // source
				"data-src",    // target
				map[string]string{
					"class": "lazy", // arrtibute
				}),
		),
	)
	if err := md.Convert(source, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
```
	

## view

```html
<p>
<img src="/loading.gif" alt="alt" title="title" class="lazy" data-src="image.png">
</p>
```

## thanks

[Goldmark](https://github.com/yuin/goldmark)

[mdigger](https://github.com/mdigger/goldmark-images)
