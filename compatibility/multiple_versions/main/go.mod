module main

go 1.24.10

require (
    example.com/module1 v1.2.3
    example.com/module2 v1.2.3
)

replace example.com/module1 => ../module1
replace example.com/module2 => ../module2
