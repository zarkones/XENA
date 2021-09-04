The bot client written in Haxe language.
Haxe transpiles to other languages, that allows us to target many platforms.

(Make sure to change Env.hx to accept build profile ID of yours.)

It's supposed to drop software delivered by Xena-Pyramid (bot-builder).
Also it uses an encoder, so the Xena-Pyramid's address is not visible when "strings" are runed over a binary file. Basic reverse engineering will reveal the address and other information encoded with app/services/Perm.hx

At the moment three targets are tested and confirmed working (at least for me on Debian):

1. C++
2. C#
3. Python3

Run "sh build.sh" in order to generate new binaries and scripts.

Run "sh run.sh" in order to do the same in an interpreted. Bare in mind that you need [Haxe](https://haxe.org/) installed.


Quick disclaimer. Some people will categorize this as a dropper, but I don't want to separate droppers from other clients. Since they all should bootstrap each other.
