# blinkdl
Video downloader for Blink home security cameras.
Uses ideas and specs from:

https://github.com/se7enack/SBlink

https://github.com/MattTW/BlinkMonitorProtocol

https://github.com/fronzbot/blinkpy

https://github.com/NGRP/node-red-contrib-viseo


##Install

	go get github.com/serge-v/blinkdl

##Usage

	blinkdl -login your@email.com
	blinkdl -list [-days=DAYS] [-page=N]
	blinkdl -download [-days=DAYS] [-page=N]

Downloads all videos for the last `DAYS` into ~/.local/blink directory.
