package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
  Use: "jobs",
  Short: "jobs service application",
  Long: "jobs service is responsible for CRUD with job entity",
  Run: func(cmd *cobra.Command, args[]string) {
    cmd.Usage()
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

var configFile string

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
  rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
}

func initConfig() {
  if configFile != "" {
    viper.SetConfigFile(configFile)
  } else {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.AddConfigPath("/etc/job")
    viper.AddConfigPath("$HOME/.job")
  }

  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // Config file not found; ignore error if desired
        fmt.Println("Config file not found: %v\n", err)
    } else {
        // Config file was found but another error was produced
        fmt.Errorf("%s \n", err)
    }
    os.Exit(1)
  }
}
