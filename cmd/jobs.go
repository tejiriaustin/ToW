package cmd

import "github.com/spf13/cobra"

var jobsCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts Tree of Wally API",
	Long:  ``,
	Run:   startJobs,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startJobs(cmd *cobra.Command, args []string) {

}
