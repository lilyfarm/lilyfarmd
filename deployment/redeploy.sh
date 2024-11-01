# We do not re-enter the credentials, we merely swap the 
# lilyfarm executable and restart the service.
sudo systemctl stop lilyfarmd
sudo cp ../binaries/lilyfarm /usr/bin/lilyfarm
sudo chmod +x /usr/bin/lilyfarm
sudo systemctl start lilyfarmd