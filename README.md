### INTRODUCTION ###
XENA is Software for Cyber-Security Automation, Adversary Simulations, and Red Team Operations.

XENA strives to be fully integrated security penetration testing framework. It is equipped with a post-exploitation agent, C2 server, and a dark-themed elegant user interface.

### SOCIAL ###
[Patreon](https://www.patreon.com/zarkones) |
[Discord](https://discord.gg/qjJwSh2TF9) |
[X.com](https://x.com/zarkones) |
[YouTube](https://www.youtube.com/channel/UCn-7I-L-ZpiELb8-6z7z_Ug) |
[Itch.io](https://zarkones.itch.io) |
[GitHub](https://github.com/zarkones)

### VIDEOS ###
[Setup & General Usage](https://youtu.be/l86krmk-YZs)

### HOW TO SETUP ###

Execute script "prod-build.sh" inside of the project's root directory.

First time compilation might take around 5-10 minutes, since UI is compiling Fyne framework.

That script would compile agents, C2 server and the user interface into folder named "export".

Run the C2 server binary in directory "export/c2". You'd need to configure a few env. variables.

Example: AUTH_TOKEN=my_api_key_for_ui HOST=127.0.0.1 PORT=8080 GIN_MODE=release ./linux_amd64

Value of AUTH_TOKEN also set in the "Settings" page in the UI, along with the C2's host.

[Setup Video Tutorial](https://youtu.be/l86krmk-YZs)

![Promo Image 1](https://raw.githubusercontent.com/zarkones/XENA/production/assets/promo/promo1.png)

![Promo Image 2](https://raw.githubusercontent.com/zarkones/XENA/production/assets/promo/promo2.png)

![Promo Image 3](https://raw.githubusercontent.com/zarkones/XENA/production/assets/promo/promo3.png)
