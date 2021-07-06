grammar NumScript;

NEWLINE: [\r\n]+;

PRINT: 'print';
FAIL: 'fail';
SEND: 'send';
OP_ADD: '+';
OP_SUB: '-';
LPAREN: '(';
RPAREN: ')';
LBRACK: '[';
RBRACK: ']';
EQ: '=';
IDENTIFIER: [a-z_:]+;
NUMBER: [0-9]+;
ASSET: [A-Z/0-9]+;

SEP: ';';
WHITESPACE: [ \n\t]+ -> skip;

monetary: LBRACK asset=ASSET number=NUMBER RBRACK;

literal
  : IDENTIFIER # LitAddress
  | ASSET # LitAsset
  | NUMBER # LitNumber
  | monetary # LitMonetary
  ;

expression
  : lhs=expression op=(OP_ADD|OP_SUB) rhs=expression # ExprAddSub
  | lit=literal # ExprLiteral
  ;

argument: name=IDENTIFIER EQ lit=literal;

statement
  : PRINT expr=expression # Print
  | FAIL # Fail
  | SEND LPAREN ((args+=argument ',')+ args+=argument?) RPAREN # Send
  ;

script:
  stmts+=statement
  (NEWLINE stmts+=statement)*
  EOF
  ;