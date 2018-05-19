package bootstrap

import (
	"context"
	"fmt"
	"html"
	"io"
)

type Nav struct {
	ID        string
	NodeName  string
	Class     string
	Tabs      bool
	Pills     bool
	Fill      bool
	Justified bool
	Yield     func()
}

func (r *Nav) Render(ctx context.Context, w io.Writer) {
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

func (r *Nav) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="nav`)
	if r.Tabs {
		io.WriteString(w, " nav-tabs")
	}
	if r.Pills {
		io.WriteString(w, " nav-pills")
	}
	if r.Fill {
		io.WriteString(w, " nav-fill")
	}
	if r.Justified {
		io.WriteString(w, " nav-justified")
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type NavItem struct {
	ID       string
	Class    string
	Active   bool
	Dropdown bool
	Yield    func()
}

func (r *NavItem) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<li`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</li>`)
}

func (r *NavItem) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="nav-item`)
	if r.Active {
		io.WriteString(w, ` active`)
	}
	if r.Dropdown {
		io.WriteString(w, ` dropdown`)
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}

type NavLink struct {
	ID       string
	Class    string
	Href     string
	Target   string
	Active   bool
	Disabled bool
	Dropdown bool
	Yield    func()
}

func (r *NavLink) Render(ctx context.Context, w io.Writer) {
	io.WriteString(w, `<a`)
	writeAttr(w, "id", r.ID)
	r.renderClass(ctx, w)
	writeAttr(w, "href", r.Href)
	writeAttr(w, "target", r.Target)
	if r.Dropdown {
		io.WriteString(w, ` data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false"`)
	}
	io.WriteString(w, `>`)

	if r.Yield != nil {
		r.Yield()
	}

	io.WriteString(w, `</a>`)
}

func (r *NavLink) renderClass(ctx context.Context, w io.Writer) {
	io.WriteString(w, ` class="nav-link`)
	if r.Dropdown {
		io.WriteString(w, ` dropdown-toggle`)
	}
	if r.Active {
		io.WriteString(w, " active")
	}
	if r.Disabled {
		io.WriteString(w, " disabled")
	}
	if s := r.Class; s != "" {
		fmt.Fprintf(w, " %s", html.EscapeString(s))
	}
	io.WriteString(w, `"`)
}
