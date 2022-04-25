package images

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var extender *Extender

type Extender struct {
	Source    string
	Target    string
	Attribute map[string]string
}

// New return initialized image render with source url replacing support.
func NewExtender(source, target string, attribute map[string]string) goldmark.Extender {
	extender = &Extender{
		Source:    source,
		Target:    target,
		Attribute: attribute,
	}
	return extender
}

func (e *Extender) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewRenderer(), 500),
		),
	)
}
