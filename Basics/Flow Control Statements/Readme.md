##For
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or Javascript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
```
> go run for.go
45
```
##For continued
The init and post statement are optional.

```
> go run for-continued.go
1024
```
##For is Go's while
At that point you can drop the semicolons: C's while is spelled for in Go.
```
> go run for-is-gos-while.go
1024
```

##Forever
If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```
> go run fovever.go
process took too long
```

##If
Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

```
> go run if.go
1.4142135623730951 2i
```

## If with a short statement
Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

(Try using v in the last return statement.)

```
> go run if-with-a-short-statement.go
9 20
```

##If and else
Variables declared inside an if short statement are also available inside any of the else blocks.

(Both calls to pow are executed and return before the call to fmt.Println in main begins.)
```
> go run if-and-else.go
27 >= 20
9 20
```

##Exercise: Loops and Functions
As a simple way to play with functions and loops, implement the square root function using Newton's method.

In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:

```
z_(n+1) = z_n - (z^2_n - x)/2z_n
```

To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

```
z := float64(1)
z := 1.0
```

## Switch
You probably knew what switch was going to look like.

A case body breaks automatically, unless it ends with a fallthrough statement.

```
> go run switch.go
Go runs on OS X.
```
