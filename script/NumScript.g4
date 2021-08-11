grammar NumScript;

NEWLINE: [\r\n];
WHITESPACE: [ \t]+ -> skip;

MULTILINE_COMMENT: '/*' (MULTILINE_COMMENT|.)*? '*/' -> skip;
LINE_COMMENT: '//' .*? NEWLINE -> skip;
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
TY_PORTION: 'portion';
STRING: '"' [a-zA-Z0-9]* '"';
PORTION:
  ( [0-9]+ [ ]? '/' [ ]? [0-9]+
  | [0-9]+     ('.'      [0-9]+)? '%'
  );
PORTION_REMAINING: 'remaining';
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

destinationAllotment: LBRACE NEWLINE (portions+=allotmentPortion  'to' dests+=expression NEWLINE)+ RBRACE;
destination
  : expression # DestAccount
  | destinationAllotment # DestAllotment
  ;

sourceInOrder: LBRACE NEWLINE (sources+=source NEWLINE)+ RBRACE;
source
  : expression # SrcAccount
  | sourceInOrder # SrcInOrder
  ;

sourceAllotment: LBRACE NEWLINE (portions+=allotmentPortion 'from' sources+=source NEWLINE)+ RBRACE;
valueAwareSource
  : source # Src
  | sourceAllotment # SrcAllotment
  ;

statement
  : PRINT expr=expression # Print
  | FAIL # Fail
  | SEND (mon=expression | monAll=monetaryAll) LPAREN NEWLINE
      ( SOURCE '=' src=valueAwareSource NEWLINE DESTINATION '=' dest=destination
      | DESTINATION '=' dest=destination NEWLINE SOURCE '=' src=valueAwareSource) NEWLINE RPAREN # Send
  ;

type_: TY_ACCOUNT | TY_ASSET | TY_NUMBER | TY_MONETARY | TY_PORTION;

origin
  : 'meta(' acc=expression ',' key=STRING ')'
  ;


varDecl: ty=type_ name=variable (EQ orig=origin)?;

varListDecl: VARS LBRACE NEWLINE+ (v+=varDecl NEWLINE+)+ RBRACE NEWLINE+;

script:
  NEWLINE*
  vars=varListDecl?
  stmts+=statement
  (NEWLINE+ stmts+=statement)*
  NEWLINE*
  EOF
  ;
