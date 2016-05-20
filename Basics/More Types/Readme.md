#More Types

## Pointers
Go has pointers. A pointer holds the memory address of a variable.

The type `*T` is a pointer to a T value. Its zero value is nil.

```
var p *int
```
The & operator generates a pointer to its operand.

```
i := 42
p = &i
```
The * operator denotes the pointer's underlying value.
```
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```
This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.

##
Structs
A struct is a collection of fields.

(And a type declaration does what you'd expect.)
```
> go run structs.go
{1 2}
```

##
Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.
```
> go run struct-pointers.go
{1000000000 2}
```

## Struct Literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.
```
> go run struct-literals.go
{1 2} &{1 2} {1 0} {0 0}
```
