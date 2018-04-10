Pastegram
===================

Send stdout to telegram chat using telegram api. You'll need to create a bot and grab their Api token.

Requirements
---------

    Telegram Bot
    Chatid of your bot

Installation
---------

    $ go get github.com/pathcl/pastegram
    $ go install github.com/pathcl/pastegram

Usage
---------

    $ export APITELEGRAM="yourtoken"
    $ export CHATID="yourchatid"
    $ uname -a | pastegram 
    Sending to telegram:  Darwin saturno 15.6.0 Darwin Kernel Version 15.6.0: Mon Nov 13 21:58:35 PST 2017; root:xnu-3248.72.11~1/RELEASE_X86_64 x86_64
