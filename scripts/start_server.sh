#!/bin/bash

MAIN_PID=`cat /usr/local/app/main.pid`
MAIN_IsAlive=`ps ${MAIN_PID} | wc -l`

if [ ${MAIN_IsAlive} = 2 ]; then
    echo "echo server is alive now"
else
    echo "echo server is dead now"

    #ENV=$(hostname | cut -d. -f2)
    start_server --pid-file main.pid -- /usr/local/app/main 3000 &

    ## drop an email
    REPORT_MAIL_TO_ADDRESS=local@local.com
    mail -s "[${ENV}] re starting echo" ${REPORT_MAIL_TO_ADDRESS} < /dev/null
fi