package ast

import (
	"./parametros"
	"fmt"
	"../funciones"
	"os"
	arrayList "github.com/colegno/arraylist"
)

type Rmdisk struct{
	path string
}

func NewRmdisk(lista *arrayList.List) Rmdisk {
	aux := Rmdisk{""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO RMDISK")
			aux.path = ""
			break
		}
	}
	return aux
}

func (d Rmdisk) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO RMDISK... ")
	if !funciones.ArchivoExiste(d.path){
		fmt.Println("ERROR EN LA RUTA...")
		return nil
	}
	var inp string
	fmt.Println("DESEA ELEMINAR EL DISCO??? (Y/N)")
	fmt.Scanln(&inp)
	if inp == "y" || inp == "Y"{
		err := os.Remove(d.path)
		if err != nil{
			fmt.Printf("ERROR AL ELIMINAR EL ARCHIVO %v\n", err)
			return nil
		}
		fmt.Println("DISCO ELIMINADO EXITOSAMENTE")
	}else if inp =="n" || inp =="N"{
		fmt.Println("HA CANCELADO LA ELIMINACION DEL DISCO")
	}
	return nil
}