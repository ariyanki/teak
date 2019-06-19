export TEAK_ENV=testing
export TEAK_APP_PATH=$(pwd)
go run main.go migrate reset
go run main.go migrate up
go run main.go seed
go test $1