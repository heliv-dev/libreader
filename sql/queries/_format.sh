#!/bin/sh

/usr/bin/find . -name "*.sql" -print | xargs -I {} sh -c 'sqlformat -r -s -o {} {} ; echo sqlformat: {}'