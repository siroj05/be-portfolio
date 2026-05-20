#!/bin/bash
# tunnel.sh - Open SSH Tunnel to Remote Portfolio Database

REMOTE_IP="43.157.248.30"
LOCAL_PORT="3307"
REMOTE_PORT="3306"

echo "=========================================================="
echo " Starting SSH Tunnel for Portfolio Database..."
echo " Forwarding local port $LOCAL_PORT to remote $REMOTE_IP:$REMOTE_PORT..."
echo "=========================================================="

# Check if port 3307 is already in use
if lsof -Pi :$LOCAL_PORT -sTCP:LISTEN -t >/dev/null ; then
    echo "Warning: Local port $LOCAL_PORT is already in use."
    echo "Checking if an active ssh tunnel is already running..."
    PID=$(lsof -t -i :$LOCAL_PORT -sTCP:LISTEN)
    if [ ! -z "$PID" ]; then
        echo "Found active process on port $LOCAL_PORT with PID: $PID"
        echo "You can close it by running: kill $PID"
    fi
    exit 1
fi

# Start SSH tunnel in the background
ssh -f -N -L $LOCAL_PORT:127.0.0.1:$REMOTE_PORT root@$REMOTE_IP

if [ $? -eq 0 ]; then
    echo "----------------------------------------------------------"
    echo " SSH Tunnel established successfully in the background!"
    echo " Local port: $LOCAL_PORT"
    echo " To close the tunnel, run: kill \$(lsof -t -i :$LOCAL_PORT)"
    echo "=========================================================="
else
    echo "Failed to establish SSH Tunnel."
    echo "=========================================================="
fi
