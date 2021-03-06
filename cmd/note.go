package cmd

import (
	"broadcastle.co/code/crm/code/note"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief description of your command",
}

var noteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note.",
	Args:  cobra.NoArgs,
	Run:   note.Create,
}

var noteEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing note.",
	Args:  cobra.ExactArgs(1),
	Run:   note.Edit,
}

var noteRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a note.",
	Args:  cobra.MinimumNArgs(1),
	Run:   note.Remove,
}

var noteViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Look at the note(s) listed.",
	Run:   note.View,
}

func init() {

	noteCmd.Aliases = append(contactCmd.Aliases, "notes")

	RootCmd.AddCommand(noteCmd)

	noteCmd.AddCommand(noteAddCmd)
	noteCmd.AddCommand(noteEditCmd)
	noteCmd.AddCommand(noteRemoveCmd)
	noteCmd.AddCommand(noteViewCmd)

}
