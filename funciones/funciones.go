package funciones

import (
	"strings"
	"os"
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
		default:
			{
				return tam * (1024*1024)
			}
	}
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
