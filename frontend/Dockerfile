FROM node:17-alpine

WORKDIR /app

#COPY package.json ./


COPY . /app

WORKDIR /app/frontend
RUN npm install
#RUN npm build


#CMD ["tail", "-f", "/dev/null"]

CMD [ "npm", "start", "--host", "0.0.0.0"]