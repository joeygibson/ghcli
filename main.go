package main

import (
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
		Short: "Display repos from a Github organization",
		Long:  "Display repos from a Github organization, sorted by different criteria.",
		Run:   CmdRoot,
	}

	starsCmd = &cobra.Command{
		Use:   "stars",
		Short: "sort by the number of stars",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: CmdStars,
	}

	forksCmd = &cobra.Command{
		Use:   "forks",
		Short: "sort by number of forks",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: CmdForks,
	}

	loginCmd = &cobra.Command{
		Use:   "login <github OAuth token>",
		Short: "Set your Github OAuth token",
		Run:   CmdLogin,
	}

	pullRequestsCmd = &cobra.Command{
		Use:   "pull-requests",
		Short: "sort by number of pull requests",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: CmdPullRequests,
	}

	contributionsCmd = &cobra.Command{
		Use:   "contributions",
		Short: "sort by number of pull requests/fork",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("verbose") {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: CmdContributions,
	}
)

func CmdRoot(cmd *cobra.Command, _ []string) {
	cmd.Help()
}

func CmdStars(_ *cobra.Command, _ []string) {
	commands.Stars(config.GetConfig())
}

func CmdForks(_ *cobra.Command, _ []string) {
	commands.Forks(config.GetConfig())
}

func CmdLogin(_ *cobra.Command, args []string) {
	commands.Login(args)
}

func CmdPullRequests(_ *cobra.Command, _ []string) {
	commands.PullRequests(config.GetConfig())
}

func CmdContributions(_ *cobra.Command, _ []string) {
	commands.Contributions(config.GetConfig())
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

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().String("org", "", "organization to use")
	rootCmd.PersistentFlags().String("token", "", "Github OAuth token")
	rootCmd.PersistentFlags().Int("top", 10, "number of results to return")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("org", rootCmd.PersistentFlags().Lookup("org"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("top", rootCmd.PersistentFlags().Lookup("top"))

	rootCmd.AddCommand(starsCmd)
	rootCmd.AddCommand(forksCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(pullRequestsCmd)
	rootCmd.AddCommand(contributionsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
