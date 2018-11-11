#!/bin/sh

echo "dump"
rm -rf ./data/pgdump
mkdir -p ./data/pgdump
PGPASSWORD="postgres" pg_dump -d postgres -h localhost -p 5432 -U postgres > ./data/pgdump/postgres.sql

echo "sort-dump"
python ./pg_dump_splitsort.py ./data/pgdump/postgres.sql

echo "clean-dump"
rm ./data/pgdump/postgres.sql ./data/pgdump/*status.sql

echo "tar-data"
rm ./state.tar
tar -cf ./state.tar ./data
