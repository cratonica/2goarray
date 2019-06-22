2goarray
========
A simple utility to encode a file (or any other data) into a Go byte slice.

Having [set up your Go environment](http://golang.org/doc/install), simply run

    go get github.com/cratonica/2goarray

Then use by piping a file into the utility and capturing the output. You must provide a name for the generated slice symbol and package name. For example:

    $GOPATH/bin/2goarray MyArray mypackage < myimage.png > myimage.go

Or in PowerShell, to loop over a folder of icons and output a .go file corresponding to the icon's name:

```powershell
$2goarray = Join-Path $Env:GOPATH '\bin\2goarray.exe'
Get-ChildItem *.ico | Foreach-Object {
  Get-Content $_.FullName | Invoke-Expression $2goarray $_.BaseName "mypackage" | Out-File -FilePath $($_.BaseName + ".go") -Encoding utf8
}
```

This will output something like:

    package mypackage

    var MyArray []byte = []byte {
      0x49, 0x20, 0x63, 0x61, 0x6e, 0x27, 0x74, 0x20, 0x62, 0x65, 0x6c, 0x69,
      0x65, 0x76, 0x65, 0x20, 0x79, 0x6f, 0x75, 0x20, 0x61, 0x63, 0x74, 0x75,
      0x61, 0x6c, 0x6c, 0x79, 0x20, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64,
      0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x20, 0x4b, 0x75, 0x64, 0x6f, 0x73,
      0x20, 0x66, 0x6f, 0x72, 0x20, 0x62, 0x65, 0x69, 0x6e, 0x67, 0x20, 0x74,
      0x68, 0x6f, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x2e, 0x0a,
    }

## Contributors
- [Clint Caywood](https://github.com/cratonica)
- [Paul Vollmer](https://github.com/paulvollmer)
