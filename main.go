package main

import (
	"flag"
	logger "github.com/lucas59356/go-logger"
	"net/http"
	"net/url"
	"os"
	"errors"
)

var (
	port     = flag.String("p", "80", "Onde eu vou escutar")
	basepath = flag.String("d", "./", "Root folder")
	// ErrFolderNotFound When the working folder does not exist
	ErrFolderNotFound = errors.New("A pasta de trabalho não existe")
	// ErrCantParseURL When server can't parse the URL
	ErrCantParseURL = errors.New("Não foi possivel parsear a url")
)

func main() {
	log := logger.New("main")
	flag.Parse()
	if !isPastaExist(*basepath) { // Será que essa pasta existe?
		log.Error(ErrFolderNotFound)
		flag.Usage()
	}
	bindTo := "0.0.0.0:" + *port
	log.Info("Iniciando servidor...")
	mux := http.NewServeMux()

	mux.HandleFunc("/", server)                                   // Handler do file server
	log.Info("Meu trabalho acontecerá na pasta: %s", *basepath) // Estou lendo na pasta!!
	log.Info("Entendido, vamos trabalhar em %s", bindTo)        // Estou bindando!!
	log.Info("Pau na máquina!")

	err := http.ListenAndServe(bindTo, mux)
	if err != nil {
		log.Error(err)
	}
}

func server(w http.ResponseWriter, r *http.Request) {
	log := logger.New("srv")
	var path = *basepath + r.RequestURI  // Junta endereço da pasta com a url pedida
	path, err := url.QueryUnescape(path) // Remover urlencode
	if err != nil {
		log.Error(ErrCantParseURL)
		log.Error(err)
	}
	log.Debug("[%s] %s %s", r.Host, r.Method, r.URL) // Debug, e tbm por que é mais legal :v
	http.ServeFile(w, r, path)
	w.Header().Set("Cache-Control", "max-age=5") // Evitar que o cache do navegador bugue a listagem, tirando a necessidade do ^F5
}

func isPastaExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
