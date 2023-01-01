# Go Faster Example Code

Work in progress.

The [Go Faster book](https://leanpub.com/gofaster) has many code examples to accompany the text.
The book is approximatly 70% complete, at 125 pages and includes almost 90 Go code examples.

I'm listing them here together with a link to the executable code on the Go Playground (where applicable).

There's no explainer text to accompany the examples, that's in the book, but the examples are designed to 
strip away the noise to highlight the principles being demonstrated.

They may well be enough as-is, but, if not please consider [purchasing my book](https://leanpub.com/gofaster). It's inexpensive and
thorough and I'm confident it will help you learn Go, faster.

One caveat, not all examples are demonstrations of best practice :)

## Contents

### Chapter 1 - Introduction to Go

1. [Visibility modifiers](chapter_1/1_visibilty_modifiers/exporting.go)
2. [Comment styles](chapter_1/2_comment_styles/main.go)
3. [Making the most of comments in documentation](chapter_1/3_making_the_most_of_comments_in_documentation/user.go)
4. [An 'example' function](chapter_1/4_an_example_function/stringutil_test.go)

### Chapter 2 - The Go command line interface (CLI)

*No Go code examples used*

### Chapter 3 - Structure of a Go program

5. [Hello World in Go](chapter_3/5_hello_world_in_go/main.go) (Go Playground: [https://go.dev/play/p/HbshOM2vDe](https://go.dev/play/p/HbshOM2vDed))
6. [How not to use aliasing](chapter_3/6_how_not_to_use_aliasing/main.go) (Go Playground: [https://go.dev/play/p/zuXJ0If5OFn](https://go.dev/play/p/zuXJ0If5OFn))
7. [The dot import prefix](chapter_3/7_the_dot_import_prefix/main.go) (Go Playground: [https://go.dev/play/p/SEUbOnOb6wH](https://go.dev/play/p/SEUbOnOb6wH))
8. [Simple package](chapter_3/8_simple_package/sideeffect.go)

### Chapter 4 - Project organisation

*No Go code examples used*

### Chapter 5 - Dependency management

*No Go code examples used*

### Chapter 6 - Variables and constants

9. Recommended variable declaration styles
10. All variable declaration styles (Go Playground: [https://go.dev/play/p/c8eFFqUfVTz](https://go.dev/play/p/c8eFFqUfVTz))
11. Discarding with the blank identifier (Go Playground: [https://go.dev/play/p/NYjfpotoOtx](https://go.dev/play/p/NYjfpotoOtx))
12. Loose typing of number constants (Go Playground: [https://go.dev/play/p/MC8HsTfnVz9](https://go.dev/play/p/MC8HsTfnVz9))
13. Managing number constants manually (Go Playground: [https://go.dev/play/p/EVMSF3BCCn8](https://go.dev/play/p/EVMSF3BCCn8))
14. Managing number constants with iota (Go Playground: [https://go.dev/play/p/h4fNoptnCJZ](https://go.dev/play/p/h4fNoptnCJZ))
15. Rebasing the numbering from 1
16. Non-linear constant sequences (Go Playground: [https://go.dev/play/p/8QypEyOoxT2](https://go.dev/play/p/8QypEyOoxT2))
17. Values used in place of named constants (Go Playground: [https://go.dev/play/p/n-POyR0-k7a](https://go.dev/play/p/n-POyR0-k7a))
18. Global and function scope (Go Playground: [https://go.dev/play/p/V7OT8HX_XMP](https://go.dev/play/p/V7OT8HX_XMP))
19. Local scope in control structure (Go Playground: [https://go.dev/play/p/edNkZtJuCPk](https://go.dev/play/p/edNkZtJuCPk))
20. Pass by value (Go Playground: [https://go.dev/play/p/uESKRz6r-tS](https://go.dev/play/p/uESKRz6r-tS))
21. Pass by reference (Go Playground: [https://go.dev/play/p/7Ne6OPnCslN](https://go.dev/play/p/7Ne6OPnCslN))
22. Obtaining a pointer directly (Go Playground: [https://go.dev/play/p/t3iJJ47laIt](https://go.dev/play/p/t3iJJ47laIt))

### Chapter 7 - Data types

23. Floating point number addition problem
24. String cut using byte slice
25. Identical strings with different lengths
26. Unicode codepoint with UTF-8 code
27. Safely obtaining length of a string
28. Array length, capacity and element initialisation
29. Array length is part of its type definition
30. Set array length with the spread operator
31. Named custom struct type and anonymous struct
32. Unnamed field properties
33. Get and set struct fields
34. Composition and struct embedding
35. Alignment and impact on memory
36. Trying to use a nil map
37. Creating an empty map
38. Map passed by value
39. Map passed by reference
40. Working with maps
41. Unsafe map access
42. Safe map access
43. Creating a slice
44. Slice not behaving as a reference type
45. Safely return the new slice to caller
46. Slice behaving as a reference type
47. Reslicing & working with slices
48. Slices of slices
49. Using append
50. Reslicing to remove a specified element from a slice
51. Using copy to create a slice with new backing array
52. New Slice backing array with additional capacity
53. Send and receive on unbuffered channel
54. The Stringer interface
55. Implementing Stringer on a custom struct type
56. Simple transposition error
57. Type safety using custom types
58. Implementing the Stringer interface on a custom type
59. Type conversion examples
60. This won't print what you expect
61. fmt.Sprintf to the rescue
62. Package *strconv* examples
63. A custom BoolToI helper function

### Chapter 8 - Managing program flow

### Chapter 9 - Digging deeper

### Chapter 10 - Concurrency

### Chapter 11 - Quality Assurance



