# kul
kul (`"kk-no's util"`) is my personal command line tool.  
More features may or may not be added.

## install
If you have go 1.16 or higher, you can use `go install` to install it.
```sh
go install github.com/kk-no/kul@latest
```

## commands
version
```shell
$ kul version
> kul 0.1.0
```

uuid
```shell
# uuid generate
$ kul uuid
> ff489297-fa0c-403c-944d-1b389ebd4eb5

# uuid generate (upper case)
$ kul uuid -u
> 675A0C14-A1EF-4E6E-8ADD-2DE2CAD18372
```

count
```shell
# string length count
$ kul count argument
> 8
```

base64
```shell
# base64 encode
$ kul base encode argument
> YXJndW1lbnQ=

# base64 decode
$ kul base decode YXJndW1lbnQ=
> argument
```