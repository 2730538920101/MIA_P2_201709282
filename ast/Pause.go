package ast

import(
	"fmt"
	"bufio"
	"os"

)

type Pause struct {
	r bool
}

func NewPause(p bool) Pause{
	e := Pause{p}
	return e
}

func (d Pause) Ejecutar() interface{} {
	fmt.Println("EJECUTANDO PAUSE PRESIONA ENTENER PARA SALIR...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	return nil
}