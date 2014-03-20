grep id syslog* >phones.log

scp root@107.170.68.93:/var/log/phones.log .

cat phones.log |awk '{print "curl -G -H \"X-UP-CALLING-LINE-ID: " $7\" }' |sort



curl -G -H "X-UP-CALLING-LINE-ID: 358451202825" -d "id=0000000" -d "site=test2.com" -d "resource=mobilephone" -d "themes=adult"  test.com/sonera?callback=lslslslsl


cat phones.log |awk '{print "curl -G -H \"X-UP-CALLING-LINE-ID: " $7 "\" -d \"id=0000000\" -d \"site=test2.com\" -d \"resource=mobilephone\" -d \"themes=adult\"  test.com/sonera?callback=lslslslsl"  }' |sort > insert.bash