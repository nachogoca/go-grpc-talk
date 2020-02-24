Compile protoc on Ubuntu

 sudo apt-get install autoconf automake libtool curl make g++ unzip

1. Download latest Protocol Buffers release (protobuf-all-[Version].tar.gz)

https://github.com/protocolbuffers/protobuf/releases

2. Extract contents

tar -zxvf protobuf-all-3.11.4.tar.gz

3. `./configure`

4. `make`

5. `make check`

6. `sudo make install`

7. `sudo ldconfig`

8. Check if it works

`protoc --version`

9. Install Go proto package

go get -u github.com/golang/protobuf/protoc-gen-go

10. Install Google grpc package

go get -u google.golang.org/grpc



