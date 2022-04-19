package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Mkusr struct{
	usuario string
	pwd string
	grp string
}

func NewMkusr(lista *arrayList.List) Mkusr {
	aux := Mkusr{"","",""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == USUARIO{
			aux.usuario = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == PWD{
			aux.pwd = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == GRP{
			aux.grp = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MKUSR")
			aux.usuario = ""
			aux.pwd = ""
			aux.grp = ""
			break
		}
	}
	return aux
}

func (d Mkusr) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MKUSR... ")
	return nil
}