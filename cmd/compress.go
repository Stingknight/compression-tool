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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/spf13/cobra"
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "use to compress the file to zip file",
	Long:  `use to compress the file to zip file`,
	Run: func(cmd *cobra.Command, args []string) {

		filename, _ := cmd.Flags().GetString("filename")
		if filename == "" {
			fmt.Println("provide a name of the file")
			return
		}

		if err := Compress(filename); err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}

		foldername, _ := cmd.Flags().GetString("foldername")
		if foldername == "" {
			fmt.Println("provide a name of the file")
		}

		if err := CompressFolder(foldername); err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}

		fmt.Println("compress of file or folder is completed--------->")
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	compressCmd.PersistentFlags().String("filename", "", "use to compress the file to zip file")

	compressCmd.PersistentFlags().String("foldername", "", "use to compress the folder to zip file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// this writes to the particular file

func Compress(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	var data = make([]byte, 2048)

	reader := bufio.NewReader(file)

	var newfilename = strings.Replace(filename, ".txt", ".gz", -1)
	file, err = os.Create(newfilename)
	if err != nil {
		return err
	}

	writer := gzip.NewWriter(file)

	defer writer.Close()

	for {
		// here what is happening it  read the small chunk from the file and if the file exceeds the limit of the buffer(4kb)
		// it will completely fille the buffer and it will pass from the buffer to a variable(here data and it will store there)
		// after that it will get again small chunk from from the file and passes from buffer 
		// untile the file reaches the end of line  
		n, err := reader.Read(data)
	
		if n > 0 {
			// here it will write the data to the buffer(4kb) and then from there if the buffer is full it automatically flushes to the file until the each and every data is passed from buffer to file 
			_, err = writer.Write(data[:n])
			if err != nil {
				return err
			}
		}

		if err==io.EOF{
			break
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// new function which compress the folder or the multiple files inside that folder


func CompressFolder(foldername string) (error){		

	zipfile,err := os.Create(foldername+".zip")
	if err!=nil{
		return err
	}

	defer zipfile.Close()

	writer := zip.NewWriter(zipfile)

	defer writer.Close()

	if err := addFilesToZip(foldername,writer,"");err!=nil{
		return err
	}

	return nil

}


func addFilesToZip(foldername string,w *zip.Writer,basepath string) (error){
	
	files,err := ioutil.ReadDir(foldername)

	if err!=nil{
		return err
	}

	for _, file := range files{
		
		var fullfilepath = filepath.Join(foldername,file.Name())
		
		if _,err := os.Stat(fullfilepath);os.IsNotExist(err){
			continue
		}

		if file.Mode() & os.ModeSymlink!=0{
			continue
		}


		if file.IsDir(){
			
			// it is directory again need to recursive way to add the files
			if err := addFilesToZip(fullfilepath,w, filepath.Join(basepath, file.Name())); err != nil {
				return err
			}

		}else if file.Mode().IsRegular(){
			data,err := os.ReadFile(fullfilepath)
			
			if err!=nil{
				return err
			}

			f,err := w.Create(filepath.Join(basepath,file.Name()))
			if err!=nil{
				return err
			}

			_,err = f.Write(data)
			if err!=nil{
				return err
			}

		}	

	}
	return nil
}