// package main is a module with raspberry pi board component.
package main

import (
	"context"

	"github.com/viam-modules/beagleboard/beaglebone"
	beagleYAi "github.com/viam-modules/beagleboard/beagleyai"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("beagleboard"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	module, err := module.NewModuleFromArgs(ctx)
	if err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, board.API, beaglebone.Model); err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, board.API, beagleYAi.Model); err != nil {
		return err
	}

	err = module.Start(ctx)
	defer module.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
