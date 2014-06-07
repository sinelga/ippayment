#! /bin/bash
#chmod +x!! must be done
#*/2 * * * * /home/juno/git/ippayment/ippayment/cronhitshandlersendsms.bash

cd /home/juno/git/ippayment/ippayment

pgrep -x hitshandler || bin/hitshandler
pgrep -x sendsms || bin/sendsms