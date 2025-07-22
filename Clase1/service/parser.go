package service

import (
	"bufio"
	"os"
	"regexp"
	"clase1/model"
)

var mkdiskRegex = regexp.MustCompile(`^mkdisk\s+-Size=\d+\s+-unit=[K|M]\s+-path=\/[\/\w\.-]+\.mia$`)


func AnalizarArchivo(path string) ([]model.LineAnalysis, error){
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var resultados []model.LineAnalysis
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		linea := scanner.Text()
		valida := mkdiskRegex.MatchString(linea)
		mensaje := "Correcta"
		if !valida {
			mensaje = "Sintaxis incorrecta"
		}
		resultados = append(resultados, model.LineAnalysis{
			Linea:	linea,
			Valida:	valida,
			Mensaje: mensaje,
		})
	}

	return resultados, scanner.Err()
}