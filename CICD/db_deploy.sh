#!/bin/bash/

#vars
dump_file=$(date  +%Y-%m-%d-%H-%M-%S)_dump.sql


#create pg_dump
docker exec -i psgr-master /bin/bash -c \
"PGPASSWORD=${PGPASSWORD} pg_dump --username postgres postgres --column-inserts --on-conflict-do-nothing" > ${dump_file}

#send to server
scp ${dump_file} vklimov@timewb-diceroll:/home/vklimov/pg_dumps/${dump_file}

#run script for unpack dump
ssh vklimov@timewb-diceroll "bash update_pgs.sh" && echo "ok"

rm ${dump_file}

