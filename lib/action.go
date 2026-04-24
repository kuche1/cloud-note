package lib

import "fmt"

type Action uint8

const (
	// IMPROVE000: Ideally we would also have some mechanism in place that checks if the content has changed in between the next 2 calls
	ActionGetNoteContent Action = iota
	ActionSetNoteContent
	ActionListNotes
	ActionCreateNewNote
	ActionDeleteExistingNote
	ActionPing
	ActionRenameNote
	Actions
)

func (self Action) FromUint8(data uint8) (Action, error) {
	if data >= uint8(Actions) {
		return 0, fmt.Errorf("Got unknown action uint8: %v", data)
	}
	return Action(data), nil
}

func (self Action) ToUint8() uint8 {
	return uint8(self)
}
