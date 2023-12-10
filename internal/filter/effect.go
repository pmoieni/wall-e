package filter

import (
	"context"
	"errors"

	"github.com/disintegration/gift"
	"github.com/urfave/cli/v2"
)

func grayscale() cli.ActionFunc {
	name := "grayscale"

	return func(cCtx *cli.Context) error {
		filters, ok := cCtx.Context.Value(CtxName).(FilterCtx)
		if !ok {
			return errors.New("filters context not defined")
		}
		filters.List[name] = gift.Grayscale()
		cCtx.Context = context.WithValue(cCtx.Context, CtxName, filters)
		return nil
	}
}
