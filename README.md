# Description

I can't include C file twice. But I know that it's real because [this library](https://github.com/gographics/imagick) can.

I have main.go that wants to call some methods from C code. Also there is "wrapper" for C code in files `g/g1.go` and `g/g2.go` and C "library" in `g/c-api`.

All goes OK when I try to use one `.go` file in `g/` with one `#include "./c-api"` statement, but with two or more files I got error like

```
# _/Users/aliksend/Documents/cgoproof/g
duplicate symbol _greet2 in:
    /Users/aliksend/Documents/cgoproof/g/_obj/g1.cgo2.o
    /Users/aliksend/Documents/cgoproof/g/_obj/g2.cgo2.o
duplicate symbol _greet1 in:
    /Users/aliksend/Documents/cgoproof/g/_obj/g1.cgo2.o
    /Users/aliksend/Documents/cgoproof/g/_obj/g2.cgo2.o
ld: 2 duplicate symbols for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

# Steps to reproduce

```bash
go run main.go
```

or

```bash
docker build .
```
