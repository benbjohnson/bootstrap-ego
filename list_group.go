package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type ListGroup struct {
	ID       string
	NodeName string
	Class    string
	Flush    bool
	Yield    func()
}

func (r *ListGroup) Render(ctx context.Context, w io.Writer) {
	nodeName := r.NodeName
	if nodeName == "" {
		nodeName = "ul"
	}

	fmt.Fprintf(w, `<%s`, nodeName)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	fmt.Fprintf(w, `</%s>`, nodeName)
}

func (r *ListGroup) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="list-group`)
	if r.Flush {
		io.WriteString(w, ` list-group-flush`)
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, ` %s`, html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type ListGroupItem struct {
	ID       string
	Class    string
	Type     string
	Style    string
	Href     string
	Form     string
	Value    string
	Target   string
	Active   bool
	Disabled bool
	Yield    func()
}

func (r *ListGroupItem) NodeName() string {
	if r.Href != "" {
		return "a"
	} else if r.Type != "" {
		return "button"
	}
	return "li"
}

func (r *ListGroupItem) Render(ctx context.Context, w io.Writer) {
	nodeName := r.NodeName()

	fmt.Fprintf(w, `<%s`, nodeName)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	r.renderType(ctx, w)

	if r.Href != "" {
		fmt.Fprintf(w, ` href="%s"`, html.EscapeString(r.Href))
	}

	if r.Disabled && nodeName == "button" {
		io.WriteString(w, " disabled")
	}
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	fmt.Fprintf(w, `</%s>`, nodeName)
}

func (r *ListGroupItem) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="list-group-item`)

	if nodeName := r.NodeName(); nodeName == "a" || nodeName == "button" {
		io.WriteString(w, ` list-group-item-action`)
	}

	switch r.Style {
	case "primary", "secondary", "success", "danger", "warning", "info", "light", "dark":
		fmt.Fprintf(w, ` list-group-item-%s`, r.Style)
	case "":
	default:
		Logger.Printf("bootstrap.ListGroupITem: Invalid style: %q", r.Style)
	}

	if r.Active {
		io.WriteString(w, " active")
	}
	if r.Disabled && r.NodeName() != "button" {
		io.WriteString(w, " disabled")
	}

	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}

	io.WriteString(w, `"`)
}

func (r *ListGroupItem) renderType(ctx context.Context, w io.Writer) {
	if r.NodeName() != "button" {
		return
	}

	switch r.Type {
	case "submit", "reset", "button":
		fmt.Fprintf(w, ` type="%s"`, r.Type)
	case "":
	default:
		Logger.Printf("bootstrap.ListGroupItem: Invalid type: %q", r.Type)
	}
}
