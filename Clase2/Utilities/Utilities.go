package Utilities

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

//función para crear el archivo binario
func CreateFile(name string) error {
	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Error creando directorio: ", err)
		return err
	}
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Error creando archivo:", err)
			return err
		}
		defer file.Close()
	}
	return nil
}

//función para abrir el archivo binario en modo lectura/escritura
func OpenFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error abriendo archivo:", err)
		return nil, err
	}
	return file, nil
}

//función para escribir el objeto en el archivo binario
func WriteObject(file *os.File, data interface{}, position int64) error {
	file.Seek(position, 0)
	err := binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Error escribiendo el archivo:", err)
		return err
	}
	return nil
}

//Función para leer los objetos desde el archivo binario
func ReadObject(file *os.File, data interface{}, position int64) error {
	file.Seek(position, 0)
	err := binary.Read(file, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("Error leyendo el objeto del archivo binario", err)
		return err
	}
	return nil
}