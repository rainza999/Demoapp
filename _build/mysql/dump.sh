#!/bin/bash

for file in $(ls /home/*.sql); do 
    mysql -uroot -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE < $file
done