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
	"time"
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

func CrearCarpetaReporte(rut string){
	carpeta := strings.Split(rut, "/")
	ruta := ""
	for _, carp := range carpeta{
		if carp == ""{
			continue
		}
		if !strings.HasSuffix(carp, ".jpg"){
			ruta = ruta + "/" + carp
			if _, err := os.Stat(ruta); os.IsNotExist(err){
				err = os.Mkdir(ruta,0777)
				if err != nil{
					panic(err)
				}
			}
		}else if !strings.HasSuffix(carp, ".png"){
			ruta = ruta + "/" + carp
			if _, err := os.Stat(ruta); os.IsNotExist(err){
				err = os.Mkdir(ruta,0777)
				if err != nil{
					panic(err)
				}
			}
		}else if !strings.HasSuffix(carp, ".pdf"){
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
func Bytes_to_struct4(s []byte) structs.TablaInodos {
	p := structs.TablaInodos{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("ERROR DE CONVERSION DE BYTE A STRUCT... ")
	}
	return p
}

func Bytes_to_struct6(s []byte) structs.Content{
	p := structs.Content{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("ERROR DE CONVERSION DE BYTE A STRUCT... ")
	}
	return p
}

func Bytes_to_struct7(s []byte) structs.BlockFile{
	p := structs.BlockFile{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("ERROR DE CONVERSION DE BYTE A STRUCT... ")
	}
	return p
}

func Bytes_to_struct5(s []byte) structs.BlockDirectory {
	p := structs.BlockDirectory{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("ERROR DE CONVERSION DE BYTE A STRUCT... ")
	}
	return p
}

func Bytes_to_struct3(s []byte) structs.SuperBlock {
	p := structs.SuperBlock{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("ERROR DE CONVERSION DE BYTE A STRUCT... ")
	}
	return p
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

func ReadSuperBlock(path string, name string, startsuperb* int) structs.SuperBlock{
	var disco structs.MasterBootRecord
	discob := Struct_to_bytes(disco)
	disco = LeerMBR(path, len(discob))
	if(disco == structs.MasterBootRecord{}){
		fmt.Println("ERROR AL LEER EL DISCO")
		return structs.SuperBlock{}
	}
	fmt.Printf("entro %d\n", &startsuperb)
	fmt.Printf("funcion %d\n", GetInicioParticion(disco, name))
	*startsuperb = GetInicioParticion(disco, name)
	fmt.Printf("entro %d\n", &startsuperb)
	if *startsuperb == -1{
		return structs.SuperBlock{}
	}
	var sbl structs.SuperBlock 
	sbl = LeerSuperBlock(path, int(unsafe.Sizeof(sbl)), *startsuperb)
	return sbl
}

func LeerSuperBlock(path string, tam int, start int) structs.SuperBlock{
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	lectura := make([]byte, tam)
	_, err = disco.ReadAt(lectura, int64(start))
	if err != nil && err != io.EOF {
		fmt.Println("ERROR AL LEER EL ARCHIVO...")
	}
	str := Bytes_to_struct3(lectura)
	fmt.Printf("S_filesystem_type: %s\n",str.S_filesystem_type)
	fmt.Printf("S_inodes_count: %s\n",str.S_inodes_count)
	fmt.Printf("S_blocks_count: %s\n",str.S_blocks_count)
	fmt.Printf("S_free_block_count: %s\n",str.S_free_blocks_count)
	fmt.Printf("S_free_inodes_count: %s\n",str.S_free_inodes_count)
	fmt.Printf("S_mtime: %s\n", str.S_mtime)
	fmt.Printf("S_mnt_count: %s\n", str.S_mnt_count)
	fmt.Printf("S_magic: %s\n", str.S_magic)
	fmt.Printf("S_inode_size: %s\n", str.S_inode_size)
	fmt.Printf("S_block_size: %s\n", str.S_block_size)
	fmt.Printf("S_first_ino: %s\n", str.S_first_ino)
	fmt.Printf("S_first_blo: %s\n", str.S_first_blo)
	fmt.Printf("S_bm_inode_start: %s\n", str.S_bm_inode_start)
	fmt.Printf("S_bm_block_start: %s\n", str.S_bm_block_start)
	fmt.Printf("S_inode_start: %s\n", str.S_inode_start)
	fmt.Printf("S_block_start: %s\n", str.S_block_start)
	fmt.Println()
	disco.Close()
	return str
}

func LeerInodo(path string, tam int, start int) structs.TablaInodos{
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	lectura := make([]byte, tam)
	_, err = disco.ReadAt(lectura, int64(start))
	if err != nil && err != io.EOF {
		fmt.Println("ERROR AL LEER EL ARCHIVO...")
	}
	str := Bytes_to_struct4(lectura)
	fmt.Printf("I_uid: %s\n",str.I_uid)
	fmt.Printf("I_gid: %s\n",str.I_gid)
	fmt.Printf("I_size: %s\n",str.I_size)
	fmt.Printf("I_atime: %s\n",str.I_atime)
	fmt.Printf("I_ctime: %s\n",str.I_ctime)
	fmt.Printf("I_mtime: %s\n", str.I_mtime)
	for i := 0; i < 15; i++{
		fmt.Printf("I_block %d: %d\n",i, str.I_block[i])
	}
	fmt.Printf("I_type: %s\n", str.I_type)
	fmt.Printf("I_perm: %s\n", str.I_perm)
	fmt.Println()
	disco.Close()
	return str
}

func LeerBlockDirectory(path string, tam int, start int) structs.BlockDirectory{
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	lectura := make([]byte, tam)
	_, err = disco.ReadAt(lectura, int64(start))
	if err != nil && err != io.EOF {
		fmt.Println("ERROR AL LEER EL ARCHIVO...")
	}
	str := Bytes_to_struct5(lectura)
	for i := 0; i < 4; i++{
		aux := Bytes_to_struct6(Struct_to_bytes(str.B_content[i]))
		fmt.Printf("B_name: %s\n", aux.B_name)
		fmt.Printf("B_inodo: %s\n", aux.B_inodo)
	}
	fmt.Println()
	disco.Close()
	return str
}

func LeerBlockFile(path string, tam int, start int) structs.BlockFile{
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	lectura := make([]byte, tam)
	_, err = disco.ReadAt(lectura, int64(start))
	if err != nil && err != io.EOF {
		fmt.Println("ERROR AL LEER EL ARCHIVO...")
	}
	str := Bytes_to_struct7(lectura)
	fmt.Printf("EL CONTENIDO DEL ARCHIVO ES: %s", string(bytes.Trim(str.B_content[:],"\x00")))
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

func EscribirSuperBlock(path string, sb* structs.SuperBlock, start int){
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(start), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBIACION DEL ARCHIVO...")
	}
	mbrByte := Struct_to_bytes(sb)
	_, _ = disco.WriteAt(mbrByte, newpos)
	disco.Close()
}

func EscribirInodo(path string, ino* structs.TablaInodos, start int){
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(start), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBIACION DEL ARCHIVO...")
	}
	mbrByte := Struct_to_bytes(ino)
	_, _ = disco.WriteAt(mbrByte, newpos)
	disco.Close()
}

func EscribirBlockDirectory(path string, bld* structs.BlockDirectory, start int){
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(start), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBIACION DEL ARCHIVO...")
	}
	mbrByte := Struct_to_bytes(bld)
	_, _ = disco.WriteAt(mbrByte, newpos)
	disco.Close()
}

func EscribirBlockFile(path string, blf* structs.BlockFile, start int){
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(start), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBIACION DEL ARCHIVO...")
	}
	mbrByte := Struct_to_bytes(blf)
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
				*inicio = nxtint 
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

func GetStartOfPartById(iden string) int{
	for j := 0; j < 27; j++{
		nomb := GetNombrePath(globales.Disk_mont[j].Path)
		if nomb != ""{
			for k := 0; k < 60; k++{
				pm := globales.Disk_mont[j].Particiones_montadas[k].Id
				if pm != ""{
					if iden == pm{
						return globales.Disk_mont[j].Particiones_montadas[k].Inicio
					}
				}
			}
		}
	}
	return -1
}

func GetPathOfPartById(iden string) string{
	for j := 0; j < 27; j++{
		nomb := GetNombrePath(globales.Disk_mont[j].Path)
		if nomb != ""{
			for k := 0; k < 60; k++{
				pm := globales.Disk_mont[j].Particiones_montadas[k].Id
				if pm != ""{
					if iden == pm{
						return globales.Disk_mont[j].Path
					}
				}
			}
		}
	}
	return ""
}

func GetNameOfPartById(iden string) string{
	for j := 0; j < 27; j++{
		nomb := GetNombrePath(globales.Disk_mont[j].Path)
		if nomb != ""{
			for k := 0; k < 60; k++{
				pm := globales.Disk_mont[j].Particiones_montadas[k].Id
				if pm != ""{
					if iden == pm{
						return globales.Disk_mont[j].Particiones_montadas[k].Name
					}
				}
			}
		}
	}
	return ""
}

func GetInicioParticion(disco structs.MasterBootRecord, nombre string) int{
	for i := 0; i < 4; i++{
		partst := string(bytes.Trim(disco.Mbr_partition[i].Part_status[:],"\x00"))
		partnom := string(bytes.Trim(disco.Mbr_partition[i].Part_name[:],"\x00"))
		partstpt, _ := strconv.Atoi(string(bytes.Trim(disco.Mbr_partition[i].Part_start[:],"\x00")))
		if partst == "ACTIVO"{
			if partnom == nombre{
				return partstpt
			}
		}
	}
	return -1
}

func GetTamanoParticion(disco structs.MasterBootRecord, nombre string) int{
	for i := 0; i < 4; i++{
		partst := string(bytes.Trim(disco.Mbr_partition[i].Part_status[:],"\x00"))
		partnom := string(bytes.Trim(disco.Mbr_partition[i].Part_name[:],"\x00"))
		partstpt, _ := strconv.Atoi(string(bytes.Trim(disco.Mbr_partition[i].Part_size[:],"\x00")))
		if partst == "ACTIVO"{
			if partnom == nombre{
				return partstpt
			}
		}
	}
	return -1
}

func CrearBitMapTable(numInodos, primero int, path string){
	cont := 0
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(primero), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBICACION DEL ARCHIVO...")
	}
	for cont < numInodos{
		_, _ = disco.WriteAt([]byte("0"), newpos+int64(cont))
		cont++
	}
	
	disco.Close()
}

func LlenarCeros(numCeros, primero int, path string){
	cont := 0
	disco, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(primero), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL LEER UBICACION DEL ARCHIVO...")
	}
	for cont < numCeros{
		_, _ = disco.WriteAt([]byte("0"), newpos+int64(cont))
		cont++
	}
	
	disco.Close()
}

