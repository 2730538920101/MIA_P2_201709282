package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Rep struct{
	ruta string
	id string
	path string
	name string
}

func NewRep(lista *arrayList.List) Rep {
	aux := Rep{"","","",""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == ID{
			aux.id = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == RUTA{
			aux.ruta = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO REP")
			aux.path = ""
			aux.ruta = ""
			aux.id = ""
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Rep) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO REP... ")
	return nil
}