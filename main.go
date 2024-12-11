package main

import (
	"fmt"
	"time"

	"github.com/bendahl/uinput"
)

func matrix() {
	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("virtual-mouse"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mouse.Close()

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("virtual-keyboard"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer keyboard.Close()

	for {
		time.Sleep(4 * time.Second)
		keyboard.KeyPress(uinput.KeyLeftalt)
		time.Sleep(500 * time.Millisecond)
		keyboard.KeyPress(uinput.KeyRightalt)
	}
}

// alternatively (to use specific version), use this:
// import "gopkg.in/bendahl/uinput.v1"
func moveMouse() {
	// initialize mouse and check for possible errors
	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("testmouse"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// always do this after the initialization in order to guarantee that the device will be properly closed
	defer mouse.Close()

	// mouse pointer will be moved up by 10 pixels
	mouse.MoveUp(10)
	// mouse pointer will be moved to the right by 10 pixels
	mouse.MoveRight(10)
	// mouse pointer will be moved down by 10 pixels
	mouse.MoveDown(10)
	// mouse pointer will be moved to the left by 10 pixels (we're back to where we started)
	mouse.MoveLeft(10)
	// move the mouse pointer by 100 pixels on the x and y axes (right and down in this case)
	mouse.Move(100, 100)

	// click left
	mouse.LeftClick()
	// click right (depending on context a context menu may appear)
	mouse.RightClick()
	// click middle (usually the scroll wheel)go get gopkg.in/bendahl/uinput.v1
	mouse.MiddleClick()

	// hold down left mouse button
	mouse.LeftPress()
	// move mouse pointer down by 100 pixels while holding down the left key
	mouse.MoveDown(100)
	// release the left mouse button
	mouse.LeftRelease()

	// wheel up
	mouse.Wheel(false, 1)
	// wheel down
	mouse.Wheel(false, -1)
	// horizontal wheel left
	mouse.Wheel(true, 1)
	// horizontal wheel right
	mouse.Wheel(true, -1)
}

func main() {
	matrix()
}
