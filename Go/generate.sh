protoc -I src/ --go_out=src/ src/simple/simple.proto
protoc -I src/ --go_out=src/ src/enum/enum.proto
protoc -I src/ --go_out=src/ src/complex/complex.proto