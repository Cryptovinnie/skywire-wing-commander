#!/bin/bash
# This Bash sript will collect the variables for $HOME\.wingcommander\config.toml
# Telegram bot. API Keys, @Username & ChatID

echo Hello, Lets get your telegram bot set up!
echo If you have not yet, Initiate chat with the @BotFather.
echo What was your Telegram API Key from the BotFather?

read APIKEYS
echo $APIKEYS

echo please go to https://api.telegram.org/bot$APIKEYS/getUpdates
echo What was the chat id?
read ChatID

echo What is your Telegram username? be sure to include the @ Symbol.
read TELEUSERNAME

echo All set up your details are
echo Api key is: "$APIKEYS"
echo Telegram username is: "$TELEUSERNAME"
echo ChatID is: $ChatID

sed -i -e "s/apikey = .*/apikey = "$APIKEYS"/" config.toml
sed -i -e "s/admin = .*/admin = "$TELEUSERNAME"/" config.toml
sed -i -e "s/chatid = .*/chatid = $ChatID/" config.toml

