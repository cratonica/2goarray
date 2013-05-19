2goarray
========

A simple utility to encode a file (or any other data) into a Go byte array.

Usage
-----
Just provide the name of the slice and the name of the package. The tool
will convert whatever is piped through stdin and output a Go file.

#### Example

    2goarray MyArray mypackage < inputfile.png > outputfile.go
