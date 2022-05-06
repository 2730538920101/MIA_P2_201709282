package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"../funciones"
	"./structs"
	"os"
	"unsafe"
	"bytes"
	"strconv"
	"math"
	"time"
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
	if d.id == ""{
		fmt.Println("EL PARAMETRO ID ES OBLIGATORIO EN ESTE COMANDO")
		return nil
	}
	switch d.tipo{
		case "FAST":
			{
				fmt.Println("SE REALIZARA UN FORMATO RAPIDO A LA PARTICION")
				break
			}
		default:
			{
				fmt.Println("SE REALIZARA UN FORMATEO COMPLETO A LA PARTICION")
				break
			}
	}
	ini := funciones.GetStartOfPartById(d.id)
	if ini < 0{
		fmt.Println("NO SE HA MONTADO LA PARTICION")
		return nil
	}
	fmt.Printf("EL SUPER BLOQUE SE INICIARA A ESCRIBIR EN: %d\n", ini)
	pa := funciones.GetPathOfPartById(d.id)
	fi, err := os.Stat(pa)
	if err != nil{
		fmt.Println("ERROR AL DETERMINAR EL TAMAÑO DEL ARCHIVO")
	}
	t := fi.Size()
	var disco structs.MasterBootRecord = funciones.LeerMBR(pa, int(t))
	if (disco == structs.MasterBootRecord{}){
		fmt.Println("ERROR EN LA LECTURA DEL DISCO")
		return nil
	}
	var mbr_sb structs.SuperBlock
	var i_tbl structs.TablaInodos
	var bl structs.BlockFile
	var sizeN int
	var cantInodos int
	var cantBloques int
	nomPar := funciones.GetNameOfPartById(d.id)
	tam, _ := strconv.Atoi(string(bytes.Trim(disco.Mbr_tamano[:],"\x00")))
	sizeN = tam - int(unsafe.Sizeof(mbr_sb)) 
	sizeN /= 4 + int(unsafe.Sizeof(i_tbl)) + 3 * int(unsafe.Sizeof(bl))
	cantInodos = int(math.Floor(float64(sizeN)))
	cantBloques = cantInodos * 3
	var initPart int = funciones.GetInicioParticion(disco, nomPar)
	if initPart < 0{
		fmt.Println("ERROR DE LECTURA DE LA PARTICION")
		return nil
	}
	var sizeInodos int = int(unsafe.Sizeof(i_tbl)) * cantInodos
	tiemp := time.Now().Format("2006-01-02 15:04:05")
	cantBloquesSt := strconv.Itoa(cantBloques)
	cantInodosSt := strconv.Itoa(cantInodos)
	ins := strconv.Itoa(int(unsafe.Sizeof(i_tbl)))
	bls := strconv.Itoa(int(unsafe.Sizeof(bl)))
	bmi_st := initPart + int(unsafe.Sizeof(mbr_sb))
	bmi_sti := strconv.Itoa(bmi_st)
	bmb_st := bmi_st + cantInodos
	bmb_sti := strconv.Itoa(bmb_st)
	in_st := bmb_st + cantBloques
	in_sti:= strconv.Itoa(in_st)
	bl_st := in_st + sizeInodos
	bl_sti := strconv.Itoa(bl_st)  
	mbr_sb = structs.NewSuperBlock("EXT2",cantInodosSt,cantBloquesSt,cantBloquesSt,cantInodosSt, tiemp, "0", "0xEF53", ins, bls, "0", "0", bmi_sti, bmb_sti, in_sti, bl_sti)
	funciones.EscribirSuperBlock(pa, &mbr_sb, initPart)
	sblo := funciones.LeerSuperBlock(pa, len(funciones.Struct_to_bytes(mbr_sb)), initPart)
	funciones.CrearBitMapTable(cantInodos, bmi_st, pa)
	funciones.CrearBitMapTable(cantBloques, bmb_st, pa)
	funciones.CrearDirectorio(&sblo, pa, "/", "/", 0)
	funciones.EscribirSuperBlock(pa, &sblo, initPart)
	var resultado []byte
	copy(resultado[:], `1,G,root,root,123\n`)
	//startPoint := funciones.GetInicioParticion(disco, nomPar)	
	//funciones.CrearArchivosEscritos("/users.txt", true, resultado, 28, pa, nomPar, startPoint, &sblo)
	return nil
}

