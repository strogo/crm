package tui

import (
	"broadcastle.co/code/crm/code/db"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
)

// App holds the application.
type App struct {
	*tview.Application
}

// Remove will display a modal.
func (a *App) Remove(v interface{}) *tview.Modal {

	modal := tview.NewModal()

	switch v.(type) {
	case db.Contact:
		removeContact(a, modal, v.(db.Contact))
	default:
		logrus.Fatal("interface not supported in tui.App.Remove()")
	}

	return modal

}

func removeContact(a *App, modal *tview.Modal, data db.Contact) {
	modal.SetText("Remove " + data.Name + "?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "Yes":
				if err := data.Remove(); err != nil {
					logrus.Fatal(err)
				}

			}

			a.Stop()
		})
}

// Table takes v and displays it as a table.
func (a *App) Table(v interface{}) *tview.Table {

	table := tview.NewTable().SetBorders(true)

	switch v.(type) {
	case []db.Contact:
		tableContacts(table, v.([]db.Contact))
	default:
		logrus.Fatal("interface not supported in tui.App.Table()")
	}

	table.Select(0, 0).
		SetFixed(1, 1).
		SetDoneFunc(func(key tcell.Key) {

			switch key {
			case tcell.KeyEscape:
				a.Stop()
			}

		})

	return table
}

// Form takes interface and return a form for it.
func (a *App) Form(v interface{}) *tview.Form {

	form := tview.NewForm()

	switch v.(type) {
	case db.Contact:
		formContacts(a, form, v.(db.Contact))
	default:
		logrus.Fatal("interface not supported")
	}

	return form

}

// ContactForm for Contact.
func (a *App) ContactForm(id uint) *tview.Form {

	contact := db.Contact{}

	if id != 0 {

		contact.ID = id

		if err := contact.Query(); err != nil {
			logrus.Fatal(err)
		}

	}

	drop := 0

	switch {
	case contact.Lead:
		drop = 0
	case contact.Advocate:
		drop = 1
	case contact.Customer:
		drop = 2
	case contact.Subscriber:
		drop = 3
	}

	return tview.NewForm().
		AddInputField("Name", contact.Name, 0, nil, func(text string) {
			contact.Name = text
		}).
		AddInputField("Email", contact.Email, 0, nil, func(text string) {
			contact.Email = text
		}).
		AddInputField("Phone Number", contact.Number, 0, nil, func(text string) {
			contact.Number = text
		}).
		AddDropDown("Relationship", []string{"Lead", "Advocate", "Customer", "Subscriber"}, drop, func(option string, optionIndex int) {

			contact.Lead = false
			contact.Advocate = false
			contact.Customer = false
			contact.Subscriber = false

			switch optionIndex {
			case 0:
				contact.Lead = true
			case 1:
				contact.Advocate = true
			case 2:
				contact.Customer = true
			case 3:
				contact.Subscriber = true
				contact.Customer = true
			}
		}).
		AddCheckbox("Contacted", contact.Contacted, func(checked bool) {
			contact.Contacted = checked
		}).
		AddButton("Save", func() {

			defer a.Stop()

			if contact.ID != 0 {
				if err := contact.Update(); err != nil {
					logrus.Fatal(err)
				}
				return
			}

			saving := contact

			if err := saving.Create(); err != nil {
				logrus.Fatal(err)
			}

		}).
		AddButton("Quit", func() {
			a.Stop()
		})

}
