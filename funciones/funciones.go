package funciones

import (
	"strings"
	"os"
	"bytes"
	"encoding/gob"
	"io"
	"fmt"
	"../ast/structs"
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