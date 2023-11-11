#!/bin/bash

host='localhost'
user='user'
password='password'
db='disc-golf'
export PGPASSWORD=$password

psql -h $host -U $user -d $db -a -f ./init.sql
unset PGPASSWORD