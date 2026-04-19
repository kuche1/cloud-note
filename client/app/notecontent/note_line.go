package notecontent

type NoteLine struct {
	contentOld            string
	contentNew            string
	contentHasBeenChanged bool
	createdFromNoLine     bool // in case the user adds a new line
}

func _NewNoteLine(content string) *NoteLine {
	return &NoteLine{
		contentOld:            content,
		contentNew:            content,
		contentHasBeenChanged: false,
		createdFromNoLine:     false,
	}
}

func _NewNoteLineFromNoLine() *NoteLine {
	return &NoteLine{
		contentOld:            "",
		contentNew:            "",
		contentHasBeenChanged: true,
		createdFromNoLine:     true,
	}
}

func (self *NoteLine) setHasNotBeenChanged() {
	self.contentOld = self.contentNew // TODO(vb): probably not effecient
	self.contentHasBeenChanged = false
	self.createdFromNoLine = false
}

func (self *NoteLine) Content() string {
	return self.contentNew
}

func (self *NoteLine) SetContent(content string) {
	self.contentNew = content

	if self.createdFromNoLine {
		self.contentHasBeenChanged = true
	} else {
		self.contentHasBeenChanged = (content != self.contentOld)
	}
}

func (self *NoteLine) ContentWithStatus() string {
	if self.contentHasBeenChanged {
		return "* " + self.contentNew
	} else {
		return "  " + self.contentNew
	}
}
