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
	"todo/service"
	"todo/types"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completedCmd = &cobra.Command{
	Use:   "completed [id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run:  completeItem,
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completeItem(cmd *cobra.Command, args []string) {

	id := args[0]
	u, uerr := user.Current()
	if uerr != nil {
		log.Fatalf("Unable to determine current user: %v", uerr)
	}
	filePath := service.GetEnv("DB_FILE", fmt.Sprintf("%s/data.csv", u.HomeDir))
	todo, err := updateCsv(filePath, id)

	if err != nil {
		log.Fatalf("Unable to update todo %q in data file %q: %v", id, filePath, err)
	}

	log.Printf("Updated todo %q (ID: %s)", todo.Name, todo.Id)
}

func updateCsv(filePath, id string) (types.Todo, error) {
	file, ferr := os.Open(filePath)

	if ferr != nil {

		return types.Todo{}, ferr
	}

	r := csv.NewReader(file)

	data, readErr := r.ReadAll()

	if readErr != nil {
		return types.Todo{}, readErr
	}

	file.Close()
	newFile, newErr := os.Create(filePath)

	if newErr != nil {
		return types.Todo{}, newErr
	}
	defer newFile.Close()

	w := csv.NewWriter(newFile)
	updatedLine := make([]string, 6)
	for _, line := range data {
		if line[0] != id {
			line[2] = "Completed"
			updatedLine = line
		}
		writeErr := w.Write(line)

		if writeErr != nil {
			return types.Todo{}, writeErr
		}

	}

	defer w.Flush()
	return types.Todo{
		Id:          updatedLine[0],
		Name:        updatedLine[1],
		Status:      updatedLine[2],
		Due:         updatedLine[3],
		Description: updatedLine[4],
		Priority:    updatedLine[5],
	}, nil
}
