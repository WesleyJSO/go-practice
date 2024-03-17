# run test and generate the cover file
go test -coverprofile=c.out -coverpkg=./src/main/... ./src/test/

# generate the html version and open
go tool cover -html=c.out