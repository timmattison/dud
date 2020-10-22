package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a Dud project",
	Long: `Initialize a Dud project by populating
a .dud directory in the working directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		cacheDir := ".dud/cache"
		if err := os.MkdirAll(cacheDir, 0755); err != nil {
			logger.Fatal(err)
		}
		viper.Set("cache", cacheDir)
		// WriteConfig() doesn't work if the file doesn't exist.
		if err := viper.WriteConfigAs(".dud/config.yaml"); err != nil {
			logger.Fatal(err)
		}
		logger.Println("Initialized .dud directory")
	},
	// Override rootCmd's PersistentPreRun which changes dir to the project
	// root. Obviously this command would fail if we're initializing said
	// directory.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
}