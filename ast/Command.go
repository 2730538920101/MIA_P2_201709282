package ast

type Command interface{
	Ejecutar() interface{}
}