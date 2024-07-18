package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa la aplicación",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hola %s\n", name)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&name, "name", "n", "", "Nombre para saludar")
	initCmd.MarkFlagRequired("name")
}
