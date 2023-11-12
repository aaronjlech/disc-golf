#!/bin/bash

host='127.0.0.1'
user='user'
password='password'
db='disc-golf'
export PGPASSWORD=$password

psql -h $host -U $user -d $db -a -f ./init.sql
unset PGPASSWORD