// Generated by ego.
// DO NOT EDIT

//line alert.ego:1
package bootstrap

import "fmt"
import "html"
import "io"
import "context"

type Alert struct {
	Yield func()
}

func (c *Alert) Render(ctx context.Context, w io.Writer) {
//line alert.ego:10
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
