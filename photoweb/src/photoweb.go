package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
	"syscall"
	"log"
	"io"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParaseFiles("upload.html")
		if err != nil {
			http.Error(w, err.Error(),
			    http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
			    http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Crete(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
			    http.StatusInternalServerError)
			    return
		}
		defer t.Close()
		if _, err := io.Copy(t, f) err != nil {
			http.Error(w, err.Error(),
			    http.StautsInternalServerError)
			return
		}

		http.Redirect(w, r, "/view?id=" + filename,
		    http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exist := isExist(imagePath); !exist {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServerFile(w, r, imagePath)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(),
		    http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range f {
		images = append(images, fileInfo.Name())
	}

	locals["images"]=images t, err := template.ParaseFiles("list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, locals)
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
