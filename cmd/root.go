package cmd

import (
  "fmt"
  "os"
	log "github.com/Sirupsen/logrus"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var ConfigFile string
var Verbose bool
var UseViper bool
var JobServiceURI string
var Environment string
var Logger = log.Logger{}

func init() {
  cobra.OnInitialize(initConfig)

  RootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "c", "", "config file (default is config.yaml)")
	RootCmd.PersistentFlags().StringVarP(&Environment, "environment", "e", "development", "Environment variable (default is development)")
  RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output (default is false)")
  RootCmd.PersistentFlags().BoolVarP(&UseViper, "viper", "r", true, "Use Viper for configuration (default is true)")

	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "severity",
			log.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})
  log.SetOutput(os.Stdout)
	if Environment == "production" {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Infof("Successfully initialized job service")
}

func initConfig() {
  if ConfigFile != "" {
    viper.SetConfigFile(ConfigFile)
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

var RootCmd = &cobra.Command{
  Use: "jobs",
  Short: "jobs service application",
  Long: "jobs service is responsible for CRUD with job entity",
  Run: func(cmd *cobra.Command, args[]string) {
    cmd.Usage()
  },
}

func Execute() {
  if err := RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
