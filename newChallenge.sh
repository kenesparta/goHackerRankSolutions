#!/bin/bash

# pass the path to create the directory challenge

cat > /tmp/fileFmt <<-EOF
/*
URL			: $2
AUTHOR		: $3
DIFFICULTY	: $4
*/

package main

func main(){

}
EOF

mkdir $1;

cd $1;

touch inputs;

touch main.go;

cat /tmp/fileFmt > main.go

go fmt main.go;