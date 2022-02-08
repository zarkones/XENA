# Build the project.
FROM node:16-alpine as builder
# Bodyparser needs git.
RUN apk add --no-cache git
# Set working directory.
WORKDIR /home/node
# Copy the versioning files.
COPY package.json ./
# Install project's dependencies.
RUN yarn install
# Copy the source-code.
COPY . .
# Create a production build.
RUN yarn build

# This step will install the node packages.
FROM node:16-alpine as installer
# Bodyparser needs git.
RUN apk add --no-cache git
# Set working directory.
WORKDIR /home/node
# Copy the versioning files.
COPY package.json ./
# Install production only node packages.
RUN yarn install --prod=true --frozen-lockfile

# Build the run-time container.
FROM node:16-alpine
# Set basic environment variables.
ENV NODE_ENV production
ENV DB_CONNECTION pg
ENV PG_PORT 5432
ENV PG_USER postgres
ENV PG_DB_NAME xena-domena
ENV DRIVE_DISK local
ENV ENV_SILENT true
ENV HOST 0.0.0.0
ENV PORT 60798
# Lower the privledges. (don't use root user)
USER node
# Make the app's directory.
RUN mkdir -p /home/node/app/
# Set working directory.
WORKDIR /home/node/app
# Copy over the app.
COPY --from=builder /home/node/build ./build
COPY --from=installer /home/node/node_modules ./node_modules
# Copy the versioning files.
COPY package.json ./
# Expose the app.
EXPOSE 60798
# Start the app.
CMD [ "node", "./build/server.js" ]