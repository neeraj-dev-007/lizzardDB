%{
package main

import (
"fmt"
)


type Expr struct {
    Type  string
    Value string
}

%}



%union {
    str  string
    expr *Expr
}

%token SELECT INSERT UPDATE DELETE CREATE TABLE FROM WHERE JOIN ON INTO VALUES SET
%token IDENTIFIER STRING_LITERAL INTEGER_LITERAL
%token EQUALS NOT_EQUALS LESS GREATER LESS_EQUALS GREATER_EQUALS
%token COMMA LPAREN RPAREN
%token INT VARCHAR TEXT DATE FLOAT

%%

statement:
select_statement
| insert_statement
| update_statement
| delete_statement
| create_statement
;

select_statement:
SELECT column_list FROM table_name optional_join_clause optional_where_clause
;

optional_join_clause:
join_clause
| /* empty */
;

join_clause:
JOIN table_name ON condition
| join_clause JOIN table_name ON condition
;

optional_where_clause:
WHERE condition
| /* empty */
;

column_list:
'*'
| IDENTIFIER column_list_tail
;

column_list_tail:
COMMA IDENTIFIER column_list_tail
| /* empty */
;

condition:
IDENTIFIER comparison_operator value
;

comparison_operator:
EQUALS
| NOT_EQUALS
| LESS
| GREATER
| LESS_EQUALS
| GREATER_EQUALS
;

table_name:
IDENTIFIER
;

value:
STRING_LITERAL
| INTEGER_LITERAL
;

insert_statement:
INSERT INTO table_name LPAREN column_list RPAREN VALUES LPAREN value_list RPAREN
;

value_list:
value value_list_tail
;

value_list_tail:
COMMA value value_list_tail
| /* empty */
;

create_statement:
CREATE TABLE table_name LPAREN column_definition_list RPAREN
;

column_definition_list:
column_definition column_definition_list_tail
;

column_definition_list_tail:
COMMA column_definition column_definition_list_tail
| /* empty */
;

column_definition:
IDENTIFIER data_type
;

data_type:
INT
| VARCHAR LPAREN INTEGER_LITERAL RPAREN
| TEXT
| DATE
| FLOAT
;

update_statement:
UPDATE table_name SET update_clause optional_where_clause
;

update_clause:
IDENTIFIER EQUALS value update_clause_tail
;

update_clause_tail:
COMMA IDENTIFIER EQUALS value update_clause_tail
| /* empty */
;

delete_statement:
DELETE FROM table_name optional_where_clause
;

%%
func main() {
    yyParse()
}

func yyError(s string) {
    fmt.Printf("Error: %s\n", s)
}