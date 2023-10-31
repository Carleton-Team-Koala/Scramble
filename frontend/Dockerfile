FROM node:17-alpine

WORKDIR /app

#COPY package.json ./


COPY . /app

WORKDIR /app/frontend
RUN npm install

CMD [ "npm", "run", "dev"]
