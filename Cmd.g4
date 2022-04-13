grammar Cmd;

// TOKENS
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
TOK_PATH:               '-'('P'|'p')('A'|'a')('T'|'t')('H'|'h');
TOK_FIT:                '-'('F'|'f')('I'|'i')('T'|'t');
TOK_SIZE:               '-'('S'|'s')('I'|'i')('Z'|'z')('E'|'e');
TOK_UNIT:               '-'('U'|'u')('N'|'n')('I'|'i')('T'|'t');
TOK_NAME:               '-'('N'|'n')('A'|'a')('M'|'m')('E'|'e');
TOK_TYPE:               '-'('T'|'t')('Y'|'y')('P'|'p')('E'|'e');
TOK_ID:                 '-'('I'|'i')('D'|'d');
TOK_USUARIO:            '-'('U'|'u')('S'|'s')('U'|'u')('A'|'a')('R'|'r')('I'|'i')('O'|'o');
TOK_PASSWORD:           '-'('P'|'p')('A'|'a')('S'|'s')('S'|'s')('W'|'w')('O'|'o')('R'|'r')('D'|'d');
TOK_PWD:                '-'('P'|'p')('W'|'w')('D'|'d');
TOK_CONT:               '-'('C'|'c')('O'|'o')('N'|'n')('T'|'t');
TOK_GRP:                '-'('G'|'g')('R'|'r')('P'|'p');
TOK_RUTA:               '-'('R'|'r')('U'|'u')('T'|'t')('A'|'a');
TOK_R:                  '-'('R'|'r');
TOK_P:                  '-'('P'|'p');

// RESPUESTAS
TOK_FIRST:              ('F'|'f')('F'|'f');
TOK_WORST:              ('W'|'w')('F'|'f');
TOK_BEST:               ('B'|'b')('F'|'f');
TOK_KB:                 ('K'|'k');
TOK_MB:                 ('M'|'m');
TOK_BYTES:              ('B'|'b');
TOK_PRIMARIA:           ('P'|'p');
TOK_EXTENDIDA:          ('E'|'e');
TOK_LOGICA:             ('L'|'l');
TOK_FAST:               ('F'|'f')('A'|'a')('S'|'s')('T'|'t');
TOK_FULL:               ('F'|'f')('U'|'u')('L'|'l')('L'|'l');
TOK_CADENA:             ["][^"\n]*["];
TOK_NUMERO:             [0-9]+;
TOK_IDENTIFICADOR:      [0-9][0-9][0-9][Aa-zZ];
TOK_CAMINO:             ([/][^\n "]+)+;
TOK_PALABRA:            [a-zA-Z0-9._-]+;

//SIMBOLOS ESPECIALES
TOK_IGUAL:              [=];
COMENTARIOS:            [#][^\n]*[\n]? ->skip;
WHITESPACE:             [ \r\n\t]+ -> skip;

// PRODUCCIONES
start : comando EOF;

comando : comando_estado param_list #AddCommand
        | com = TOK_PAUSE #Pause
        | com = TOK_LOGOUT #Logout
        ;

comando_estado : com = TOK_MKDISK #Mkdisk
               | com = TOK_RMDISK #Rmdisk
               | com = TOK_FDISK #Fdisk
               | com = TOK_MOUNT #Mount
               | com = TOK_MKFS #Mkfs
               | com = TOK_LOGIN #Login
               | com = TOK_MKGRP #Mkgrp
               | com = TOK_RMGRP #Rmgrp
               | com = TOK_MKUSR #Mkusr
               | com = TOK_RMUSR #Rmusr
               | com = TOK_MKFILE #Mkfile
               | com = TOK_MKDIR #Mkdir
               | com = TOK_EXEC #Exec
               | com = TOK_REP #Rep
               ;

param_list : param_list param    
           | param
           ;

param : par = TOK_SIZE TOK_IGUAL res = TOK_NUMERO #Size
      | par = TOK_PATH TOK_IGUAL res = TOK_CAMINO #Path_R
      | par = TOK_PATH TOK_IGUAL res = TOK_CADENA #Path_S
      | par = TOK_FIT TOK_IGUAL res = TOK_FIRST #FitFF
      | par = TOK_FIT TOK_IGUAL res = TOK_WORST #FitWF
      | par = TOK_FIT TOK_IGUAL res = TOK_BEST #FitBF
      | par = TOK_UNIT TOK_IGUAL res = TOK_BYTES #Unit_B
      | par = TOK_UNIT TOK_IGUAL res = TOK_KB #Unit_K
      | par = TOK_UNIT TOK_IGUAL res = TOK_MB #Unit_M
      | par = TOK_NAME TOK_IGUAL res = TOK_PALABRA #Name_P
      | par = TOK_NAME TOK_IGUAL res = TOK_CADENA #Name_S
      | par = TOK_USUARIO TOK_IGUAL res = TOK_PALABRA #Usr_P
      | par = TOK_USUARIO TOK_IGUAL res = TOK_CADENA #Usr_S
      | par = TOK_GRP TOK_IGUAL res = TOK_PALABRA #Grp_P
      | par = TOK_GRP TOK_IGUAL res = TOK_CADENA #Grp_S
      | par = TOK_PASSWORD TOK_IGUAL res = TOK_PALABRA #Pass_P
      | par = TOK_PWD TOK_IGUAL res = TOK_PALABRA #Pwd_P
      | par = TOK_TYPE TOK_IGUAL res = TOK_PRIMARIA #Type_P
      | par = TOK_TYPE TOK_IGUAL res = TOK_LOGICA #Type_L
      | par = TOK_TYPE TOK_IGUAL res = TOK_EXTENDIDA #Type_E
      | par = TOK_TYPE TOK_IGUAL res = TOK_FAST #Type_Fast
      | par = TOK_TYPE TOK_IGUAL res = TOK_FULL #Type_Full
      | par = TOK_ID TOK_IGUAL res = TOK_IDENTIFICADOR #Id
      | par = TOK_CONT TOK_IGUAL res = TOK_CADENA #Cont_S
      | par = TOK_CONT TOK_IGUAL res = TOK_CAMINO #Cont_R
      | par = TOK_RUTA TOK_IGUAL res = TOK_CADENA #Ruta_S
      | par = TOK_RUTA TOK_IGUAL res = TOK_CAMINO #Ruta_R
      | par = TOK_P #Pp
      | par = TOK_R #Rr
      ;

