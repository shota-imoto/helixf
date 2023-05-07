FROM node:19
COPY client/package.json ./client/package.json
WORKDIR client
RUN yarn install
COPY client ./client
CMD ["yarn", "dev"]