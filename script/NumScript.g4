grammar NumScript;

NEWLINE: [\r\n]+;

VARS: 'vars';
PRINT: 'print';
FAIL: 'fail';
SEND: 'send';
OP_ADD: '+';
OP_SUB: '-';
LPAREN: '(';
RPAREN: ')';
LBRACK: '[';
RBRACK: ']';
LBRACE: '{';
RBRACE: '}';
EQ: '=';
TY_ACCOUNT: 'account';
TY_ASSET: 'asset';
TY_NUMBER: 'number';
TY_MONETARY: 'monetary';
NUMBER: [0-9]+;
IDENTIFIER: [a-z0-9_:]+;
VARIABLE_NAME: '$' [a-z_]+ [a-z0-9_]+;
ACCOUNT: '@' [a-z_]+ [a-z0-9_:]+;
ASSET: [A-Z0-9/]+;

SEP: ';';
WHITESPACE: [ \t\r\n]+ -> skip;

monetary: LBRACK asset=ASSET amount=NUMBER RBRACK;

literal
  : ACCOUNT # LitAccount
  | ASSET # LitAsset
  | NUMBER # LitNumber
  | monetary # LitMonetary
  ;

expression
  : lhs=expression op=(OP_ADD|OP_SUB) rhs=expression # ExprAddSub
  | lit=literal # ExprLiteral
  | variable=VARIABLE_NAME # ExprVariable
  ;

argument: name=IDENTIFIER EQ val=expression;

statement
  : PRINT expr=expression # Print
  | FAIL # Fail
  | SEND LPAREN ((args+=argument ',')+ args+=argument?) RPAREN # Send
  ;

type_: TY_ACCOUNT | TY_ASSET | TY_NUMBER | TY_MONETARY;

varDecl: ty=type_ name=VARIABLE_NAME;

varListDecl: VARS LBRACE v+=varDecl+ RBRACE;

script:
  vars=varListDecl?
  stmts+=statement
  (NEWLINE stmts+=statement)*
  NEWLINE?
  EOF
  ;
