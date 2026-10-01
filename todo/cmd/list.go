/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/user"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"todo/service"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Aliases: []string{"get"},
	Run:     listItem,
}

func init() {
	rootCmd.AddCommand(listCmd)

}

func listItem(cmd *cobra.Command, args []string) {

	w := tabwriter.NewWriter(os.Stdout, 1, 2, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tNAME\tSTATUS\tDUE")
	fmt.Fprintln(w, "--\t----\t------\t---")
	u, err := user.Current()
	if err != nil {
		panic("User not found")
	}
	dataFile := service.GetEnv("DB_FILE", fmt.Sprintf("%s/data.csv", u.HomeDir))

	file, ferr := os.Open(dataFile)

	if ferr != nil {
		newFile, crerr := os.Create(dataFile)

		if crerr != nil {

			log.Fatal("Unable to create "+dataFile, crerr)
		}
		file = newFile

	}

	defer file.Close()
	csvRedar := csv.NewReader(file)
	csvData, cerr := csvRedar.ReadAll()

	if cerr != nil {
		log.Fatal("Unable to parse file as CSV for "+dataFile, err)
	}

	for _, line := range csvData {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", line[0], line[1], line[2], line[3])

	}

	w.Flush()

}
