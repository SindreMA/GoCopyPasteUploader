package Types

type ClipboardType string

const (
	Text  ClipboardType = "text"
	Image ClipboardType = "image"
	File  ClipboardType = "file"
)

type IClipboard struct {
	Type ClipboardType
	Data string // text, base64
	Ext  string // file extension
}
