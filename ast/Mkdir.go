package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Mkdir struct{
	path string
	p bool
}

func NewMkdir(lista *arrayList.List) Mkdir {
	aux := Mkdir{"",false}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == P{
			aux.p = true
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MKDIR")
			aux.path = ""
			aux.p = false
			break
		}
	}
	return aux
}

func (d Mkdir) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MKDIR... ")
	return nil
}