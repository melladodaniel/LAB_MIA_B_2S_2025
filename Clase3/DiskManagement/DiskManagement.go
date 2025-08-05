package DiskManagement

import (
	"clase3/Structs"
	"clase3/Utilities"
	"fmt"
)

func Mkdisk(size int, fit string, unit string) {
	fmt.Println("======Inicio MKDISK======")
    fmt.Println("======Parámetros Recibidos======")
	fmt.Println("Size:", size)
	fmt.Println("Fit:", fit)
	fmt.Println("Unit:", unit)

	// validar fit = bf/ff/wf
	if fit != "bf" && fit != "ff" && fit != "wf" {
		fmt.Println("Error: Fit debe ser bf, ff, o wf")
		return
	}

	// validar que el tamaño sea mayor a 0
	if size <= 0 {
		fmt.Println("Error: Tamaño debe ser mayor a 0")
		return
	}

	// validar que unidad sea igual a k o m
	if unit != "k" && unit != "m" {
		fmt.Println("Error: Unidad debe ser k o m")
		return
	}

	// Crear Archivo
	err := Utilities.CreateFile("./test/A.bin")
	if err != nil {
		fmt.Println("Error creando archivo", err)
		return
	}

	// Definir el tamaño del archivo
	if unit == "k" {
		size *= 1024
	} else {
		size *= 1024 * 1024
	}

	// Abrir archivo binario
	file, err := Utilities.OpenFile("./test/A.bin")
	if err != nil {
		return
	}

	// Buffer 1024 bytes
	zeroBuffer := make([]byte, 1024)

	// escribir 0 binarios en el archivo
	for i := 0; i < size/1024; i++ {
		err := Utilities.WriteObject(file, zeroBuffer, int64(i*1024))
		if err != nil {
			return
		}
	}

	// crear nueva instancia de MBR
	var newMBR Structs.MBR
	newMBR.MbrSize = int32(size)
	newMBR.Signature = 10                      // random
	copy(newMBR.Fit[:], fit)                   // fit del MRB
	copy(newMBR.CreationDate[:], "2025-08-04") // fecha actual del MBR

	// Escribir MBR al archivo
	if err := Utilities.WriteObject(file, newMBR, 0); err != nil {
		fmt.Println("Error escribiendo MBR al archivo:", err)
		return
	}

	var tempMBR Structs.MBR

	// Leer MBR del archivo para verificar
	if err := Utilities.ReadObject(file, &tempMBR, 0); err != nil {
		fmt.Println("Error leyendo MBR del archivo:", err)
		return
	}

	// Imprimir MBR para verificar
	fmt.Println("===Data recuperada===")
	fmt.Println("Tamaño del MBR:", tempMBR.MbrSize)
	fmt.Println("Fit:", string(tempMBR.Fit[:]))
	fmt.Println("Fecha de creación:", string(tempMBR.CreationDate[:]))
	fmt.Println("Firma:", tempMBR.Signature)

	// Cerrar el archivo binario
	defer file.Close()

	fmt.Println("======Fin MKDISK======")
}