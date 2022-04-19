package ast

import (
	"./parametros"
	"fmt"
	"strconv"
	arrayList "github.com/colegno/arraylist"
)

type Fdisk struct{
	size int
	fit string
	unit string
	path string
	tipo string
	name string
}

func NewFdisk(lista *arrayList.List) Fdisk {
	aux := Fdisk{0,"WF","K","","P",""}
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
		} else if res.(parametros.Parametro).Tipo == TYPE{
			aux.tipo = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO FDISK")
			aux.path = ""
			aux.unit = ""
			aux.fit = ""
			aux.size = 0
			aux.tipo = ""
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Fdisk) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO FDISK... ")
	return nil
}