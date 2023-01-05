package elapsing

import (
	"fmt"
	"time"

	"github.com/nekomeowww/elapsing/pkg/utils"
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

type Steps []Step

func (r Steps) Indexes() ([]string, int) {
	indexes := lo.Map(r, func(item Step, i int) string { return fmt.Sprintf("%d", i+1) })
	indexesMaxLength := utils.StringsMaxLength(indexes)

	return indexes, indexesMaxLength
}

func (r Steps) Names() ([]string, int) {
	pointNames := lo.Map(r, func(item Step, _ int) string {
		if point, ok := item.(Point); ok {
			return point.Name
		} else if point, ok := item.(*Point); ok {
			return point.Name
		} else {
			return ""
		}
	})
	pointNamesMaxLength := utils.StringsMaxLength(pointNames)

	return pointNames, pointNamesMaxLength
}

func (r Steps) Lasts() ([]string, int) {
	elapsedLasts := lo.Map(r, func(item Step, _ int) string {
		if point, ok := item.(Point); ok {
			return fmt.Sprintf("%v", point.SinceLast)
		} else if point, ok := item.(*Point); ok {
			return fmt.Sprintf("%v", point.SinceLast)
		} else {
			return "0s"
		}
	})
	elapsedLastsMaxLength := utils.StringsMaxLength(elapsedLasts)

	return elapsedLasts, elapsedLastsMaxLength
}

func (r Steps) Totals() ([]string, int) {
	elapsedTotals := lo.Map(r, func(item Step, _ int) string {
		if point, ok := item.(Point); ok {
			return fmt.Sprintf("%v", point.SinceInitial)
		} else if point, ok := item.(*Point); ok {
			return fmt.Sprintf("%v", point.SinceInitial)
		} else {
			return "0s"
		}
	})
	elapsedTotalsMaxLength := utils.StringsMaxLength(elapsedTotals)

	return elapsedTotals, elapsedTotalsMaxLength
}
