package main

// Define constants used by the application
const (
	version = "v0.2.0-beta.1"

	// Bot command messages:
	// Help message
	msgHelpShort = "*Usage:*\n" +
		"- /help - show this message\n" +
		"- /about - show information and credits about my creator and any contributors\n" +
		"- /status - ask how I'm going.. and if I'm still running\n" +
		"- /start - start activly monitoring your Skyminer. Once started, notifications will be sent to you for events that occur. A heartbeat will also be initiated to let you know if the bot and the Miner are still running.\n" +
		"- /stop - stop monitoring your Skyminer. Once stopped, I won't send any more notifications\n"

	msgHelp = "*Wing Commander* here. I will help you to manage and monitor your Skyminer and its Nodes.\n\n" +
		msgHelpShort +
		"\n" +
		"\n" +
		"Note: I am bound to this chat. I will only respond to commands from my configured Admin (%s)."

	// About cmd message
	msgAbout = "*Wing Commander (" + version + ")*\n" +
		"A Telegram bot written in *Go* designed to help the *Skyfleet* community monitor and manage their *Skyminers*.\n" +
		"\n" +
		"*Created by:* @BigOokie *2018*\n" +
		"*GitHub:* https://github.com/BigOokie/skywire-wing-commander\n" +
		"*Twitter:* https://twitter.com/BigOokie\n" +
		"\n" +
		"Issues and feature requests must be logged via [GitHub](https://github.com/BigOokie/skywire-wing-commander/issues/new/choose)\n" +
		"\n" +
		"*Skycoin*: https://www.skycoin.net/\n" +
		"\n" +
		"*Donations most welcome* 👍\n" +
		"*Skycoin:* ES5LccJDhBCK275APmW9tmQNEgiYwTFKQF"

	msgConnectedNodes = "*Connected Nodes:* %v"
	// Status cmd message
	msgStatus = "*Wing Commander* Ready and reporting for duty 👍\n" + msgConnectedNodes
	// Heartbeat message
	msgHeartbeat = "*Wing Commander Heatbeat* ❣️\n" + msgConnectedNodes

	// Node Connect/Disconnect Event Messages
	msgNodeConnected    = "*Node Connected:* %s\n\n" + msgConnectedNodes
	msgNodeDisconnected = "‼ *Node Disconnected:* %s\n\n" + msgConnectedNodes

	// Start cmd messages
	msgMonitorAlreadyStarted = "️️*Wing Commander* Monitoring has already been started."
	msgMonitorStart          = "*Wing Commander* Monitoring starting..."

	// Stop cmd message
	msgMonitorStop       = "*Wing Commander* Monitoring stopping..."
	msgMonitorStopped    = "*Wing Commander* Monitoring stopped..."
	msgMonitorNotRunning = "*Wing Commander* Monitoring is not running..."

	// Default cmd message (unhandled)
	msgDefault = "Sorry. I don't know that command."

	// OS Interupt Signals
	msgOSInteruptSig = "*Wing Commander* OS Interupt Signal Received. Exiting."
	msgOSKillSig     = "*Wing Commander* OS Kill Signal Received. Exiting."
)
