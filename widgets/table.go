/**
 * Created by Goland
 * @file   grid.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/30 17:18
 * @desc   grid.go
 */

package widgets

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
	"gioui.org/x/outlay"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/utils"
	"golang.org/x/exp/maps"
	"image"
)

type Table struct {
	theme       *theme.Theme
	height      unit.Dp
	grid        component.GridState
	headerFun   layout.ListElement
	dataFun     outlay.Cell
	headers     []string
	data        []map[string]any
	dataContent []widget.Bool
	keys        []string
}

func NewTable(th *theme.Theme) *Table {
	table := &Table{
		theme:  th,
		height: unit.Dp(30),
	}
	return table
}

func (t *Table) SetHeader(header []string) *Table {
	t.headers = header
	return t
}
func (t *Table) SetData(data []map[string]any) *Table {
	t.data = data
	for range data {
		t.dataContent = append(t.dataContent, widget.Bool{})
	}
	t.keys = maps.Keys(t.data[0])
	return t
}

func (t *Table) SetHeaderFun(headerFun layout.ListElement) *Table {
	t.headerFun = headerFun
	return t
}
func (t *Table) SetDataFun(dataFun outlay.Cell) *Table {
	t.dataFun = dataFun
	return t
}

func (t *Table) LayoutTable(gtx layout.Context) layout.Dimensions {
	if len(t.data) == 0 {
		return layout.Dimensions{}
	}
	inset := layout.UniformInset(unit.Dp(2))
	orig := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	macro := op.Record(gtx.Ops)
	dims := inset.Layout(gtx, layout.Spacer{Height: t.height}.Layout)
	_ = macro.Stop()
	gtx.Constraints = orig
	if t.headerFun == nil {
		t.headerFun = func(gtx layout.Context, index int) layout.Dimensions {
			utils.DrawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.TableHeaderBgColor)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return Label(t.theme, t.headers[index], true).Layout(gtx)
			})
		}
	}
	if t.dataFun == nil {
		t.dataFun = func(gtx layout.Context, row, col int) layout.Dimensions {
			return t.dataContent[row].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if t.dataContent[row].Hovered() {
					utils.DrawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.DefaultContentBgGrayColor)
				} else {
					utils.DrawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.DefaultWindowBgGrayColor)
				}
				NewLine(t.theme).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
				labelDims := layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return Label(t.theme, fmt.Sprint(t.data[row][t.keys[col]])).Layout(gtx)
				})
				return labelDims
			})
		}
	}
	return component.Table(t.theme.Material(), &t.grid).Layout(gtx, len(t.data), len(t.data[0]),
		func(axis layout.Axis, index, constraint int) int {
			switch axis {
			case layout.Horizontal:
				return constraint / len(t.headers)
			default:
				return dims.Size.Y
			}
		},
		t.headerFun,
		t.dataFun,
	)
}
