
CC=6g
L=6l
AR=gopack

mkdir lib


echo "collection.go"
6g -o lib/gocollection.6 src/gocollection/gocollection.go

6g -I lib  -o lib/golists.6 src/gocollection/golists/*.go


mkdir lib/test

echo "Test1"
6g -I lib -o lib/test/p1.6 test/p1.go
6l -L lib -o lib/test/p1 lib/test/p1.6

