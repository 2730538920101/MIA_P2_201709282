package ast

import (
	"./parametros"
	"fmt"
	"strconv"
	arrayList "github.com/colegno/arraylist"
)

type Mkdisk struct{
	size int
	fit string
	unit string
	path string
}

func NewMkdisk(lista *arrayList.List) Mkdisk {
	aux := Mkdisk{0,"FF","M",""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == UNIT{
			aux.unit = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == FIT{
			aux.fit = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == SIZE{
			aux.size,_ = strconv.Atoi(res.(parametros.Parametro).Valor)
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MKDISK")
			aux.path = ""
			aux.unit = ""
			aux.fit = ""
			aux.size = 0
			break
		}
	}
	return aux
}

func (d Mkdisk) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MKDISK... ")
	return nil
}