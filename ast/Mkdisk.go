package ast

import (
	"./parametros"
	"fmt"
	"strconv"
	arrayList "github.com/colegno/arraylist"
	"../funciones"
	"os"
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
	if d.size <= 0{
		fmt.Println("EL PARAMETRO SIZE ES OBLIGATORIO Y DEBE SER MAYOR QUE CERO...")
	}else if d.path == ""{
		fmt.Println("EL PARAMETRO PATH ES OBLIGATORIO Y DEBE SER UNA RUTA VALIDA")
	}
	a := funciones.TamByte(d.unit, d.size)
	fmt.Println(a)
	funciones.CrearCarpetaDisco(d.path)
	lim := 0
	block := make([]byte, a)
	for j := 0; j < a; j++{
		block[j] = 0
	}
	disco, err := os.Create(d.path)
	if err != nil{
		fmt.Println(err)
	}
	for lim < a{
		_, err := disco.Write(block)
		if err != nil{
			fmt.Println(err)
		}
		lim++
	}
	disco.Close()
	return nil
}




