![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/xena-face/static/xena-logo.png)

### CONTENT MAP ###

1. DISCLAIMER
2. SUPPORT MY RESEARCH
3. DESCRIPTION
4. REQUIREMENTS
4. POSTMAN COLLECTIONS
5. RUN STEPES
6. BUILD STEPS
7. DEVELOPMENT LOGS
8. WHOAMI?

### DISCLAIMER ###

You bear the full responsibility of your actions.
This is an open-source project. This is not a commercial project.
This project is not related to my current nor past employers.
This is my hobby and I'm developing this tool for the sake of learning and understanding, meaning educational purpose. You shall not hold viable its creator/s nor any contributor to the project for any damage you may have done. If you contribute to the project bare in mind that your code or data may be changed in the future without a notice.
This software is not allowed to be used for political purpose.
This software is not allowed to be used in a commercial purpose of any kind, shape or form without a written permission.
Selling and redistributing this software is not permited without a written permission.

### SUPPORT MY RESEARCH ###

Working full-time and developing a hobby project is stressful. Gin-Tonic can ease the pain, consider supporting my blackouts.

Monero: **87RZRSuASU4cdVXWSXvxLUQ84MpnbpxNHfkucSP9gjNE4TzCUSWT5H7fYunk7JLGPpE9QwHXjZQtaNpeyuKsB8WWLGz13ZJ**

Or you can Star this project. It is free to do so.

### DESCRIPTION ###

XENA is the managed remote administration platform. Aiming to create an ecosystem which provides logistical support to the bot clients. Each service exposes an API in a JSON format delivered by HTTP protocol. Goal is to have a hybrid between centralized and decentralized network, depending on your preference, since it should be custimazible.

The software is not production ready. This are ealry stages of development, feel free to contribute, but prior to that, please, read the disclaimer.
For any questions or help, reach out to me at zarkones.xena@gmail.com

> **Xena-Face**

Web user interface powered by Nuxt.ts and TypeScript. The reason I’ve chosen the web is pure convenience. If a user has to download binaries, that would take some time, I needed the tool to be accessible instantaneously.
Plus that way no traces are left onto the machine while performing a penetration campaign, AND it is available through the Tor browser, what more could we ask for?

> **Xena-Atila**

Message broker powered Adonis.ts and TypeScript. Before you ask, I know, maybe a preexisting solution for message distribution would be a better choice, but I argue that this use case is very domain specific.
I’ve went with Postgres as the storage solution, but you can easily install different SQL drivers in the Adonis framework and go with your favorite database engine. SQL driver installation process is a piece of cake.
Adonis provides support for PostgreSQL, MySQL, MSSQL, MariaDB and SQLite out of the box.

> **Xena-Apep**

The bot client written in Golang. Why Golang you might be asking, well, cross-platform + the convenience of development. Keep in mind that I do not use Go in the professional capacity, so the code can and will be improved a lot.
This can be used to drop other clients onto the environment.

> **Xena-Ra**

The bot client written in Adonis.ts and TypeScript.
Has ability to detect machine's environment and hardware.
Able to detects if it's running as Windows Subsystem for Linux, as root user, and recognize Docker containers.

> **Xena-Pyramid**

Bot builder. We need a manged way of bot building and distribution, plus this service will support binary encoding out of the box in the future. Thus making hash based detection useless.
I intent this to also be used for building of other kind of software.

> **Xena-Anaconda** [coming soon]

Post exploitation bot client written in Python3 (typed, but can be improved for sure). Code is there, I'm just not realising it now. Which can also be said for the following service bellow. Now, back to the topic. Anaconda! Psszzt!@#! Modular, extandable base code for writting a bot client. It has services available to you, as well as a light-weight core which handles multi-processing. That way you write separate python scripts, which are going to be executed each into its own process. 

> **Xena-Axe** [coming soon]

Oh, boy... Where to begin. Let me first introduce you to Haxe. An open source high-level strictly-typed programming language with a fast optimizing cross-compiler. It transpiles to: JavaScript, HashLink, Eval, JVM, PHP7, C, Lua, C++, Python, Java, Flash, Neko, ActionScript, PHP5.
With that all covered, I'm not sure which machine cannot run at least one bot clinet of this framework.

> **Xena-Sensi** [coming soon]

Sensi, polite and brilient assistent. It there to supervise the network and protect it. Utilizes GPT-2 and GPT-3, if you have a key.

Now when you've got a bit more familiar with the ecosystem, let's dive into a hypotetical visualization of what needs to be implemented. 

![Diagram of the network](https://miro.medium.com/max/875/1*LRCSF5nna9FhVm77Oc1Q7Q.jpeg)

### REQUIREMENT ###

Node.js >= 14.15.4 for running Xena-Atila, Xena-Sensi, Xena-Ra.
Golang for running Xena-Apep. (tested with golang version 1.15.9)

And a Linux operating system.

### POSTMAN COLLECTIONS ###

I’ve exported my collections from postman which I use for development purpose. That should allow you to better understand how to create your own bot for XENA platform. You can read Xena-Apep for reference, but Apep isn’t complete yet, I’ll probably focus on it in the following development logs. So, until then, feel free to reach out to me at zarkones.xena@gmail.com.

Collections may be found in a JSON format under ./postman-collections

Environment veriables:

xena-atila-url  ===  http://127.0.0.1:60666

xena-pyrmid-url  ===  http://127.0.0.1:60667

xena-ra-url  ===  http://127.0.0.1:60696

xena-sensi-url  ===  http://127.0.0.1:60699

### RUN STEPES ###

xena-atila
> cd xena-atila && yarn && node ace migration:run && yarn dev

xena-apep
> cd xena-apep && go run *.go http://127.0.0.1:60666

xena-pyramid
> cd xena-pyramid && yarn && node ace migration:run && yarn dev

xena-ra
> cd xena-ra && yarn && yarn dev

xena-face
> cd xena-face && yarn && yarn dev

### BUILD STEPS ###

Give me a bit more time, I'm implementing build support via Xena-Pyramid.
A user should not be bothered to do this manually.
It needs to be available through the web UI. (Xena-Face)

...we can learn some things from the cloud. :)

### DEVELOPMENT LOGS ###

#2 - Back to the future.
https://zarkones.medium.com/xena-devlog-2-back-to-the-future-866fe6f23ad6

#1 - Baby steps.
https://zarkones.medium.com/xena-r-a-t-devlog-1-7010468588b9

### WHOAMI? ###

Zarkones. Full-stack software developer with an accent on back-end services.
Current solutions are overpriced, when it comes to red-team software.
I need to understand the topic and share my knowladge in a appropriate manner.
