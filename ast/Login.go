package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Login struct{
	usuario string
	password string
	id string
}

func NewLogin(lista *arrayList.List) Login {
	aux := Login{"","",""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == USUARIO{
			aux.usuario = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == PASSWORD{
			aux.password = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == ID{
			aux.id = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO LOGIN")
			aux.usuario = ""
			aux.password = ""
			aux.id = ""
			break
		}
	}
	return aux
}

func (d Login) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO LOGIN... ")
	return nil
}