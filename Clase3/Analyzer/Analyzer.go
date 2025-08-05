package Analyzer

import (
	"clase3/DiskManagement"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
    "bufio"
)

var re = regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

func Analyze() {
	var input string
	fmt.Println("Ingresa un comando para analizar:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	// input = "Mkdisk -size=50 -unit=M -fit=WF"

	command, params := getCommandAndParams(input)

	fmt.Println("Comando:", command, "Parámetros:", params)

	AnalyzeCommnad(command, params)
}

func getCommandAndParams(input string) (string, string) {
	parts := strings.Fields(input)
	if len(parts) > 0 {
		command := strings.ToLower(parts[0])
		params := strings.Join(parts[1:], " ")
		return command, params
	}
	return "", input
}

func AnalyzeCommnad(command string, params string) {
	switch command {
	case "mkdisk":
		fn_mkdisk(params)
	case "exit":
		fmt.Println("Saliendo del programa")
		os.Exit(0)
	default:
		fmt.Println("Error: Comando no reconocido.")
	}
}


func fn_mkdisk(params string) {
	// Definiendo banderas
	fs := flag.NewFlagSet("mkdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Size")
	fit := fs.String("fit", "f", "Fit")
	unit := fs.String("unit", "m", "Unit")

	// obtener valores
	managementFlags(fs, params)

	// Llamar a la función
	DiskManagement.Mkdisk(*size, *fit, *unit)

}

func managementFlags(fs *flag.FlagSet, params string) {
	// Parser de las banderas
	fs.Parse(os.Args[1:])

	// Encontrar las banderas en el input
	matches := re.FindAllStringSubmatch(params, -1)

	// Obtener los nombres de todas las banderas
	var flagNames []string
	fs.VisitAll(func(f *flag.Flag) {
		flagNames = append(flagNames, f.Name)
	})

	// Procesar el comando ingresado
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		if contains(flagNames, flagName) {
			fs.Set(flagName, flagValue)
		} else {
			fmt.Println("Error: Bandera no encontrada:", flagName)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
