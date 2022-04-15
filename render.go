package images

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// Renderer struct is a renderer.NodeRenderer implementation for the extension.
type Renderer struct {
	html.Config
}

// NewRenderer builds a new Renderer with given options and returns it.
func NewRenderer() renderer.NodeRenderer {
	return &Renderer{}
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs interface.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindImage, r.renderImage)
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Image)

	for k, v := range extender.Attribute {
		n.SetAttributeString(k, util.StringToReadOnlyBytes(v))
	}
	n.SetAttributeString(extender.Target, n.Destination)
	n.Destination = util.StringToReadOnlyBytes(extender.Source)

	w.WriteString(`<img src="`)
	if r.Unsafe || !html.IsDangerousURL(n.Destination) {
		w.Write(util.EscapeHTML(util.URLEscape(n.Destination, true)))
	}
	w.WriteString(`" alt="`)
	w.Write(n.Text(source))
	w.WriteString(`"`)

	if n.Title != nil {
		w.WriteString(` title="`)
		w.Write(n.Title)
		w.WriteString(`"`)
	}

	if n.Attributes() != nil {
		html.RenderAttributes(w, n, html.ImageAttributeFilter)
	}

	if r.XHTML {
		w.WriteString(" />")
	} else {
		w.WriteString(">")
	}

	return ast.WalkSkipChildren, nil
}
