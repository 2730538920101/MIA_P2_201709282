package ast

import (
	"./parametros"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"../funciones"
	"strings"
	"./structs"
	"bytes"
	"strconv"
	"os"
	"os/exec"
)

type Rep struct{
	ruta string
	id string
	path string
	name string
}

func NewRep(lista *arrayList.List) Rep {
	aux := Rep{"","","",""}
	for i := 0; i < lista.Len(); i++{
		res := lista.GetValue(i)
		if res.(parametros.Parametro).Tipo == PATH{
			aux.path = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == ID{
			aux.id = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == RUTA{
			aux.ruta = res.(parametros.Parametro).Valor
		} else if res.(parametros.Parametro).Tipo == NAME{
			aux.name = res.(parametros.Parametro).Valor
		}else{
			fmt.Println("EL PARAMETRO INGRESADO NO PERTENECE AL COMANDO REP")
			aux.path = ""
			aux.ruta = ""
			aux.id = ""
			aux.name = ""
			break
		}
	}
	return aux
}

func (d Rep) Ejecutar() interface{}{
	fmt.Println("EJECUTANDO REP... ")
	if d.name == "" || d.path == "" || d.id == ""{
		fmt.Println("LOS PARAMETROS NAME, PATH & ID SON OBLIGATORIOS PARA EL COMANDO REP")
		return nil
	}
	path := funciones.GetPathOfPartById(d.id)
	if path == ""{
		fmt.Println("NO HA SIDO MONTADA LA PARTICION")
		return nil
	}
	var m structs.MasterBootRecord
	mbr := funciones.Struct_to_bytes(m)
	m = funciones.LeerMBR(path, len(mbr))
	if(m == structs.MasterBootRecord{}){
		fmt.Println("ERROR AL LEER EL DISCO")
		return nil
	}	
	funciones.CrearCarpetaReporte(d.path)
	if strings.EqualFold(d.name, "tree"){
		fmt.Println("CREANDO REPORTE TREE")
	}else if strings.EqualFold(d.name, "disk"){
		fmt.Println("CREANDO REPORTE DISK")
		auxtotal := string(bytes.Trim(m.Mbr_tamano[:],"\x00"))
		auxtotali,_ := strconv.Atoi(auxtotal)
		auxp1 := string(bytes.Trim(m.Mbr_partition[0].Part_size[:],"\x00"))
		auxp2 := string(bytes.Trim(m.Mbr_partition[1].Part_size[:],"\x00"))
		auxp3 := string(bytes.Trim(m.Mbr_partition[2].Part_size[:],"\x00"))
		auxp4 := string(bytes.Trim(m.Mbr_partition[3].Part_size[:],"\x00"))
		t1 := string(bytes.Trim(m.Mbr_partition[0].Part_type[:],"\x00"))
		t2 := string(bytes.Trim(m.Mbr_partition[1].Part_type[:],"\x00"))
		t3 := string(bytes.Trim(m.Mbr_partition[2].Part_type[:],"\x00"))
		t4 := string(bytes.Trim(m.Mbr_partition[3].Part_type[:],"\x00"))
		auxp1i,_ := strconv.Atoi(auxp1)
		auxp2i,_ := strconv.Atoi(auxp2)
		auxp3i,_ := strconv.Atoi(auxp3)
		auxp4i,_ := strconv.Atoi(auxp4)
		porp1 := (float64(auxp1i) * 100)/float64(auxtotali)
		porp2 := (float64(auxp2i) * 100)/float64(auxtotali)
		porp3 := (float64(auxp3i) * 100)/float64(auxtotali)
		porp4 := (float64(auxp4i) * 100)/float64(auxtotali)
		libre := auxtotali - (auxp1i + auxp2i + auxp3i + auxp4i)
		librest := strconv.Itoa(libre)
		plibre := (float64(libre) * 100)/float64(auxtotali)
		p1 := fmt.Sprintf("%.2v",porp1)
		p2 := fmt.Sprintf("%.2v",porp2)
		p3 := fmt.Sprintf("%.2v",porp3)
		p4 := fmt.Sprintf("%.2v",porp4)
		pl := fmt.Sprintf("%.2v", plibre)
		fmt.Printf("TAMAÑO DEL DISCO: %s\n",auxtotal)
		fmt.Printf("TAMAÑO DE PARTICION 1: %s\n PORCENTAJE: %s\n", auxp1,p1)
		fmt.Printf("TAMAÑO DE PARTICION 2: %s\n PORCENTAJE: %s\n", auxp2,p2)
		fmt.Printf("TAMAÑO DE PARTICION 3: %s\n PORCENTAJE: %s\n", auxp3,p3)
		fmt.Printf("TAMAÑO DE PARTICION 4: %s\n PORCENTAJE: %s\n", auxp4,p4)
		fmt.Printf("TAMAÑO DE ESPACIO LIBRE: %s\n PORCENTAJE: %s\n", librest, pl)
		CrearReporteDisk(auxp1,auxp2,auxp3,auxp4,auxtotal,p1,p2,p3,p4,t1,t2,t3,t4,librest,pl, d.path)
	}else if strings.EqualFold(d.name, "file"){
		if d.ruta == ""{
			fmt.Println("EL PARAMETRO RUTA ES OBLIGATORIO CON EL REPORTE FILE")
			return nil
		}
		fmt.Println("CREANDO REPORTE FILE...")
	}
	
	return nil
}

func CrearReporteDisk(t1, t2, t3, t4, tt, p1, p2, p3, p4, tp1, tp2, tp3, tp4, tl, pl, path string){
	disco, err := os.Create("/home/ec2-user/MIA_P2/reportesDot/reporte_disk.dot")
	if err != nil{
		fmt.Println(err)
	}
	archivo := ""
	archivo += "digraph G {\n"
	archivo +=	"parent[shape = plaintext label=<\n"
	archivo +=	"<table border='1' cellborder='1'>\n"
	archivo +=	"<tr>\n"
	archivo += "<td rowspan = \"2\" bgcolor =\"#dd8703\"> MBR <br/> TAMAÑO: "
	archivo += tt
	archivo += "<td/>\n"
	archivo += "<td rowspan = \"2\" bgcolor =\"#50b104\">"
	archivo += tp1 
	archivo += "<br/> TAMAÑO: "
	archivo += t1
	archivo += "<br/> PORCENTAJE: "
	archivo += p1
	archivo += "% </td>\n" 
	archivo	+= "<td rowspan = \"2\" bgcolor =\"#50b104\">"
	archivo += tp2 
	archivo += "<br/> TAMAÑO: "
	archivo += t2
	archivo += "<br/> PORCENTAJE: "
	archivo += p2
	archivo += "% </td>\n" 
	archivo += "<td rowspan = \"2\" bgcolor =\"#50b104\">"
	archivo += tp3 
	archivo += "<br/> TAMAÑO: "
	archivo += t3
	archivo += "<br/> PORCENTAJE: "
	archivo += p3
	archivo += "% </td>\n" 
	archivo += "<td rowspan = \"2\" bgcolor =\"#50b104\">"
	archivo += tp4 
	archivo += "<br/> TAMAÑO: "
	archivo += t4
	archivo += "<br/> PORCENTAJE: "
	archivo += p4
	archivo +="% </td>\n"
	archivo += "<td rowspan = \"2\" bgcolor =\"#50b104\"> LIBRES <br/> TAMAÑO: "
	archivo += tl
	archivo += "<br/> PORCENTAJE: "
	archivo += pl
	archivo += "% </td>\n"  
	archivo	+= "</tr>\n"
	archivo +=	"</table>\n"
	archivo += ">];\n"
	archivo += "}\n"
	disco.WriteString(archivo)
	cmd := exec.Command("dot", "-Tjpg", "reporte_disk.dot", "-o", path)
	err = cmd.Run()
	if err != nil {
		fmt.Println("ERROR AL CREAR EL REPORTE")
	}
	disco.Close()
}