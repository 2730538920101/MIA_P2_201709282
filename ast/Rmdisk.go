package ast

import (
	"./parametros"
	"fmt"
	
	arrayList "github.com/colegno/arraylist"
)

type Rmdisk struct{
	path string
}

func NewRmdisk(lista *arrayList.List) Rmdisk {
	aux := Rmdisk{""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO RMDISK")
			aux.path = ""
			break
		}
	}
	return aux
}

func (d Rmdisk) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO RMDISK... ")
	return nil
}