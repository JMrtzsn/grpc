package main

// "makefile" to generate grpc files

//go:generate mkdir -p pb
//go:generate protoc --go_out=plugins=grpc:pb --go_opt=paths=source_relative outliers.proto

// go_out:  place the output files in the pb dir
// go_opts: generate code in the relative source - PB
// outliers.proto - name of the protocol buffers grpc def file
