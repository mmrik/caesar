# caesar
A tiny Caesar cipher to be invoked from the command line. Just for fun

# Requirements
- Go 1.15

# Installation
Clone repository, and from the commandline run `go install`. Make sure that your Go bin directory is part of your PATH variable

# Usage
`caesar <word or sentencence to shift>`

Example: `caesar "uryyb jbeyq"` outputs `hello world`. By default caesar defaults to a ROT13 cipher (thirteen character shift in the english language).
Other supported languages are German and Swedish

To change language, pass the flag `-language` with either the value `de` or `sv`

Example: `caesar -language=sv "Xuz Ioeätua"` outputs `Hej Världen`

To change the amount of character to shift, pass the `-shift=<amount>` flag. Both positive and negative values work.
Example: `caesar -shift=-4 giewev"` outputs `caesar`
