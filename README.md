This project for now just run in ***LINUX UBUNTU***

Please install dependencie to notify in linux desktop 

```sh
sudo apt-get install libnotify-bin
```

The installer will create this service ```/etc/systemd/system/price-alerter.service``` and copy the compiled price-alerter file on dist/price-alerter to /usr/bin/price-alerter.

sudo systemctl daemon-reload
sudo systemctl start price-alerter
sudo systemctl enable myservice // for auto reload 
sudo journalctl -u price-alerter -f