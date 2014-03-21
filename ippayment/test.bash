grep id syslog* >phones.log

scp root@107.170.68.93:/var/log/phones.log .
scp juno@107.170.68.93:git/ippayment/ippayment/syslog .

cat phones.log |awk '{print "curl -G -H \"X-UP-CALLING-LINE-ID: " $7\" }' |sort


curl -G -H "X-UP-CALLING-LINE-ID: 358451202825" -d "id=0000000" -d "site=test2.com" -d "resource=mobilephone" -d "themes=adult" -d "provider=sonera" test.com/sonera?callback=lslslslsl


cat phones.log |awk '{print "curl -G -H \"X-UP-CALLING-LINE-ID: " $7 "\" -d \"id=0000000\" -d \"site=test2.com\" -d \"resource=mobilephone\" -d \"themes=adult\"  test.com/sonera?callback=lslslslsl"  }' |sort > insert.bash

go/bin/tiedot -mode=httpd -dir=git/ippayment/ippayment/tiedotDB -port=8080

 curl --data-ascii q='{"eq": "358451202801", "in": ["ClPhonenum"]}' "http://localhost:8080/query?col=mobclients"
 
 curl "http://localhost:8080/delete?col=mobclients&id=4371515570365599128"