# A Tour of Go

## Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

## Variables

The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.

## Variables with initializers

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

## Short variable declarations

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
##Basic types

Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

The example shows variables of several types, and also that variable declarations may be factored into blocks, as with import statements.

```
go run basic-types.go
bool(false)
uint64(18446744073709551615)
complex128((2+3i))
```
## Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

    0 for numeric types,
    false for the boolean type, and
     (the empty string) for strings.
```
$go run zero.go
0 0 false ""
```
##Type conversions

The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

Or, put more simply:

i := 42
f := float64(i)
u := uint(f)

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.
```
go run type-conversion.go
3 4 5
```
##Type inference

When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

var i int
j := i // j is an int

But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128

Try changing the initial value of v in the example code and observe how its type is affected.
```
go run type-inference.go
v is of type float64
w is of type int
```
##Constants

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

```
go run constants.go
Hello 世界
Happy 3.14 Day
Go rules? true
```
##Numeric Constants

Numeric constants are high-precision values.

An untyped constant takes the type needed by its context.

Try printing needInt(Big) too.

(An int can store at maximum a 64-bit integer, and sometimes less.)

```
go run numeric-constants.go
21
0.2
1.2676506002282295e+29
```
