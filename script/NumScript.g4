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
ALL: '*';
EQ: '=';
TY_ACCOUNT: 'account';
TY_ASSET: 'asset';
TY_NUMBER: 'number';
TY_MONETARY: 'monetary';
RATIO: [0-9]+ [ ]* '/' [ ]* [0-9]+;
NUMBER: [0-9]+;
PERCENT: '%';
IDENTIFIER: [a-z0-9_:]+;
VARIABLE_NAME: '$' [a-z_]+ [a-z0-9_]*;
ACCOUNT: '@' [a-z_]+ [a-z0-9_:]*;
ASSET: [A-Z/0-9]+;

amount
  : num=NUMBER # AmountSpecific
  | ALL # AmountAll
  ;

monetary: LBRACK asset=ASSET amt=amount RBRACK;

frac
  : r=RATIO # Ratio
  | pint=NUMBER ('.' pfrac=NUMBER)? PERCENT # Percentage
  ;

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

allocationPart: fr=frac 'to' dest=expression;

allocationBlock: LBRACE NEWLINE (parts+=allocationPart NEWLINE)+ RBRACE;

allocation
  : allocationBlock # AllocBlock
  | expression # AllocAccount
  ;

sourceBlock: LBRACE NEWLINE (sources+=expression NEWLINE)+ RBRACE;

source
  : sourceBlock # SrcBlock
  | expression # SrcAccount
  ;

statement
  : PRINT expr=expression # Print
  | FAIL # Fail
  | SEND mon=expression LPAREN NEWLINE
      ( SOURCE '=' src=source NEWLINE DESTINATION '=' dest=allocation
      | DESTINATION '=' dest=allocation NEWLINE SOURCE '=' src=source) NEWLINE RPAREN # Send
  ;

type_: TY_ACCOUNT | TY_ASSET | TY_NUMBER | TY_MONETARY;

varDecl: ty=type_ name=VARIABLE_NAME;

varListDecl: VARS LBRACE NEWLINE+ (v+=varDecl NEWLINE+)+ RBRACE NEWLINE+;

script:
  NEWLINE*
  vars=varListDecl?
  stmts+=statement
  (NEWLINE+ stmts+=statement)*
  NEWLINE*
  EOF
  ;
