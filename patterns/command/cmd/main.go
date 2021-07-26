package main

import (
	"command/pkg"
	"fmt"
)

func main() {

	editor := &pkg.Editor{Text: "Hello"}

	copyCom := &pkg.CopyCommand{Editor: editor, CopyText: ", Oleg"}
	pasteCom := &pkg.PasteCommand{Editor: editor}
	buttonForCopy := pkg.Button{Command: copyCom}
	buttonForPaste := pkg.Button{Command: pasteCom}
	buttonForCopy.Press()
	buttonForPaste.Press()

	fmt.Println(editor.Text)
}
