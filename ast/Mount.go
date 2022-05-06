package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"../funciones"
	"../globales"
	"./structs"
	"strconv"
	"os"
	"bytes"
)



type Mount struct{
	name string
	path string
}

func NewMount(lista *arrayList.List) Mount {
	aux := Mount{"",""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO MOUNT")
			aux.path = ""
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Mount) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO MOUNT... ")
	if d.path == "" || d.name == ""{
		fmt.Println("LOS PARAMETROS PATH Y NAME SON OBLIGATORIOS EN EL COMANDO MOUNT")
		return nil
	}
	var cont int = 0
	var flag bool = false
	for (globales.Disk_mont[cont] != structs.DiscosMontados{}){
		if globales.Disk_mont[cont].Path == d.path{
			flag = true
			break
		}
		cont++	
	}
	if !flag{
		globales.Disk_mont[cont].Path = d.path
		globales.Disk_mont[cont].Letra = string(rune(97 + cont))
	}
	var flag2 bool = false
	var cont2 int = 0
	disco_r := globales.Disk_mont[cont]
	for (disco_r.Particiones_montadas[cont2] != structs.ParticionesMontadas{}){
		if disco_r.Particiones_montadas[cont2].Name == d.name{
			fmt.Println("LA PARTICION YA ESTA MONTADA")
			return nil
		}
		cont2++
	}
	fi, err := os.Stat(d.path)
	if err != nil{
		fmt.Println("ERROR AL DETERMINAR EL TAMAÑO DEL ARCHIVO")
	}
	t := fi.Size()
	disk := funciones.LeerMBR(d.path, int(t))
	
	if (disk == structs.MasterBootRecord{}){
		fmt.Println("ERROR AL LEER EL MBR")
		return nil
	}
	var i int
	var init int
	for i = 0; i < 4; i++{
		nom := string(bytes.Trim(disk.Mbr_partition[i].Part_name[:],"\x00"))
		if nom == d.name{
			flag2 = true
			initcp, _ := strconv.Atoi(string(bytes.Trim(disk.Mbr_partition[i].Part_start[:],"\x00")))
			init = initcp
			break
		}
	}
	if flag2{
		ide := funciones.GetPartId(disco_r.Letra,cont2)
		globales.Disk_mont[cont].Particiones_montadas[cont2].Id = ide
		globales.Disk_mont[cont].Particiones_montadas[cont2].Name = d.name
		globales.Disk_mont[cont].Particiones_montadas[cont2].Inicio = init
	}else{
		fmt.Println("LA PARTICION NO EXISTE")
		return nil
	}
	fmt.Println("SE HA MONTADO LA PARTICION EXITOSAMENTE")
	//ESCRIBIR EL SUPER BLOQUE
	return nil
}