func GetBloqueInicial(bloque* structs.SuperBlock, indice* int) int{
	aux1,_ := strconv.Atoi(string(bytes.Trim(bloque.S_block_start[:],"\x00")))
	aux2,_ := strconv.Atoi(string(bytes.Trim(bloque.S_block_size[:],"\x00")))
	return aux1 + (aux2 * *indice)
}

func GetInodoInicial(bloque* structs.SuperBlock, indice* int) int{
	aux1,_ := strconv.Atoi(string(bytes.Trim(bloque.S_inode_start[:],"\x00")))
	aux2,_ := strconv.Atoi(string(bytes.Trim(bloque.S_inode_size[:],"\x00")))
	return aux1 + (aux2 * *indice)
}

func LimpiarInodo(inodo* structs.TablaInodos){
	if (inodo != &structs.TablaInodos{}){
		for i:=0; i<15; i++{
			inodo.I_block[i] = -1
		}
	}
}

func GetBloqueCarpetasLibre(nomDir string, path string, bloque* structs.SuperBlock, indiceBloqueActual* int, indiceInodoActual* int, indiceLibre* int){
	var inodo structs.TablaInodos
	ino := Struct_to_bytes(inodo)
	aux :=  GetInodoInicial(bloque, indiceInodoActual)
	inodo = LeerInodo(path, len(ino), aux)
	if (inodo == structs.TablaInodos{}){
		fmt.Println("NO SE HA CREADO LA CARPETA")
		return
	}
	var idPoint int = 0
	var bld structs.BlockDirectory
	var encontrado bool = false
	for idPoint < 15 && !encontrado{
		if inodo.I_block[idPoint] != -1{
			aux := inodo.I_block[idPoint]
			var ptaux* int = &aux
			bld2 := Struct_to_bytes(bld)
			bld = LeerBlockDirectory(path, len(bld2), GetBloqueInicial(bloque, ptaux))
			if string(bytes.Trim(bld.B_content[0].B_name[:],"\x00")) == nomDir{
				*indiceBloqueActual = inodo.I_block[idPoint]
				indiceBloque := 2
				for indiceBloque < 4{
					if string(bytes.Trim(bld.B_content[indiceBloque].B_inodo[:],"\x00")) == "-1"{
						*indiceLibre = indiceBloque
						encontrado = true
						break
					}
					indiceBloque++
				}
			}
		}
		idPoint++
	}
}

