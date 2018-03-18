package main

import (
	"fmt"
	"os"

	"github.com/benbjohnson/egobook"
)

//go:generate ego .

func main() {
	app := egobook.App{
		Head: Head,
		Modules: []*egobook.Module{
			{
				Name: "Alert",
				Stories: []*egobook.Story{
					{Name: "Primary", Render: AlertRenderer("primary")},
					{Name: "Secondary", Render: AlertRenderer("secondary")},
					{Name: "Success", Render: AlertRenderer("success")},
					{Name: "Danger", Render: AlertRenderer("danger")},
					{Name: "Warning", Render: AlertRenderer("warning")},
					{Name: "Info", Render: AlertRenderer("info")},
					{Name: "Light", Render: AlertRenderer("light")},
					{Name: "Dark", Render: AlertRenderer("dark")},
				},
			},
			{
				Name: "Badge",
				Stories: []*egobook.Story{
					{Name: "Primary", Render: BadgeRenderer("primary")},
					{Name: "Secondary", Render: BadgeRenderer("secondary")},
					{Name: "Success", Render: BadgeRenderer("success")},
					{Name: "Danger", Render: BadgeRenderer("danger")},
					{Name: "Warning", Render: BadgeRenderer("warning")},
					{Name: "Info", Render: BadgeRenderer("info")},
					{Name: "Light", Render: BadgeRenderer("light")},
					{Name: "Dark", Render: BadgeRenderer("dark")},
				},
			},
			{
				Name: "Breadcrumb",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderBreadcrumb},
				},
			},
		},
	}

	m := &egobook.Main{App: app}
	if err := m.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
