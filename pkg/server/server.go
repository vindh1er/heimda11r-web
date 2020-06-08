package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"bytes"
	"html/template"

	"github.com/gorilla/mux"
	"github.com/markbates/pkger"
	"k8s.io/klog"
)

// StartServer sets up the web server.
func StartServer(serverPort int, serverHost string) {
	r := mux.NewRouter()
	r.StrictSlash(true)

	// main page
	r.HandleFunc("/", indexHandler)

	// health endpoint
	r.HandleFunc("/health", healthHandler)

	// Handle /static for static files
	fileServer := http.FileServer(pkger.Dir("/public/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	address := fmt.Sprintf("%s:%d", serverHost, serverPort)
	klog.V(3).Infof("Starting dashboard server on %s", address)
	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	klog.Infof("webserver listening on port %s:%d", serverHost, serverPort)
	klog.Fatal(srv.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	klog.V(5).Info("Handling request using projectsHandler.")

	templateFile, err := pkger.Open("/pkg/server/templates/index.gohtml")
	if err != nil {
	        klog.Error(err)
	}
	defer templateFile.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(templateFile)
	if err != nil {
	        klog.Error(err)
	}
	tmpl := template.Must(template.New("index.gohtml").Parse(buf.String()))
	err = tmpl.Execute(w, nil)
	if err != nil {
	        klog.Error(err)
	}


}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, `{"healthy": true}`)
	if err != nil {
		klog.Error(err)
	}
}
