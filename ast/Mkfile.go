package ast

import (
	"./parametros"
	"fmt"
	"strconv"
	arrayList "github.com/colegno/arraylist"
)

type Mkfile struct{
	path string
	r bool
	size int
	cont string
}

func NewMkfile(lista *arrayList.List) Mkfile {
	aux := Mkfile{"",false,0,""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == R{
			aux.r = true
		} else if res.(parametros.Parametro).Tipo == CONT{
			aux.cont = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == SIZE{
			aux.size,_ = strconv.Atoi(res.(parametros.Parametro).Valor)
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MKFILE")
			aux.path = ""
			aux.r = false
			aux.cont = ""
			aux.size = 0
			break
		}
	}
	return aux
}

func (d Mkfile) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MKFILE... ")
	return nil
}