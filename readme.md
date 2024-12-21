# Memento Server
The Memento web service written in Go.


## Env Variables
Environment variables are loaded from a `.env` which must be created in the root of the project dir.

These are the environment variables required to run `memento-server`
* __DB:__ path to sqlite database
* __TOKEN_SECRET:__ key used to sign JWT