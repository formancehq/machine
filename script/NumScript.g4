grammar NumScript;

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

script:
  expression*
  EOF
  ;