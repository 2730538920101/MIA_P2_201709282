package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Mkgrp struct{
	name string
}

func NewMkgrp(lista *arrayList.List) Mkgrp {
	aux := Mkgrp{""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MKGRP")
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Mkgrp) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MKGRP... ")
	return nil
}