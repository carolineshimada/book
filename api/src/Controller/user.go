package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Criando usuàrio")))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Buscando usuàrio")))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Buscando usuàrios")))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Atualizando usuàrio")))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Deletando usuàrio")))
}
