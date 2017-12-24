FILE STRUCTURE
==============
I chose to design this program as a "triangles" package accompanied with a "main" program. I make an effort to write
idiomatic code, and as far as I've gathered, this layout is a common way to structure Go projects. It could be argued that
this small project would be easier written in just a single file, but I find it easier to test when the business logic
is kept in a separate package. Especially as Go has pretty nice unit testing functionality built in.

IMPLEMENTATION
==============
First of all I try to familiarize myself with the problem domain (in this case triangles)
For this assignment, I assume that Wikipedia's information on triangles is good enough.

The triangle inequality states that the sum of two sides must be greater than the remaining side.
Here we have our first test case. For this assignment, I will consider degenerate triangles (triangles with no area) as invalid.

My original triangle inequality check looked like this:
```
if t.a+t.b <= t.c || t.a+t.c <= t.b || t.b+t.c <= t.a {
  return false
}
```
However, I realized that if the program input was really large integer numbers, there is a possibility that the
integer addition would overflow. I therefore refactored the above code to the following:
```
if t.a <= t.c-t.b || t.a <= t.b-t.c || t.b <= t.a-t.c {
  return false
}
```

As for classifying the triangle (once I know it's valid) I looked at the Wikipedia description of the different types and
wrote tests for the different types. The implementation was straightforward.

When writing the main program, I chose to use the standard library's flag package since it would validate the input for me and
allow me to add some usage text easily.

