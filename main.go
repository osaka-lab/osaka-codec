package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var output string
var isString bool

func main() {
	var rootCmd = &cobra.Command{
		Use:   "osaka-codec",
		Short: "Encode your data with osaka-codec",
	}

	var encodeCmd = &cobra.Command{
		Use:   "encode",
		Short: "Encode data with osaka-codec",
		Args:  cobra.MinimumNArgs(1),
		RunE:  encode,
	}

	var decodeCmd = &cobra.Command{
		Use:   "decode",
		Short: "Decode data with osaka-codec",
		Args:  cobra.MinimumNArgs(1),
		RunE:  decode,
	}

	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Specify an output file to save the result. (Only on file encoding/decoding)")
	rootCmd.PersistentFlags().BoolVarP(&isString, "string", "s", false, "Specify that the input arguments are a string rather than a file path.")
	rootCmd.AddCommand(encodeCmd, decodeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
	}
}

func encode(cmd *cobra.Command, args []string) error {
	if len(args) > 1 || isString {
		return encodeString(cmd, args)
	}

	input := args[0]

	file, err := os.Open(input)
	if os.IsNotExist(err) {
		return fmt.Errorf("file '%s' does not exist", input)
	}

	b, readErr := io.ReadAll(file)
	if readErr != nil {
		return readErr
	}

	encoded, encodeErr := Encode(b)
	if encodeErr != nil {
		return encodeErr
	}

	if output == "" {
		output = input + ".osaka"
	}

	newFile, newErr := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0644)
	if newErr != nil {
		return newErr
	}

	_, writeErr := newFile.Write([]byte(encoded))
	if writeErr != nil {
		return writeErr
	}

	fmt.Printf("Encoded your data '%s'\n", output)

	return nil
}

func decode(cmd *cobra.Command, args []string) error {
	if len(args) > 1 || isString {
		return decodeString(cmd, args)
	}
	input := args[0]

	file, err := os.Open(input)
	if os.IsNotExist(err) {
		return fmt.Errorf("file '%s' does not exist", input)
	}

	b, readErr := io.ReadAll(file)
	if readErr != nil {
		return readErr
	}

	decoded, decodeErr := Decode(string(b))
	if decodeErr != nil {
		return decodeErr
	}

	if output == "" {
		output = input + ".decoded_osaka"
	}

	newFile, newErr := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0644)
	if newErr != nil {
		return newErr
	}

	_, writeErr := newFile.Write([]byte(decoded))
	if writeErr != nil {
		return writeErr
	}

	fmt.Printf("Decoded your data '%s'\n", output)

	return nil
}

func encodeString(_ *cobra.Command, args []string) error {
	text := strings.Join(args, " ")
	encoded, decodeErr := Encode([]byte(text))
	if decodeErr != nil {
		return decodeErr
	}

	fmt.Printf("%s\n", encoded)

	return nil
}

func decodeString(_ *cobra.Command, args []string) error {
	text := strings.Join(args, " ")
	decoded, decodeErr := Decode(text)
	if decodeErr != nil {
		return decodeErr
	}

	fmt.Printf("%s\n", decoded)

	return nil
}
