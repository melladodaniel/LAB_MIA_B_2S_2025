package controller

import (
	"encoding/json"
	"net/http"
	"clase1/service"
)

func AnalizarArchivoHandler(w http.ResponseWriter, r *http.Request){
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		http.Error(w, "Falta el parámetro `file`", http.StatusBadRequest)
		return
	}

	resultados,  err := service.AnalizarArchivo(filePath)
	if err != nil {
		http.Error(w, "Error al procesar el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)

}