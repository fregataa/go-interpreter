# go-interpreter

## Composition

The interpreter will have those components.
the lexer, the parser, the AST, the internal object system, and the evaluator.

### Lexer

In order for us to work with source code we need to turn it into a more accessible form.
We're going to change the representation of our source code two times before we evaluate it.

> Source Code -> Tokens -> Abstract Syntax Tree

The firs transformation is called "lexical analysis", or "lexing".
It's done by a lexer. Tokens itself are small, easily categorizable data structures that are then fed to the parser,
which does the second transformation and turns the tokens into an "AST".
