package main

import (
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

type dotFileHidingFile struct {
	http.File
}

type dotFileHidingFileSystem struct {
	http.FileSystem
}

func containsDotFile(name string) bool {
	parts := strings.SplitSeq(name, "/")
	for part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

func (f dotFileHidingFile) Readdir(n int) (fis []fs.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files { // Filters out the dot files
		if !strings.HasPrefix(file.Name(), ".") {
			fis = append(fis, file)
		}
	}
	if err == nil && n > 0 && len(fis) == 0 {
		err = io.EOF
	}
	return
}

func (fsys dotFileHidingFileSystem) Open(name string) (http.File, error) {
	if containsDotFile(name) { // If dot file, return 403 response
		return nil, fs.ErrPermission
	}

	file, err := fsys.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return dotFileHidingFile{file}, nil
}

func main() {

	fs := dotFileHidingFileSystem{http.Dir("public/")}
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(fs))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
