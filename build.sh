idl_dir=${IDL_DIR:=../service-idl}
module_name=$(grep '^module' go.mod | awk '{print $2}')
files=$(find $idl_dir -type f -name '*.proto')
if [ $? != 0 ];then
  exit 1
fi
for file in $files
do
  proto_path=$(dirname "$file")
    file_base=$(basename "$file" .proto)
    goctl rpc protoc --proto_path=$proto_path $file --go_out=./go-zero-gen/$file_base --go-grpc_out=./go-zero-gen/$file_base --zrpc_out=./go-zero-gen/$file_base --style=goZero
    rm -rf go-zero-gen/$file_base/etc
    rm -rf go-zero-gen/$file_base/internal
    rm -rf go-zero-gen/$file_base/$file_base.go
    if [ $? != 0 ];then
      exit 2
    fi
done