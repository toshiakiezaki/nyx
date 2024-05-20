package extractors

import (
	"github.com/eidolons/nyx/models"
	"github.com/spf13/cobra"
)

func InfoSettings(cmd cobra.Command, args []string) models.InfoSettings {
	return models.NewInfoSettings(args)
}
