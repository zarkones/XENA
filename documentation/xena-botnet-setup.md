# Requirements

Node.js >= 14.15.4

Golang >= 1.15.9

Python >= 3.9

Haxe >= 4.1.5

C++ >= 11

Docker

Postgres >= 12 (or MySQL, MSSQL, MariaDB and SQLite, this requires additional configuration, see: https://docs.adonisjs.com/guides/database/introduction)

Not all of the languages are required. Node.js is used for the back-end infrastructure, thus it's pretty much mandatory. Other languages are used by the bot clients and thus use accordingly.

# Setup Postgres

## Description

Postgres is a free and open-soruce database which we'll use to store data of our back-end services.

## Requirements

Docker

## Setup

First make sure that you install docker using the following 'apt' package menager for unix systems.
> sudo apt install docker-ce docker-ce-cli containerd.io

If you don't have 'apt' as your package manager, please refer to Docker's official documentation on how to install it: https://docs.docker.com/engine/install

Then bring up the docker container using the following command.
> docker run -d --name xena-postgres -e POSTGRES_DB=xena-atila -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword postgres

If you're going to run other services as containers as well, create a network for them:
> docker network create xena

But then do not forget to add Postgres container to that network. Command which accounts for that is:
> docker run -d --name xena-postgres --net xena -e POSTGRES_DB=xena-atila -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword postgres

# Setup Xena-Service-Atila

## Description

This service is acting as Command & Control Center. It holds information about the bots, such is their ID and public key. This is the primary mechanism for interacting with bots.

## Requirements

Node.js >= 14.15.4

## Setup

Navigate to the folder where the service is located:
> cd services/xena-service-atila

Then install required dependencies:
> yarn

Or if you aren't using 'yarn' as your package manager, you could do the equivalent with 'npm' manager:
> npm install

This will install all the necessary node modules, aka. packages, it might take a few minutes depending on your internet connection.

## Configuration

Rename '.env.example' into '.env'.

Then run the following command in order to generate a random app key.
> node ace generate:key

Then place that secret as the value of APP_KEY environment variable inside of the '.env' file.

For the value of PG_HOST set 'postgres' or any other name that you've given to the Postgres account. PG_PASSWORD should be populated with your password that you've also set for the Postgres account.

TRUSTED_PUBLIC_KEY environment variable should be set to your public key. If you do not have a key-pair refer to 'key-pair-generation.md'.

## Starting Up

In order to run the service in the development mode run:
> yarn dev

Or if you aren't using 'yarn' as your package manager, you could do the equivalent with 'npm' manager:
> npm run dev

In order to build the service you need to run:
> yarn build

Or if you aren't using 'yarn' as your package manager, you could do the equivalent with 'npm' manager:
> npm run build

Then you will need to copy over the '.env' file, and run it using:
> cp .env build/.env

> node build/server.js

But if you wish to deploy it as a container run:
> docker build -t xena-service-atila .

> docker run -it --net xena --name='xena-atila' -e PG_HOST='ip-address-of-the-postgres-continer' -e CORS_POLICY_ALLOWED_ORIGINS='http://127.0.0.1:3000' -e PG_PASSWORD='your-db-password' -e APP_KEY='your-app-secret' -e TRUSTED_PUBLIC_KEY='yout-public-key' -p 60666:60666 xena-service-atila

# Setup Xena-Service-Face

## Description

This service allows us to interact with the botnet through an elegant, dark-themed web user interface.

## Requirements

Node.js >= 14.15.4

## Setup

Navigate to the folder where the service is located:
> cd user-interfaces/xena-service-face

Then install required dependencies:
> yarn

Or if you aren't using 'yarn' as your package manager, you could do the equivalent with 'npm' manager:
> npm install

This will install all the necessary node modules, aka. packages, it might take a few minutes depending on your internet connection.

## Configuration

Rename '.env.example' into '.env'.

You do not need to change the default urls of back-end services in '.env'. Since you can personalize these in the U.I. in the page Settings under the Connections tab.

## Starting Up

In order to run the service in the development mode run:
> yarn dev

Or if you aren't using 'yarn' as your package manager, you could do the equivalent with 'npm' manager:
> npm run dev

In order to build the service you need to run:
> yarn build

Or if you aren't using 'yarn' as your package manager, you could do the equivalent with 'npm' manager:
> npm run build

Then run it using:
> ./main

But if you wish to deploy it as a container run:
> docker build -t xena-service-face .

> docker run -it -p 3000:3000 --net xena --name='xena-face' xena-service-face

Now you can access the service through your browser at http://127.0.0.1:3000