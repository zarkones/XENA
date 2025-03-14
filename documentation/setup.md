### HOW TO SETUP ###

Execute script "prod-build.sh" inside of the project's root directory.

First time compilation might take around 5-10 minutes, since UI is compiling Fyne framework.

That script would compile agents, C2 server and the user interface into folder named "export".

Run the C2 server binary in directory "export/c2". You'd need to configure a few env. variables.

Example: AUTH_TOKEN=my_api_key_for_ui HOST=127.0.0.1 PORT=8080 GIN_MODE=release ./linux_amd64

Value of AUTH_TOKEN also set in the "Settings" page in the UI, along with the C2's host.