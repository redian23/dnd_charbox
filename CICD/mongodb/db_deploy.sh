#!/bin/bash/
source .var
#vars
dump_file=$(date  +%Y-%m-%d-%H-%M-%S)_data.dump

#create mongo_dump
docker exec -i mongodb-master /usr/bin/mongodump --username ${user} --password ${passwd} --authenticationDatabase admin --db data --archive > ${dump_file}

#send to server
scp ${dump_file} vklimov@timewb-diceroll:/home/vklimov/mongo_dumps/${dump_file}

#run script for unpack dump
ssh vklimov@timewb-diceroll "bash update_db.sh" && echo "ok"

rm ${dump_file}

