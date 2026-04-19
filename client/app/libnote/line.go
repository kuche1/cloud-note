package libnote

import (
	"fmt"
	"log"
	"strings"
)

type Line struct {
	contentOriginal string
	existsOriginal  bool // used to represent if a line is new

	contentCurrent string
	existsCurrent  bool // used to represent if a line has been deleted
}

func _NewLine(content string, existsOriginal bool) *Line {
	if strings.Contains(content, "\n") {
		log.Print("Warning: Found \\n in line")
	}

	return &Line{
		contentOriginal: content,
		existsOriginal:  existsOriginal,
		contentCurrent:  content,
		existsCurrent:   true,
	}
}

func (self *Line) Content() (_content string, _exists bool) {
	return self.contentCurrent, self.existsCurrent
}

func (self *Line) SetContent(newContent string) {
	self.contentCurrent = newContent
	self.existsCurrent = true
}

func (self *Line) ContentOriginal() (_content string, _exists bool) {
	return self.contentOriginal, self.existsOriginal
}

func (self *Line) StatusAndContent() (_repr string, _twoLines bool) {
	if (!self.existsOriginal) && (self.existsCurrent) {
		return fmt.Sprintf("%c %v", LineStatusAdded, self.contentCurrent), false
	}

	if (!self.existsOriginal) && (!self.existsCurrent) {
		return fmt.Sprintf("%c %v", LineStatusError, "Error: Note has neither existed before nor does it exist now"), false
	}

	if (self.existsOriginal) && (!self.existsCurrent) {
		return fmt.Sprintf("%c %v", LineStatusRemoved, self.contentOriginal), false
	}

	// if (self.existsOriginal) && (self.existsCurrent) {

	if self.contentOriginal == self.contentCurrent {
		return fmt.Sprintf("%c %v", LineStatusUpToDate, self.contentCurrent), false
	}

	return fmt.Sprintf(
		"%c %v\n%c %v",
		LineStatusRemoved, self.contentOriginal,
		LineStatusAdded, self.contentCurrent,
	), true
}

func (self *Line) Delete() (_hasAlsoNotExistedBefore bool) {
	if !self.existsCurrent {
		log.Print("Warning: Attempt to delete and already deleted note")
	}

	self.existsCurrent = false
	self.contentCurrent = ""

	return self.existsOriginal
}

func (self *Line) HasBeenChanged() bool {
	return (self.existsOriginal != self.existsCurrent) || (self.contentOriginal != self.contentCurrent)
}