func CrearBloqueCarpetas(name, indiceDir, namePad, indicePad string) structs.BlockDirectory{
	var block structs.BlockDirectory
	copy(block.B_content[0].B_name[:], name)
	copy(block.B_content[0].B_inodo[:], indiceDir)
	copy(block.B_content[1].B_name[:], namePad)
	copy(block.B_content[1].B_name[:], indicePad)
	return block
}


func CrearDirectorio(bloque* structs.SuperBlock, path, nomDir, nomPad string, indicePad int) int{
	tiemp := time.Now().Format("2006-01-02 15:04:05")
	var aux structs.TablaInodos = structs.NewTablaInodos("1","1","-1", tiemp, tiemp, tiemp ,"DIRECTORIO","777")
	indiceI := string(bytes.Trim(bloque.S_first_ino[:],"\x00"))
	indicePadst := strconv.Itoa(indicePad)
	var auxbl structs.BlockDirectory = CrearBloqueCarpetas(nomDir, indiceI, nomPad, indicePadst)
	ax,_ := strconv.Atoi(string(bytes.Trim(bloque.S_first_blo[:],"\x00")))
	aux.I_block[0] = ax
	indiceIi,_ := strconv.Atoi(indiceI)
	EscribirInodo(path, &aux, GetInodoInicial(bloque, &indiceIi))
	fsa,_ := strconv.Atoi(string(bytes.Trim(bloque.S_first_blo[:],"\x00")))
	EscribirBlockDirectory(path, &auxbl, GetBloqueInicial(bloque, &fsa))
	auxc1,_:= strconv.Atoi(string(bytes.Trim(bloque.S_free_blocks_count[:],"\x00")))
	auxc2,_:= strconv.Atoi(string(bytes.Trim(bloque.S_free_inodes_count[:],"\x00")))
	ac1 := auxc1 - 1
	ac2 := auxc2 - 1 
	copy(bloque.S_free_blocks_count[:], strconv.Itoa(ac1))
	copy(bloque.S_free_inodes_count[:], strconv.Itoa(ac2))
	return indiceIi
}

