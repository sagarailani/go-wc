/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type FunctionalFlags struct {
	AllEnabled bool
	LineBreaks bool
	WordCounts bool
	CharCounts bool
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-wc",
	Short: "Short description",
	Long:  `Longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {

		flags := FunctionalFlags{}
		var err error
		flags.LineBreaks, err = cmd.Flags().GetBool("line-breaks")
		if err != nil {
			log.Fatal("Error while retrieving flag from command")
			return err
		}

		flags.WordCounts, err = cmd.Flags().GetBool("word-count")
		if err != nil {
			log.Fatal("Error while retrieving flag from command")
			return err
		}

		flags.CharCounts, err = cmd.Flags().GetBool("char-count")
		if err != nil {
			log.Fatal("Error while retrieving flag from command")
			return err
		}
		flags.AllEnabled = !flags.LineBreaks && !flags.WordCounts && !flags.CharCounts
		var totalLines, totalWords, totalChars int
		for _, fileName := range args {
			if !fileValidations(fileName) {
				continue
			}

			if flags.AllEnabled {
				lineCount, charCount := countLinesAndChars(fileName)
				wordCount := countWords(fileName)
				totalLines += lineCount
				totalWords += wordCount
				totalChars += charCount
				fmt.Printf("%8d %8d %8d %v\n", lineCount, wordCount, charCount, fileName)
			} else {
				var output string
				if flags.LineBreaks {
					lineCount, _ := countLinesAndChars(fileName)
					totalLines += lineCount
					output = fmt.Sprintf("%8d ", lineCount)
				}
				if flags.WordCounts {
					wordCount := countWords(fileName)
					totalWords += wordCount
					output += fmt.Sprintf("%8d ", wordCount)
				}
				if flags.CharCounts {
					_, charCount := countLinesAndChars(fileName)
					totalChars += charCount
					output += fmt.Sprintf("%8d ", charCount)
				}
				output += fmt.Sprintf("%v", fileName)
				fmt.Println(output)
			}
		}
		if len(args) > 1 {
			fmt.Printf("%8d %8d %8d total\n", totalLines, totalWords, totalChars)
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("line-breaks", "l", false, "show line count for files")
	rootCmd.Flags().BoolP("word-count", "w", false, "show word count for files")
	rootCmd.Flags().BoolP("char-count", "c", false, "show char count for files")
}

func countLinesAndChars(fileName string) (int, int) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineCount, charCount int
	for scanner.Scan() {
		data := scanner.Bytes()
		charCount += len(data) + 1
		lineCount += 1
	}
	return lineCount, charCount
}

func countWords(fileName string) int {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var wordCount int
	for scanner.Scan() {
		wordCount += 1
	}
	return wordCount
}

func fileValidations(fileName string) bool {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("go-wc: %v: read: No such file or directory\n", fileName)
		return false
	}

	if fileInfo.IsDir() {
		fmt.Printf("go-wc: %v: read: Is a directory\n", fileName)
		return false
	}
	if fileInfo.Mode().Perm()&0444 != 0444 {
		fmt.Printf("go-wc: %v: open: Permission denied\n", fileName)
		return false
	}
	return true
}
