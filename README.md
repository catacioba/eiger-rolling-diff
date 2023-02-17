Eiger Rolling Diff
------------------

This repository contains a solution for [this](https://github.com/eqlabs/recruitment-exercises/blob/8e49a7b8cf9c415466876e852fbd862f74105ec6/rolling-hash.md)
coding challenge. It tries to implement an algorithm similar to RDiff signature/delta but in human readable form.

It uses 2 hash functions: a weaker rolling hash function that is fast to compute and a more expensive MD5.
The rolling function used is a very simple one: the base 256 representation of the bytes modulo 65521.
This was a good start and the code can easily be adapted to support more hash functions like Adler-32.
MD5 is used when we have a match for the rolling function to reduce the chance of false positives.

The algorithm prints in human readable form what to copy from the old file and what to add
to the new file in order to obtain the new file. This is enough to handle deleted data because what
was deleted should be neither copied nor added from/to the old file to obtain the new file.

Usage:

    go run cmd/main.go --chunkSize=<chunk_size> <old_file> <new_file>  


TODO: 
* Add CLI subcommands
* Change chunk aggregation from `map[weaksum]->metadata` to `map[weaksum]->list[metadata]` to handle
  weaksum collisions.
* Change Signature representation
* Change Delta representation
* Add E2E tests
* Add other rolling functions like Adler-32
