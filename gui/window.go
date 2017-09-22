package gui

import (
	"github.com/sciter-sdk/go-sciter/window"
	"github.com/sciter-sdk/go-sciter"
	"fmt"
	"github.com/lxn/win"
)

const (
	EW_CHILD        = sciter.SW_CHILD        // child window only, if this flag is set all other flags ignored
	EW_TITLEBAR     = sciter.SW_TITLEBAR     // toplevel window, has titlebar
	EW_RESIZEABLE   = sciter.SW_RESIZEABLE   // has resizeable frame
	EW_TOOL         = sciter.SW_TOOL         // is tool window
	EW_CONTROLS     = sciter.SW_CONTROLS     // has minimize / maximize buttons
	EW_GLASSY       = sciter.SW_GLASSY       // glassy window ( DwmExtendFrameIntoClientArea on windows )
	EW_ALPHA        = sciter.SW_ALPHA        // transparent window ( e.g. WS_EX_LAYERED on Windows )
	EW_MAIN         = sciter.SW_MAIN         // main window of the app, will terminate the app on close
	EW_POPUP        = sciter.SW_POPUP        // the window is created as topmost window.
	EW_ENABLE_DEBUG = sciter.SW_ENABLE_DEBUG // make this window inspector ready
	EW_OWNS_VM      = sciter.SW_OWNS_VM      // it has its own script VM
)

type Window struct {
	width  int
	height int
	x      int
	y      int
	flag   sciter.WindowCreationFlag

	file  string
	title string

	callback       *sciter.CallbackHandler
	windowHandler  func(window *window.Window)
	elementHandler func(root *sciter.Element)
	window         *window.Window
}

func NewLoginDialod(flag sciter.WindowCreationFlag, width, height int, title string) *Window {
	cxScreen := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	cyScreen := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	x := (cxScreen - width) / 2
	y := (cyScreen - height) / 2

	wnd := new(Window)
	wnd.x = x
	wnd.y = y
	wnd.width = width
	wnd.height = height
	wnd.flag = flag

	wnd.title = title
	return wnd
}

func (d *Window) SetWidth(width int) {
	d.width = width
}
func (d *Window) SetHeight(height int) {
	d.height = height
}
func (d *Window) SetX(x int) {
	d.x = x
}
func (d *Window) SetY(y int) {
	d.y = y
}
func (d *Window) SetFile(file string) {
	d.file = file
}
func (d *Window) SetTitle(title string) {
	d.title = title
}
func (d *Window) SetCallBackHandlers(handler *sciter.CallbackHandler) {
	d.callback = handler
}
func (d *Window) SetWindowHandlers(handler func(window *window.Window)) {
	d.windowHandler = handler
}
func (d *Window) SetElementHandlers(handler func(root *sciter.Element)) {
	d.elementHandler = handler
}

func (d *Window) Width() int {
	return d.width
}
func (d *Window) Height() int {
	return d.height
}
func (d *Window) X() int {
	return d.x
}
func (d *Window) Y() int {
	return d.y
}
func (d *Window) File() string {
	return d.file
}
func (d *Window) Title() string {
	return d.title
}

func (d *Window) Show() {

	wnd, err := window.New(d.flag, sciter.NewRect(d.y, d.x, d.width, d.height));
	if err != nil {
		fmt.Println("gui.Window.NewLoginDialog.window.New:", err.Error())
	}
	d.window = wnd

	d.window.LoadFile(d.file)
	d.window.SetTitle(d.title)

	if d.callback != nil {
		d.window.SetCallback(d.callback)
	}

	if d.windowHandler != nil {
		d.windowHandler(d.window)
	}

	root, err := d.window.GetRootElement()
	if err != nil {
		fmt.Println("gui.Window.Show.d.window.GetRootElement:", err.Error())
	}
	if d.elementHandler != nil {
		d.elementHandler(root)
	}

	d.window.Show()
}

func (d *Window) Run() {
	d.window.Run()
}
