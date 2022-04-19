package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Mkfs struct{
	id string
	tipo string
}

func NewMkfs(lista *arrayList.List) Mkfs {
	aux := Mkfs{"","FULL"}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == ID{
			aux.id = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == TYPE{
			aux.tipo = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MKFS")
			aux.id = ""
			aux.tipo = ""
			break
		}
	}
	return aux
}

func (d Mkfs) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MKFS... ")
	return nil
}