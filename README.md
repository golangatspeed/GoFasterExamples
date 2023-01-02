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

9. [Recommended variable declaration styles](chapter_6/9_recommended_variable_declaration_styles/main.go)
10. [All variable declaration styles](chapter_6/10_all_variable_declaration_styles/main.go) (Go Playground: [https://go.dev/play/p/c8eFFqUfVTz](https://go.dev/play/p/c8eFFqUfVTz))
11. [Discarding with the blank identifier](chapter_6/11_discarding_with_the_blank_identifier/main.go) (Go Playground: [https://go.dev/play/p/NYjfpotoOtx](https://go.dev/play/p/NYjfpotoOtx))
12. [Loose typing of number constants](chapter_6/12_loose_typing_of_number_constants/main.go) (Go Playground: [https://go.dev/play/p/MC8HsTfnVz9](https://go.dev/play/p/MC8HsTfnVz9))
13. [Managing number constants manually](chapter_6/13_managing_number_constants_manually/main.go) (Go Playground: [https://go.dev/play/p/EVMSF3BCCn8](https://go.dev/play/p/EVMSF3BCCn8))
14. [Managing number constants with iota](chapter_6/14_managing_number_constants_with_iota/main.go) (Go Playground: [https://go.dev/play/p/h4fNoptnCJZ](https://go.dev/play/p/h4fNoptnCJZ))
15. [Rebasing the numbering from 1](chapter_6/15_rebasing_the_numbering_from_1/main.go)
16. [Non-linear constant sequences](chapter_6/16_non-linear_constant_sequences/main.go) (Go Playground: [https://go.dev/play/p/8QypEyOoxT2](https://go.dev/play/p/8QypEyOoxT2))
17. [Values used in place of named constants](chapter_6/17_values_used_in_place_of_named_constants/main.go) (Go Playground: [https://go.dev/play/p/n-POyR0-k7a](https://go.dev/play/p/n-POyR0-k7a))
18. [Global and function scope](chapter_6/18_global_and_function_scope/main.go) (Go Playground: [https://go.dev/play/p/V7OT8HX_XMP](https://go.dev/play/p/V7OT8HX_XMP))
19. [Local scope in control structure](chapter_6/19_local_scope_in_control_structure/main.go) (Go Playground: [https://go.dev/play/p/edNkZtJuCPk](https://go.dev/play/p/edNkZtJuCPk))
20. [Pass by value](chapter_6/20_pass_by_value/main.go) (Go Playground: [https://go.dev/play/p/uESKRz6r-tS](https://go.dev/play/p/uESKRz6r-tS))
21. [Pass by reference](chapter_6/21_pass_by_reference/main.go) (Go Playground: [https://go.dev/play/p/7Ne6OPnCslN](https://go.dev/play/p/7Ne6OPnCslN))
22. [Obtaining a pointer directly](chapter_6/22_obtaining_a_pointer_directly/main.go) (Go Playground: [https://go.dev/play/p/t3iJJ47laIt](https://go.dev/play/p/t3iJJ47laIt))

### Chapter 7 - Data types

23. Floating point number addition problem (Go Playground: [https://go.dev/play/p/_uoAHQeoh9h](https://go.dev/play/p/_uoAHQeoh9h))
24. String cut using byte slice (Go Playground: [https://go.dev/play/p/egXJA0gm97x](https://go.dev/play/p/egXJA0gm97x))
25. Identical strings with different lengths (Go Playground: [https://go.dev/play/p/ujUnmx-LsWu](https://go.dev/play/p/ujUnmx-LsWu))
26. Unicode codepoint with UTF-8 code (Go Playground: [https://go.dev/play/p/luDIj6DwPAG](https://go.dev/play/p/luDIj6DwPAG))
27. Safely obtaining length of a string (Go Playground: [https://go.dev/play/p/wTnoddQnjvJ](https://go.dev/play/p/wTnoddQnjvJ))
28. Array length, capacity and element initialisation (Go Playground: [https://go.dev/play/p/D_hs2NHHoAs](https://go.dev/play/p/D_hs2NHHoAs))
29. Array length is part of its type definition (Go Playground: [https://go.dev/play/p/xgNuJjZQHZM](https://go.dev/play/p/xgNuJjZQHZM))
30. Set array length with the spread operator (Go Playground: [https://go.dev/play/p/bxX2YplIPa8](https://go.dev/play/p/bxX2YplIPa8))
31. Named custom struct type and anonymous struct (Go Playground: [https://go.dev/play/p/gWnMhWF_OVK](https://go.dev/play/p/gWnMhWF_OVK))
32. Unnamed field properties (Go Playground: [https://go.dev/play/p/hun_q6dx-VC](https://go.dev/play/p/hun_q6dx-VC))
33. Get and set struct fields (Go Playground: [https://go.dev/play/p/gM5fADxC9gB](https://go.dev/play/p/gM5fADxC9gB))
34. Composition and struct embedding (Go Playground: [https://go.dev/play/p/6DSmzLgW3Xo](https://go.dev/play/p/6DSmzLgW3Xo))
35. Alignment and impact on memory (Go Playground: [https://go.dev/play/p/q7GJI1xfEqb](https://go.dev/play/p/q7GJI1xfEqb))
36. Trying to use a nil map (Go Playground: [https://go.dev/play/p/dOBBmtyVYuV](https://go.dev/play/p/dOBBmtyVYuV))
37. Creating an empty map (Go Playground: [https://go.dev/play/p/kDhy20FhD6B](https://go.dev/play/p/kDhy20FhD6B))
38. Map passed by value (Go Playground: [https://go.dev/play/p/i20LKvkaC0J](https://go.dev/play/p/i20LKvkaC0J))
39. Map passed by reference (Go Playground: [https://go.dev/play/p/lgiCDOYzSrj](https://go.dev/play/p/lgiCDOYzSrj))
40. Working with maps (Go Playground: [https://go.dev/play/p/6-ht0hCoQFj](https://go.dev/play/p/6-ht0hCoQFj))
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



