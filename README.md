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

23. [Floating point number addition problem](chapter_7/23_floating_point_number_addition_problem/main.go) (Go Playground: [https://go.dev/play/p/_uoAHQeoh9h](https://go.dev/play/p/_uoAHQeoh9h))
24. [String cut using byte slice](chapter_7/24_string_cut_using_byte_slice/main.go) (Go Playground: [https://go.dev/play/p/egXJA0gm97x](https://go.dev/play/p/egXJA0gm97x))
25. [Identical strings with different lengths](chapter_7/25_identical_strings_with_different_lengths/main.go) (Go Playground: [https://go.dev/play/p/ujUnmx-LsWu](https://go.dev/play/p/ujUnmx-LsWu))
26. [Unicode codepoint with UTF-8 code](chapter_7/26_unicode_codepoint_with_utf-8-code/main.go) (Go Playground: [https://go.dev/play/p/luDIj6DwPAG](https://go.dev/play/p/luDIj6DwPAG))
27. [Safely obtaining length of a string](chapter_7/27_safely_obtaining_length_of_a_string/main.go) (Go Playground: [https://go.dev/play/p/wTnoddQnjvJ](https://go.dev/play/p/wTnoddQnjvJ))
28. [Array length, capacity and element initialisation](chapter_7/28_array_length_capacity_and_element_initialisation/main.go) (Go Playground: [https://go.dev/play/p/D_hs2NHHoAs](https://go.dev/play/p/D_hs2NHHoAs))
29. [Array length is part of its type definition](chapter_7/29_array_length_is_part_of_its_type_definition/main.go) (Go Playground: [https://go.dev/play/p/xgNuJjZQHZM](https://go.dev/play/p/xgNuJjZQHZM))
30. [Set array length with the spread operator](chapter_7/30_set_array_length_with_the_spread_operator/main.go) (Go Playground: [https://go.dev/play/p/bxX2YplIPa8](https://go.dev/play/p/bxX2YplIPa8))
31. [Named custom struct type and anonymous struct](chapter_7/31_named_custom_struct_type_and_anonymous_struct/main.go) (Go Playground: [https://go.dev/play/p/gWnMhWF_OVK](https://go.dev/play/p/gWnMhWF_OVK))
32. [Unnamed field properties](chapter_7/32_unnamed_field_properties/main.go) (Go Playground: [https://go.dev/play/p/hun_q6dx-VC](https://go.dev/play/p/hun_q6dx-VC))
33. [Get and set struct fields](chapter_7/33_get_and_set_struct_fields/main.go) (Go Playground: [https://go.dev/play/p/gM5fADxC9gB](https://go.dev/play/p/gM5fADxC9gB))
34. [Composition and struct embedding](chapter_7/34_composition_and_struct_embedding/main.go) (Go Playground: [https://go.dev/play/p/6DSmzLgW3Xo](https://go.dev/play/p/6DSmzLgW3Xo))
35. [Alignment and impact on memory](chapter_7/35_alignment_and_impact_on_memory/main.go) (Go Playground: [https://go.dev/play/p/q7GJI1xfEqb](https://go.dev/play/p/q7GJI1xfEqb))
36. [Trying to use a nil map](chapter_7/36_trying_to_use_a_nil_map/main.go) (Go Playground: [https://go.dev/play/p/dOBBmtyVYuV](https://go.dev/play/p/dOBBmtyVYuV))
37. [Creating an empty map](chapter_7/37_creating_an_empty_map/main.go) (Go Playground: [https://go.dev/play/p/kDhy20FhD6B](https://go.dev/play/p/kDhy20FhD6B))
38. [Map passed by value](chapter_7/38_map_passed_by_value/main.go) (Go Playground: [https://go.dev/play/p/i20LKvkaC0J](https://go.dev/play/p/i20LKvkaC0J))
39. [Map passed by reference](chapter_7/39_map_passed_by_reference/main.go) (Go Playground: [https://go.dev/play/p/lgiCDOYzSrj](https://go.dev/play/p/lgiCDOYzSrj))
40. [Working with maps](chapter_7/40_working_with_maps/main.go) (Go Playground: [https://go.dev/play/p/6-ht0hCoQFj](https://go.dev/play/p/6-ht0hCoQFj))
41. [Unsafe map access](chapter_7/41_unsafe_map_access/main.go) (Go Playground: [https://go.dev/play/p/9jX2B7x1eha](https://go.dev/play/p/9jX2B7x1eha))
42. [Safe map access](chapter_7/42_safe_map_access/main.go) (Go Playground: [https://go.dev/play/p/F_HX9A1ExNG](https://go.dev/play/p/F_HX9A1ExNG))
43. [Creating a slice](chapter_7/43_creating_a_slice/main.go) (Go Playground: [https://go.dev/play/p/922R9tr-3Aq](https://go.dev/play/p/922R9tr-3Aq))
44. [Slice not behaving as a reference type](chapter_7/44_slice_not_behaving_as_a_reference_type/main.go) (Go Playground: [https://go.dev/play/p/0cOx0y5Xc_L](https://go.dev/play/p/0cOx0y5Xc_L))
45. [Safely return the new slice to caller](chapter_7/45_safely_return_the_new_slice_to_caller/main.go) (Go Playground: [https://go.dev/play/p/A-Lg98S0GRJ](https://go.dev/play/p/A-Lg98S0GRJ))
46. [Slice behaving as a reference type](chapter_7/46_slice_behaving_as_a_reference_type/main.go) (Go Playground: [https://go.dev/play/p/CvKjbcRyOAr](https://go.dev/play/p/CvKjbcRyOAr))
47. [Reslicing & working with slices](chapter_7/47_reslicing_&_working_with_slices/main.go) (Go Playground: [https://go.dev/play/p/Aiw7WhYGePH](https://go.dev/play/p/Aiw7WhYGePH))
48. [Slices of slices](/chapter_7/48_slices_of_slices/main.go) (Go Playground: [https://go.dev/play/p/NAWtoteoj5x](https://go.dev/play/p/NAWtoteoj5x))
49. [Using append](chapter_7/49_using_append/main.go) (Go Playground: [https://go.dev/play/p/WUPIBwbVCEL](https://go.dev/play/p/WUPIBwbVCEL))
50. [Reslicing to remove a specified element from a slice](chapter_7/50_reslicing_to_remove_a_specified_element_from_a_slice/main.go) (Go Playground: [https://go.dev/play/p/e_IrhpQHBXE](https://go.dev/play/p/e_IrhpQHBXE))
51. [Using copy to create a slice with new backing array](chapter_7/51_using_copy_to_create_a_slice_with_new_backing_array/main.go) (Go Playground: [https://go.dev/play/p/YHVH-gE9lYW](https://go.dev/play/p/YHVH-gE9lYW))
52. [New Slice backing array with additional capacity](chapter_7/52_new_slice_backing_array_with_additional_capacity/main.go) (Go Playground: [https://go.dev/play/p/JMbFUxzyNzK](https://go.dev/play/p/JMbFUxzyNzK))
53. [Send and receive on unbuffered channel](chapter_7/53_send_and_receive_on_unbuffered_channel/main.go) (Go Playground: [https://go.dev/play/p/f8M2OFSLnJA](https://go.dev/play/p/f8M2OFSLnJA))
54. [The Stringer interface](chapter_7/54_the_stringer_interface/main.go) (Go Playground: [https://go.dev/play/p/xl1phIcBej1](https://go.dev/play/p/xl1phIcBej1))
55. [Implementing Stringer on a custom struct type](chapter_7/55_implementing_stringer_on_a_custom_struct_type/main.go) (Go Playground: [https://go.dev/play/p/Oa4sBfwsx78](https://go.dev/play/p/Oa4sBfwsx78))
56. [Simple transposition error](chapter_7/56_simple_transposition_error/main.go) (Go Playground: [https://go.dev/play/p/XhTIz_A-dxd](https://go.dev/play/p/XhTIz_A-dxd))
57. [Type safety using custom types](chapter_7/57_type_safety_using_custom_types/main.go) (Go Playground: [https://go.dev/play/p/6UQcR54_wYX](https://go.dev/play/p/6UQcR54_wYX))
58. [Implementing the Stringer interface on a custom type](chapter_7/58_implementing_the_stringer_interface_on_a_custom_type/main.go) (Go Playground: [https://go.dev/play/p/RSGm_FPhf6L](https://go.dev/play/p/RSGm_FPhf6L))
59. [Type conversion examples](chapter_7/59_type_conversion_examples/main.go) (Go Playground: [https://go.dev/play/p/FOkyY6XJp1q](https://go.dev/play/p/FOkyY6XJp1q))
60. [This won't print what you expect](chapter_7/60_this_wont_print_what_you_expect/main.go) (Go Playground: [https://go.dev/play/p/GHVwfyD-mlH](https://go.dev/play/p/GHVwfyD-mlH))
61. [fmt.Sprintf to the rescue](chapter_7/61_fmt_sprintf_to_the_rescue/main.go) (Go Playground: [https://go.dev/play/p/b6APk-xVK9X](https://go.dev/play/p/b6APk-xVK9X))
62. [Package strconv examples](chapter_7/62_package_strconv_examples/main.go) (Go Playground: [https://go.dev/play/p/5kW2FyxBGgU](https://go.dev/play/p/5kW2FyxBGgU))
63. [A custom BoolToI helper function](chapter_7/63_a_custom_booltoi_helper_function) (Go Playground: [https://go.dev/play/p/gseJ2vzyDSc](https://go.dev/play/p/gseJ2vzyDSc))

### Chapter 8 - Managing program flow

### Chapter 9 - Digging deeper

### Chapter 10 - Concurrency

### Chapter 11 - Quality Assurance



