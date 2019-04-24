#!/usr/bin/env bash

COLUMNS=12
  
isEmptyFileOrDir() {
	dir=$1
	if [[ -z "$(ls -A ${dir} 2>/dev/null)" ]]; then
	  echo "true"
  else
    echo "false"
  fi
}

emptySdk=$(isEmptyFileOrDir "/home/bcchain/.build/sdk")
emptyThird=$(isEmptyFileOrDir "/home/bcchain/.build/thirdparty")
emptyClean=$(isEmptyFileOrDir "/home/bcchain/clean.sh")
emptyData=$(isEmptyFileOrDir "/home/bcchain/.appstate.db")
emptyBcchain=$(isEmptyFileOrDir "/usr/local/bcchain/bin/bcchain")

isEmptyChain() {
	if [[ ${emptySdk} == "true" ]] && [[ ${emptyThird} == "true" ]] && [[ ${emptyClean} == "true" ]] && [[ ${emptyData} == "true" ]] && [[ ${emptyBcchain} == "true" ]]; then
	  echo "true"
	  return
	else
	  echo "false"
	  return
	fi
}

isCompletedChain() {
	if [[ ${emptySdk} == "false" ]] && [[ ${emptyThird} == "false" ]] && [[ ${emptyClean} == "false" ]] && [[ ${emptyData} == "false" ]] && [[ ${emptyBcchain} == "false" ]]; then
	  echo "true"
	  return
	else
	  echo "false"
	  return
	fi
}

isCorruptedChain() {
	t=$(isEmptyChain)
	if [[ ${t} = "true" ]]; then
	  echo "false"
	  return
	fi
	t=$(isCompletedChain)
	if [[ ${t} = "true" ]]; then
	  echo "false"
	  return
	fi
	echo "true"
	return
}

doCopyFiles() {
  chainID=$1
  echo "Start copying files ..."
	
	mkdir -p /etc/bcchain /home/bcchain/{log,.appstate.db,.build} /usr/local/bcchain/bin
	echo ${chainID} > /etc/bcchain/genesis
	
	getent group bcchain  >/dev/null 2>&1 || groupadd -r bcchain
	getent passwd bcchain  >/dev/null 2>&1 || useradd -r -g bcchain \
	    -d /home/bcchain -s /sbin/nologin -c "BlockChain application System User" bcchain
	usermod -d /home/bcchain -g bcchain bcchain 2>/dev/null

	systemctl stop docker
	rm -f /run/docker.* 2>/dev/null
	systemctl start docker
	usermod -G $(ls -g /run/docker.sock|awk '{print $3}') bcchain 2>/dev/null
		
  cp ${chainID}/.config/abci-forks.json* /usr/local/bcchain/bin
  cp ${chainID}/.config/{bcchain.yaml,*.tar.gz} /etc/bcchain
	cp bcchain runApp.sh /usr/local/bcchain/bin
	cp start.sh stop.sh clean.sh /home/bcchain 
	cp bcchain.service /usr/lib/systemd/system
	
	chmod 644 /etc/bcchain/*
	chmod 755 /etc/bcchain
	chmod 775 /home/bcchain
	chmod 775 /home/bcchain/log
	chmod 644 /usr/lib/systemd/system/bcchain.service
	chmod 755 /usr/local/bcchain /usr/local/bcchain/bin /usr/local/bcchain/bin/*
	
	diff version_sdk /home/bcchain/.build/sdk/version >/dev/null 2>/dev/null
	if [[ "$?" != "0" ]]; then
  	oldVer=$(cat /home/bcchain/.build/sdk/version 2>/dev/null | tr -d "\r")
	  newVer=$(cat version_sdk|tr -d "\r")
    if [[ -z ${oldVer} ]]; then
      echo install sdk with version ${newVer}
    else
      echo update sdk from version ${oldVer} to ${newVer}
    fi
    rm -fr /home/bcchain/.build/sdk 2>/dev/null
	  tar xvf sdk_${newVer}.tar.gz -C /home/bcchain/.build >/dev/null
	fi
	
	diff version_thirdparty /home/bcchain/.build/thirdparty/version >/dev/null 2>/dev/null
	if [[ "$?" != "0" ]]; then
  	oldVer=$(cat /home/bcchain/.build/thirdparty/version 2>/dev/null | tr -d "\r")
 	  newVer=$(cat version_thirdparty|tr -d "\r")
    if [[ -z ${oldVer} ]]; then
      echo install thirdparty packages with version ${newVer}
    else
      echo update thirdparty packages from version ${oldVer} to ${newVer}
    fi
    rm -fr /home/bcchain/.build/thirdparty 2>/dev/null
	  tar xvf thirdparty_${newVer}.tar.gz -C /home/bcchain/.build >/dev/null
	fi
	
	chgrp -R bcchain /etc/bcchain
	chown -R bcchain:bcchain /home/bcchain
	
	echo "End of copy files."
	
	systemctl daemon-reload
}
