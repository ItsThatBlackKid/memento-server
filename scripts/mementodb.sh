#!/usr/bin/env bash
cd "$(dirname "$0")" || exit
cd ../ # pwd = memento root

# create the db
sqlite3 db/memento.db < db/mementodb.sql

