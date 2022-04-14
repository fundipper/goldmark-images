package images

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var extender *Extender

type ReplaceFunc = func(src string) (string, map[string]string)

type Extender struct {
	ReplaceFunc
}

// New return initialized image render with source url replacing support.
func NewExtender(opt ReplaceFunc) goldmark.Extender {
	extender = &Extender{
		ReplaceFunc: opt,
	}
	return extender
}

func (e *Extender) Extend(m goldmark.Markdown) {
	if e.ReplaceFunc != nil {
		m.Renderer().AddOptions(
			renderer.WithNodeRenderers(
				util.Prioritized(NewRenderer(), 500),
			),
		)
	}
}
