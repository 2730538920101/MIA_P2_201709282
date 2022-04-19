package ast

type Parametro struct{
	Valor string
	Tipo string
}

func NewParametro(tipo string, val string) Parametro{
	e := Parametro{val, tipo}
	return e 
}