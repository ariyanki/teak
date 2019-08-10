### add key
### makesure key file exists
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa

cd /ariyanki/teak

# Restart teak service.
sudo service teak stop
mv teak_new teak
sudo service teak start