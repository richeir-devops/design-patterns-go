package memento

///////////////////////////////////////////
// 备忘录模式（Memento）：
// 实际上就是保存对象的当前形态，这样以后任何时候都可以恢复
// 一般情况下用的应该是用 JSON 来保存，这样一般比较纯粹
///////////////////////////////////////////

///////////////////////////////////////////

type EditorMemento struct {
	content string
}

///////////////////////////////////////////

///////////////////////////////////////////

type Editor struct {
	content string
}

func (e *Editor) write(words string) {
	e.content = e.content + words
}

func (e *Editor) save() *EditorMemento {
	return &EditorMemento{content: e.content}
}

func (e *Editor) resotre(memento *EditorMemento) {
	e.content = memento.content
}

///////////////////////////////////////////
