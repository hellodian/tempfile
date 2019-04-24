#!/usr/bin/env bash

set -uo pipefail
export PATH=/usr/bin:bin:/usr/sbin:sbin
. common.sh

usage() {
  echo ""
  echo "---------------------------CAUTION---------------------------"
  echo "Please make a install request because bcchain is clean.      "
  echo ""
  exit 1
}

getChainID() {
    echo $(cat /etc/bcchain/genesis 2>/dev/nul)
    return
}

removeOldData() {
  bash clean.sh
}

uid=$(id -u)
if [[ "${uid:-}" != "0" ]]; then
  echo "must be root user"
  exit 1
fi

if $(systemctl -q is-active bcchain.service) ; then
  systemctl stop bcchain.service
fi

uid=$(id -u bcchain 2>/dev/null)
if [[ "${uid:-}" != "0" ]];then
  pid=$(ps -fubcchain|grep 'bcchain'|awk '$0 !~/grep/ {print $2}'|sed -e 's/\n/ /')
  if [ "${pid:-}" != "" ]; then
      echo "kill old process. ${pid}"
      kill -9 ${pid}
  fi
fi

isEmpty=$(isEmptyChain)
if [ $isEmpty = "true" ] ; then
  usage
fi

oldVersion=0.0.0.0
if [ $emptyBcchain = "false" ]; then
	oldVersion=$(/usr/local/bcchain/bin/bcchain version)
fi

isCorrupted=$(isCorruptedChain)
if [ $isCorrupted = "true" ] ; then
  echo ""
  echo Old chain or data is corrupted, do you want to remove all of this chain to reinstall?
  options=("yes" "no")
  select opt in "${options[@]}" ; do
  case ${opt} in
    "yes")
      echo "Yes, remove all of this chain to reinstall"		  
		  rm -rf /home/bcchain /usr/local/bcchain /etc/bcchain
			echo ""
			echo "Select which CHAINID to install"
			dirs=$(ls -d */ | tr -d "\/")
			choices1=($dirs)
			select chainID in "${choices1[@]}"; do
			  [[ -n ${chainID} ]] || { echo "Invalid choice." >&2; continue; }
			  echo "You selected CHAINID=${chainID}"
        echo ""
			  doCopyFiles ${chainID}
				echo ""
				version=$(./bcchain version | tr -d "\r")
				echo "Congratulation !!! BCCHAIN is successfully installed with version ${version}."
				echo ""
			  exit 0
			  break
			done
      break
      ;;
    "no")
      echo "No, keep old chain, need to be repaired manually"
      echo ""
      exit 1
      break
      ;;
    *) echo "Invalid choice.";;
    esac
  done
fi

isCompleted=$(isCompletedChain)
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
echo "Congratulation !!! BCCHAIN is successfully updated from ${oldVersion} to version ${version}."
echo ""
