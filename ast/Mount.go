package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Mount struct{
	name string
	path string
}

func NewMount(lista *arrayList.List) Mount {
	aux := Mount{"",""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MOUNT")
			aux.path = ""
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Mount) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MOUNT... ")
	return nil
}