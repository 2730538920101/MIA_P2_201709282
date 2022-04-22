package ast

import (
	"./parametros"
	"fmt"
	"strconv"
	arrayList "github.com/colegno/arraylist"
	"../funciones"
	"os"
	"math/rand"
	"time"
	"./structs"
	"io"
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
	block := make([]byte, a)
	for j := 0; j < a; j++{
		block[j] = 0
	}
	disco, err := os.Create(d.path)
	if err != nil{
		fmt.Println(err)
	}
	disco.Write(block)
	disco.Close()
	
	//MBR
	tam := strconv.Itoa(a)
	tiemp := time.Now().Format("2006-01-02 15:04:05")
	rnd := strconv.Itoa(rand.Intn(501))
	ft := d.fit
	mbr := structs.NewMBR(tam, tiemp, rnd, ft)
	disco, err = os.OpenFile(string(d.path), os.O_RDWR, 0777)
	if err != nil{
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	newpos, err := disco.Seek(int64(0), os.SEEK_SET)
	if err != nil {
		fmt.Println("ERROR AL ABRIR EL ARCHIVO...")
	}
	mbrByte := funciones.Struct_to_bytes(mbr)
	
	_, _ = disco.WriteAt(mbrByte, newpos)
	
	lectura := make([]byte, a)
	_, err = disco.ReadAt(lectura, int64(0))
		if err != nil && err != io.EOF {
			fmt.Println("ERROR AL LEER EL ARCHIVO...")
		}
	str := funciones.Bytes_to_struct(lectura)
	fmt.Printf("TAMANO: %s\n",str.Mbr_tamano)
	fmt.Printf("FECHA DE CREACION: %s\n",str.Mbr_fecha_creacion)
	fmt.Printf("SIGNATURE: %s\n",str.Mbr_disk_signature)
	fmt.Printf("FIT: %s\n",str.Disk_fit)
	disco.Close()
	return nil


}




