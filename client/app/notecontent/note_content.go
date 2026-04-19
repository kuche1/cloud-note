package notecontent

import (
	"fmt"
	"strings"
)

type NoteContent struct {
	lines []*NoteLine
}

func NewNoteContent(content string) *NoteContent {
	contentSplit := strings.Split(content, "\n")

	lines := make([]*NoteLine, 0, len(contentSplit))

	for _, line := range contentSplit {
		lines = append(lines, _NewNoteLine(line))
	}

	return &NoteContent{
		lines: lines,
	}
}

func (self *NoteContent) Len() int {
	return len(self.lines)
}

func (self *NoteContent) Line(index int) *NoteLine {
	return self.lines[index]
}

func (self *NoteContent) HasBeenChanged() bool {
	// TODO: not effecient but errorless-prone

	for _, line := range self.lines {
		if line.contentHasBeenChanged {
			return true
		}
	}

	return false
}

func (self *NoteContent) SetHasNotBeenChanged() {
	for _, line := range self.lines {
		line.setHasNotBeenChanged()
	}
}

// TODO: not optimal, we can instead update the function to take the whole object
// and send the data line by line
func (self *NoteContent) AsString() (string, error) {
	// TODO: not effecient but errorless-prone

	var sb strings.Builder

	if len(self.lines) == 0 {
		return "", nil
	}

	for _, line := range self.lines[:len(self.lines)-1] {
		_, err := sb.WriteString(fmt.Sprintf("%v\n", line.contentNew))
		if err != nil {
			return "", err
		}
	}

	_, err := sb.WriteString(self.lines[len(self.lines)-1].contentNew)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

func (self *NoteContent) IsEmpty() bool {
	if len(self.lines) == 0 {
		return true
	}

	if len(self.lines) > 1 {
		return false
	}

	line := self.lines[0]
	if (line.contentNew == "") && (line.contentOld == "") {
		return true
	}

	return false
}

func (self *NoteContent) AddLineTop() {
	self.lines = append([]*NoteLine{_NewNoteLineFromNoLine()}, self.lines...)
}

func (self *NoteContent) AddLineBot() {
	self.lines = append(self.lines, _NewNoteLineFromNoLine())
}
