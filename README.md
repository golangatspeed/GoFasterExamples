# Go Faster Example Code

The [Go Faster book](https://leanpub.com/gofaster) has many code examples to accompany the text.
The book is completed, at 200+ pages, 48,000+ words and includes over 145 Go code examples.

I'm including all the examples here, together with a link to the executable code on the Go Playground (where applicable).

There's no explainer text to accompany the examples, that's in the book, but the examples are designed to 
strip away the noise to highlight the principles being demonstrated.

They may well be enough as-is, but if not, please consider [purchasing my book](https://leanpub.com/gofaster). It's inexpensive and
thorough, and I'm confident it will help you learn Go, faster.

One caveat. Not all examples are demonstrations of best practice and some are specifically set up to generate compile time and runtime
errors :)

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
42. [Safe map access](chapter_7/42_safe_map_access/main.go) (Go Playground: [https://go.dev/play/p/a4zZyEM8viF](https://go.dev/play/p/a4zZyEM8viF))
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

64. [Using goto to restart function execution](chapter_8/64_using_goto_to_restart_function_execution/main.go) (Go Playground: [https://go.dev/play/p/5dsPU41vPQ7](https://go.dev/play/p/5dsPU41vPQ7))
65. [Simple function with conditional if/else logic](chapter_8/65_simple_function_with_conditional_if_else_logic/main.go) (Go Playground: [https://go.dev/play/p/hUPNR_4gRy_6](https://go.dev/play/p/hUPNR_4gRy_6))
66. [Eliminating the *else* statements](chapter_8/66_eliminating_the_else_statements/main.go) (Go Playground: [https://go.dev/play/p/lJm4fJ3REPa](https://go.dev/play/p/lJm4fJ3REPa))
67. [Short-form if statement](chapter_8/67_short_form_if_statement/main.go) (Go Playground: [https://go.dev/play/p/qmoU-gAtzCw](https://go.dev/play/p/qmoU-gAtzCw))
68. [Simple switch with default](chapter_8/68_simple_switch_with_default) (Go Playground:[ https://go.dev/play/p/3FiTf9lGEDd](https://go.dev/play/p/3FiTf9lGEDd))
69. [Expressionless switch statement](chapter_8/69_expressionless_switch_statement/main.go) (Go Playground: [https://go.dev/play/p/cx8jV4QXHRI](https://go.dev/play/p/cx8jV4QXHRI))
70. [Short-form switch with multiple match tests](chapter_8/70_short_form_switch_with_multiple_match_tests/main.go) (Go Playground: [https://go.dev/play/p/DjPtZbdigfm](https://go.dev/play/p/DjPtZbdigfm))
71. [Fallthrough to execute the next case](chapter_8/71_fallthrough_to_execute_the_next_case/main.go) (Go Playground: [https://go.dev/play/p/hw1WLkjYPUb](https://go.dev/play/p/hw1WLkjYPUb))
72. [Infinite loop with for](chapter_8/72_infinite_loop_with_for/main.go) (Go Playground: [https://go.dev/play/p/ZBcHX0L11Qq](https://go.dev/play/p/ZBcHX0L11Qq))
73. [Three component loop with for](chapter_8/73_three_component_loop_with_for/main.go) (Go Playground: [https://go.dev/play/p/IFE1-HGEGkQ](https://go.dev/play/p/IFE1-HGEGkQ))
74. [While equivalent using for](chapter_8/74_while_equivalent_using_for/main.go) (Go Playground: [https://go.dev/play/p/oC25p-YNNBx](https://go.dev/play/p/oC25p-YNNBx))
75. [Do while equivalent with for](chapter_8/75_do_while_equivalent_with_for/main.go) (Go Playground: [https://go.dev/play/p/SZUMlRfW3Fn](https://go.dev/play/p/SZUMlRfW3Fn))
76. [For each performed using idiomatic for range](chapter_8/76_for_each_performed_using_idiomatic_for_range/main.go) (Go Playground: [https://go.dev/play/p/CVjj4pUtc2c](https://go.dev/play/p/CVjj4pUtc2c))
77. [Using break to exit a loop](chapter_8/77_using_break_to_exit_a_loop/main.go) (Go Playground: [https://go.dev/play/p/o2vXkdAYNk3](https://go.dev/play/p/o2vXkdAYNk3))
78. [Using continue to advance to next loop iteration](chapter_8/78_using_continue_to_advance_to_next_loop_iteration/main.go) (Go Playground: [https://go.dev/play/p/qbVGN5S3hLY](https://go.dev/play/p/qbVGN5S3hLY))
79. [Using errors.Is to handle different error values](chapter_8/79_using_errors_is_to_handle_different_error_values/main.go) (Go Playground: [https://go.dev/play/p/gb2eikJ9Hp6](https://go.dev/play/p/gb2eikJ9Hp6))
80. [A custom error type](chapter_8/80_a_custom_error_type/main.go) (Go Playground: [https://go.dev/play/p/x7pZDTqda5O](https://go.dev/play/p/x7pZDTqda5O))
81. [Error wrapping and unwrapping](chapter_8/81_error_wrapping_and_unwrapping/main.go) (Go Playground: [https://go.dev/play/p/UzPyJIGc1DH](https://go.dev/play/p/UzPyJIGc1DH))
82. [Basic panic and recover](chapter_8/82_basic_panic_and_recover/main.go) (Go Playground: [https://go.dev/play/p/9vmgVqKTbU5](https://go.dev/play/p/9vmgVqKTbU5))
83. [Embedding log.Logger to augment its features](chapter_8/83_embedding_log_logger_to_augment_its_features/main.go) (Go Playground: [https://go.dev/play/p/ca17kF91BtL](https://go.dev/play/p/ca17kF91BtL))

### Chapter 9 - Digging deeper

84. [Passing values as variadic arguments](chapter_9/84_passing_values_as_variadic_arguments/main.go) (Go Playground: [https://go.dev/play/p/NPutwf2aF4l](https://go.dev/play/p/NPutwf2aF4l))
85. [Anonymous and named return parameters](chapter_9/85_anonymous_and_named_return_parameters/main.go) (Go Playground: [https://go.dev/play/p/RdVlD9CwJS4](https://go.dev/play/p/RdVlD9CwJS4))
86. [Bug risk or not?](chapter_9/86_bug_risk_or_not/main.go) (Go Playground: [https://go.dev/play/p/gHxfLdGgTaX](https://go.dev/play/p/gHxfLdGgTaX))
87. [Callback style functions](chapter_9/87_callback_style_functions/main.go) (Go Playground: [https://go.dev/play/p/KXSxlwcWJB4](https://go.dev/play/p/KXSxlwcWJB4))
88. [Closure functions](chapter_9/88_closure_functions/main.go) (Go Playground: [https://go.dev/play/p/x8XnvAM6O9C](https://go.dev/play/p/x8XnvAM6O9C))
89. [Pointer return for database connection](chapter_9/89_pointer_return_for_database_connection/main.go) (Go Playground: [https://go.dev/play/p/0afFBh0GaFG](https://go.dev/play/p/0afFBh0GaFG))
90. [Escape analysis with benchmark](chapter_9/90_escape_analysis_with_benchmark) (Go Playground: [https://go.dev/play/p/rqTcK20eVqV](https://go.dev/play/p/rqTcK20eVqV))
91. [Escape analysis on function returns](chapter_9/91_escape_analysis_on_function_returns) (Go Playground: [https://go.dev/play/p/rC5x05MpKP5](https://go.dev/play/p/rC5x05MpKP5))
92. [Value and pointer receivers](chapter_9/92_value_and_pointer_receivers/main.go) (Go Playground: [https://go.dev/play/p/qTBFdsIDdCL](https://go.dev/play/p/qTBFdsIDdCL))
93. [Simple file store](chapter_9/93_94_simple_file_store/filestore/filestore.go)
94. [Using the store](chapter_9/93_94_simple_file_store/main.go)
95. [Interface design](chapter_9/95_96_97_98_interface_design/store/store.go) 
96. [Implementing store.NewStore](chapter_9/95_96_97_98_interface_design/store/store.go)
97. [Implementing store.ReadWriteDeleter](chapter_9/95_96_97_98_interface_design/filestore/filestore.go)
98. [Refactoring to use the interface](chapter_9/95_96_97_98_interface_design/main.go)
99. [SendNotification function](chapter_9/99_100_sendnotification_function/main.go)
100. [A basic unit test](chapter_9/99_100_sendnotification_function/main_test.go)
101. [Notifier service interface](chapter_9/101_102_103_notifier_service_interface/service/service.go)
102. [Modifying the notification code](chapter_9/101_102_103_notifier_service_interface/notifications/notifications.go)
103. [Modify main](chapter_9/101_102_103_notifier_service_interface/main.go)
104. [Creating a mock notifications service](chapter_9/104_105_creating_a_mock_notifications_service/notifications/mock.go)
105. [A better unit test](chapter_9/104_105_creating_a_mock_notifications_service/main_test.go)
106. [Injecting makeCallToAPI](chapter_9/106_107_injecting_makecalltoapi/main.go)
107. [An alternative mock and unit test](chapter_9/106_107_injecting_makecalltoapi/main_test.go)
108. [Type assertion provides the concrete type](chapter_9/108_type_assertion_provides_the_concrete_type/main.go) (Go Playground: [https://go.dev/play/p/n9njIiD7Kk1](https://go.dev/play/p/n9njIiD7Kk1))
109. [Invalid type assertion, compiles but will panic](chapter_9/109_invalid_type_assertion_code_compiles_but_will_panic/main.go) (Go Playground: [https://go.dev/play/p/j3egltwQKPo](https://go.dev/play/p/j3egltwQKPo))
110. [Safely performing a type assertion](chapter_9/110_safely_performing_the_type_assertion/main.go) (Go Playground: [https://go.dev/play/p/3jfW7_hkRjM](https://go.dev/play/p/3jfW7_hkRjM))
111. [Type switch style assertion with default case](chapter_9/111_type_switch_with_default_case/main.go) (Go Playground: [https://go.dev/play/p/0ovjw0QFJtL](https://go.dev/play/p/0ovjw0QFJtL)) 
112. [Using reflection to discover a variable's type](chapter_9/112_using_reflection_to_discover_type_of_variable/main.go) (Go Playground: [https://go.dev/play/p/HFJswkAEnB5](https://go.dev/play/p/HFJswkAEnB5))
113. [Using reflection to determine if variable is a value or pointer](chapter_9/113_using_reflection_to_determine_if_value_or_pointer/main.go) (Go Playground: [https://go.dev/play/p/Y7p3BSuVGjN](https://go.dev/play/p/Y7p3BSuVGjN))
114. [Using reflection to work with a struct](chapter_9/114_using_reflection_to_work_with_a_struct/main.go) (Go Playground: [https://go.dev/play/p/D1GKlMeHs0V](https://go.dev/play/p/D1GKlMeHs0V))
115. [Performance overhead of reflection](chapter_9/115_performance_overhead_of_reflection/main.go)
116. [Adding two *int64* integers](chapter_9/116_adding_two_int64_integers/main.go) (Go Playground: [https://go.dev/play/p/XDU49bgxRs_j](https://go.dev/play/p/XDU49bgxRs_j))
117. [Adding other integer types](chapter_9/117_adding_other_integer_types/main.go) (Go Playground: [https://go.dev/play/p/Q3SzagyF-CS](https://go.dev/play/p/Q3SzagyF-CS))
118. [A single "sumIntAny" implementation](chapter_9/118_a_single_sumintany_implementation/main.go) (Go Playground: [https://go.dev/play/p/GZ0JOBSwu9_n](https://go.dev/play/p/GZ0JOBSwu9_n))
119. [A generic implementation of SumAny (Go 1.18+)](chapter_9/119_a_generic_implementation_of_sumany/main.go) (Go Playground: [https://go.dev/play/p/jbzggpV7uM8](https://go.dev/play/p/jbzggpV7uM8))
120. [Using a type constraint list (Go 1.18+)](chapter_9/120_using_a_type_constraint_list/main.go) (Go Playground: [https://go.dev/play/p/yZemwfvY6ST](https://go.dev/play/p/yZemwfvY6ST))
121. [Creating a custom constraint (Go 1.18+)](chapter_9/121_creating_a_custom_constraint/main.go) (Go Playground: [https://go.dev/play/p/IxpNJBWPkZo](https://go.dev/play/p/IxpNJBWPkZo))
122. [Using composition with constraints (Go 1.18+)](chapter_9/122_using_composition_with_constraints/main.go) (Go Playground: [https://go.dev/play/p/2yARr7xNBG8](https://go.dev/play/p/2yARr7xNBG8))
123. [Allow any type with underlying type int (Go 1.18+)](chapter_9/123_allow_any_type_with_underlying_type_int/main.go) (Go Playground: [https://go.dev/play/p/b8jNMrEtUcL](https://go.dev/play/p/b8jNMrEtUcL))
124. [Multiple type parameters (Go 1.18+)](chapter_9/124_multiple_type_parameters/main.go) (Go Playground: [https://go.dev/play/p/DFeMUaFqBzG](https://go.dev/play/p/DFeMUaFqBzG))

### Chapter 10 - Concurrency

125. [Setting GOMAXPROCS during runtime](chapter_10/125_setting_GOMAXPROCS_during_runtime/main.go) (Go Playground:[https://go.dev/play/p/ptFRl7NAoT_6](https://go.dev/play/p/ptFRl7NAoT_6))
126. [Where is the output?](chapter_10/126_where_is_the_output/main.go) (Go Playground:[https://go.dev/play/p/nd5nwv5b3ok](https://go.dev/play/p/nd5nwv5b3ok))
127. [Using context.WithCancel()](chapter_10/127_using_context_withcancel/main.go) (Go Playground:[https://go.dev/play/p/aAbmie_rtVV](https://go.dev/play/p/aAbmie_rtVV))
128. [Using context.WithTimeout()](chapter_10/128_uising_context_withtimeout/main.go) (Go Playground:[https://go.dev/play/p/6y1hYI5pSdM](https://go.dev/play/p/6y1hYI5pSdM))
129. [Using context.WithDeadline()](chapter_10/129_using_context_withdeadline/main.go) (Go Playground:[https://go.dev/play/p/MyjT3EVGkus](https://go.dev/play/p/MyjT3EVGkus))
130. [Sharing data via context](chapter_10/130_sharing_data_via_context/main.go) (Go Playground:[https://go.dev/play/p/7p504omEWsq](https://go.dev/play/p/7p504omEWsq))
131. [Waitgroup updating shared state](chapter_10/131_waitgroup_updating_shared_state) (Go Playground:[https://go.dev/play/p/Z3Rh0ufgiVr](https://go.dev/play/p/Z3Rh0ufgiVr))
132. [WaitGroup updating shared state unsafely](chapter_10/132_waitgroup_updating_shared_state_unsafely/main.go) (Go Playground:[https://go.dev/play/p/dW1kdCrqmCu](https://go.dev/play/p/dW1kdCrqmCu))
133. [Mutex to guarantee exclusivity of access](chapter_10/133_mutex_to_guarantee_exclusivity_of_access/main.go) (Go Playground:[https://go.dev/play/p/ZObOF_af_W9](https://go.dev/play/p/ZObOF_af_W9))
134. [Signalling channel for mutual exclusion](chapter_10/134_signalling_channel_for_mutual_exclusion/main.go) (Go Playground:[https://go.dev/play/p/JHiaz5Rjvii](https://go.dev/play/p/JHiaz5Rjvii))
135. [Sizing the buffer to store results](chapter_10/135_sizing_the_buffer_to_store_results/main.go) (Go Playground:[https://go.dev/play/p/YVTWPzUYht9](https://go.dev/play/p/YVTWPzUYht9))
136. [Unidirectional channel types](chapter_10/136_unidirectional_channel_types/main.go) (Go Playground:[https://go.dev/play/p/ztggMJsU0cj](https://go.dev/play/p/ztggMJsU0cj))
137. [Unbuffered channel for signalling](chapter_10/137_unbuffered_channel_for_signalling/main.go) (Go Playground:[https://go.dev/play/p/ro8OHTv4G4x](https://go.dev/play/p/ro8OHTv4G4x))
138. [Signalling by closing the channel](chapter_10/138_signalling_by_closing_the_channel/main.go) (Go Playground:[https://go.dev/play/p/ciX1rKK-XAV](https://go.dev/play/p/ciX1rKK-XAV))
139. [Signalling multiple goroutines by closing the channel](chapter_10/139_signalling_multiple_goroutines_by_closing_the_channel/main.go) (Go Playground:[https://go.dev/play/p/vHFbaV5Tyg3](https://go.dev/play/p/vHFbaV5Tyg3))
140. [Reading from a closed channel](chapter_10/140_reading_from_a_closed_channel/main.go) (Go Playground:[https://go.dev/play/p/iedP9__fCzR](https://go.dev/play/p/iedP9__fCzR))
141. [Reading from a channel with range](chapter_10/141_reading_from_a_channel_with_range/main.go) (Go Playground:[https://go.dev/play/p/Qdo3mWvKmmg](https://go.dev/play/p/Qdo3mWvKmmg))
142. [Reading from a closed channel with range](chapter_10/142_reading_from_a_closed_channel_with_range/main.go) (Go Playground:[https://go.dev/play/p/X5kcrsmOpQW](https://go.dev/play/p/X5kcrsmOpQW))
143. [Monitoring multiple channels with select](chapter_10/143_monitoring_multiple_channels_with_select/main.go) (Go Playground:[https://go.dev/play/p/L8d4dFaVta3](https://go.dev/play/p/L8d4dFaVta3))

### Chapter 11 - Quality Assurance

144. [Example test with subtests](chapter_11/144_example_test_with_subtests/main.go) (Go Playground:[https://go.dev/play/p/2NNWfknTbXC](https://go.dev/play/p/2NNWfknTbXC))
145. [Examples with ordered and unordered output](chapter_11/145_examples_with_ordered_and_unordered_output/main.go) (Go Playground:[https://go.dev/play/p/RRE5gXK3A87](https://go.dev/play/p/RRE5gXK3A87))
146. [Modify code to obtain a CPU profile](chapter_11/146_modify_code_to_obtain_a_cpu_profile/main.go)
147. [Realtime profiling during program execution](chapter_11/147_realtime_profiling_during_program_execution/main.go)



