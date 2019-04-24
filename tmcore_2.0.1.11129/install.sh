#!/bin/bash

sudo systemctl stop    tmcore 2>/dev/null
sudo systemctl disable tmcore 2>/dev/null
cd pieces
bash install.sh
if [ $? -eq 1 ]; then
    exit
fi

su - tmcore -s /bin/bash -c "/usr/local/tmcore/bin/run.sh init"
if [ $? -eq 0 ]; then
	sudo systemctl enable tmcore 2>/dev/null
	sudo systemctl start  tmcore 2>/dev/null
fi
