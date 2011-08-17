
CC=6g
L=6l
AR=gopack

mkdir lib

echo "collection.go"
6g -o lib/collection.6 src/gocollection/collection.go
echo "Add to archive"
gopack gr lib/gocollection.a lib/collection.6

echo "sllist.go"
6g -I lib -o lib/sllist.6 src/gocollection/sllist.go

echo "dllist.go"
6g -I lib -o lib/dllist.6 src/gocollection/dllist.go

echo "Add to archive"
gopack gr lib/gocollection.a lib/sllist.6 lib/dllist.6


