#!/bin/bash

sudo systemctl stop    bcchain 2>/dev/null
sudo systemctl disable bcchain 2>/dev/null

cd pieces
bash install.sh
if [ $? -eq 1 ]; then
    exit
fi

sudo systemctl enable bcchain 2>/dev/null
sudo systemctl start  bcchain 2>/dev/null
