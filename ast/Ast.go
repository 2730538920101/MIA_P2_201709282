package ast

import arrayList "github.com/colegno/arraylist"

type Ast struct{
	ComList *arrayList.List
}

func NewAst(lista *arrayList.List) Ast{
	ast := Ast{ComList: lista}
	return ast

}