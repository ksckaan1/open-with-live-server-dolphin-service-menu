#! /usr/bin/bash
echo "Starting uninstallation..."
sudo rm -rf /usr/bin/go-live-server 
sudo rm -rf $HOME/.local/share/kservices5/ServiceMenus/openWithLiveServer.desktop
echo "Uninstalled Successfully!"