go mod tidy 
sudo systemctl start sqlite
sudo mysql -u root < sqlsetup.sql
