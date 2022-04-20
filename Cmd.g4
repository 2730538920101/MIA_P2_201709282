grammar Cmd;

options {
        tokenVocab = CmdLex;
}

@header{
        import "../ast"
        import "../ast/parametros"
        import arrayList "github.com/colegno/arraylist"
        import "strings"
}

// PRODUCCIONES
start returns [ast.Ast ast] 
           : comandList{$ast = ast.NewAst($comandList.lista)} EOF;

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
               :  TOK_MKDISK param_list {
                       $com = ast.NewMkdisk($param_list.lista)
                }
               |  TOK_RMDISK param_list {
                       $com = ast.NewRmdisk($param_list.lista)
                }
               |  TOK_FDISK param_list {
                       $com = ast.NewFdisk($param_list.lista)
                }
               |  TOK_MOUNT param_list {
                       $com = ast.NewMount($param_list.lista)
                }
               |  TOK_MKFS param_list {
                       $com = ast.NewMkfs($param_list.lista)
                }
               |  TOK_LOGIN param_list {
                       $com = ast.NewLogin($param_list.lista)
                } 
               |  TOK_MKGRP param_list {
                       $com = ast.NewMkgrp($param_list.lista)
                }
               |  TOK_RMGRP param_list {
                       $com = ast.NewRmgrp($param_list.lista)
                }
               |  TOK_MKUSR param_list {
                       $com = ast.NewMkusr($param_list.lista)
                }
               |  TOK_RMUSR param_list {
                       $com = ast.NewRmusr($param_list.lista)
                }
               |  TOK_MKFILE param_list {
                       $com = ast.NewMkfile($param_list.lista)
                }
               |  TOK_MKDIR param_list {
                       $com = ast.NewMkdir($param_list.lista)
                }
               |  TOK_EXEC param_list {
                       $com = ast.NewExec($param_list.lista)
                }
               |  TOK_REP param_list {
                       $com = ast.NewRep($param_list.lista)
                }
               |  TOK_PAUSE {
                       $com = ast.NewPause(true)
                }
               |  TOK_LOGOUT {
                       $com = ast.NewLogout(true)
                }
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

param returns [parametros.Parametro par]
      :  TOK_SIZE TOK_IGUAL TOK_NUMERO {
              $par = parametros.NewParametro(ast.SIZE, $TOK_NUMERO.text)
        }
      |  TOK_PATH TOK_IGUAL TOK_CAMINO {
              $par = parametros.NewParametro(ast.PATH, $TOK_CAMINO.text)
        }
      |  TOK_PATH TOK_IGUAL TOK_CADENA {
              str := strings.Trim($TOK_CADENA.text, "\"")
              $par = parametros.NewParametro(ast.PATH, str)
        }
      |  TOK_FIT TOK_IGUAL TOK_FIRST {
              $par = parametros.NewParametro(ast.FIT, ast.FF)
        }
      |  TOK_FIT TOK_IGUAL TOK_WORST {
              $par = parametros.NewParametro(ast.FIT, ast.WF)
        }
      |  TOK_FIT TOK_IGUAL TOK_BEST {
              $par = parametros.NewParametro(ast.FIT, ast.BF)
        }
      |  TOK_UNIT TOK_IGUAL TOK_BYTES {
              $par = parametros.NewParametro(ast.UNIT, ast.B)
        }
      |  TOK_UNIT TOK_IGUAL TOK_KB {
              $par = parametros.NewParametro(ast.UNIT, ast.K)
        }
      |  TOK_UNIT TOK_IGUAL TOK_MB {
              $par = parametros.NewParametro(ast.UNIT, ast.M)
        }
      |  TOK_NAME TOK_IGUAL TOK_PALABRA {
              $par = parametros.NewParametro(ast.NAME, $TOK_PALABRA.text)
        }
      |  TOK_NAME TOK_IGUAL TOK_CADENA {
              str := strings.Trim($TOK_CADENA.text, "\"")
              $par = parametros.NewParametro(ast.NAME, str)
        }
      |  TOK_USUARIO TOK_IGUAL TOK_PALABRA {
              $par = parametros.NewParametro(ast.USUARIO, $TOK_PALABRA.text)
        }
      |  TOK_USUARIO TOK_IGUAL TOK_CADENA {
              str := strings.Trim($TOK_CADENA.text, "\"")
              $par = parametros.NewParametro(ast.USUARIO, str)
        }
      |  TOK_GRP TOK_IGUAL TOK_PALABRA {
              $par = parametros.NewParametro(ast.GRP, $TOK_PALABRA.text)
        }
      |  TOK_GRP TOK_IGUAL TOK_CADENA {
              str := strings.Trim($TOK_CADENA.text, "\"")
              $par = parametros.NewParametro(ast.GRP, str)
        }
      |  TOK_PASSWORD TOK_IGUAL TOK_PALABRA {
              $par = parametros.NewParametro(ast.PASSWORD, $TOK_PALABRA.text)
        }
      |  TOK_PWD TOK_IGUAL TOK_PALABRA {
              $par = parametros.NewParametro(ast.PWD, $TOK_PALABRA.text)
        }
      |  TOK_TYPE TOK_IGUAL TOK_PRIMARIA {
              $par = parametros.NewParametro(ast.TYPE, ast.PP)
        }
      |  TOK_TYPE TOK_IGUAL TOK_LOGICA {
              $par = parametros.NewParametro(ast.TYPE, ast.L)
        }
      |  TOK_TYPE TOK_IGUAL TOK_EXTENDIDA {
              $par = parametros.NewParametro(ast.TYPE, ast.E)
        }
      |  TOK_TYPE TOK_IGUAL TOK_FAST {
              $par = parametros.NewParametro(ast.TYPE, ast.FAST)
        }
      |  TOK_TYPE TOK_IGUAL TOK_FULL {
              $par = parametros.NewParametro(ast.TYPE, ast.FULL)
        }
      |  TOK_ID TOK_IGUAL TOK_IDENTIFICADOR {
              $par = parametros.NewParametro(ast.ID, $TOK_IDENTIFICADOR.text)
        }
      |  TOK_CONT TOK_IGUAL TOK_CADENA {
              str := strings.Trim($TOK_CADENA.text, "\"")
              $par = parametros.NewParametro(ast.CONT, str)
        }
      |  TOK_CONT TOK_IGUAL TOK_CAMINO {
              $par = parametros.NewParametro(ast.CONT, $TOK_CAMINO.text)
        }
      |  TOK_RUTA TOK_IGUAL TOK_CADENA {
              str := strings.Trim($TOK_CADENA.text, "\"")
              $par = parametros.NewParametro(ast.RUTA, str)
        }
      |  TOK_RUTA TOK_IGUAL TOK_CAMINO {
              $par = parametros.NewParametro(ast.RUTA, $TOK_CAMINO.text)
        }
      |  TOK_P {
              $par = parametros.NewParametro(ast.P, "true")
        }
      |  TOK_R {
              $par = parametros.NewParametro(ast.R, "true")
        }
      ;



