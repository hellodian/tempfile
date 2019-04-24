#!/usr/bin/env bash

COLUMNS=12

isEmptyFileOrDir() {
	dir=$1
	if [ -z "$(ls -A $dir 2>/dev/null)" ]; then
	  echo "true"
  else
    echo "false"
  fi
}

emptyClean=$(isEmptyFileOrDir "/home/tmcore/clean.sh")
emptyTendermint=$(isEmptyFileOrDir "/usr/local/tmcore/bin/tendermint")
emptyNodeKey=$(isEmptyFileOrDir "/etc/tmcore/config/node_key.json")
emptyPrivKey=$(isEmptyFileOrDir "/etc/tmcore/config/priv_validator.json")

isEmptyNode() {
	if [ $emptyClean == "true" ] && [ $emptyTendermint == "true" ] && [ $emptyNodeKey == "true" ] && [ $emptyPrivKey == "true" ]; then
	  echo "true"
	  return
	else
	  echo "false"
	  return
	fi
}

isCompletedNode() {
	if [ $emptyClean == "false" ] && [ $emptyTendermint == "false" ] && [ $emptyNodeKey == "false" ] && [ $emptyPrivKey == "false" ]; then
	  echo "true"
	  return
	else
	  echo "false"
	  return
	fi
}

isCorruptedNode() {
	t=$(isEmptyNode)
	if [ $t = "true" ]; then
	  echo "false"
	  return
	fi
	t=$(isCompletedNode)
	if [ $t = "true" ]; then
	  echo "false"
	  return
	fi
	echo "true"
	return
}

doCopyFiles() {
  chainID_=$1
  echo "Start copying files ..."
	mkdir -p /etc/tmcore/genesis /home/tmcore/{data,log} /usr/local/tmcore/bin
	mkdir -p /etc/systemd/system/tmcore.service.d
	
	getent group tmcore  >/dev/null 2>&1 || groupadd -r tmcore
	getent passwd tmcore  >/dev/null 2>&1 || useradd -r -g tmcore \
	  -d /etc/tmcore -s /sbin/nologin -c "BlockChain.net tendermint core System User" tmcore
	
	cp jq tendermint p2p_ping run.sh rutaller.bash /usr/local/tmcore/bin
	cp start.sh stop.sh clean.sh /home/tmcore
	cp tmcore.service /usr/lib/systemd/system
	cp override.conf /etc/systemd/system/tmcore.service.d
  cp ${chainID_}/tendermint-forks.json* /usr/local/tmcore/bin
  tar cvf - ${chainID_} 2>/dev/null |(cd /etc/tmcore/genesis;tar xvf - >/dev/null )
	
	touch /var/spool/cron/root
	sed -i '/rutaller.bash/d' /var/spool/cron/root
	echo "* * * * * /usr/local/tmcore/bin/rutaller.bash > /dev/null" >> /var/spool/cron/root
	chown -R tmcore:tmcore /etc/tmcore
	chown -R tmcore:tmcore /home/tmcore/data
	chown tmcore:tmcore /home/tmcore/log
	chown tmcore:tmcore /home/tmcore
	
	chmod 600 /var/spool/cron/root
	chmod 755 /home/tmcore/data
	chmod 775 /home/tmcore
	chmod 775 /home/tmcore/log
	chmod 644 /usr/lib/systemd/system/tmcore.service
	chmod 755 /etc/tmcore /usr/local/tmcore /usr/local/tmcore/bin /usr/local/tmcore/bin/*
	
	echo "End of copy files."
	
	systemctl daemon-reload
}
