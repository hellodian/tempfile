#!/bin/bash

sudo systemctl stop    tmcore 2>/dev/null
sudo systemctl disable tmcore 2>/dev/null
cd pieces
bash update.sh
if [ $? -eq 0 ]; then
	sudo systemctl enable tmcore 2>/dev/null
	sudo systemctl start  tmcore 2>/dev/null
fi
