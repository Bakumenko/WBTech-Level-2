package pkg

type CopyCommand struct {
	Editor   *Editor
	CopyText string
}

func (c *CopyCommand) execute() {
	c.Editor.memory = c.CopyText
}
