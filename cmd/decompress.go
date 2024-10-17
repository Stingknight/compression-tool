/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	// "path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// decompressCmd represents the decompress command
var decompressCmd = &cobra.Command{
	Use:   "decompress",
	Short: "To decompress the file",
	Long:  `To decompress the file`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")
		if filename == "" {
			fmt.Println("please provide the filename")
		}

		if err := Decompress(filename); err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		foldername, _ := cmd.Flags().GetString("foldername")
		if foldername == "" {
			fmt.Println("provide a name of the folder")
		}

		if err := DecompressFolder(foldername,"./ngnix-chart"); err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}

		fmt.Println("decompressing completed------------------>>>>>>>>>>>>")
	},
}

func init() {
	rootCmd.AddCommand(decompressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	decompressCmd.PersistentFlags().String("filename", "", "filename to be passed")
	decompressCmd.PersistentFlags().String("foldername", "", "foldername to be passed")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decompressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Decompress(filename string) error {
	
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	gzipreader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	
	defer gzipreader.Close()

	var newfilename = strings.Replace(filename, ".gz", ".txt", -1)
	
	outfile, err := os.Create(newfilename)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(outfile)
	
	var data = make([]byte, 2048)

	defer outfile.Close()

	for {
		n, err := gzipreader.Read(data)
		
		if n > 0 {
			_, err = writer.Write(data[:n])

			if err != nil {
				return err
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	
	defer writer.Flush()
	return nil

}


func DecompressFolder(source string,destination string) (error){

	zipreader,err := zip.OpenReader(source)
	if err!=nil{
		return err
	}

	defer zipreader.Close()

	extracFIleAndWrite := func (z *zip.ReadCloser) error{
		for _, file := range z.File{

			fmt.Printf("Unzipping %s:\n", file.Name)
	
			arc,err := file.Open()
			if err!=nil{
				return err
			}
	
			defer arc.Close()
	
			var newfilepath string = filepath.Join(destination,file.Name)
			// case 1 if it is a directory
			fmt.Println("newfilepath",newfilepath)
	
			if file.FileInfo().IsDir(){
				continue
			}
	
			err = os.MkdirAll(filepath.Dir(newfilepath), os.ModePerm)
			if err!=nil{
				return err
			}
	
			// CASE 2 : we have a file
			// create new uncompressed file
			uncompressedfile,err := os.OpenFile(newfilepath,os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err!=nil{
				return err
			}
			
			// Copy the contents
			_,err = io.Copy(uncompressedfile,arc)
			if err!=nil{
				return err
			}
	
			defer uncompressedfile.Close()
		}
		return nil
	}

	if err := extracFIleAndWrite(zipreader);err!=nil{
		return nil
	}

	return nil
}