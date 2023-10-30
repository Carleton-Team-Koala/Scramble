FROM node:17-alpine

WORKDIR /app

COPY package.json .
RUN npm install

COPY . .
RUN npm build

EXPOSE 80

CMD [ "npm", "start" ]