#!/usr/bin/env bash

set -uo pipefail

export PATH=/usr/bin:/bin:/usr/sbin:/sbin

if ! $(systemctl -q is-enabled tmcore.service 2>/dev/null) ; then
    exit 0
fi
protocol="http://"
[[ -f /etc/tmcore/config/STAR.bcbchain.io.crt ]] && protocol="https://"

myPort=$(grep -m 1 "^laddr =" /etc/tmcore/config/config.toml|awk -F: '{print $3}'|sed -e 's/"//')

myStatusTempFile=$(mktemp /tmp/tmcore-myStatus-XXXXXXXX)
(curl --max-time 3 --connect-timeout 2 --silent -k ${protocol}localhost:${myPort}/status 2>&1) > ${myStatusTempFile}
imSyncing=$(cat ${myStatusTempFile} | awk '/syncing/{print $2}')

myHeight=$(cat ${myStatusTempFile} | awk '/latest_block_height/{gsub(",","");print $2}')
myHeight=${myHeight:--1}
rm -f "${myStatusTempFile}"

now=$(date +%s)

stopIt() {
    ps -futmcore|grep "tendermint node"|grep -v grep|awk '{print "kill -HUP "$2}'|sh
    >&2 echo restarted
}

iGrow() {
    echo ${myHeight} > /tmp/tmcore.monit.${now}
    for f in $(ls /tmp/tmcore.monit.*); do
        t=$(echo ${f} | cut -d'.' -f 3)
        lifeOfSecond=$(expr ${now} - ${t})
        if [ ${lifeOfSecond} -gt 600 ]; then
            rm -f ${f}
            continue
        elif [ ${lifeOfSecond:-0} -gt 400 ]; then
            historyHeight=$(cat ${f})
            grow=$(expr ${myHeight} - ${historyHeight})
            if [ ${grow} -eq 0 ]; then
                rm -f /tmp/tmcore.monit.*
                # echo 100
                stopIt
                return
            fi
        fi
    done
    # echo 0
}

if [ ${myHeight} -eq -1 ]; then
    # echo -1 -1
    stopIt
fi

if ! ${imSyncing:-false} ; then
    PEERS=$(curl --max-time 3 --connect-timeout 2 --silent -k ${protocol}localhost:${myPort}/net_info | awk -F\" '/listen_addr/{print $4}')
    peerHeight=0

    for PEER in ${PEERS}; do
        P_IP="$(echo ${PEER}|cut -d':' -f1)"
        P_PORT="$(echo ${PEER}|cut -d':' -f2)"
        P_PORT=$((P_PORT + 1))
        peerStatusTempFile=$(mktemp /tmp/tmcore-peerStatus-XXXXXXXX)
        (curl --max-time 3 --connect-timeout 2 --silent -k ${protocol}${P_IP}:${P_PORT}/status 2>&1) > ${peerStatusTempFile}
        peerSyncing=$(expr `cat ${peerStatusTempFile} | awk '/syncing/{print $2}'`)
        if ${peerSyncing:-false} ; then
            rm -f "${peerStatusTempFile}"
            continue
        fi
        peerHeight=$(cat ${peerStatusTempFile} | awk '/latest_block_height/{gsub(",","");print $2}')
        peerHeight=${peerHeight:--1}
        rm -f "${peerStatusTempFile}"
        if [ ${peerHeight} -ne -1 ]; then
            break
        fi
    done

    if [ ${peerHeight} -ne -1 ]; then
        subVal=$(expr ${peerHeight:-0} - ${myHeight})
        # echo $subVal
        if [ ${subVal} -gt 10 ]; then
            # echo gt 10
            stopIt
        fi
    fi
fi

iGrow
# iii=$(iGrow)
# if [ $iii -ne 0 ]; then
#    exit -1
# fi

exit 0
