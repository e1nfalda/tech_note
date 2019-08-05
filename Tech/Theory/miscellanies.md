## miscellanies

### EBNF（扩展的巴科斯范式）

> *metasyntax*一种。和`BNF`语法稍有差异。python使用的BNF,golang使用的EBNF。[wike]([https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form))

|                            Usage                             | Notation  |
| :----------------------------------------------------------: | :-------: |
|                          definition                          |     =     |
| [concatenation](https://en.wikipedia.org/wiki/Concatenation) |     ,     |
|                         termination                          |     ;     |
| [alternation](https://en.wikipedia.org/wiki/Alternation_(formal_language_theory)) |    \|     |
|                           optional                           |  [ ... ]  |
|                          repetition                          |  { ... }  |
|                           grouping                           |  ( ... )  |
|                       terminal string                        |  " ... "  |
|                       terminal string                        |  ' ... '  |
|                           comment                            | (* ... *) |
|                       special sequence                       |  ? ... ?  |
|                          exception                           |     -     |