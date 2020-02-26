package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	port     = flag.String("p", "80", "Onde eu vou escutar")
	basepath = flag.String("d", "./", "Root folder")
)

func main() {
	flag.Parse()
	if !isPastaExist(*basepath) { // Será que essa pasta existe?
		log.Println("[main] erro: A pasta de trabalho não existe")
		flag.Usage()
		os.Exit(1)
	}
	bindTo := "0.0.0.0:" + *port
	log.Println("[main] info: Iniciando servidor...")
	mux := http.NewServeMux()

	mux.HandleFunc("/", server)                                                  // Handler do file server
	log.Printf("[main] info: Meu trabalho acontecerá na pasta: %s\n", *basepath) // Estou lendo na pasta!!
	log.Printf("[main] info: Entendido, vamos trabalhar em %s\n", bindTo)        // Estou bindando!!
	log.Println("[main] info: Pau na máquina!")

	err := http.ListenAndServe(bindTo, mux)
	if err != nil {
		log.Printf("[main] erro: %s", err.Error())
		os.Exit(1)
	}
}

func server(w http.ResponseWriter, r *http.Request) {
	var path = *basepath + r.RequestURI  // Junta endereço da pasta com a url pedida
	path, err := url.QueryUnescape(path) // Remover urlencode
	if err != nil {
		log.Println("[srv] warn: Não foi possivel parsear a url")
		log.Printf("[srv] error: %s", err.Error())
	}
	log.Printf("[srv] info: [%s] %s %s", r.Host, r.Method, r.URL) // Debug, e tbm por que é mais legal :v
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
