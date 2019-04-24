#!/usr/bin/env bash

set -uo pipefail
export PATH=/usr/bin:/bin:/usr/sbin:/sbin
. common.sh

usage() {
  echo ""
  echo "---------------------------CAUTION---------------------------"
  echo "Please make a install request because tmcore is clean.       "
  echo ""
  exit 1
}

getChainID() {
    cd /etc/tmcore/genesis
    for d in */; do
      echo $d
      return
    done
}

removeOldData() {
  bash clean.sh
}

uid=$(id -u)
if [[ "${uid:-}" != "0" ]]; then
  echo "must be root user"
  exit 1
fi

if $(systemctl -q is-active tmcore.service 2>/dev/null) ; then
  systemctl stop tmcore.service
fi

uid=$(id -u tmcore 2>/dev/null)
if [[ "${uid:-}" != "0" ]];then
  pid=$(ps -futmcore|grep 'tendermint'|awk '$0 !~/grep/ {print $2}'|sed -e 's/\n/ /')
  if [ "${pid:-}" != "" ]; then
      echo "kill old process. ${pid}"
      kill -9 ${pid}
  fi
fi

isEmpty=$(isEmptyNode)
if [ $isEmpty = "true" ] ; then
  usage
fi

oldVersion=0.0.0.0
if [ $emptyTendermint = "false" ]; then
	oldVersion=$(/usr/local/tmcore/bin/tendermint version | grep "build version: " | tr -d "build version: ")
fi

isCorrupted=$(isCorruptedNode)
if [ $isCorrupted = "true" ] ; then
  echo ""
  echo Old node or data is corrupted, do you want to remove all of this node to reinstall?
  options=("yes" "no")
  select opt in "${options[@]}" ; do
  case ${opt} in
    "yes")
      echo "Yes, remove all of this node to reinstall"		  
		  rm -rf /home/tmcore /usr/local/tmcore /etc/tmcore
			echo ""
			echo "Select which CHAINID to install"
			dirs=$(ls -d */ | tr -d "\/")
			choices1=($dirs)
			select chainID in "${choices1[@]}"; do
			  [[ -n ${chainID} ]] || { echo "Invalid choice." >&2; continue; }
			  echo "You selected CHAINID=${chainID}"
        echo ""
			  doCopyFiles ${chainID}
			  su - tmcore -s /bin/bash -c "/usr/local/tmcore/bin/run.sh init" 
			  exit 0
			  break
			done
      break
      ;;
    "no")
      echo "No, keep old node and data, need to be repaired manually"
      echo ""
      exit 1
      break
      ;;
    *) echo "Invalid choice.";;
    esac
  done
fi

isCompleted=$(isCompletedNode)
if [ $isCompleted = "true" ] ; then
  echo ""
  echo Old data exists, do you want to remove all data to re-sync?
  options=("yes" "no")
  select opt in "${options[@]}" ; do
  case ${opt} in
    "yes")
      echo "Yes, remove old data"
      echo ""
      removeOldData
      break
      ;;
    "no")
      echo "No, keep old data"
      echo ""
      break
      ;;
    *) echo "Invalid choice.";;
    esac
  done
fi

chainID=$(getChainID)
doCopyFiles ${chainID}

version=$(head -1 version | tr -d "\r")

echo ""
echo "Congratulation !!! TENDERMINT is successfully updated from ${oldVersion} to version ${version}."
echo ""
