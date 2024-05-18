package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "nyx",
	Short: "Increase reliability of deployments by versioning your database",
	Long: `Nyx is based on the Flyway concepts and inspired to be fully compatible with it,
being the Golang implementation of this Java utility. It is an open-source
database migration tool. It strongly favors simplicity and convention over
configuration.

It is based on 7 basic commands: migrate, clean, info, validate, undo, baseline,
and repair.

Currently, its main goal is to serve as a command-line tool for pipelines,
as it simplifies the needed dependencies and utility size. In addition, it only
supports SQL-based migrations at this moment, since Java-based migrations are
unavailable due to language restrictions.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ${HOME}/nyx.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("nyx")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "config file:", viper.ConfigFileUsed())
	}
}
