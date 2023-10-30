FROM node:17-alpine

WORKDIR /app

# COPY package.json ./

COPY . /app
RUN npm install
RUN npm build

EXPOSE 3000

CMD [ "npm", "start" ]