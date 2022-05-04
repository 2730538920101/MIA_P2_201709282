package ast

import (
	"./parametros"
	"fmt"
	"strconv"
	arrayList "github.com/colegno/arraylist"
	"../funciones"
	"os"
	"./structs"
	"bytes"


)

type Fdisk struct{
	size int
	fit string
	unit string
	path string
	tipo string
	name string
}

func NewFdisk(lista *arrayList.List) Fdisk {
	aux := Fdisk{0,"WF","K","","P",""}
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
		} else if res.(parametros.Parametro).Tipo == TYPE{
			if (res.(parametros.Parametro).Valor == PP )|| (res.(parametros.Parametro).Valor == E )|| (res.(parametros.Parametro).Valor == L){
				aux.tipo = res.(parametros.Parametro).Valor
			}else{
				fmt.Println("EL VALOR ASIGNADO AL PARAMETRO TYPE NO ES ACEPTADO")
				aux.path = ""
				aux.unit = ""
				aux.fit = ""
				aux.size = 0
				aux.tipo = ""
				aux.name = ""
				break
			}
		} else if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO FDISK")
			aux.path = ""
			aux.unit = ""
			aux.fit = ""
			aux.size = 0
			aux.tipo = ""
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Fdisk) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO FDISK... ")
	a := funciones.TamByte(d.unit, d.size)
	fmt.Printf("EL TAMAÑO DE LA PARTICION SERA DE: %d\n", a)
	if d.size <= 0 || d.path == "" || d.name == ""{
		fmt.Println("LOS PARAMETROS SIZE, PATH Y NAME SON OBLIGATORIOS PARA EL COMANDO FDISK")
		return nil
	}
	fi, err := os.Stat(d.path)
	if err != nil{
		fmt.Println("ERROR AL DETERMINAR EL TAMAÑO DEL ARCHIVO")
	}
	t := fi.Size()
	disk := funciones.LeerMBR(d.path, int(t))
	var part structs.Partition
	var flag_esp bool = false
	var flag_ext bool = false
	var i int
	var eb structs.ExtendedBootRecord
	for i = 0; i < 4; i++{
		aux := &disk.Mbr_partition[i]
		st := string(bytes.Trim(aux.Part_status[:], "\x00"))
		t := string(bytes.Trim(aux.Part_type[:], "\x00"))
		if  st == "INACTIVO"{
			flag_esp = true
			break
		}else{
			if string(bytes.Trim(aux.Part_name[:], "\x00")) == d.name { //validar nombres en la extendida
				fmt.Printf("LA PARTICION %s YA FUE CREADA ANTERIORMENTE\n", d.name)
				return nil
			}else{
				if d.tipo == "L" && t == "E"{
					break
				}
			}
		}
	}
	pos, _ := strconv.Atoi(string(bytes.Trim(disk.Mbr_partition[i].Part_start[:],"\x00")))
	ta, _ := strconv.Atoi(string(bytes.Trim(disk.Mbr_partition[i].Part_size[:],"\x00")))
	eb = funciones.GetPrimerEBR(d.path, pos, ta)
	var inicio int
	var flag_log bool = false
	if d.tipo == "L"{
		flag_log = true
	}
	funciones.GetDireccionInicio(disk, d.fit, a, &inicio, i, eb, flag_log)
	tam := strconv.Itoa(a)
	iniciost := strconv.Itoa(inicio) 
	part = structs.NewPartition("ACTIVO", d.tipo, d.fit, iniciost, tam, d.name)
	fmt.Printf("LA DIRECCION INICIAL PARA LA PARTICION ES: %d\n", inicio)
	switch d.tipo{
		case PP:
			{
				fmt.Println("SE CREARA UNA PARTICION PRIMARIA")
				if inicio > 0{
					if flag_esp{
						disk.Mbr_partition[i] = part
						funciones.EscribirMBR(d.path, &disk)
						
					}else{
						fmt.Println("ERROR AL CREAR LA PARTICION")
					} 
				}else{
					fmt.Println("ERROR AL CREAR LA PARTICION")
				} 
				break
			}
		case E:
			{
				fmt.Println("SE CREARA UNA PARTICION EXTENDIDA")
				if inicio > 0{
					if flag_esp{
						if ValidarExtendida(&disk, flag_ext){
							fmt.Println("YA EXISTE UNA PARTICION EXTENDIDA EN EL DISCO")
							return nil
						}else{
							disk.Mbr_partition[i] = part
							funciones.EscribirMBR(d.path, &disk)
						
						}		
					}else{
						fmt.Println("ERROR AL CREAR LA PARTICION")
					} 
				}else{
					fmt.Println("ERROR AL CREAR LA PARTICION")
				} 
				break
			}
		case L:
			{
				fmt.Println("SE CREARA UNA PARTICION LOGICA")
				if inicio > 0{
					if ValidarExtendida(&disk, flag_ext){
						stat := string(bytes.Trim(part.Part_status[:],"\x00"))
						ft := string(bytes.Trim(part.Part_fit[:],"\x00"))
						st := string(bytes.Trim(part.Part_start[:],"\x00"))
						stst, _ := strconv.Atoi(st)
						tam := string(bytes.Trim(part.Part_size[:],"\x00"))
						nom := string(bytes.Trim(part.Part_name[:],"\x00"))
						l := structs.NewEBR(stat, ft, st, tam, "-1", nom)
						funciones.EscribirEBR(d.path, &l, stst)	
						if (eb != structs.ExtendedBootRecord{}){
							copy(eb.Part_next[:], st)
							st2 := string(bytes.Trim(eb.Part_start[:],"\x00"))
							stst2, _ := strconv.Atoi(st2)
							funciones.EscribirEBR(d.path, &eb, stst2)
						}
						dkst, _ := strconv.Atoi(string(bytes.Trim(disk.Mbr_partition[i].Part_start[:], "\x00")))
						funciones.LeerEBR(d.path, int(t), dkst)
					}else{
						fmt.Println("NO EXISTE UNA PARTICION EXTENDIDA EN EL DISCO")
						return nil
					}
				}else{
					fmt.Println("ERROR AL CREAR LA PARTICION")
				}
				break
			}
	}
	fi2, err := os.Stat(d.path)
	if err != nil{
		fmt.Println("ERROR AL DETERMINAR EL TAMAÑO DEL ARCHIVO")
	}
	p := fi2.Size()
	funciones.LeerMBR(d.path, int(p))
	return nil
}

func ValidarExtendida(disk* structs.MasterBootRecord, flag_ext bool)bool{
	for i := 0; i < 4; i++{
		aux := &disk.Mbr_partition[i]
		aux_t := string(bytes.Trim(aux.Part_type[:],"\x00"))
		if aux_t == "E"{
			flag_ext = true
			return true
		}
	}
	return false
}



