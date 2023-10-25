### IMPORTANT UPDATE ###

Our team has moved its attention to XENA v2. Version 1 of XENA (this open-source project, aka. alpha version) is composed of 5 programming languages. That has proven costly to us. We decided to rewrite it into completely new codebase, utilizing only Golang for front-end, back-end, and systems programming. We're also looking for someone to maintain XENA v1. Reach out.

Download XENA Beta: https://zarkones.itch.io/xena

![Logo of the XENA project.](https://img.itch.zone/aW1nLzExODEwMDI5LnBuZw==/original/fS0sJs.png)

### CONTENT MAP ###

1. [DISCLAIMER](#disclaimer)
2. [SUPPORT MY RESEARCH](#support-my-research)
3. [HOW TO SETUP](#how-to-setup)
4. [DESCRIPTION](#description)
5. [GALLERY](#gallery)
6. [SERVICES AND BOT CLIENTS](#services-and-bot-clients)
7. [CONTRACTS](#contracts)
8. [REQUIREMENTS](#requirements)
9. [POSTMAN COLLECTIONS](#postman-collections)
10. [WHOAMI?](#whoami)

### DISCLAIMER ###

You bear the full responsibility of your actions.
This is an open-source project. This is not a commercial project.
This project is not related to my current nor past employers.
This is my hobby and I'm developing this tool for the sake of learning and understanding, meaning educational purpose. You shall not hold viable its creator/s nor any contributor to the project for any damage you may have done. If you contribute to the project bare in mind that your code or data may be changed in the future without a notice and that it belongs to the project, not you. 
This software is not allowed to be used for political purpose.
This software is not allowed to be used for commercial purpose of any kind, shape or form without a written permission.
Selling and redistributing this software is not permitted without a written permission.
This software is not allowed to be used for training of algorithms without a written permission.

### SUPPORT OUR RESEARCH ###

[BECOME A PATREON](https://www.patreon.com/zarkones)

**Monero:** 87RZRSuASU4cdVXWSXvxLUQ84MpnbpxNHfkucSP9gjNE4TzCUSWT5H7fYunk7JLGPpE9QwHXjZQtaNpeyuKsB8WWLGz13ZJ

**Ethereum:** 0x787Ba8EF8d75489160C6687296839F947DC62736

**Or you can Star this project. It is free to do so.**

### HOW TO SETUP ###

Execute inside of the root folder of the project.

> sh setup.sh

[Setup Video Tutorial](https://youtu.be/i5Ct7qg_qVE)

### DESCRIPTION ###

XENA is the managed remote administration platform for botnet creation & development powered by blockchain and machine learning. Favoring secrecy and resiliency over performance. Aiming to provide an ecosystem which serves the bot herders. It's micro-service oriented allowing for specialization and lower footprint. Ese of deployment is enabled via Docker containers.

You interact with the botnet through an elegant dark-themed web user interface. All communication is signed with bot herder's private key, and verified by the bots using hardcoded public key.

Bot clients are really diverse, powered by: Golang, Haxe, C++ and Python3.
With this in mind, bots are able to run on Windows, MacOS, Linux and FreeBSD.

### GALLERY ###

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/login-page.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/disabled-functionality.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/interaction-with-bot-clients.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/recon.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/author-studio-aka-bot-builder.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/settings-page.png)

### SERVICES AND BOT CLIENTS ###

**Xena-Service-Face**

Web user interface powered by Nuxt.ts and TypeScript. The reason I’ve chosen the web is pure convenience, accessible instantaneously via browser.
Plus that way no traces are left onto the machine while performing a penetration campaign, AND it is available through the Tor browser, AND is accessible across devices, what more could we ask for?

**Xena-Service-Atila**

Command & Control (C2) server powered Adonis.ts and TypeScript. In regards to the storage solution I’ve gone with Postgres, but you can easily install different SQL drivers in the Adonis framework and use with your favorite database engine. SQL driver installation process is a piece of cake.
Adonis provides support for PostgreSQL, MySQL, MSSQL, MariaDB and SQLite out of the box. Its main responsibilities are keeping track of bots, providing analytics and delivering messages to the bots.

**Xena-Bot-Apep**

The bot client written in Golang. Why Golang you might be asking, well, cross-platform + the convenience of development. Keep in mind that I do not use Go in the professional capacity, so the code can and will be improved a lot.
This can be used to drop other clients onto the environment.

Features:
+ Cross platform. Native static & dynamic binaries.
+ Executes shell command.
+ Domain Generation Alogithm. (DGA)
+ Get operating system's details.
+ Persistent on Linux. (requires root)
+ Grabs Chromium's web history. (linux)
+ Grabs Chromium's search history. (linux)
+ Search duckduckgo and return results.
+ Gets the currently active CNC from Gettr's user profile's website.
+ SSH brute-forcer.
+ Hibernate on start up for ~15 minutes to avoid analysis.
+ P2P networking.
+ Control over Discord server.

**Xena-Bot-Monolitic**

The bot client written in C++.

Features:
+ Execute hardcoded payload.
+ Telnet bruteforcer.
+ Obfuscated binary.

**Xena-Service-Ra**

The API written in Adonis.ts and TypeScript. Provides command line hacking tools via HTTP endpoints, making them available for use through the web user interface.

Features:
+ Cross platform. Requires NodeJS.
+ Nmap.
+ Sublist3r.
+ SqlMap.

**Xena-Service-Pyramid**

Bot builder API written in Adonis.ts and TypeScript. Leverages server-side polymorphism, thus making hash based detection useless.

Features:
+ Cross platform. Requires NodeJS.
+ Outsources building of Xena-Bot-Apep.

**Xena-Bot-Anaconda**

Post exploitation bot client written in Python3 (typed, but can be improved for sure). Extandable multi-processing engine allowing for writting a custom modules and usage of already available modules. That way you write separate python scripts, which are going to be executed each into its own process. 

Features:
+ Cross platform. Requires Python3.
+ Modular.
+ Executes shell command.
+ MicroTik (Winbox) exploit.
+ Enumerate running processes.
+ Return Bash history.
+ Check local proxy settings.
+ Get operating system's details.

This features (below) are implemented, but not exposed. You can write modules in order to consume them.

+ Basic web parsing.
+ Take a screenshot.
+ Subdomain bruteforcing.
+ Take a camera shot.

**Xena-Bot-Varvara**

Bot client powered by Haxe language. It which drops other bot clients provided by Xena-Service-Pyramid. It makes them persistent, but the mechanism is super simple. It modifies the .bashrc file. Feel free to open a Pull-Request in order to add other persistency methods.

Haxe transpiles to other languages, of which at the moment are tested C++, C#, Python3.

Features:
+ Cross platform. Transpiles into multiple targets; PHP, Python, C++, C-Sharp. Meaning that native static & dynamic binaries are possible.
+ Downloads Xena-Bot-Apep from Xena-Service-Pyramid and persists it within the operating system.

**Xena-Service-Sensi**

Sensi, polite and brilient assistent. It exists to supervise the network and protect it. Utilizes GPT-2 and GPT-3 (if you have a key).
At the moment GPT-2 is not released, it needs some work and training. You can utilize GPT-3 in the current version.

Features:
+ Cross platform. Requires NodeJS.
+ GPT 3 gateway. Allowing you to chat with AI in the web user interface.

### CONTRACTS ###

**Xena-Contract-Botchain**

This contract allows the bootstrapping of the botnet by acting as a smart contract on Ethereum blockchain. The creator of the contract has the highest authority and is able to assign an administrator of the contract and reassign it. Only the administrator can change the currently promoted configuration, messages and sponsored peers.

The contract's creator cannot change any properties of the contract apart from the currently active administrator. That way we avoid a single point of failure, where the creator connects to the blockchain and discloses his IP address thus risk getting attacked and his wallet key stolen. Also all configuration, messages and sponsored peers MUST be signed in order to lower the attack area. Thus a successful takeover of the contract requires the creator's wallet key and the bot master's private key.

There is no need to publish this contract on a private blockchain.

### REQUIREMENTS ###

- Docker

And an operating system of your choice. The framework is fully cross-platform.

- Windows can support everything native or through WSL support.
- Linux distributions can easily meet all requirements. This is the recommended environment, but it's not a requirement.
- MacOS is a capable Unix family member, I've never used it, but I assume the platform should run without a major issue.
- FreeBSD's story goes the same way as for MacOS, never used it, but I do not see a reason why it couldn't support the framework.

### POSTMAN COLLECTIONS ###

I’ve exported my collections from postman which I use for development purpose. That should allow you to better understand how to create your own bot for XENA platform. You can read Xena-Bot-Apep for reference.

Collections may be found in a JSON format under ./postman-collections

Environment veriables:

xena-atila-url  ===  http://127.0.0.1:60666

xena-pyrmid-url  ===  http://127.0.0.1:60667

xena-ra-url  ===  http://127.0.0.1:60696

xena-sensi-url  ===  http://127.0.0.1:60699

Social links:
1. [YouTube](https://www.youtube.com/channel/UCn-7I-L-ZpiELb8-6z7z_Ug)
2. [Medium](https://medium.com/@zarkones)
3. [GitHub](https://github.com/zarkones)
5. [Twitter](https://twitter.com/zarkones)
