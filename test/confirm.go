package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/widgets"
	"github.com/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable
	var th = theme.NewTheme()
	confirm := widgets.NewConfirm(th)
	confirm.Confirm(func() {
		fmt.Println("确定...")
	})
	confirm.Cancel(func() {
		fmt.Println("取消...")
	})

	win := window.NewInitialize()
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		if clickable.Clicked(gtx) {
			confirm.Message("确定退出吗?")
		}
		layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(400))
						return widgets.Label(th, "&clickable, nil, 0,  unit.Dp(100)").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.DefaultButton(th, &clickable, "click me", unit.Dp(100)).Layout(gtx)
					}),
				)
			}),
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if confirm.Visible() {
					return confirm.Layout(gtx)
				}
				return layout.Dimensions{}
			}),
		)
	})
	win.Run()
}
