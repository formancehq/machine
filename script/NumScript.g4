grammar NumScript;

NEWLINE: [\r\n];
WHITESPACE: [ \t]+ -> skip;

VARS: 'vars';
PRINT: 'print';
FAIL: 'fail';
SEND: 'send';
SOURCE: 'source';
DESTINATION: 'destination';
ALLOCATE: 'allocate';
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
RATIO: [0-9]+ [ ]* '/' [ ]* [0-9]+;
NUMBER: [0-9]+;
PERCENT: '%';
IDENTIFIER: [a-z0-9_:]+;
VARIABLE_NAME: '$' [a-z_]+ [a-z0-9_]+;
ACCOUNT: '@' [a-z_]+ [a-z0-9_:]*;
ASSET: [A-Z/0-9]+;

monetary: LBRACK asset=ASSET amount=NUMBER RBRACK;

frac
  : r=RATIO # Ratio
  | p=NUMBER PERCENT # Percentage
  ;

allocationPart: fr=frac 'to' dest=expression;

allocation: LBRACE NEWLINE (parts+=allocationPart NEWLINE)+ RBRACE;

literal
  : ACCOUNT # LitAccount
  | ASSET # LitAsset
  | NUMBER # LitNumber
  | monetary # LitMonetary
  | allocation # LitAllocation
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
  | SEND mon=expression LPAREN NEWLINE
      ( SOURCE '=' src=expression NEWLINE DESTINATION '=' dest=expression
      | DESTINATION '=' dest=expression NEWLINE SOURCE '=' src=expression) NEWLINE RPAREN # Send
  ;

type_: TY_ACCOUNT | TY_ASSET | TY_NUMBER | TY_MONETARY;

varDecl: ty=type_ name=VARIABLE_NAME;

varListDecl: VARS LBRACE NEWLINE+ (v+=varDecl NEWLINE+)+ RBRACE NEWLINE+;

script:
  vars=varListDecl?
  stmts+=statement
  (NEWLINE+ stmts+=statement)*
  NEWLINE*
  EOF
  ;
