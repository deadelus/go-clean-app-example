package cmd

import (
	"fmt"
	"live-semantic/src/domain/uc"
	"os"

	"github.com/deadelus/go-clean-app/src/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	useCases  uc.UseCases
	appLogger logger.Logger
	verbose   bool
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "live-semantic",
	Short: "🚀 Live Semantic CLI Application",
	Long: `Live Semantic - A modern CLI application with clean architecture.

Built with ❤️  using Go, Cobra, and Clean Architecture principles.
Supports CLI, Web API, and WebSocket modes.`,
	Version: "1.0.0",
}

// Execute executes the root command
func Execute(uc uc.UseCases, logger logger.Logger) {
	useCases = uc
	appLogger = logger

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialize the root command
func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.live-semantic.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Bind flags to viper
	if err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		fmt.Println("Error binding verbose flag:", err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".live-semantic")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
