package elapsing

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/nekomeowww/elapsing/pkg/utils"
	"github.com/samber/lo"
)

func alterRenderResultWithColor(renderResult string, style *list.Style, color color.Color) string {
	renderResult = strings.ReplaceAll(renderResult, style.CharItemSingle, color.Render(style.CharItemSingle))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemTop, color.Render(style.CharItemTop))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemMiddle, color.Render(style.CharItemMiddle))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemVertical, color.Render(style.CharItemVertical))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemBottom, color.Render(style.CharItemBottom))
	return renderResult
}

var (
	// format of [%s ( %s total)]
	singleLineFormat = fmt.Sprintf("%%s %%s %s%%v %s%%v %s%s%s",
		color.FgGray.Render("["),
		color.FgGray.Render("("),
		color.FgGray.Render("total"),
		color.FgGray.Render(")"),
		color.FgGray.Render("]"),
	)
)

func (e *Elapsing) appendStatsDataToList(list list.Writer, parentIndex int) {
	var name string

	switch e.ElapsingType {
	case ElapsingTypeBase:
		name = color.FgCyan.Render(lo.Ternary(e.Name == "", "(unknown name)", e.Name))
	case ElapsingTypeFunc:
		name = fmt.Sprintf("%s %s",
			color.FgGray.Render(fmt.Sprintf("#%d", parentIndex+1)),
			color.FgCyan.Render(lo.Ternary(e.Name == "", "(unknown function name)", e.Name)),
		)
	}

	list.AppendItem(name)
	list.Indent()

	indexes, indexesMaxLength := e.Steps.Indexes()
	names, namesMaxLength := e.Steps.Names()
	lasts, lastsMaxLength := e.Steps.Lasts()
	totals, totalsMaxLength := e.Steps.Totals()

	lo.ForEach(indexes, func(_ string, i int) {
		if e.Steps[i].Type() == StepTypeElapsing {
			e.Steps[i].(*Elapsing).appendStatsDataToList(list, i)
			return
		}

		list.AppendItem(fmt.Sprintf(singleLineFormat,
			color.FgGray.Render(fmt.Sprintf("#%s", utils.StringPadStart(indexes[i], indexesMaxLength))),
			color.FgYellow.Render(utils.StringPadEnd(names[i], namesMaxLength)),
			color.FgGreen.Render(utils.StringPadStart(lasts[i], lastsMaxLength)),
			color.FgGray.Render(utils.StringPadStart(totals[i], totalsMaxLength)),
		))
	})

	list.UnIndent()
}

func (e *Elapsing) Stats() string {
	e.stepsLock.Lock()
	defer e.stepsLock.Unlock()

	newList := list.NewWriter()
	newList.SetStyle(list.StyleConnectedLight)
	e.appendStatsDataToList(newList, 0)

	renderResult := newList.Render()
	return alterRenderResultWithColor(renderResult, newList.Style(), color.FgGray)
}
