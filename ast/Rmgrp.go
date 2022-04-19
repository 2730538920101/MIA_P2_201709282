package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Rmgrp struct{
	name string
}

func NewRmgrp(lista *arrayList.List) Rmgrp {
	aux := Rmgrp{""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO RMGRP")
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Rmgrp) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO RMGRP... ")
	return nil
}