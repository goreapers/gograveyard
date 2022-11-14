module replace

go 1.19

require (
	golang.org/x/net v1.2.3
    golang.org/x/http v1.2.1
    golang.org/x/http v1.3.4
)

replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
replace golang.org/x/http => example.com/fork/http v1.19.0
