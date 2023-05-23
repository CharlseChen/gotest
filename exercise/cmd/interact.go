package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "My awesome command line tool",
		Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from mycli!")
		},
	}

	// Define flags and configuration options using viper
	viper.SetConfigName(".myclirc") // Name of config file (without extension)
	viper.AddConfigPath("$HOME")    // Search config in home directory
	viper.AutomaticEnv()            // Read in environment variables that match

	// Define a flag using cobra
	rootCmd.Flags().StringP("config", "c", "", "config file (default is $HOME/.myclirc)")

	// Bind the flag to viper
	viper.BindPFlag("config", rootCmd.Flags().Lookup("config"))

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
