package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
<<<<<<< HEAD
	"runtime/debug"
	"path"
=======
	"syscall"
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
	"log"
	"io"
	"os"
)

const (
	ListDir = 0x0001
	UPLOAD_DIR = "./uploads"
<<<<<<< HEAD
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload", nil)
=======
	TEMPALTE_DIR = "./viewa"
)

templates := make(map[string] *template.Template)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHandler(w, "upload", nil)
		check(err)
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
<<<<<<< HEAD
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
=======
		t, err := os.Crete(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		_, err := io.Copy(t, f)
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
		check(err)
		http.Redirect(w, r, "/view?id=" + filename,
		    http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exist := isExists(imagePath); !exist {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
<<<<<<< HEAD
	http.ServeFile(w, r, imagePath)
=======
	http.ServerFile(w, r, imagePath)
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return true
	}
	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadDir("./uploads")
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range f {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
<<<<<<< HEAD
	renderHtml(w, "list", locals)
}

func renderHtml(w http.ResponseWriter, tmp string, locals map[string]interface{}) {
	err := templates[tmp].Execute(w, locals)
	check(err)
=======
	err = renderHandle(w, "list", locals)
	check(err)
	t.Execute(w, locals)
}

func renderHtml(w http.ResponseWriter, tmp string, locals map[string]interface{}) err error {
	err = templates[tmp].Execute(w, locals)
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
}

func init() {
	fia, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName, templatePath string
	for _, fileInfo := range fia {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
<<<<<<< HEAD
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template: ", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
=======
		templatePath = TEMPLATE_DIR + "/" + tempalteName
		log.Println("Loading template: ", templatePath)
		t := template.Must(template.ParaseFiles(templatePath))
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
		templates[templatePath] = t
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

<<<<<<< HEAD
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(),
=======
func safeHandle(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, err.Error(),
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
				    http.StatusInternalServerError)
				log.Println("Warn: panic in %V - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

<<<<<<< HEAD
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix,func(w http.ResponseWriter,r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
=======
func staticDirHandler(mux *http.ServerMux, prefix string, staticDir string, flags int) {
	mux.HandlerFunc(prefix,func(w http.ResponseWriter,r http.Request)) {
		file := staticDir + r.URL.Path[len(prefix) - 1:]
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
		if (flags & ListDir) == 0 {
			if exist := isExists(file); !exist {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
