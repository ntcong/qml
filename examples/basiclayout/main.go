// Basic Layout Example
// http://qt-project.org/doc/qt-5.1/qtquickcontrols/basiclayouts.html

package main

import (
	"fmt"
	"github.com/niemeyer/qml"
	"os"
)

func main() {
	qml.Init(nil)
	engine := qml.NewEngine()
	engine.On("quit", func() { os.Exit(0) })
	component, err := engine.LoadFile("basiclayout.qml")
	if err != nil {
		fmt.Println(err)
		return
	}

	window := component.CreateWindow(nil)
	window.Show()
	window.Wait()
}
