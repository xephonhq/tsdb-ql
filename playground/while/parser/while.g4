grammar while;
prog: cexp+;
aexp: aexp '*' aexp # mul
    | aexp '/' aexp # div
    | aexp '+' aexp # sum
    | aexp '-' aexp # sub
    | INT # num
    | ID  # var
    ;
bexp: aexp '<' aexp # les
    | aexp '>' aexp # gt
    | aexp '==' aexp # eq
    | aexp '!=' aexp # neq
    | BOOL # bool
    ;
cexp: 'skip' ';'  # skip
    | ID ':=' aexp ';' # assign
    | 'if' bexp 'then' '{' cexp+ '}' 'else' '{' cexp+ '}' # if
    | 'while' bexp 'do' '{' cexp+ '}' # while
    ;
ID  : [a-zA-Z]+;
INT : [0-9]+;
BOOL: 'true' | 'false';
WS  : [ \t\r\n]+ -> skip ;
COMMENT: '//' .*? '\n' -> skip;
