### CONTENT MAP ###

1. DISCLAIMER
2. DESCRIPTION
3. RUN STEPES
4. BUILD STEPS
5. DEVELOPMENT LOGS

### DISCLAIMER ###

You bear the full responsibility of your actions.
This is an open-source project. This is not a commercial project.
This project is not related to my current nor past employers.
This is my hobby and I'm developing this tool for the sake of learning and understanding, meaning educational purpose. You shall not hold viable its creator/s nor any contributor to the project for any damage you may have done. If you contribute to the project bare in mind that your code or data may be changed in the future without a notice.

### DESCRIPTION ###

The software is not production ready. This are ealry stages of development, feel free to contribute, but prior to that, please, read the disclaimer.

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

### RUN STEPES ###

xena-atila
> cd xena-atila && yarn && node ace migration:run && yarn dev

xena-apep
> cd xena-apep && go run main.go http://127.0.0.1:60666

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

#1 - Baby steps.
https://zarkones-xena.medium.com/xena-r-a-t-devlog-1-7010468588b9