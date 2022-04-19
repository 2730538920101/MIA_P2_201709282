package ast

import(
	"fmt"
)

type Logout struct {
	r bool
}

func NewLogout(p bool) Logout{
	e := Logout{p}
	return e
}

func (d Logout) Ejecutar() interface{} {
	fmt.Println("EJECUTANDO LOGOUT...")
	return nil
}
