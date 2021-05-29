#! /usr/bin/bash
echo "Starting installation..."
sudo cp ./go-live-server /usr/bin/
sudo cp ./openWithLiveServer.desktop $HOME/.local/share/kservices5/ServiceMenus/
echo "Installed Successfully"