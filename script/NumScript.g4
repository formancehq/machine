grammar NumScript;

NEWLINE: [\r\n]+;

CALC: 'calc';
FAIL: 'fail';
IDENTIFIER: [a-z]+;
ASSET: [A-Z]+;
NUMBER: [0-9]+;
OP_ADD: '+';
OP_SUB: '-';

SEP: ';';
WHITESPACE: [ \n\t]+ -> skip;

expression
  : expression op=(OP_ADD|OP_SUB) expression # AddSub
  | NUMBER # Number
  ;

statement
  : CALC expr=expression # Calc
  | FAIL # Fail
  ;

script:
  (statement NEWLINE)*
  statement?
  EOF
  ;