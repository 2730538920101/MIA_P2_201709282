package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"./analizador"
	
)
func main(){
	fmt.Println("BIENVENIDO AL SISTEMA DE ARCHIVOS, PROYECTO 2 MIA 1S2022")
	fmt.Println("CARLOS JAVIER MARTINEZ POLANCO 201709282")
	fmt.Println()
	for true{
		fmt.Print("Ingrese un comando: ")
		prueba := bufio.NewReader(os.Stdin)
		input, _ := prueba.ReadString('\n')
		if input !="\n"{
			switch input {
				case "exit\n":
					{
						os.Exit(0)
					}
				case "cls\n":
					{
						cmd := exec.Command("clear")
						cmd.Stdout = os.Stdout
						cmd.Run()
					}
				default:
					{

						analizador.Analizar(input)
						fmt.Println()
						fmt.Println("ANALISIS REALIZADO")
						fmt.Println()	
					}					
			}
		}
	}
}

