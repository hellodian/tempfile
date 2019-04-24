#!/usr/bin/env bash

set -uo pipefail
export PATH=/usr/bin:/bin:/usr/sbin:/sbin
. common.sh 

usage1() {
  echo ""
  echo "---------------------------CAUTION----------------------------"
  echo "Please make a update request because bcchain exists, but chain"
  echo "or data is corrupted."
  echo ""
  exit 1
}

usage2() {
  echo ""
  echo "---------------------------CAUTION----------------------------"
  echo "Please make a update request because bcchain application data "
  echo "exists, not a clean install."
  echo ""
  exit 1
}

uid=$(id -u)
if [[ "${uid:-}" != "0" ]]; then
  echo "must be root user"
  exit 1
fi

if $(systemctl -q is-active bcchain.service) ; then
  systemctl stop bcchain
  usage2
fi

uid=$(id -u bcchain 2>/dev/null)
if [ "${uid:-}" != "" ]; then
  pid=$(ps -fubcchain|grep 'bcchain'|awk '$0 !~/grep/ {print $2}'|sed -e 's/\n/ /')
  if [ "${pid:-}" != "" ]; then
      echo "kill old process. ${pid}"
      kill -9 ${pid}
  fi
fi

isCorrupted=$(isCorruptedChain)
if [ $isCorrupted = "true" ] ; then
  usage1
fi

isCompleted=$(isCompletedChain)
if [ $isCompleted = "true" ] ; then
  usage2
fi

echo ""
echo "Select which CHAINID to install"
dirs=$(ls -d */ | tr -d "\/")
choices1=($dirs)
select chainID in "${choices1[@]}"; do
  [[ -n ${chainID} ]] || { echo "Invalid choice." >&2; continue; }
  echo "You selected CHAINID=${chainID}"
  echo ""
  doCopyFiles ${chainID}
  break
done

echo ""
version=$(./bcchain version | tr -d "\r")
echo "Congratulation !!! BCCHAIN is successfully installed with version ${version}."
echo ""

exit 0
