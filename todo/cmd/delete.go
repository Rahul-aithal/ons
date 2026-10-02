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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run:  deleteItem,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteItem(cmd *cobra.Command, args []string) {

	id := args[0]
	u, uerr := user.Current()
	if uerr != nil {
		log.Fatalf("Unable to determine current user: %v", uerr)
	}
	filePath := service.GetEnv("DB_FILE", fmt.Sprintf("%s/data.csv", u.HomeDir))
	todo, err := delteFromCSV(filePath, id)

	if err != nil {
		log.Fatalf("Unable to delete todo %q from data file %q: %v", id, filePath, err)
	}

	log.Printf("Deleted todo %q (ID: %s)", todo.Name, todo.Id)
}

func delteFromCSV(filePath, id string) (types.Todo, error) {
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
	deletedLine := make([]string, 6)
	for _, line := range data {
		if line[0] != id {
			w.Write(line)
		} else {
			deletedLine = line
		}
	}

	defer w.Flush()
	return types.Todo{
		Id:          deletedLine[0],
		Name:        deletedLine[1],
		Status:      deletedLine[2],
		Due:         deletedLine[3],
		Description: deletedLine[4],
		Priority:    deletedLine[5],
	}, nil
}
