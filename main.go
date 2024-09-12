package main

import (
	"log"
	"os"
	"strings"
)

type File struct {
	name		string
	content		string
}

func main() {
	// Read all the HTML files
	files, err := html_finder_reader()
	if err != nil {
		log.Fatalln("No HTML file found in the ./html directory! Error message: ", err)
	}
	log.Println("Files in the HTML directory: ")
	log.Print(files)
}

func html_finder_reader() ([]File, error) {
	// Open the directory
	d, err := os.Open("./html")
	if err != nil {
		log.Fatalln("Unable to open the ./html directory. Error message: ", err)
		return nil, err
	}
	// Read files from the directory
	fs_files, err := d.ReadDir(0)
	if err != nil {
		log.Fatalln("Unable to read ./html directory. Error message: ", err)
		return nil, err
	}

	var files[]File
	for _, file := range fs_files {
		// Check for NOT a directory and ".html" extension
		if !file.IsDir() && strings.Contains(file.Name(), ".html") {
			file_content, err := os.ReadFile("./html/" + file.Name())
			if err != nil {
				log.Fatalf("\nError reading the file: %s. Error message: %s\n", file.Name(), err)
				
			}
			f := File {
				name: file.Name(),
				content: string(file_content),
			}
			files = append(files, f)
		}
	}
	
	return files, nil
}

func html_2_md_parser(html string) (string, error) {
	return "", nil
}

func file_writer(file_name string, path string) (error) {
	return nil
}