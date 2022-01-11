![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/user-interfaces/xena-face/static/xena-logo.png)

### CONTENT MAP ###

1. [DISCLAIMER](#disclaimer)
2. [SUPPORT MY RESEARCH](#support-my-research)
3. [RUN STEPS](#run-steps)
4. [GALLERY](#gallery)
5. [DESCRIPTION](#description)
6. [REQUIREMENTS](#requirements)
7. [POSTMAN COLLECTIONS](#postman-collections)
8. [BUILD STEPS](#build-steps)
9. [DEVELOPMENT LOGS](#development-logs)
10. [WHOAMI?](#whoami)

### DISCLAIMER ###

You bear the full responsibility of your actions.
This is an open-source project. This is not a commercial project.
This project is not related to my current nor past employers.
This is my hobby and I'm developing this tool for the sake of learning and understanding, meaning educational purpose. You shall not hold viable its creator/s nor any contributor to the project for any damage you may have done. If you contribute to the project bare in mind that your code or data may be changed in the future without a notice.
This software is not allowed to be used for political purpose.
This software is not allowed to be used in a commercial purpose of any kind, shape or form without a written permission.
Selling and redistributing this software is not permited without a written permission.
This software is not allowed to be used for training of algorithms without a written permission.

### SUPPORT MY RESEARCH ###

Working full-time and developing a hobby project is stressful. Gin-Tonic can ease the pain, consider supporting my blackouts.

**Monero:** 87RZRSuASU4cdVXWSXvxLUQ84MpnbpxNHfkucSP9gjNE4TzCUSWT5H7fYunk7JLGPpE9QwHXjZQtaNpeyuKsB8WWLGz13ZJ

**Etherium:** 0x787Ba8EF8d75489160C6687296839F947DC62736

**Or you can Star this project. It is free to do so.**

### RUN STEPS ###

xena-face
> cd user-interfaces/xena-service-face && yarn && yarn dev

xena-atila
> cd services/xena-service-atila && yarn && node ace migration:run && yarn dev

xena-pyramid
> cd services/xena-service-pyramid && yarn && node ace migration:run && yarn dev

xena-ra
> cd services/xena-service-ra && yarn && yarn dev

xena-sensi
> cd services/xena-service-sensi && yarn && yarn dev

xena-apep
> cd bot-clients/xena-bot-apep && go run *.go

xena-anaconda
> cd bot-clients/xena-bot-anaconda && python3 main.py

xena-varvara
> cd bot-clients/xena-bot-varvara && sh run.sh

### GALLERY ###

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/login-page.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/disabled-functionality.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/interaction-with-bot-clients.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/recon.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/author-studio-aka-bot-builder.png)

![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/promotional-materials/settings-page.png)

### DESCRIPTION ###

XENA is the managed remote administration platform for botnet creation & development. Aiming to provide an ecosystem which serves the bot clients. Each service exposes an API in a JSON format delivered by HTTP protocol. Goal is to have a hybrid between centralized and decentralized network, depending on your preference, since it should be custimazible.

The software is not production ready. This are ealry stages of development, feel free to contribute, but prior to that, please, read the disclaimer.
For any questions or help, reach out to me at zarkones.xena@gmail.com

Bot clients are very diverse, powered by Golang, TypeScript, Haxe and Python3.
With this in mind, I'm confident that the framework covers a large surface area.

**SERVCES & BOT CLIENTS**

**Xena-Face**

Web user interface powered by Nuxt.ts and TypeScript. The reason I’ve chosen the web is pure convenience. If a user has to download binaries, that would take some time, I needed the tool to be accessible instantaneously.
Plus that way no traces are left onto the machine while performing a penetration campaign, AND it is available through the Tor browser, what more could we ask for?

Features:
+ Browse the bot clients.
+ Issue messages to the bot clients.
+ Display graphs containing data about the platform's state.

**Xena-Atila**

Message broker powered Adonis.ts and TypeScript. Before you ask, I know, maybe a preexisting solution for message distribution would be a better choice, but I argue that this use case is very domain specific.
I’ve went with Postgres as the storage solution, but you can easily install different SQL drivers in the Adonis framework and go with your favorite database engine. SQL driver installation process is a piece of cake.
Adonis provides support for PostgreSQL, MySQL, MSSQL, MariaDB and SQLite out of the box.

Features:
+ Share information about the clients.
+ Share messages between clients.

**Xena-Apep**

The bot client written in Golang. Why Golang you might be asking, well, cross-platform + the convenience of development. Keep in mind that I do not use Go in the professional capacity, so the code can and will be improved a lot.
This can be used to drop other clients onto the environment.

Feature:
+ Cross platform. Native static & dynamic binaries.
+ Executes shell command.

**Xena-Ra**

The bot client written in Adonis.ts and TypeScript.
Has ability to detect machine's environment and hardware.
Able to detects if it's running as Windows Subsystem for Linux, as root user, and recognize Docker containers.

Features:
+ Corss platform. Requires NodeJS.
+ Provides a deep insight into the environment's details.
+ Capable of parsing web traffic in order to find emails, numbers, keywords, etc...

**Xena-Pyramid**

Bot builder. We need a manged way of bot building and distribution, plus this service will support binary encoding out of the box in the future. Thus making hash based detection useless.
I intent this to also be used for building of other kind of software.

Features:
+ Corss platform. Requires NodeJS.
+ Outsources building of bot client binaries.

**Xena-Anaconda**

Post exploitation bot client written in Python3 (typed, but can be improved for sure). Anaconda! Psszzt!@#! Modular, extandable base code for writting a bot client. It has services available to you, as well as a light-weight core which handles multi-processing. That way you write separate python scripts, which are going to be executed each into its own process. 

Features:
+ Corss platform. Requires Python3.
+ Modular.
+ Executes shell command.

This features (below) are implemented, but not exposed, which was done on purpose. They are useful only under penetration testing campaigns (hopefully authorized). But the platform doesn't have an authorization strategy inplemented, meaning that messages are not signed and encrypted. So, a professional pen. testing campaign is not possible at the present.

+ Basic web parsing.
+ Take a screenshot.
+ Reads bash history.
+ Checks the proxy settings.
+ Enumerate running processes.
+ Provides a slight insight into the environment's details.
+ Subdomain bruteforcing.
+ Take a camera shot.

**Xena-Varvara**

Bot client powered by Haxe language. It which drops other bot clients provided by Xena-Pyramid. It makes them persistent, but the mechanism is super simple. It modifies the .bashrc file.
Feel free to open a Pull-Request in order to add other persistency methods.

Haxe transpiles to other languages, of which at the moment are tested C++, C#, Python3.

Features:
+ Corss platform. Transpiles into multiple targets; PHP, Python, C++, C-Sharp. Meaning that native static & dynamic binaries are possible.
+ Downloads Apep from Pyramid and persists it on the operating system.

**Xena-Sensi**

Sensi, polite and brilient assistent. It exists to supervise the network and protect it. Utilizes GPT-2 and GPT-3 (if you have a key).
At the moment GPT-2 is not released, it needs some work and training. You can utilize GPT-3 in the current version.

Features:
+ Corss platform. Requires NodeJS.
+ GPT 3 gateway.

**Xena-Axe** [coming soon]

Oh, boy... Where to begin. Let me first introduce you to Haxe. An open source high-level strictly-typed programming language with a fast optimizing cross-compiler. It transpiles to: JavaScript, HashLink, Eval, JVM, PHP7, C, Lua, C++, Python, Java, Flash, Neko, ActionScript, PHP5.
With that all covered, I don't know which machine cannot run at least one of the bot client of this framework.

### REQUIREMENTS ###

- Node.js >= 14.15.4 for running Xena-Atila, Xena-Sensi, Xena-Ra, Xena-Pyramid.

- Golang for running Xena-Apep. (tested with golang version 1.15.9)

- Python version 3.9.2 is recommended, because of Xena-Anaconda.

- Haxe for Xena-Varvara, recommended version 4.1.5.

And an operating system of your choice. The framework is fully cross-platform.

- Windows can support everything native or through WSL support.
- Linux distributions can easily meet all requirements. This is the recommended environment, but it's not a requirement.
- MacOS is a capable Unix family member, I've never used it, but I assume the platform should run without a major issue.
- FreeBSD's story goes the same way as for MacOS, never used it, but I do not see a reason why it couldn't support the framework.

### POSTMAN COLLECTIONS ###

I’ve exported my collections from postman which I use for development purpose. That should allow you to better understand how to create your own bot for XENA platform. You can read Xena-Apep for reference, but Apep isn’t complete yet, I’ll probably focus on it in the following development logs. So, until then, feel free to reach out to me at zarkones.xena@gmail.com.

Collections may be found in a JSON format under ./postman-collections

Environment veriables:

xena-atila-url  ===  http://127.0.0.1:60666

xena-pyrmid-url  ===  http://127.0.0.1:60667

xena-ra-url  ===  http://127.0.0.1:60696

xena-sensi-url  ===  http://127.0.0.1:60699

xena-apep-url === Unknow, but can be predicted, it uses the time in order to pick a different port every day.

### BUILD STEPS ###

Build support is provided by Xena-Pyramid.
You can interact with it by Xena-Face. (web user interface)

Bare in mind that this feature is still work in progress.

...we can learn some things from the cloud. :)

### DEVELOPMENT LOGS ###

#2 - Back to the future.
https://zarkones.medium.com/xena-devlog-2-back-to-the-future-866fe6f23ad6

#1 - Baby steps.
https://zarkones.medium.com/xena-r-a-t-devlog-1-7010468588b9

#0 - Botnets, design & implementation.
https://www.youtube.com/watch?v=24hGJjgRfUI

### WHOAMI? ###

Zarkones. Software developer with an accent on back-end services.
Current solutions are overpriced, when it comes to red-team software.
I need to understand the topic and share my knowladge in a appropriate manner.

Social links:
1. [YouTube](https://www.youtube.com/channel/UCn-7I-L-ZpiELb8-6z7z_Ug)
2. [Medium](https://medium.com/@zarkones)
3. [GitHub](https://github.com/zarkones)
4. [Reddit](https://www.reddit.com/r/xenarat)
5. [Twitter](https://twitter.com/zarkones)