func CrearNodoCarpeta(dirPad, dirName, path string, bloque* structs.SuperBlock, indiceInodoPadre* int, indiceBloqueActual* int){
	indiceLibre := -1
	GetBloqueCarpetasLibre(dirPad, path, bloque, indiceBloqueActual, indiceInodoPadre, &indiceLibre)
	if indiceLibre == -1{
		fmt.Println("ERROR AL CREAR CARPETA")
		return
	}
	var auxbl structs.BlockDirectory
	var newIndice int = CrearDirectorio(bloque, path, dirName, dirPad, *indiceInodoPadre)
	var bl structs.BlockDirectory = LeerBlockDirectory(path,int(unsafe.Sizeof(auxbl)),  GetBloqueInicial(bloque, indiceBloqueActual)) 
	newIndst := strconv.Itoa(newIndice)
	copy(bl.B_content[indiceLibre].B_inodo[:],newIndst)
	copy(bl.B_content[indiceLibre].B_name[:], dirName)
	EscribirBlockDirectory(path, &bl, GetBloqueInicial(bloque, indiceBloqueActual))
	*indiceInodoPadre = newIndice
}

func BuscarCarpeta(nameDir, path string, indiceInodoActual* int, bloque* structs.SuperBlock) int{
	var inodo structs.TablaInodos
	inodo = LeerInodo(path, int(unsafe.Sizeof(inodo)), GetInodoInicial(bloque, indiceInodoActual))
	if (inodo == structs.TablaInodos{}){
		return -1
	}
	var indexBlock int
	var bl structs.BlockDirectory
	for indexBlock = 0; indexBlock < 15; indexBlock++{
		if inodo.I_block[indexBlock] != -1{
			if string(bytes.Trim(inodo.I_type[:],"\x00")) == "DIRECTORIO"{
				auxb := inodo.I_block[indexBlock]
				bl = LeerBlockDirectory(path, int(unsafe.Sizeof(bl)), GetBloqueInicial(bloque, &auxb))
				if (bl == structs.BlockDirectory{}){
					return -1
				}
				var i int
				for i = 2; i < 4; i++{
					if string(bytes.Trim(bl.B_content[i].B_inodo[:],"\x00")) == "-1"{
						if string(bytes.Trim(bl.B_content[i].B_name[:],"\x00")) == nameDir{
							aux2, _ := strconv.Atoi(string(bytes.Trim(bl.B_content[i].B_inodo[:],"\x00")))
							*indiceInodoActual = aux2
							inodo = LeerInodo(path, int(unsafe.Sizeof(inodo)), GetInodoInicial(bloque, indiceInodoActual))
							if(inodo == structs.TablaInodos{}){
								return -1
							}else{
								return inodo.I_block[0]
							}
						}
					}
				}
			}
		}
	}
	return -1
}


