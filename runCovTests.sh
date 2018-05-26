cd cache
go test -coverprofile=coverage.out
sleep 15
cd ..

cd circuitbreaker
go test -coverprofile=coverage.out
sleep 15
cd ..

cd database
go test -coverprofile=coverage.out
sleep 15

cd mysqldb
go test -coverprofile=coverage.out
sleep 15
cd ../..

cd environment
go test -coverprofile=coverage.out
sleep 15
cd ..

cd gwerrors
go test -coverprofile=coverage.out
sleep 15
cd ..

cd handlers
go test -coverprofile=coverage.out
sleep 15
cd ..

cd managers
go test -coverprofile=coverage.out
sleep 15
cd ..

cd monitor
go test -coverprofile=coverage.out
sleep 15
cd ..