#!/bin/bash
sudo apt-get update && sudo apt-get upgrade
wget https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
tar -xzvf go1.10.2.linux-amd64.tar.gz
rm -r go1.10.2.linux-amd64.tar.gz
sudo mv go /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin' >> .bash_profile
source ~/.bash_profile
