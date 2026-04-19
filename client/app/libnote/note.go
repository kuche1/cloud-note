package libnote

import (
	"slices"
	"strings"
)

type Note struct {
	lines            []*Line
	linesLenOriginal int
}

func NewNote(content string) *Note {
	contentSplit := strings.Split(content, "\n")

	lines := []*Line{}

	if len(content) == 0 {
		lines = []*Line{}
	} else {
		lines = make([]*Line, 0, len(contentSplit))
		for _, line := range contentSplit {
			lines = append(lines, _NewLine(line, true))
		}
	}

	return &Note{
		lines:            lines,
		linesLenOriginal: len(lines),
	}
}

func (self *Note) LineContent(index int) (_content string, _exists bool) {
	return self.lines[index].Content()
}

func (self *Note) SetLineContent(index int, content string) {
	self.lines[index].SetContent(content)
}

func (self *Note) LineStatusAndContent(index int) (_repr string, _twoLines bool) {
	return self.lines[index].StatusAndContent()
}

func (self *Note) LineDelete(index int) {
	hasAlsoExistedBefore := self.lines[index].Delete()
	if !hasAlsoExistedBefore {
		self.lines = slices.Delete(self.lines, index, index+1)
	}
}

func (self *Note) Len() int {
	return len(self.lines)
}

func (self *Note) HasBeenChanged() bool {
	if len(self.lines) != self.linesLenOriginal {
		return true
	}

	for _, line := range self.lines {
		if line.HasBeenChanged() {
			return true
		}
	}

	return false
}

func (self *Note) Content() string {
	if len(self.lines) == 0 {
		return ""
	}

	var sb strings.Builder

	for lineIdx, line := range self.lines {
		content, exists := line.Content()
		if !exists {
			continue
		}

		sb.WriteString(content)

		// omit the new line for the last line
		if lineIdx == len(self.lines)-1 {
			continue
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

func (self *Note) AddLineTop() {
	self.lines = append(
		[]*Line{_NewLine("", false)},
		self.lines...,
	)
}

func (self *Note) AddLineBot() {
	self.lines = append(
		self.lines,
		_NewLine("", false),
	)
}
