# Build AdonisJS
FROM node:16-alpine as builder
# Set directory for all files
WORKDIR /home/node
# Copy over package.json files
COPY package*.json ./
# Install all packages
RUN npm install
# Copy over source code
COPY . .
# Build AdonisJS for production
RUN npm run build --production


# Build final runtime container
FROM node:16-alpine
# Set environment variables
ENV NODE_ENV=production
# Disable .env file loading
ENV ENV_SILENT=true
# Listen to external network connections
# Otherwise it would only listen in-container ones
ENV HOST=0.0.0.0
# Set port to listen
ENV PORT=60666
# Set app key at start time
ENV APP_KEY=
# Set home dir
WORKDIR /home/node
# Copy over built files
COPY --from=builder /home/node/build .
# Install only required packages
RUN npm ci --production
# Expose port to outside world
EXPOSE 60666
# Start server up
CMD [ "node", "server.js" ]