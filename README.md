Pastegram
===================

Send stdout to telegram chat using telegram api. You'll need to create a bot and grab their Api token.

Requirements
---------

You'll need two things:

    A Telegram Bot Api token
    The Chatid of your bot

[Howto create a telegram bot](https://www.sohamkamani.com/blog/2016/09/21/making-a-telegram-bot/)

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

Please keep in mind your $GOPATH.
