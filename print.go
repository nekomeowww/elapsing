package elapsing

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/nekomeowww/elapsing/internal/utils"
	"github.com/samber/lo"
)

func alterRenderedResultWithColor(renderResult string, style *list.Style, color color.Color) string {
	renderResult = strings.ReplaceAll(renderResult, style.CharItemSingle, color.Render(style.CharItemSingle))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemTop, color.Render(style.CharItemTop))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemMiddle, color.Render(style.CharItemMiddle))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemVertical, color.Render(style.CharItemVertical))
	renderResult = strings.ReplaceAll(renderResult, style.CharItemBottom, color.Render(style.CharItemBottom))
	return renderResult
}

const (
	defaultSingleLineFormat = "%%s %%s %s%%v%s"
)

var (
	// format of [%s ( %s total)]
	singleLineFormat = fmt.Sprintf(defaultSingleLineFormat,
		color.FgGray.Render("["),
		color.FgGray.Render("]"),
	)
)

func (e *Elapsing) appendStatsDataToList(list list.Writer, parentIndex int) {
	var name string

	switch e.elapsingType {
	case elapsingTypeBase:
		name = color.FgCyan.Render(lo.Ternary(e.name == "", defaultUnknownName, e.name))
	case elapsingTypeFunc:
		name = fmt.Sprintf("%s %s",
			color.FgGray.Render(fmt.Sprintf("#%d", parentIndex+1)),
			color.FgCyan.Render(lo.Ternary(e.name == "", defaultUnknownFunctionName, e.name)),
		)
	}

	list.AppendItem(name)
	list.Indent()

	indexes, indexesMaxLength := e.steps.Indexes()
	names, namesMaxLength := e.steps.Names()
	lasts, lastsMaxLength := e.steps.Lasts()
	// totals, totalsMaxLength := e.steps.Totals()

	for i := range indexes {
		if e.steps[i].Type() == StepTypeElapsing {
			e.steps[i].(*Elapsing).appendStatsDataToList(list, i)
			continue
		}

		list.AppendItem(fmt.Sprintf(singleLineFormat,
			color.FgGray.Render(fmt.Sprintf("#%s", utils.StringPadStart(indexes[i], indexesMaxLength))),
			color.FgYellow.Render(utils.StringPadEnd(names[i], namesMaxLength)),
			color.FgGreen.Render(utils.StringPadStart(lasts[i], lastsMaxLength)),
		))
	}

	list.UnIndent()
}

func (e *Elapsing) Stats() string {
	e.stepsLock.Lock()
	defer e.stepsLock.Unlock()

	newList := list.NewWriter()
	newList.SetStyle(list.StyleConnectedLight)
	e.appendStatsDataToList(newList, 0)

	renderResult := newList.Render()
	return alterRenderedResultWithColor(renderResult, newList.Style(), color.FgGray)
}
