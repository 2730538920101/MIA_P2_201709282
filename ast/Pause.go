package ast

import(
	"fmt"
)

type Pause struct {
	r bool
}

func NewPause(p bool) Pause{
	e := Pause{p}
	return e
}

func (d Pause) Ejecutar() interface{} {
	fmt.Println("EJECUTANDO PAUSE...")
	return nil
}