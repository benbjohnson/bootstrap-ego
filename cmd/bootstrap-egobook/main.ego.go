// Generated by ego.
// DO NOT EDIT

//line main.ego:1
package main

import "fmt"
import "html"
import "io"
import "context"

import "github.com/benbjohnson/bootstrap-ego"

func Head(ctx context.Context, w io.Writer) {
//line main.ego:7
	_, _ = io.WriteString(w, "<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.3/css/bootstrap.min.css\" integrity=\"sha384-Zug+QiDoJOrZ5t4lssLdxGhVrurbmBWopoEl+M6BdEfwnCJZtKxi1KgxUyJq13dy\" crossorigin=\"anonymous\">\n\t<script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\n\t<script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js\" integrity=\"sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q\" crossorigin=\"anonymous\"></script>\n\t<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.3/js/bootstrap.min.js\" integrity=\"sha384-a5N7Y/aK3qNeh15eJKGWxsqtnX/wWdSZSKp+81YjTmS15nvnvxKHuzaWwXHDli+4\" crossorigin=\"anonymous\"></script>\n")
//line main.ego:11
}

func AlertRenderer(style string) func(ctx context.Context, w io.Writer) {
	return func(ctx context.Context, w io.Writer) {
//line main.ego:15
		_, _ = io.WriteString(w, "\n")
//line main.ego:15
		{
			egoComponent := bootstrap.Alert{
				Style: style,
				Yield: func() {
//line main.ego:16
					_, _ = io.WriteString(w, "This is a ")
//line main.ego:16
					_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(style)))
//line main.ego:16
					_, _ = io.WriteString(w, " alert with <a href=\"#\" class=\"alert-link\">an example link</a>. Give it a click if you like.\n")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:18
		_, _ = io.WriteString(w, "\n\n")
//line main.ego:19
		{
			egoComponent := bootstrap.Alert{
				Style: style,
				Yield: func() {
//line main.ego:20
					_, _ = io.WriteString(w, "<h4 class=\"alert-heading\">Well done!</h4>\n\t<p>Aww yeah, you successfully read this important alert message. This example text is going to run a bit longer so that you can see how spacing within an alert works with this kind of content.</p>\n\t<hr/>\n\t<p class=\"mb-0\">Whenever you need to, be sure to use margin utilities to keep things nice and tidy.</p>\n")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:25
		_, _ = io.WriteString(w, "\n\n")
//line main.ego:26
		{
			egoComponent := bootstrap.Alert{
				Style:       "warning",
				Class:       "fade show",
				Dismissable: true,
				Yield: func() {
//line main.ego:27
					_, _ = io.WriteString(w, "<strong>Holy guacamole!</strong> You should check in on some of those fields below.\n\t")
//line main.ego:28
					{
						egoComponent := bootstrap.CloseButton{}
						egoComponent.Render(ctx, w)
					}
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:30
		_, _ = io.WriteString(w, "\n")
//line main.ego:30
	}
}

func BadgeRenderer(style string) func(ctx context.Context, w io.Writer) {
	return func(ctx context.Context, w io.Writer) {
//line main.ego:35
		_, _ = io.WriteString(w, "\n<h1>Example heading ")
//line main.ego:35
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Yield: func() {
//line main.ego:35
					_, _ = io.WriteString(w, "New")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:35
		_, _ = io.WriteString(w, "</h1>\n<h2>Example heading ")
//line main.ego:36
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Yield: func() {
//line main.ego:36
					_, _ = io.WriteString(w, "New")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:36
		_, _ = io.WriteString(w, "</h2>\n<h3>Example heading ")
//line main.ego:37
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Yield: func() {
//line main.ego:37
					_, _ = io.WriteString(w, "New")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:37
		_, _ = io.WriteString(w, "</h3>\n<h4>Example heading ")
//line main.ego:38
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Yield: func() {
//line main.ego:38
					_, _ = io.WriteString(w, "New")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:38
		_, _ = io.WriteString(w, "</h4>\n<h5>Example heading ")
//line main.ego:39
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Yield: func() {
//line main.ego:39
					_, _ = io.WriteString(w, "New")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:39
		_, _ = io.WriteString(w, "</h5>\n<h6>Example heading ")
//line main.ego:40
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Yield: func() {
//line main.ego:40
					_, _ = io.WriteString(w, "New")
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:40
		_, _ = io.WriteString(w, "</h6>\n\n<p>\n\t")
//line main.ego:43
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Yield: func() {
//line main.ego:43
					_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(style)))
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:44
		_, _ = io.WriteString(w, "\n\t")
//line main.ego:44
		{
			egoComponent := bootstrap.Badge{
				Style: style,
				Pill:  true,
				Yield: func() {
//line main.ego:44
					_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(style)))
				},
			}
			egoComponent.Render(ctx, w)
		}
//line main.ego:45
		_, _ = io.WriteString(w, "\n</p>\n")
//line main.ego:46
	}
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
