#! /bin/bash
#chmod +x!! must be done
#*/2 * * * * /home/juno/git/ippayment/ippayment/crongofastscript.bash

cd /home/juno/git/ippayment/ippayment

pgrep hitshandler || bin/hitshandler
pgrep sendsms || bin/sendsms