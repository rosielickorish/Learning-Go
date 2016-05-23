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

## Structs
A struct is a collection of fields.

(And a type declaration does what you'd expect.)
```
> go run structs.go
{1 2}
```

## Pointers to structs
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

##Arrays
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
```
> go run array.go
Hello World
[Hello World]
[2 3 5 7 11 13]
```

## Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

This expression creates a slice of the first five elements of the array a:

a[0:5]

```
> go run slices.go
[3 5 7]
```

## Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.
```
> go run slices-pointers.go
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

## Slice literals
A slice literal is like an array literal without the length.

This is an array literal:

[3]bool{true, true, false}
And this creates the same array as above, then builds a slice that references it:

[]bool{true, true, false}
