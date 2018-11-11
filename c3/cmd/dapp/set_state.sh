#!/bin/sh

#echo "untar-data"
rm -rf ./data
tar -xf ./state.tar

#echo "pg-restore"
cat ./data/pgdump/*.sql | PGPASSWORD="postgres" psql -d postgres -h localhost -p 5432 -U postgres

echo "done setting state"
