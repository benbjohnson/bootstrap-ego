package main

import (
	"fmt"
	"os"

	"github.com/benbjohnson/egobook"
)

//go:generate ego .

func main() {
	m := &egobook.Main{
		Modules: []*egobook.Module{
			{
				Name: "Alert",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderAlert},
				},
			},
			{
				Name: "Badge",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderBadge},
				},
			},
			{
				Name: "Breadcrumb",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderBreadcrumb},
				},
			},
			{
				Name: "Button",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderButton},
				},
			},
			{
				Name: "ButtonGroup",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderButtonGroup},
				},
			},
			{
				Name: "Card",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderCard},
				},
			},
			{
				Name: "CardColumns",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderCardColumns},
				},
			},
			{
				Name: "CardDeck",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderCardDeck},
				},
			},
			{
				Name: "Carousel",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderCarousel},
				},
			},
			{
				Name: "Collapse",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderCollapse},
				},
			},
			{
				Name: "Dropdown",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderDropdown},
				},
			},
			{
				Name: "DropdownHeader",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderDropdownHeader},
				},
			},
			{
				Name: "DropdownItem",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderDropdownItem},
				},
			},
			{
				Name: "DropdownMenu",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderDropdownMenu},
				},
			},
			{
				Name: "Form",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderForm},
				},
			},
			{
				Name: "FormCheck",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderFormCheck},
				},
			},
			{
				Name: "FormCheckInput",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderFormCheckInput},
				},
			},
			{
				Name: "FormControl",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderFormControl},
				},
			},
			{
				Name: "FormControlFile",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderFormControlFile},
				},
			},
			{
				Name: "FormGroup",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderFormGroup},
				},
			},
			{
				Name: "FormRow",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderFormRow},
				},
			},
			{
				Name: "FormText",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderFormText},
				},
			},
			{
				Name: "InputGroup",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderInputGroup},
				},
			},
			{
				Name: "Jumbotron",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderJumbotron},
				},
			},
			{
				Name: "ListGroup",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderListGroup},
				},
			},
			{
				Name: "ListGroupItem",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderListGroupItem},
				},
			},
			{
				Name: "Modal",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderModal},
				},
			},
			{
				Name: "Nav",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderNav},
				},
			},
			{
				Name: "Navbar",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderNavbar},
				},
			},
			{
				Name: "Pagination",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderPagination},
				},
			},
			{
				Name: "Pills",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderPills},
				},
			},
			{
				Name: "Popover",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderPopover},
				},
			},
			{
				Name: "Progress",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderProgress},
				},
			},
			{
				Name: "Tabs",
				Stories: []*egobook.Story{
					{Name: "Default", Render: RenderTabs},
				},
			},
		},
	}

	if err := m.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
