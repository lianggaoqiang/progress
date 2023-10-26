package progress

type TextBar struct {
	Setting       TextBarSetting
	isInit        bool
	staticContent string // the content of TextBar
}

// TextBarSetting describes the setting of TextBar
type TextBarSetting struct {
	Inline bool
	Hidden bool
}

// enable TextBar implement interface Bar
func (b *TextBar) kind() string {
	return "text"
}
func (b *TextBar) Hide() {
	b.Setting.Hidden = true
}
func (b *TextBar) Show() {
	b.Setting.Hidden = false
}
func (b *TextBar) IsHidden() bool {
	return b.Setting.Hidden
}

// NewTextBar returns an instance pointer of a static content bar
func NewTextBar(content string) *TextBar {
	return &TextBar{
		isInit:        true,
		staticContent: content,
	}
}

// Custom changes the setting of TextBar
func (b *TextBar) Custom(setting TextBarSetting) *TextBar {
	b.Setting = setting
	return b
}
