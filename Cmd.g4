grammar Cmd;

options {
        tokenVocab = CmdLex;
}

@header{
        import "../ast"
        import arrayList "github.com/colegno/arraylist"
}

// PRODUCCIONES
start returns [ast.Ast ast] 
           : comandList{$ast = ast.NewAst($comandList.lista)};

comandList returns[*arrayList.List lista] @init{ $lista = arrayList.New()}
           : AUXCLIST = comandList comando {
                        $AUXCLIST.lista.Add($comando.com)
                        $lista = $AUXCLIST.lista
                   }
           | comando {
                        $lista.Add($comando.com)
                   }
           ;

comando returns [ast.Command com]
               :  TOK_MKDISK param_list 
               |  TOK_RMDISK param_list 
               |  TOK_FDISK param_list 
               |  TOK_MOUNT param_list 
               |  TOK_MKFS param_list 
               |  TOK_LOGIN param_list 
               |  TOK_MKGRP param_list 
               |  TOK_RMGRP param_list 
               |  TOK_MKUSR param_list 
               |  TOK_RMUSR param_list 
               |  TOK_MKFILE param_list 
               |  TOK_MKDIR param_list 
               |  TOK_EXEC param_list 
               |  TOK_REP param_list 
               |  TOK_PAUSE 
               |  TOK_LOGOUT  
               ;

param_list returns [*arrayList.List lista] @init{$lista = arrayList.New()}
           :AUXPLIST = param_list param {
                   $AUXPLIST.lista.Add($param.par)
                   $lista = $AUXPLIST.lista
                }    
           | param {
                   $lista.Add($param.par)
                }
           ;

param returns [ast.Parametro par]
      :  TOK_SIZE TOK_IGUAL  TOK_NUMERO 
      |  TOK_PATH TOK_IGUAL  TOK_CAMINO 
      |  TOK_PATH TOK_IGUAL  TOK_CADENA 
      |  TOK_FIT TOK_IGUAL  TOK_FIRST 
      |  TOK_FIT TOK_IGUAL  TOK_WORST 
      |  TOK_FIT TOK_IGUAL  TOK_BEST 
      |  TOK_UNIT TOK_IGUAL  TOK_BYTES 
      |  TOK_UNIT TOK_IGUAL  TOK_KB 
      |  TOK_UNIT TOK_IGUAL  TOK_MB 
      |  TOK_NAME TOK_IGUAL  TOK_PALABRA 
      |  TOK_NAME TOK_IGUAL  TOK_CADENA 
      |  TOK_USUARIO TOK_IGUAL  TOK_PALABRA 
      |  TOK_USUARIO TOK_IGUAL  TOK_CADENA 
      |  TOK_GRP TOK_IGUAL  TOK_PALABRA 
      |  TOK_GRP TOK_IGUAL  TOK_CADENA 
      |  TOK_PASSWORD TOK_IGUAL  TOK_PALABRA 
      |  TOK_PWD TOK_IGUAL  TOK_PALABRA 
      |  TOK_TYPE TOK_IGUAL  TOK_PRIMARIA 
      |  TOK_TYPE TOK_IGUAL  TOK_LOGICA 
      |  TOK_TYPE TOK_IGUAL  TOK_EXTENDIDA 
      |  TOK_TYPE TOK_IGUAL  TOK_FAST 
      |  TOK_TYPE TOK_IGUAL  TOK_FULL 
      |  TOK_ID TOK_IGUAL  TOK_IDENTIFICADOR 
      |  TOK_CONT TOK_IGUAL  TOK_CADENA 
      |  TOK_CONT TOK_IGUAL  TOK_CAMINO 
      |  TOK_RUTA TOK_IGUAL  TOK_CADENA 
      |  TOK_RUTA TOK_IGUAL  TOK_CAMINO 
      |  TOK_P 
      |  TOK_R 
      ;



