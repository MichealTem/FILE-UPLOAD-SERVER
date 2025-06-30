package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "uploading File\n")

	//implementing File-Uploads.
	//1. parse input, type multipart/form-data.
	r.ParseMultipartForm(10 << 20)

	//2. retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//3. write temporary file on our server
	tempFile, err := os.CreateTemp("temp-images", "upload-*png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	//4. return whether or not this has been successful
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)

}

func main() {
	fmt.Println("Go File Upload Server")
	setupRoutes()
}
