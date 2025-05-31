//go:build linux

// Package beaglebone implements a BeagleY-AI based board.
package beagley

import (
	"context"

	"github.com/pkg/errors"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/board/genericlinux"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const modelName = "beagley"

// Model for viam supported beagleboard beaglebone board.
var Model = resource.NewModel("viam", "beagleboard", "beagle-y-ai")

func init() {
	gpioMappings, err := genericlinux.GetGPIOBoardMappings(modelName, boardInfoMappings)
	var noBoardErr genericlinux.NoBoardFoundError
	if errors.As(err, &noBoardErr) {
		logging.Global().Debugw("error getting beagley GPIO board mapping", "error", err)
	}

	resource.RegisterComponent(
		board.API,
		Model,
		resource.Registration[board.Board, *genericlinux.Config]{
			Constructor: func(
				ctx context.Context,
				_ resource.Dependencies,
				conf resource.Config,
				logger logging.Logger,
			) (board.Board, error) {
				return genericlinux.NewBoard(ctx, conf, genericlinux.ConstPinDefs(gpioMappings), logger)
			},
		})
}
