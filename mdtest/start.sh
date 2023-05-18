#!/bin/bash

if [ $# -ne 1 ]; then
        echo "please specify 1 command line arguments"
		exit 1
fi


screen -S $1 -X kill

screen -dmS $1 

screen -r $1 -X stuff "./cpumanager $1\n"

screen -ls

echo "Running programs w/ screen name: $1"