func CrearNodoArchivo(size int, text []byte, path string, dirpad, name string, bloque* structs.SuperBlock, indiceBloqueActual, indiceInodoPadre int){
	var indiceLibre int = -1
	GetBloqueCarpetasLibre(dirpad, path, bloque, &indiceBloqueActual, &indiceInodoPadre, &indiceLibre)
	tiemp := time.Now().Format("2006-01-02 15:04:05")
	var inodo structs.TablaInodos = structs.NewTablaInodos("1","1",strconv.Itoa(size), tiemp, tiemp, tiemp ,"DIRECTORIO","664")
	var bl structs.BlockFile = structs.NewBlockFile("")
	var indiceInodo int = 0
	var indiceCar int = 0
	var contCar int = 0
	var guardado bool = true
	for guardado{
		if contCar >=64 || indiceCar >= size{
			if indiceInodo < 15{
				a, _ := strconv.Atoi(string(bytes.Trim(bloque.S_first_blo[:],"\x00"))) 
				inodo.I_block[indiceInodo] = a 
				GuardarBloqueArchivos(&bl, bloque, path)
				indiceInodo++
			}
			bl = structs.NewBlockFile("")
			contCar = 0;
		}
		bl.B_content[contCar] = text[indiceCar]
		if indiceCar >= size{
			break
		}
		contCar++
		indiceCar++
	}
	b,_ := strconv.Atoi(string(bytes.Trim(bloque.S_first_ino[:],"\x00"))) 
	EscribirInodo(path, &inodo, GetInodoInicial(bloque, &b))
	var aux_carpetas structs.BlockDirectory 
	aux2 := Struct_to_bytes(aux_carpetas)
	aux_carpetas = LeerBlockDirectory(path, len(aux2), GetBloqueInicial(bloque, &indiceBloqueActual))
	s:= string(bytes.Trim(bloque.S_first_ino[:],"\x00"))
	copy(aux_carpetas.B_content[indiceLibre].B_inodo[:], s)
	copy(aux_carpetas.B_content[indiceLibre].B_name[:], name)
	EscribirBlockDirectory(path, &aux_carpetas, GetBloqueInicial(bloque, &indiceBloqueActual))
	v,_ := strconv.Atoi(string(bytes.Trim(bloque.S_free_inodes_count[:],"\x00")))
	v--
	m := strconv.Itoa(v)
	copy(bloque.S_free_inodes_count[:], m) 
}

func GuardarBloqueArchivos(bloqueA* structs.BlockFile, bloqueS* structs.SuperBlock, path string){
	aux,_ :=strconv.Atoi(string(bytes.Trim(bloqueS.S_first_blo[:],"\x00")))
	EscribirBlockFile(path, bloqueA, GetBloqueInicial(bloqueS, &aux))
	at,_ := strconv.Atoi(string(bytes.Trim(bloqueS.S_free_blocks_count[:],"\x00")))
	at--
	fn := strconv.Itoa(at)
	copy(bloqueS.S_free_blocks_count[:], fn)
}

func CrearArchivosEscritos(newPath string, createPath bool, text []byte, size int, path string, namePartition string, startp int, sblo* structs.SuperBlock){
	var indexInodoPadre int = 0
	var dirPad string = "/"
	var token string
	var indexBloqueActual int = 0
	aux_Np := strings.Split(newPath, "/")
	for i:=0; i < len(aux_Np); i++{
		token = aux_Np[i]
		if i == len(aux_Np) -1{
			CrearNodoArchivo(size, text, path, dirPad, token, sblo, indexBloqueActual, indexInodoPadre)
		}else{
			var indexBloque int = BuscarCarpeta(token, path, &indexInodoPadre, sblo)
			if indexBloque != -1{
				indexBloqueActual = indexBloque
			}else{
				if createPath{
					CrearNodoCarpeta(dirPad, token, path, sblo, &indexInodoPadre, &indexBloqueActual)
				}else{
					fmt.Println("ERROR EL DIRECTORIO NO EXISTE")
					return
				}
			}
			dirPad = token
		}
	}
	EscribirSuperBlock(path, sblo, startp)
}

