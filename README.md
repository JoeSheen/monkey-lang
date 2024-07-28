# Monkey Lang

This repository contains my implementation of the monkey programming language defined by [Thorsten Ball](https://thorstenball.com/) in his books [Writing An Interpreter In Go](https://interpreterbook.com/) and [Writing A Compiler In Go](https://compilerbook.com/).

## Running the REPL

1. Open a terminal
2. `git clone https://github.com/JoeSheen/monkey-lang.git`
3. `cd .\monkey-lang\`
4. `go run main.go`

## Syntax

```
let fibonacci = fn(x) {
  if (x == 0) {
    return 0;
  } else {
    if (x == 1) {
      return 1;
    } else {
      return fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
let val = fibonacci(10);
puts(val) // prints 55 to console
```
