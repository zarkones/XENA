![Logo of the XENA project.](https://raw.githubusercontent.com/zarkones/XENA/production/xena-face/static/xena-logo.png)

### CONTENT MAP ###

1. DISCLAIMER
2. SUPPORT MY RESEARCH
3. DESCRIPTION
4. POSTMAN COLLECTIONS
5. RUN STEPES
6. BUILD STEPS
7. DEVELOPMENT LOGS

### DISCLAIMER ###

You bear the full responsibility of your actions.
This is an open-source project. This is not a commercial project.
This project is not related to my current nor past employers.
This is my hobby and I'm developing this tool for the sake of learning and understanding, meaning educational purpose. You shall not hold viable its creator/s nor any contributor to the project for any damage you may have done. If you contribute to the project bare in mind that your code or data may be changed in the future without a notice.

### SUPPORT MY RESEARCH ###

Working full-time and developing a hobby project is stressful. Gin-Tonic can ease the pain, consider supporting my blackouts.

Monero: **87RZRSuASU4cdVXWSXvxLUQ84MpnbpxNHfkucSP9gjNE4TzCUSWT5H7fYunk7JLGPpE9QwHXjZQtaNpeyuKsB8WWLGz13ZJ**

### DESCRIPTION ###

XENA is the managed remote administration platform. Aiming to create an ecosystem which provides logistical support to the bot clients. Each service exposes an API in a JSON format delivered by HTTP protocol and other fallback channels. Goal is to have everything be a hybrid between centralized and decentralized network, depending on your preference, since it should be custimazible.

The software is not production ready. This are ealry stages of development, feel free to contribute, but prior to that, please, read the disclaimer.
For any questions or help, reach out to me at zarkones.xena@gmail.com

> Xena-Face

Web user interface powered by Nuxt.ts and TypeScript. The reason I’ve chosen the web is pure convenience. If a user has to download binaries, that would take some time, I needed the tool to be accessible instantaneously.
Plus that way no traces are left onto the machine while performing a penetration campaign, AND it is available through the Tor browser, what more could we ask for?

> Xena-Atila

Message broker powered Adonis.ts and TypeScript. Before you ask, I know, maybe a preexisting solution for message distribution would be a better choice, but I argue that this use case is very domain specific.
I’ve went with Postgres as the storage solution, but you can easily install different SQL drivers in the Adonis framework and go with your favorite database engine. SQL driver installation process is a piece of cake.
Adonis provides support for PostgreSQL, MySQL, MSSQL, MariaDB and SQLite out of the box.

> Xena-Apep 

The bot client written in Golang. Why Golang you might be asking, well, cross-platform + the convenience of development. Keep in mind that I do not use Go in the professional capacity, so the code can and will be improved a lot.
This can be used to drop other clients onto the environment.

> Xena-Ra

The bot client written in Adonis.ts and TypeScript.
Has ability to detect machine's environment and hardware.
Able to detects if it's running as Windows Subsystem for Linux, as root user, and recognize Docker containers.

> Xena-Pyramid

Bot builder. We need a manged way of bot building and distribution, plus this service will support binary encoding out of the box in the future. Thus making hash based detection useless.
I intent this to also be used for building of other kind of software.


Now when you've got a bit more familiar with the ecosystem, let's dive into a hypotetical visualization of what needs to be implemented. 

![Diagram of the network](https://miro.medium.com/max/875/1*LRCSF5nna9FhVm77Oc1Q7Q.jpeg)

### POSTMAN COLLECTIONS ###

I’ve exported my collections from postman which I use for development purpose. That should allow you to better understand how to create your own bot for XENA platform. You can read Xena-Apep for reference, but Apep isn’t complete yet, I’ll probably focus on it in the following development logs. So, until then, feel free to reach out to me at zarkones.xena@gmail.com.

Collections may be found in a JSON format under ./postman-collections

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
