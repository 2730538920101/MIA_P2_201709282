package funciones

import (
	"strings"
	"os"
	"bytes"
	"encoding/gob"
	"io"
	"fmt"
	"../ast/structs"
	"strconv"
	"unsafe"
	"../globales"
)

func TamByte(unidad string, tam int) int{
	switch unidad{
		case "B":
			{
				return tam
			}
		case "K":
			{
				return tam * 1024
			}
		case "M":
			{
				return tam * (1024*1024)
			}
	}
	return 0
}

func CrearCarpetaDisco(rut string){
	carpeta := strings.Split(rut, "/")
	ruta := ""
	for _, carp := range carpeta{
		if carp == ""{
			continue
		}
		if !strings.HasSuffix(carp, ".dk"){
			ruta = ruta + "/" + carp
			if _, err := os.Stat(ruta); os.IsNotExist(err){
				err = os.Mkdir(ruta,0777)
				if err != nil{
					panic(err)
				}
			}
		}
	}
}

func Struct_to_bytes(p interface{})[]byte{
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil && err == io.EOF {
		fmt.Println("ERROR DE CONVERSION DE STRUCT A BYTE... ")
	}
	return buf.Bytes()
}

func Bytes_to_struct2(s []byte) structs.ExtendedBootRecord {
	p := structs.ExtendedBootRecord{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("ERROR DE CONVERSION DE BYTE A STRUCT... ")
	}
	return p
}

func Bytes_to_struct(s []byte) structs.MasterBootRecord {
	p := structs.MasterBootRecord{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("ERROR DE CONVERSION DE BYTE A STRUCT... ")
	}
	return p
}

func ArchivoExiste(ruta string) bool{
	if _, err := os.Stat(ruta);os.IsNotExist(err){
		return false
	}
	return true
}

func LeerMBR(path string, tam int) structs.MasterBootRecord{
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	lectura := make([]byte, tam)
	_, err = disco.ReadAt(lectura, int64(0))
	if err != nil && err != io.EOF {
		fmt.Println("ERROR AL LEER EL ARCHIVO...")
	}
	str := Bytes_to_struct(lectura)
	fmt.Printf("TAMAÑO: %s\n",str.Mbr_tamano)
	fmt.Printf("FECHA DE CREACION: %s\n",str.Mbr_fecha_creacion)
	fmt.Printf("SIGNATURE: %s\n",str.Mbr_disk_signature)
	fmt.Printf("FIT: %s\n",str.Disk_fit)
	for i := 0; i <4; i++{
		fmt.Println()
		fmt.Printf("PART %d\n", i+1)
		fmt.Println()
		fmt.Printf("STATUS: %s\n", str.Mbr_partition[i].Part_status)
		fmt.Printf("TYPE: %s\n", str.Mbr_partition[i].Part_type)
		fmt.Printf("FIT: %s\n", str.Mbr_partition[i].Part_fit)
		fmt.Printf("START: %s\n", str.Mbr_partition[i].Part_start)
		fmt.Printf("SIZE: %s\n", str.Mbr_partition[i].Part_size)
		fmt.Printf("NAME: %s\n", str.Mbr_partition[i].Part_name)
		fmt.Println()
	}
	disco.Close()
	return str
}

func LeerEBR(path string, tam int, start int) structs.ExtendedBootRecord{
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	lectura := make([]byte, tam)
	_, err = disco.ReadAt(lectura, int64(start))
	if err != nil && err != io.EOF {
		fmt.Println("ERROR AL LEER EL ARCHIVO...")
	}
	str := Bytes_to_struct2(lectura)
	fmt.Printf("ESTADO: %s\n",str.Part_status)
	fmt.Printf("FIT: %s\n",str.Part_fit)
	fmt.Printf("START: %s\n",str.Part_start)
	fmt.Printf("SIZE: %s\n",str.Part_size)
	fmt.Printf("NEXT: %s\n",str.Part_next)
	fmt.Printf("NAME: %s\n", str.Part_name)
	fmt.Println()
	disco.Close()
	return str
}

func EscribirMBR(path string, master* structs.MasterBootRecord){
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(0), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBICACION DEL ARCHIVO...")
	}
	mbrByte := Struct_to_bytes(master)
	
	_, _ = disco.WriteAt(mbrByte, newpos)
	disco.Close()
}

