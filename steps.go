package elapsing

import (
	"fmt"
	"time"

	"github.com/nekomeowww/elapsing/internal/utils"
	"github.com/samber/lo"
)

type StepType int

const (
	_ StepType = iota
	StepTypePoint
	StepTypeElapsing
)

type Step interface {
	Type() StepType
	On() time.Time
}

type steps []Step

func (r steps) Indexes() ([]string, int) {
	indexes := lo.Map(r, func(item Step, i int) string { return fmt.Sprintf("%d", i+1) })
	indexesMaxLength := utils.StringsMaxLength(indexes)

	return indexes, indexesMaxLength
}

func (r steps) Names() ([]string, int) {
	pointNames := lo.Map(r, func(item Step, _ int) string {
		if p, ok := item.(point); ok {
			return p.name
		} else if p, ok := item.(*point); ok {
			return p.name
		} else {
			return ""
		}
	})
	pointNamesMaxLength := utils.StringsMaxLength(pointNames)

	return pointNames, pointNamesMaxLength
}

func (r steps) Lasts() ([]string, int) {
	elapsedLasts := lo.Map(r, func(item Step, _ int) string {
		if p, ok := item.(point); ok {
			return fmt.Sprintf("%v", p.sinceLast)
		} else if p, ok := item.(*point); ok {
			return fmt.Sprintf("%v", p.sinceLast)
		} else {
			return "0s"
		}
	})
	elapsedLastsMaxLength := utils.StringsMaxLength(elapsedLasts)

	return elapsedLasts, elapsedLastsMaxLength
}

func (r steps) Totals() ([]string, int) {
	elapsedTotals := lo.Map(r, func(item Step, _ int) string {
		if p, ok := item.(point); ok {
			return fmt.Sprintf("%v", p.sinceInitial)
		} else if p, ok := item.(*point); ok {
			return fmt.Sprintf("%v", p.sinceInitial)
		} else {
			return "0s"
		}
	})
	elapsedTotalsMaxLength := utils.StringsMaxLength(elapsedTotals)

	return elapsedTotals, elapsedTotalsMaxLength
}
