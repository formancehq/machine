grammar NumScript;

NEWLINE: [\r\n]+;
WHITESPACE: [ \t]+ -> skip;

MULTILINE_COMMENT: '/*' (MULTILINE_COMMENT|.)*? '*/' -> skip;
LINE_COMMENT: '//' .*? NEWLINE -> skip;
VARS: 'vars';
META: 'meta';
SET_TX_META: 'set_tx_meta';
PRINT: 'print';
FAIL: 'fail';
IF: 'if';
SEND: 'send';
SOURCE: 'source';
FROM: 'from';
MAX: 'max';
DESTINATION: 'destination';
TO: 'to';
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
TY_PORTION: 'portion';
TY_BOOLEAN: 'boolean';
TY_STRING: 'string';
STRING: '"' [a-zA-Z0-9_\- ]* '"';
PORTION:
  ( [0-9]+ [ ]? '/' [ ]? [0-9]+
  | [0-9]+     ('.'      [0-9]+)? '%'
  );
PORTION_REMAINING: 'remaining';
BOOLEAN: 'true' | 'false';
NUMBER: [0-9]+;
PERCENT: '%';
VARIABLE_NAME: '$' [a-z_]+ [a-z0-9_]*;
ACCOUNT: '@' [a-z_]+ [a-z0-9_:]*;
ASSET: [A-Z/0-9]+;

monetary: LBRACK asset=ASSET amt=NUMBER RBRACK;

monetaryAll: LBRACK asset=ASSET '*' RBRACK;

literal
  : ACCOUNT # LitAccount
  | ASSET # LitAsset
  | NUMBER # LitNumber
  | STRING # LitString
  | monetary # LitMonetary
  ;

variable: VARIABLE_NAME;

expression
  : lhs=expression op=(OP_ADD|OP_SUB) rhs=expression # ExprAddSub
  | lit=literal # ExprLiteral
  | var_=variable # ExprVariable
  ;

allotmentPortion
  : PORTION # allotmentPortionConst
  | por=variable # allotmentPortionVar
  | PORTION_REMAINING # allotmentPortionRemaining
  ;

destinationInOrder: LBRACE NEWLINE (dests+=destination NEWLINE)+ RBRACE;
destinationMaxed: MAX max=expression TO dest=destination;
destinationAllotment: LBRACE NEWLINE (portions+=allotmentPortion TO dests+=destination NEWLINE)+ RBRACE;
destination
  : expression # DestAccount
  | destinationMaxed # DestMaxed
  | destinationInOrder # DestInOrder
  | destinationAllotment # DestAllotment
  ;

sourceInOrder: LBRACE NEWLINE (sources+=source NEWLINE)+ RBRACE;
sourceMaxed: MAX max=expression FROM src=source;
source
  : expression # SrcAccount
  | sourceMaxed # SrcMaxed
  | sourceInOrder # SrcInOrder
  ;

sourceAllotment: LBRACE NEWLINE (portions+=allotmentPortion FROM sources+=source NEWLINE)+ RBRACE;
valueAwareSource
  : source # Src
  | sourceAllotment # SrcAllotment
  ;

ifStatement:
  IF condition=variable '{' (NEWLINE stmts+=statement)+ NEWLINE '}'
  ;

statement
  : PRINT expr=expression # Print
  | SET_TX_META '(' key=STRING ',' value=expression ')' #SetTxMeta
  | FAIL # Fail
  | SEND (mon=expression | monAll=monetaryAll) LPAREN NEWLINE
      ( SOURCE '=' src=valueAwareSource NEWLINE DESTINATION '=' dest=destination
      | DESTINATION '=' dest=destination NEWLINE SOURCE '=' src=valueAwareSource) NEWLINE RPAREN # Send
  | stmt=ifStatement #IfStmt
  ;


type_: TY_ACCOUNT | TY_ASSET | TY_NUMBER | TY_STRING | TY_MONETARY | TY_PORTION | TY_BOOLEAN;

origin
  : META '(' acc=expression ',' key=STRING ')'
  ;


varDecl: ty=type_ name=variable (EQ orig=origin)?;

varListDecl: VARS LBRACE NEWLINE (v+=varDecl NEWLINE+)+ RBRACE NEWLINE;

script:
  NEWLINE*
  vars=varListDecl?
  stmts+=statement
  (NEWLINE stmts+=statement)*
  NEWLINE*
  EOF
  ;
