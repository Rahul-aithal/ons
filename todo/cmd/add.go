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
	"time"
	"todo/service"
	"todo/types"

	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [dd-mm-yyy]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(2),
	Run:  addItem,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addItem(cmd *cobra.Command, args []string) {
	todo := &types.Todo{
		Id:          uuid.New().String(),
		Name:        args[0],
		Status:      "Pending",
		Due:         args[1],
		Description: "",
		Priority:    "low",
	}
	if todo.Status == "Pending" {
		timeFormat := "02-01-2006"
		dueTime, derr := time.Parse(timeFormat, todo.Due)

		if derr != nil {
			log.Fatalf("Due date parsing error %s", derr)
		}
		timeLeft := time.Until(dueTime)
		if timeLeft <= 0 {
			todo.Status = "Late"
		}
	}
	u, uerr := user.Current()
	if uerr != nil {
		log.Fatalf("User not found %s", uerr)
	}
	filePath := service.GetEnv("DB_FILE", fmt.Sprintf("%s/data.csv", u.HomeDir))
	err := appendToCSV(filePath, []string{todo.Id, todo.Name, todo.Status, todo.Due, todo.Description, todo.Priority})

	if err != nil {
		log.Fatalf("Append Failed %s", err)
	}

	log.Printf("Added new todo %s\n", todo.Name)
}

func appendToCSV(filePath string, row []string) error {
	file, ferr := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if ferr != nil {
		return ferr
	}
	defer file.Close()
	w := csv.NewWriter(file)

	if werr := w.Write(row); werr != nil {
		return werr
	}
	defer w.Flush()
	return nil
}
