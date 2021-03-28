# Golang AST sample

## Source file to parse

[sample/sample.go](sample/sample.go)

## How to run
```
go run ./main.go
```

## Result
```
     0  *ast.File {
     1  .  Package: sample/sample.go:1:1
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: sample/sample.go:1:9
     4  .  .  Name: "sample"
     5  .  }
     6  .  Decls: []ast.Decl (len = 5) {
     7  .  .  0: *ast.GenDecl {
     8  .  .  .  TokPos: sample/sample.go:3:1
     9  .  .  .  Tok: import
    10  .  .  .  Lparen: sample/sample.go:3:8
    11  .  .  .  Specs: []ast.Spec (len = 2) {
    12  .  .  .  .  0: *ast.ImportSpec {
    13  .  .  .  .  .  Path: *ast.BasicLit {
    14  .  .  .  .  .  .  ValuePos: sample/sample.go:4:2
    15  .  .  .  .  .  .  Kind: STRING
    16  .  .  .  .  .  .  Value: "\"errors\""
    17  .  .  .  .  .  }
    18  .  .  .  .  .  EndPos: -
    19  .  .  .  .  }
    20  .  .  .  .  1: *ast.ImportSpec {
    21  .  .  .  .  .  Name: *ast.Ident {
    22  .  .  .  .  .  .  NamePos: sample/sample.go:5:2
    23  .  .  .  .  .  .  Name: "r"
    24  .  .  .  .  .  }
    25  .  .  .  .  .  Path: *ast.BasicLit {
    26  .  .  .  .  .  .  ValuePos: sample/sample.go:5:4
    27  .  .  .  .  .  .  Kind: STRING
    28  .  .  .  .  .  .  Value: "\"math/rand\""
    29  .  .  .  .  .  }
    30  .  .  .  .  .  EndPos: -
    31  .  .  .  .  }
    32  .  .  .  }
    33  .  .  .  Rparen: sample/sample.go:6:1
    34  .  .  }
    35  .  .  1: *ast.GenDecl {
    36  .  .  .  TokPos: sample/sample.go:9:1
    37  .  .  .  Tok: type
    38  .  .  .  Lparen: -
    39  .  .  .  Specs: []ast.Spec (len = 1) {
    40  .  .  .  .  0: *ast.TypeSpec {
    41  .  .  .  .  .  Name: *ast.Ident {
    42  .  .  .  .  .  .  NamePos: sample/sample.go:9:6
    43  .  .  .  .  .  .  Name: "RandGenerator"
    44  .  .  .  .  .  .  Obj: *ast.Object {
    45  .  .  .  .  .  .  .  Kind: type
    46  .  .  .  .  .  .  .  Name: "RandGenerator"
    47  .  .  .  .  .  .  .  Decl: *(obj @ 40)
    48  .  .  .  .  .  .  }
    49  .  .  .  .  .  }
    50  .  .  .  .  .  Assign: -
    51  .  .  .  .  .  Type: *ast.InterfaceType {
    52  .  .  .  .  .  .  Interface: sample/sample.go:9:20
    53  .  .  .  .  .  .  Methods: *ast.FieldList {
    54  .  .  .  .  .  .  .  Opening: sample/sample.go:9:30
    55  .  .  .  .  .  .  .  List: []*ast.Field (len = 1) {
    56  .  .  .  .  .  .  .  .  0: *ast.Field {
    57  .  .  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
    58  .  .  .  .  .  .  .  .  .  .  0: *ast.Ident {
    59  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:10:2
    60  .  .  .  .  .  .  .  .  .  .  .  Name: "RandInt"
    61  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
    62  .  .  .  .  .  .  .  .  .  .  .  .  Kind: func
    63  .  .  .  .  .  .  .  .  .  .  .  .  Name: "RandInt"
    64  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 56)
    65  .  .  .  .  .  .  .  .  .  .  .  }
    66  .  .  .  .  .  .  .  .  .  .  }
    67  .  .  .  .  .  .  .  .  .  }
    68  .  .  .  .  .  .  .  .  .  Type: *ast.FuncType {
    69  .  .  .  .  .  .  .  .  .  .  Func: -
    70  .  .  .  .  .  .  .  .  .  .  Params: *ast.FieldList {
    71  .  .  .  .  .  .  .  .  .  .  .  Opening: sample/sample.go:10:9
    72  .  .  .  .  .  .  .  .  .  .  .  List: []*ast.Field (len = 1) {
    73  .  .  .  .  .  .  .  .  .  .  .  .  0: *ast.Field {
    74  .  .  .  .  .  .  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 2) {
    75  .  .  .  .  .  .  .  .  .  .  .  .  .  .  0: *ast.Ident {
    76  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:10:10
    77  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "min"
    78  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
    79  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Kind: var
    80  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "min"
    81  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 73)
    82  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  }
    83  .  .  .  .  .  .  .  .  .  .  .  .  .  .  }
    84  .  .  .  .  .  .  .  .  .  .  .  .  .  .  1: *ast.Ident {
    85  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:10:15
    86  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "max"
    87  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
    88  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Kind: var
    89  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "max"
    90  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 73)
    91  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  }
    92  .  .  .  .  .  .  .  .  .  .  .  .  .  .  }
    93  .  .  .  .  .  .  .  .  .  .  .  .  .  }
    94  .  .  .  .  .  .  .  .  .  .  .  .  .  Type: *ast.Ident {
    95  .  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:10:19
    96  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "int"
    97  .  .  .  .  .  .  .  .  .  .  .  .  .  }
    98  .  .  .  .  .  .  .  .  .  .  .  .  }
    99  .  .  .  .  .  .  .  .  .  .  .  }
   100  .  .  .  .  .  .  .  .  .  .  .  Closing: sample/sample.go:10:22
   101  .  .  .  .  .  .  .  .  .  .  }
   102  .  .  .  .  .  .  .  .  .  .  Results: *ast.FieldList {
   103  .  .  .  .  .  .  .  .  .  .  .  Opening: sample/sample.go:10:24
   104  .  .  .  .  .  .  .  .  .  .  .  List: []*ast.Field (len = 2) {
   105  .  .  .  .  .  .  .  .  .  .  .  .  0: *ast.Field {
   106  .  .  .  .  .  .  .  .  .  .  .  .  .  Type: *ast.Ident {
   107  .  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:10:25
   108  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "int"
   109  .  .  .  .  .  .  .  .  .  .  .  .  .  }
   110  .  .  .  .  .  .  .  .  .  .  .  .  }
   111  .  .  .  .  .  .  .  .  .  .  .  .  1: *ast.Field {
   112  .  .  .  .  .  .  .  .  .  .  .  .  .  Type: *ast.Ident {
   113  .  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:10:30
   114  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "error"
   115  .  .  .  .  .  .  .  .  .  .  .  .  .  }
   116  .  .  .  .  .  .  .  .  .  .  .  .  }
   117  .  .  .  .  .  .  .  .  .  .  .  }
   118  .  .  .  .  .  .  .  .  .  .  .  Closing: sample/sample.go:10:35
   119  .  .  .  .  .  .  .  .  .  .  }
   120  .  .  .  .  .  .  .  .  .  }
   121  .  .  .  .  .  .  .  .  }
   122  .  .  .  .  .  .  .  }
   123  .  .  .  .  .  .  .  Closing: sample/sample.go:11:1
   124  .  .  .  .  .  .  }
   125  .  .  .  .  .  .  Incomplete: false
   126  .  .  .  .  .  }
   127  .  .  .  .  }
   128  .  .  .  }
   129  .  .  .  Rparen: -
   130  .  .  }
   131  .  .  2: *ast.FuncDecl {
   132  .  .  .  Name: *ast.Ident {
   133  .  .  .  .  NamePos: sample/sample.go:14:6
   134  .  .  .  .  Name: "NewRandGenerator"
   135  .  .  .  .  Obj: *ast.Object {
   136  .  .  .  .  .  Kind: func
   137  .  .  .  .  .  Name: "NewRandGenerator"
   138  .  .  .  .  .  Decl: *(obj @ 131)
   139  .  .  .  .  }
   140  .  .  .  }
   141  .  .  .  Type: *ast.FuncType {
   142  .  .  .  .  Func: sample/sample.go:14:1
   143  .  .  .  .  Params: *ast.FieldList {
   144  .  .  .  .  .  Opening: sample/sample.go:14:22
   145  .  .  .  .  .  Closing: sample/sample.go:14:23
   146  .  .  .  .  }
   147  .  .  .  .  Results: *ast.FieldList {
   148  .  .  .  .  .  Opening: -
   149  .  .  .  .  .  List: []*ast.Field (len = 1) {
   150  .  .  .  .  .  .  0: *ast.Field {
   151  .  .  .  .  .  .  .  Type: *ast.Ident {
   152  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:14:25
   153  .  .  .  .  .  .  .  .  Name: "RandGenerator"
   154  .  .  .  .  .  .  .  .  Obj: *(obj @ 44)
   155  .  .  .  .  .  .  .  }
   156  .  .  .  .  .  .  }
   157  .  .  .  .  .  }
   158  .  .  .  .  .  Closing: -
   159  .  .  .  .  }
   160  .  .  .  }
   161  .  .  .  Body: *ast.BlockStmt {
   162  .  .  .  .  Lbrace: sample/sample.go:14:39
   163  .  .  .  .  List: []ast.Stmt (len = 1) {
   164  .  .  .  .  .  0: *ast.ReturnStmt {
   165  .  .  .  .  .  .  Return: sample/sample.go:15:2
   166  .  .  .  .  .  .  Results: []ast.Expr (len = 1) {
   167  .  .  .  .  .  .  .  0: *ast.UnaryExpr {
   168  .  .  .  .  .  .  .  .  OpPos: sample/sample.go:15:9
   169  .  .  .  .  .  .  .  .  Op: &
   170  .  .  .  .  .  .  .  .  X: *ast.CompositeLit {
   171  .  .  .  .  .  .  .  .  .  Type: *ast.Ident {
   172  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:15:10
   173  .  .  .  .  .  .  .  .  .  .  Name: "randGenerator"
   174  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
   175  .  .  .  .  .  .  .  .  .  .  .  Kind: type
   176  .  .  .  .  .  .  .  .  .  .  .  Name: "randGenerator"
   177  .  .  .  .  .  .  .  .  .  .  .  Decl: *ast.TypeSpec {
   178  .  .  .  .  .  .  .  .  .  .  .  .  Name: *ast.Ident {
   179  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:18:6
   180  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "randGenerator"
   181  .  .  .  .  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 174)
   182  .  .  .  .  .  .  .  .  .  .  .  .  }
   183  .  .  .  .  .  .  .  .  .  .  .  .  Assign: -
   184  .  .  .  .  .  .  .  .  .  .  .  .  Type: *ast.StructType {
   185  .  .  .  .  .  .  .  .  .  .  .  .  .  Struct: sample/sample.go:18:20
   186  .  .  .  .  .  .  .  .  .  .  .  .  .  Fields: *ast.FieldList {
   187  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Opening: sample/sample.go:18:27
   188  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Closing: sample/sample.go:19:1
   189  .  .  .  .  .  .  .  .  .  .  .  .  .  }
   190  .  .  .  .  .  .  .  .  .  .  .  .  .  Incomplete: false
   191  .  .  .  .  .  .  .  .  .  .  .  .  }
   192  .  .  .  .  .  .  .  .  .  .  .  }
   193  .  .  .  .  .  .  .  .  .  .  }
   194  .  .  .  .  .  .  .  .  .  }
   195  .  .  .  .  .  .  .  .  .  Lbrace: sample/sample.go:15:23
   196  .  .  .  .  .  .  .  .  .  Rbrace: sample/sample.go:15:24
   197  .  .  .  .  .  .  .  .  .  Incomplete: false
   198  .  .  .  .  .  .  .  .  }
   199  .  .  .  .  .  .  .  }
   200  .  .  .  .  .  .  }
   201  .  .  .  .  .  }
   202  .  .  .  .  }
   203  .  .  .  .  Rbrace: sample/sample.go:16:1
   204  .  .  .  }
   205  .  .  }
   206  .  .  3: *ast.GenDecl {
   207  .  .  .  TokPos: sample/sample.go:18:1
   208  .  .  .  Tok: type
   209  .  .  .  Lparen: -
   210  .  .  .  Specs: []ast.Spec (len = 1) {
   211  .  .  .  .  0: *(obj @ 177)
   212  .  .  .  }
   213  .  .  .  Rparen: -
   214  .  .  }
   215  .  .  4: *ast.FuncDecl {
   216  .  .  .  Recv: *ast.FieldList {
   217  .  .  .  .  Opening: sample/sample.go:21:6
   218  .  .  .  .  List: []*ast.Field (len = 1) {
   219  .  .  .  .  .  0: *ast.Field {
   220  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
   221  .  .  .  .  .  .  .  0: *ast.Ident {
   222  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:21:7
   223  .  .  .  .  .  .  .  .  Name: "g"
   224  .  .  .  .  .  .  .  .  Obj: *ast.Object {
   225  .  .  .  .  .  .  .  .  .  Kind: var
   226  .  .  .  .  .  .  .  .  .  Name: "g"
   227  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 219)
   228  .  .  .  .  .  .  .  .  }
   229  .  .  .  .  .  .  .  }
   230  .  .  .  .  .  .  }
   231  .  .  .  .  .  .  Type: *ast.StarExpr {
   232  .  .  .  .  .  .  .  Star: sample/sample.go:21:9
   233  .  .  .  .  .  .  .  X: *ast.Ident {
   234  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:21:10
   235  .  .  .  .  .  .  .  .  Name: "randGenerator"
   236  .  .  .  .  .  .  .  .  Obj: *(obj @ 174)
   237  .  .  .  .  .  .  .  }
   238  .  .  .  .  .  .  }
   239  .  .  .  .  .  }
   240  .  .  .  .  }
   241  .  .  .  .  Closing: sample/sample.go:21:23
   242  .  .  .  }
   243  .  .  .  Name: *ast.Ident {
   244  .  .  .  .  NamePos: sample/sample.go:21:25
   245  .  .  .  .  Name: "RandInt"
   246  .  .  .  }
   247  .  .  .  Type: *ast.FuncType {
   248  .  .  .  .  Func: sample/sample.go:21:1
   249  .  .  .  .  Params: *ast.FieldList {
   250  .  .  .  .  .  Opening: sample/sample.go:21:32
   251  .  .  .  .  .  List: []*ast.Field (len = 1) {
   252  .  .  .  .  .  .  0: *ast.Field {
   253  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 2) {
   254  .  .  .  .  .  .  .  .  0: *ast.Ident {
   255  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:21:33
   256  .  .  .  .  .  .  .  .  .  Name: "min"
   257  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
   258  .  .  .  .  .  .  .  .  .  .  Kind: var
   259  .  .  .  .  .  .  .  .  .  .  Name: "min"
   260  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 252)
   261  .  .  .  .  .  .  .  .  .  }
   262  .  .  .  .  .  .  .  .  }
   263  .  .  .  .  .  .  .  .  1: *ast.Ident {
   264  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:21:38
   265  .  .  .  .  .  .  .  .  .  Name: "max"
   266  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
   267  .  .  .  .  .  .  .  .  .  .  Kind: var
   268  .  .  .  .  .  .  .  .  .  .  Name: "max"
   269  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 252)
   270  .  .  .  .  .  .  .  .  .  }
   271  .  .  .  .  .  .  .  .  }
   272  .  .  .  .  .  .  .  }
   273  .  .  .  .  .  .  .  Type: *ast.Ident {
   274  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:21:42
   275  .  .  .  .  .  .  .  .  Name: "int"
   276  .  .  .  .  .  .  .  }
   277  .  .  .  .  .  .  }
   278  .  .  .  .  .  }
   279  .  .  .  .  .  Closing: sample/sample.go:21:45
   280  .  .  .  .  }
   281  .  .  .  .  Results: *ast.FieldList {
   282  .  .  .  .  .  Opening: sample/sample.go:21:47
   283  .  .  .  .  .  List: []*ast.Field (len = 2) {
   284  .  .  .  .  .  .  0: *ast.Field {
   285  .  .  .  .  .  .  .  Type: *ast.Ident {
   286  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:21:48
   287  .  .  .  .  .  .  .  .  Name: "int"
   288  .  .  .  .  .  .  .  }
   289  .  .  .  .  .  .  }
   290  .  .  .  .  .  .  1: *ast.Field {
   291  .  .  .  .  .  .  .  Type: *ast.Ident {
   292  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:21:53
   293  .  .  .  .  .  .  .  .  Name: "error"
   294  .  .  .  .  .  .  .  }
   295  .  .  .  .  .  .  }
   296  .  .  .  .  .  }
   297  .  .  .  .  .  Closing: sample/sample.go:21:58
   298  .  .  .  .  }
   299  .  .  .  }
   300  .  .  .  Body: *ast.BlockStmt {
   301  .  .  .  .  Lbrace: sample/sample.go:21:60
   302  .  .  .  .  List: []ast.Stmt (len = 2) {
   303  .  .  .  .  .  0: *ast.IfStmt {
   304  .  .  .  .  .  .  If: sample/sample.go:22:2
   305  .  .  .  .  .  .  Cond: *ast.BinaryExpr {
   306  .  .  .  .  .  .  .  X: *ast.Ident {
   307  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:22:5
   308  .  .  .  .  .  .  .  .  Name: "min"
   309  .  .  .  .  .  .  .  .  Obj: *(obj @ 257)
   310  .  .  .  .  .  .  .  }
   311  .  .  .  .  .  .  .  OpPos: sample/sample.go:22:9
   312  .  .  .  .  .  .  .  Op: >
   313  .  .  .  .  .  .  .  Y: *ast.Ident {
   314  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:22:11
   315  .  .  .  .  .  .  .  .  Name: "max"
   316  .  .  .  .  .  .  .  .  Obj: *(obj @ 266)
   317  .  .  .  .  .  .  .  }
   318  .  .  .  .  .  .  }
   319  .  .  .  .  .  .  Body: *ast.BlockStmt {
   320  .  .  .  .  .  .  .  Lbrace: sample/sample.go:22:15
   321  .  .  .  .  .  .  .  List: []ast.Stmt (len = 1) {
   322  .  .  .  .  .  .  .  .  0: *ast.ReturnStmt {
   323  .  .  .  .  .  .  .  .  .  Return: sample/sample.go:23:3
   324  .  .  .  .  .  .  .  .  .  Results: []ast.Expr (len = 2) {
   325  .  .  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
   326  .  .  .  .  .  .  .  .  .  .  .  ValuePos: sample/sample.go:23:10
   327  .  .  .  .  .  .  .  .  .  .  .  Kind: INT
   328  .  .  .  .  .  .  .  .  .  .  .  Value: "0"
   329  .  .  .  .  .  .  .  .  .  .  }
   330  .  .  .  .  .  .  .  .  .  .  1: *ast.CallExpr {
   331  .  .  .  .  .  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
   332  .  .  .  .  .  .  .  .  .  .  .  .  X: *ast.Ident {
   333  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:23:13
   334  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "errors"
   335  .  .  .  .  .  .  .  .  .  .  .  .  }
   336  .  .  .  .  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
   337  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:23:20
   338  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "New"
   339  .  .  .  .  .  .  .  .  .  .  .  .  }
   340  .  .  .  .  .  .  .  .  .  .  .  }
   341  .  .  .  .  .  .  .  .  .  .  .  Lparen: sample/sample.go:23:23
   342  .  .  .  .  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
   343  .  .  .  .  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
   344  .  .  .  .  .  .  .  .  .  .  .  .  .  ValuePos: sample/sample.go:23:24
   345  .  .  .  .  .  .  .  .  .  .  .  .  .  Kind: STRING
   346  .  .  .  .  .  .  .  .  .  .  .  .  .  Value: "\"min should be smaller than max\""
   347  .  .  .  .  .  .  .  .  .  .  .  .  }
   348  .  .  .  .  .  .  .  .  .  .  .  }
   349  .  .  .  .  .  .  .  .  .  .  .  Ellipsis: -
   350  .  .  .  .  .  .  .  .  .  .  .  Rparen: sample/sample.go:23:56
   351  .  .  .  .  .  .  .  .  .  .  }
   352  .  .  .  .  .  .  .  .  .  }
   353  .  .  .  .  .  .  .  .  }
   354  .  .  .  .  .  .  .  }
   355  .  .  .  .  .  .  .  Rbrace: sample/sample.go:24:2
   356  .  .  .  .  .  .  }
   357  .  .  .  .  .  }
   358  .  .  .  .  .  1: *ast.ReturnStmt {
   359  .  .  .  .  .  .  Return: sample/sample.go:26:2
   360  .  .  .  .  .  .  Results: []ast.Expr (len = 2) {
   361  .  .  .  .  .  .  .  0: *ast.BinaryExpr {
   362  .  .  .  .  .  .  .  .  X: *ast.CallExpr {
   363  .  .  .  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
   364  .  .  .  .  .  .  .  .  .  .  X: *ast.Ident {
   365  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:26:9
   366  .  .  .  .  .  .  .  .  .  .  .  Name: "r"
   367  .  .  .  .  .  .  .  .  .  .  }
   368  .  .  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
   369  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:26:11
   370  .  .  .  .  .  .  .  .  .  .  .  Name: "Intn"
   371  .  .  .  .  .  .  .  .  .  .  }
   372  .  .  .  .  .  .  .  .  .  }
   373  .  .  .  .  .  .  .  .  .  Lparen: sample/sample.go:26:15
   374  .  .  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
   375  .  .  .  .  .  .  .  .  .  .  0: *ast.BinaryExpr {
   376  .  .  .  .  .  .  .  .  .  .  .  X: *ast.Ident {
   377  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:26:16
   378  .  .  .  .  .  .  .  .  .  .  .  .  Name: "max"
   379  .  .  .  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 266)
   380  .  .  .  .  .  .  .  .  .  .  .  }
   381  .  .  .  .  .  .  .  .  .  .  .  OpPos: sample/sample.go:26:19
   382  .  .  .  .  .  .  .  .  .  .  .  Op: -
   383  .  .  .  .  .  .  .  .  .  .  .  Y: *ast.Ident {
   384  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:26:20
   385  .  .  .  .  .  .  .  .  .  .  .  .  Name: "min"
   386  .  .  .  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 257)
   387  .  .  .  .  .  .  .  .  .  .  .  }
   388  .  .  .  .  .  .  .  .  .  .  }
   389  .  .  .  .  .  .  .  .  .  }
   390  .  .  .  .  .  .  .  .  .  Ellipsis: -
   391  .  .  .  .  .  .  .  .  .  Rparen: sample/sample.go:26:23
   392  .  .  .  .  .  .  .  .  }
   393  .  .  .  .  .  .  .  .  OpPos: sample/sample.go:26:25
   394  .  .  .  .  .  .  .  .  Op: +
   395  .  .  .  .  .  .  .  .  Y: *ast.Ident {
   396  .  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:26:27
   397  .  .  .  .  .  .  .  .  .  Name: "min"
   398  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 257)
   399  .  .  .  .  .  .  .  .  }
   400  .  .  .  .  .  .  .  }
   401  .  .  .  .  .  .  .  1: *ast.Ident {
   402  .  .  .  .  .  .  .  .  NamePos: sample/sample.go:26:32
   403  .  .  .  .  .  .  .  .  Name: "nil"
   404  .  .  .  .  .  .  .  }
   405  .  .  .  .  .  .  }
   406  .  .  .  .  .  }
   407  .  .  .  .  }
   408  .  .  .  .  Rbrace: sample/sample.go:27:1
   409  .  .  .  }
   410  .  .  }
   411  .  }
   412  .  Scope: *ast.Scope {
   413  .  .  Objects: map[string]*ast.Object (len = 3) {
   414  .  .  .  "RandGenerator": *(obj @ 44)
   415  .  .  .  "NewRandGenerator": *(obj @ 135)
   416  .  .  .  "randGenerator": *(obj @ 174)
   417  .  .  }
   418  .  }
   419  .  Imports: []*ast.ImportSpec (len = 2) {
   420  .  .  0: *(obj @ 12)
   421  .  .  1: *(obj @ 20)
   422  .  }
   423  .  Unresolved: []*ast.Ident (len = 9) {
   424  .  .  0: *(obj @ 94)
   425  .  .  1: *(obj @ 106)
   426  .  .  2: *(obj @ 112)
   427  .  .  3: *(obj @ 273)
   428  .  .  4: *(obj @ 285)
   429  .  .  5: *(obj @ 291)
   430  .  .  6: *(obj @ 332)
   431  .  .  7: *(obj @ 364)
   432  .  .  8: *(obj @ 401)
   433  .  }
   434  }
```