# Template Engine Syntax

The template engine renders the following:
- Variables (STRING) - Implemented
- Conditionals Statements - Pending Implementation
- Loops (either for, while, range) - Pending Implementation

## Variables
The template engine allows rendering Strings.

### Syntax
- A variable can be named using the following characteres `a-z,A-Z,0-9,_`. You can this regular expression to test if your variable name is valid. `{{*.[a-zA-Z0-9_].*}}`
- Variable must be enclosed within two curly brackets is considered a variable
- Space arround the variable is not required but recommended for readability


``` bash
{{ variable }}
{{variable}}
{{i_am_a_valid_variable}}
{{i_am_a_valid_variable_1}}
{{_i_am_a_valid_variable_2}}
{{_I_Am_A_Valid_Variable_3}}
{{_iAmAValidVariable4}}
```

## Conditional Statements
{% condition %}

## Loops
{% range %} - if no limit end of variable
{% range limit %}

<!-- EOF -->

