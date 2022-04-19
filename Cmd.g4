grammar Cmd;

@header{
        import "./parser"
        import "./util"
        import "./ast"
        import arrayList "github.com/colegno/arraylist"
}


// TOKENS
fragment SLASH:                  '/';
fragment COMDOB:                 '"';
fragment HASHTAG:                '#';
fragment DASH:                   '-';
fragment TOK_IGUAL:              '=';

// COMANDOS PRINCIPALES
TOK_MKDISK:             ('M'|'m')('K'|'k')('D'|'d')('I'|'i')('S'|'s')('K'|'k');
TOK_RMDISK:             ('R'|'r')('M'|'m')('D'|'d')('I'|'i')('S'|'s')('K'|'k');
TOK_FDISK:              ('F'|'f')('D'|'d')('I'|'i')('S'|'s')('K'|'k');
TOK_MOUNT:              ('M'|'m')('O'|'o')('U'|'u')('N'|'n')('T'|'t');
TOK_MKFS:               ('M'|'m')('K'|'k')('F'|'f')('S'|'s');
TOK_LOGIN:              ('L'|'l')('O'|'o')('G'|'g')('I'|'i')('N'|'n');
TOK_LOGOUT:             ('L'|'l')('O'|'o')('G'|'g')('O'|'o')('U'|'u')('T'|'t');
TOK_MKGRP:              ('M'|'m')('K'|'k')('G'|'g')('R'|'r')('P'|'p');
TOK_RMGRP:              ('R'|'r')('M'|'m')('G'|'g')('R'|'r')('P'|'p');
TOK_MKUSR:              ('M'|'m')('K'|'k')('U'|'u')('S'|'s')('R'|'r');
TOK_RMUSR:              ('R'|'r')('M'|'m')('U'|'u')('S'|'s')('R'|'r');
TOK_MKFILE:             ('M'|'m')('K'|'k')('F'|'f')('I'|'i')('L'|'l')('E'|'e');
TOK_MKDIR:              ('M'|'m')('K'|'k')('D'|'d')('I'|'i')('R'|'r');
TOK_EXEC:               ('E'|'e')('X'|'x')('E'|'e')('C'|'c');
TOK_PAUSE:              ('P'|'p')('A'|'a')('U'|'u')('S'|'s')('E'|'e');
TOK_REP:                ('R'|'r')('E'|'e')('P'|'p');

// PARAMETROS
TOK_PATH:               DASH('P'|'p')('A'|'a')('T'|'t')('H'|'h');
TOK_FIT:                DASH('F'|'f')('I'|'i')('T'|'t');
TOK_SIZE:               DASH('S'|'s')('I'|'i')('Z'|'z')('E'|'e');
TOK_UNIT:               DASH('U'|'u')('N'|'n')('I'|'i')('T'|'t');
TOK_NAME:               DASH('N'|'n')('A'|'a')('M'|'m')('E'|'e');
TOK_TYPE:               DASH('T'|'t')('Y'|'y')('P'|'p')('E'|'e');
TOK_ID:                 DASH('I'|'i')('D'|'d');
TOK_USUARIO:            DASH('U'|'u')('S'|'s')('U'|'u')('A'|'a')('R'|'r')('I'|'i')('O'|'o');
TOK_PASSWORD:           DASH('P'|'p')('A'|'a')('S'|'s')('S'|'s')('W'|'w')('O'|'o')('R'|'r')('D'|'d');
TOK_PWD:                DASH('P'|'p')('W'|'w')('D'|'d');
TOK_CONT:               DASH('C'|'c')('O'|'o')('N'|'n')('T'|'t');
TOK_GRP:                DASH('G'|'g')('R'|'r')('P'|'p');
TOK_RUTA:               DASH('R'|'r')('U'|'u')('T'|'t')('A'|'a');
TOK_R:                  DASH('R'|'r');
TOK_P:                  DASH('P'|'p');

// RESPUESTAS

TOK_FIRST:              ('F'|'f')('F'|'f');
TOK_WORST:              ('W'|'w')('F'|'f');
TOK_BEST:               ('B'|'b')('F'|'f');

TOK_FAST:               ('F'|'f')('A'|'a')('S'|'s')('T'|'t');
TOK_FULL:               ('F'|'f')('U'|'u')('L'|'l')('L'|'l');
TOK_CADENA:             COMDOB(~([\n]|'"'))*COMDOB;
TOK_NUMERO:             [0..9]+;
TOK_IDENTIFICADOR:      [0..9][0..9][0..9]([a..z]|[A..Z]);
TOK_CAMINO:             (SLASH(~([\n]|'"'|' '))*)+;
TOK_PALABRA:            ([a..z]|[A..Z]|[0..9]|[._-])+;

//SIMBOLOS ESPECIALES


COMENTARIOS:            HASHTAG~([\n])*[\n]* -> skip;
WHITESPACE:             [ \r\n\t]+ -> skip;

// PRODUCCIONES
start returns [ast.Ast ast] 
           : comandlist{$ast = ast.NewAst($comandlist.lista)};

comandList returns[*arrayList.List lista] @init{ $lista = arrayList.New()}
           : AUXCLIST = comandlist comando {
                        $AUXCLIST.lista.Add($comando.com)
                        $lista.Add($comando.com)
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

param_list : param_list param    
           | param
           ;

param :  TOK_SIZE TOK_IGUAL  TOK_NUMERO 
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

