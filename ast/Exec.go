package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Exec struct{
	path string
}

func NewExec(lista *arrayList.List) Exec {
	aux := Exec{""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO EXEC")
			aux.path = ""
			break
		}
	}
	return aux
}

func (d Exec) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO EXEC... ")
	return nil
}