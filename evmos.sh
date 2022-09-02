#!/bin/sh

go install
cd E2E/interE2E
go install
go install kt/exe/...

tail -f /dev/null