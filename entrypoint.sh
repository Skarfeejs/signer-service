#!/bin/sh +x

/opt/orbs/orbs-signer $@ | multilog t s16777215 n3 '!tai64nlocal' /opt/orbs/logs 2>&1