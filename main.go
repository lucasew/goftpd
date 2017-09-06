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
		log.Fatal("ERRO: A pasta de trabalho não existe: " + *basepath)
		flag.Usage()
	}
	bindTo := "0.0.0.0:" + *port
	log.Println("Iniciando servidor...")
	mux := http.NewServeMux()

	mux.HandleFunc("/", server)                                   // Handler do file server
	log.Printf("Meu trabalho acontecerá na pasta:  " + *basepath) // Estou lendo na pasta!!
	log.Printf("Entendido, vamos trabalhar em %s", bindTo)        // Estou bindando!!
	log.Println("Pau na máquina!")

	err := http.ListenAndServe(bindTo, mux)
	if err != nil {
		log.Fatal("ERRO: " + err.Error())
	}
}

func server(w http.ResponseWriter, r *http.Request) {
	var path = *basepath + r.RequestURI  // Junta endereço da pasta com a url pedida
	path, err := url.QueryUnescape(path) // Remover urlencode
	if err != nil {
		log.Println("ERRO: Não foi possivel parsear a url: ", path)
	}
	log.Printf("[%s] %s %s", r.Host, r.Method, r.URL) // Debug, e tbm por que é mais legal :v
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
