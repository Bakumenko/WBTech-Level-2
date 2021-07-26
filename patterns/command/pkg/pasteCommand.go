package pkg

type PasteCommand struct {
	Editor *Editor
}

func (p *PasteCommand) execute() {
	p.Editor.Text += p.Editor.memory
}
