package libnote

type LineStatus byte

const (
	LineStatusUpToDate LineStatus = ' '
	LineStatusRemoved  LineStatus = '-'
	LineStatusAdded    LineStatus = '+'
	LineStatusError    LineStatus = 'E'
)