func EscribirEBR(path string, part_ext* structs.ExtendedBootRecord, start int){
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(start), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBIACION DEL ARCHIVO...")
	}
	mbrByte := Struct_to_bytes(part_ext)
	_, _ = disco.WriteAt(mbrByte, newpos)
	disco.Close()
}

func GetDireccionInicio(master structs.MasterBootRecord, fit string, tam int, inicio* int, pos int, ebr structs.ExtendedBootRecord, logica bool){
	*inicio = -1
	if logica {
		if(ebr != structs.ExtendedBootRecord{}){
			stpt, _ := strconv.Atoi(string(bytes.Trim(ebr.Part_start[:],"\x00"))) 
			sipt, _ := strconv.Atoi(string(bytes.Trim(ebr.Part_size[:],"\x00")))
			nxt := string(bytes.Trim(ebr.Part_next[:],"\x00"))
			nxtint, _ := strconv.Atoi(nxt) 
			if nxt == "-1"{
				*inicio = stpt + sipt
			}else{
				*inicio = nxtint +sipt
			}
		}else{
			dir, _ := strconv.Atoi(string(bytes.Trim(master.Mbr_partition[pos].Part_start[:], "\x00")))
			*inicio = dir 
		}
		
	}else{
		dtam := int(unsafe.Sizeof(master))
		fmt.Printf("TAMAÑO EN BYTES DEL MBR: %d\n", dtam)
		dt := string(bytes.Trim(master.Mbr_tamano[:], "\x00"))
		dtint, _ := strconv.Atoi(dt)
		fmt.Printf("TAMAÑO DEL DISCO: %d\n", dtint)
		fmt.Printf("TAMAÑO EN BYTES DE LA PARTICION: %d\n", tam)
		switch fit{
			case "FF":
				{
					fmt.Println("SE REALIZARA EL PRIMER AJUSTE")
					break
				}
			case "WF":
				{
					fmt.Println("SE REALIZARA EL PEOR AJUSTE")
					break
				}
			case "BF":
				{
					fmt.Println("SE REALIZARA EL MEJOR AJUSTE")
					break
				}
		}
		if pos == 0{
			*inicio = dtam 
		}else{
			partsize, _ := strconv.Atoi(string(bytes.Trim(master.Mbr_partition[pos - 1].Part_size[:],"\x00")))
			partst, _ := strconv.Atoi(string(bytes.Trim(master.Mbr_partition[pos - 1].Part_start[:],"\x00")))
			*inicio = partsize + partst  
		}	
	}
	if(*inicio == -1){
		fmt.Println("ERROR, ESPACIO INSUFICIENTE PARA CREAR UNA PARTICION")
	}
}

func GetPrimerEBR(path string, pos int, t int) structs.ExtendedBootRecord{
	disco, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("ERROR AL ABRIR EL ARCHIVO")
	}
	lectura := make([]byte, t)
	_, err = disco.ReadAt(lectura, int64(pos))
	if err != nil && err != io.EOF {
		fmt.Println("ERROR AL LEER EL ARCHIVO")
	}
	ejm := Bytes_to_struct2(lectura)
	return ejm
}

func GetNombrePath(rut string) string{
	carpeta := strings.Split(rut, "/")
	for _, carp := range carpeta{
		if carp == ""{
			continue
		}
		if strings.HasSuffix(carp, ".dk"){
			carpt := strings.Split(carp, ".")
			return carpt[0]
		}
	}
	return ""
}

func GetPartId(letra string, numero int) string{
	var st string = "82"
	nu := strconv.Itoa(numero)
	st = st + nu
	st = st + string(letra)
	return st
}

func ShowMounts(){
	for j := 0; j < 27; j++{
		nomb := GetNombrePath(globales.Disk_mont[j].Path)
		if nomb != ""{
			fmt.Printf("DISCO MONTADO %d: %s\n", j+1, nomb)
			for k := 0; k < 60; k++{
				pm := globales.Disk_mont[j].Particiones_montadas[k].Id
				if pm != ""{
					fmt.Printf("PARTICION MONTADA ID: %s\n", pm)
					fmt.Println()
				}
			}
		}
	}
}