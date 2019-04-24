#!/bin/bash

sudo systemctl stop    bcchain 2>/dev/null
sudo systemctl disable bcchain 2>/dev/null
cd pieces
bash update.sh
if [ $? -eq 0 ]; then
	sudo systemctl enable bcchain 2>/dev/null
	sudo systemctl start  bcchain 2>/dev/null
fi 