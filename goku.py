import lark
import rich

# Gramática Corrigida
grammar = r"""
start: program

program: package imports* function*

package: "package" NAME
imports: "import" STRING

function: "func" NAME "(" arg_list? ")" "{" statement* "}"

arg_list: parameter ("," parameter)*
parameter: NAME NAME // Ex: x int (Nome do argumento e Nome do tipo)

?statement: call
call: NAME "()"

NAME: /\w+/
STRING: /"[^"]*"/
%ignore /[ \t\n\r]+/
"""

program_code = """
package main

import "fmt"

func main() {
a()
}

func a(x int) {

}
"""

parser = lark.Lark(grammar, start="start")
tree = parser.parse(program_code)
rich.print(tree)
