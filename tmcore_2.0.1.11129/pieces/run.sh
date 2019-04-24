#!/usr/bin/env bash

set -euo pipefail

if [ "${TMHOME:-}" == "" ] && [ "$HOME" != "/home/$(whoami)" ]; then
    export TMHOME=/etc/tmcore
fi
export PATH=/usr/local/tmcore/bin:/usr/bin:/bin:/usr/sbin:/sbin

clean_old() {
    pid=$(ps -futmcore|grep 'tendermint node'|awk '$0 !~/grep/ {print $2}'|sed -e 's/\n/ /')
    if [ "${pid:-}" != "" ]; then
        echo "kill old process."
        kill -9 ${pid}
    fi
}

start() {
    echo "run node ->->->->->"
    [ -d "/home/tmcore/log" ] || ( mkdir /home/tmcore/log; chmod 775 /home/tmcore/log )
    while true
    do
     tendermint node | (TS=`date "+%F %T"`;sed -e "s/^/[${TS}] tmcore - /") >>/home/tmcore/log/tmcore.out 2>>/home/tmcore/log/tmcore.out
     sleep 1
    done
    date >> /home/tmcore/log/tmcore.out
    exit 1
}

getChainID() {
    cd /etc/tmcore/genesis
    dirs=$(ls -d */ | tr -d "\/")
    for d in $dirs; do
      echo $d
      return
    done
}

if [ "${1:-}" == "" ] || [ "${1:-}" == "start" ] || [ "${1:-}" == "restart" ]; then
    clean_old
    start
    exit 0
fi

if [ "${1:-}" == "stop" ]; then
    clean_old
    exit 0
fi

if [[ "${1:-}" == "init" ]]; then
    chainID=$(getChainID)
    genesisPath="/etc/tmcore/genesis/${chainID}"
    nodeCfg="${genesisPath}/${chainID}-nodes.json"
    followCfg="${genesisPath}/${chainID}-followers.json"
    
	  echo ""
    echo "Select which role that the node will run"
		cmain="VALIDATOR -- be sure all the validator nodes are installing at the same time"
		cofficial="OFFICIAL FOLLOWER"
		cunofficial="UNOFFICIAL FOLLOWER"

		idx=0
		choices2=("")
    if [ -f ${nodeCfg} ]; then
      choices2[$idx]=$cmain
      idx=$(($idx+1))
    fi
    if [ -f ${followCfg} ]; then
      choices2[$idx]=$cofficial
      idx=$(($idx+1))
    fi
    choices2[$idx]=$cunofficial
    select nodeType in "${choices2[@]}"; do
        case ${nodeType} in
        $cmain)
            echo "You selected VALIDATOR node"
            echo ""
            echo "Initializing all genesis node..."
            tendermint init --genesis_path ${genesisPath}
						echo ""
						echo ""
						version=$(tendermint version | grep "build version: " | tr -d "build version: ")
						echo "Congratulation !!! TENDERMINT is successfully installed with version ${version}."
						echo ""
            exit 0
            ;;
        $cofficial)
            echo "You selected OFFICIAL follower"
            echo ""
            echo "Select your node's name"
            echo
            choices3=($(jq '.[].name' ${followCfg}))
            select followerName in "${choices3[@]}"; do
                [[ -n ${followerName} ]] || { echo "Invalid choice." >&2; continue; }

                privateIp=$(jq -r ".[${REPLY}-1].ip_priv" ${followCfg})
                if [[ "${privateIp}" != "null" && "${privateIp}" != "" ]]; then
                    myIPs=$(ip -4 -o addr | awk '{split($4,a,"/");print a[1]}')
                    ok=0
                    for ip in ${myIPs}; do
                        if [[ "${ip}" == "${privateIp}" ]]; then
                        ok=1
                        fi
                    done
                    if [[ "${ok}" == "0" ]]; then
                        echo "${followerName}'s privateIp not your's, maybe you made a mistake." >&2; continue
                    fi
                else
                    echo "${followerName} has no ip_priv in config file, what a horrible..." >&2; continue
                fi
                voters=$(jq -r '.[].public' ${nodeCfg}|paste -d, -s)
                portN=$(jq -r ".[${REPLY}-1].listen_port" ${followCfg})
                aAddr=$(jq -r ".[${REPLY}-1].announce" ${followCfg})
                proxy_app=$(jq -r ".[${REPLY}-1].apps[0]" ${followCfg})
                CMD="tendermint init --genesis_path ${genesisPath} --follow ${voters}"
                if [[ "${portN}" != "null" && "${portN}" != "" ]]; then
                    CMD="${CMD} --listen_port ${portN}"
                fi
                if [[ "${aAddr}" != "null" && "${aAddr}" != ""  ]]; then
                    CMD="${CMD} --a_address ${aAddr}"
                fi
                if [[ "${proxy_app}" != "null" && "${proxy_app}" != "" ]]; then
                    CMD="${CMD} --proxy_app ${proxy_app}"
                fi
                ${CMD}
                exit 0
            done
            ;;
        $cunofficial)
            echo ""
            echo "You selected UNOFFICIAL follower"
            echo ""
            echo "Please input the OFFICIAL follower's name[:port] you want to follow"
            echo "Multi nodes can be separated by comma ,"
            echo "for example \"earth.bcbchain.io,mar.bcbchain.io:46657\" or \"venus.bcbchain.io\""
            read -p "officials to follow: " officials
            tendermint init --genesis_path ${genesisPath} --follow ${officials}
            exit 0
            ;;
        *) echo "Invalid choice.";;
        esac
    done

fi
