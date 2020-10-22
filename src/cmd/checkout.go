package cmd

import (
	"github.com/kevin-hanselman/dud/src/cache"
	"github.com/kevin-hanselman/dud/src/index"
	"github.com/kevin-hanselman/dud/src/strategy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().BoolVarP(
		&useCopyStrategy,
		"copy",
		"c",
		false,
		"copy artifacts instead of linking",
	)
	checkoutCmd.Flags().BoolVarP(
		&checkoutSingleStage,
		"single-stage",
		"s",
		false,
		"don't recursively operate on dependencies",
	)
}

var useCopyStrategy, checkoutSingleStage bool

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkout all artifacts from the cache",
	Long:  "checkout all artifacts from the cache",
	Run: func(cmd *cobra.Command, args []string) {

		strat := strategy.LinkStrategy
		if useCopyStrategy {
			strat = strategy.CopyStrategy
		}

		ch, err := cache.NewLocalCache(viper.GetString("cache"))
		if err != nil {
			logger.Fatal(err)
		}

		// TODO: forcing a checkout will require a "force load lock"
		// flag in index.FromFile
		idx, err := index.FromFile(".dud/index")
		if err != nil {
			logger.Fatal(err)
		}

		if len(args) == 0 {
			// Ignore checkoutSingleStage flag when no args passed.
			checkoutSingleStage = false
			for path := range idx {
				args = append(args, path)
			}
		}

		checkedOut := make(map[string]bool)
		for _, path := range args {
			inProgress := make(map[string]bool)
			if err := idx.Checkout(
				path,
				ch,
				strat,
				!checkoutSingleStage,
				checkedOut,
				inProgress,
				logger,
			); err != nil {
				logger.Fatal(err)
			}
		}
	},
}