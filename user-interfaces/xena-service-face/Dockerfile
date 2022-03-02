FROM node:16-alpine

# Create destination directory.
RUN mkdir -p /usr/src/xena-service-face
WORKDIR /usr/src/xena-service-face

# Update and install dependency.
RUN apk update && apk upgrade
RUN apk add git

# Copy the app, note .dockerignore
COPY . /usr/src/xena-service-face/
RUN yarn install
RUN yarn build

EXPOSE 3000

ENV NUXT_HOST=0.0.0.0
ENV NUXT_PORT=3000
ENV HOST=0.0.0.0
ENV PORT=3000


CMD [ "yarn", "nuxt", "start" ]