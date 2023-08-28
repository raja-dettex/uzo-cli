/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"uzo/util"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:                   "code <zip_file_name>",
	Short:                 "used to open the vs code editor in the current folder path",
	Long:                  ``,
	Args:                  cobra.ExactArgs(1),
	Example:               fmt.Sprintf("uzo code filename.zip\nuzo code .\\folder\\filename.zip"),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var argument string
		var filename string
		argument = args[0]
		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}
		if fileExists {
			filename, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("file %v does not exist\n", argument)
			return
		}
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		util.Unzip(filename, wd)
		os.Chdir(util.FileNameWithoutExtension(filename))
		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		commandCode := exec.Command("code", wd)
		err = commandCode.Run()
		if err != nil {
			fmt.Printf("vs code executable file not found at %PATH%\n")
		}

	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
