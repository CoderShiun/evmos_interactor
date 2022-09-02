#!/bin/sh

go install
cd E2E/interE2E
go install

tail -f /dev/null