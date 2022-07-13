package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var emailThread = make(map[string]bool)
var mu sync.Mutex

func main() {

	// Recebe a requisição
	// Adiciona o e-mail da requisição na thread
	// Se tem thread com o e-mail aguarda até ele ser false
	// Se não tem faz o verificar

	http.HandleFunc("/verify", procesarVerify)
	http.ListenAndServe(":8080", nil)

}

func addEmailRequest(email string) bool {
	mu.Lock()
	defer mu.Unlock()

	_, exists := emailThread[email]
	if !exists {
		emailThread[email] = true
	}

	return exists
}

func removeEmail(email string) {
	log.Printf("[%s] Remover", email)
	delete(emailThread, email)
}

func procesarVerify(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	thread := r.URL.Query().Get("source")

	if len(email) == 0 {
		http.Error(w, "[2] Precisa declarar o email", http.StatusInternalServerError)
		return
	}

	log.Printf("[%s] inicio", thread)
	go func() {

		for emailThread[email] {
			log.Printf("[%s] esperar", thread)
			time.Sleep(10 * time.Microsecond)
		}

		defer func() {
			removeEmail(email)
		}()

		processImage(email, thread)
	}()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("email: " + email))

}

func processImage(email string, thread string) {
	log.Printf("[%s] Validar", thread)
	time.Sleep(3 * time.Second)
	log.Printf("[%s] Validado", thread)
}
