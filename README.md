### INTRODUCTION ###
XENA is Software for Cyber-Security Automation, Adversary Simulations, and Red Team Operations.

XENA strives to be fully integrated security penetration testing framework. It is equipped with a post-exploitation agent, C2 server, and a dark-themed elegant user interface.

### VIDEOS ###
[Setup & General Usage](https://youtu.be/l86krmk-YZs)

### HOW TO SETUP ###

Execute script "prod-build.sh" inside of the project's root directory.

First time compilation might take around 5-10 minutes, since UI is compiling Fyne framework.

That script would compile agents, C2 server and the user interface into folder named "export".

Run the C2 server binary in directory "export/c2". You'd need to configure a few env. variables.

Example: AUTH_TOKEN=my_api_key_for_ui HOST=127.0.0.1 PORT=8080 GIN_MODE=release ./linux_amd64


[Setup Video Tutorial](https://youtu.be/l86krmk-YZs)

### SOCIAL ###

[Patreon](https://www.patreon.com/zarkones)
[YouTube](https://www.youtube.com/channel/UCn-7I-L-ZpiELb8-6z7z_Ug)
[X](https://x.com/zarkones)
[GitHub](https://github.com/zarkones)
[Itch.io](https://zarkones.itch.io)