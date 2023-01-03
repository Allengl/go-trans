version=$(git describe --tags $(git rev-list --tags='v[0-9].[0-9]*' --max-count=1))
cd ..
go generate 
go build -ldflags "-s -w" -o sync-${version}.exe
