package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Rmusr struct{
	usuario string
}

func NewRmusr(lista *arrayList.List) Rmusr {
	aux := Rmusr{""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == USUARIO{
			aux.usuario = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO RMUSR")
			aux.usuario = ""
			break
		}
	}
	return aux
}

func (d Rmusr) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO RMUSR... ")
	return nil
}