package main

import (
	"fmt"
	"github.com/joeygibson/ghcli/pkg/commands"
	"github.com/joeygibson/ghcli/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ghcli",
		Short: "Command-line tool to get stats about Github organizations",
		Long:  "Command-line tool to get stats about Github organizations",
		Run:   CmdRoot,
	}

	starsCmd = &cobra.Command{
		Use:   "stars",
		Short: "Display info about the top-n repos based on number of stars",
		Long:  "Display info about the top-n repos based on number of stars",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: CmdStars,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ghcli",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s (build: %s)\n", version, build)
		},
	}

	version string
	build   string
)

func CmdRoot(cmd *cobra.Command, _ []string) {
	cmd.Help()
}

func CmdStars(_ *cobra.Command, _ []string) {
	commands.Stars(config.GetConfig())
}

func init() {
	viper.SetEnvPrefix("GH_CLI")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigName(".ghcli")
	viper.SetConfigType("yml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	viper.ReadInConfig()

	defaultTokenFile := config.GetDefaultTokenFileName()

	rootCmd.PersistentFlags().Bool("verbose", false, "verbose output")
	rootCmd.PersistentFlags().String("org", "", "organization to use")
	rootCmd.PersistentFlags().String("user", "", "username for authorization")
	rootCmd.PersistentFlags().String("password", "", "password for authorization")
	rootCmd.PersistentFlags().String("token-file", defaultTokenFile, "file containing OAuth token")
	rootCmd.PersistentFlags().Int("top", 10, "number of results to return")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("org", rootCmd.PersistentFlags().Lookup("org"))
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("token.file", rootCmd.PersistentFlags().Lookup("token-file"))
	viper.BindPFlag("top", rootCmd.PersistentFlags().Lookup("top"))

	rootCmd.AddCommand(starsCmd)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
