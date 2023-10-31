FROM node:17-alpine

WORKDIR /app

#COPY package.json ./


COPY . /app

WORKDIR /app/frontend
RUN npm install
#RUN npm build
# EXPOSE 3000

#CMD ["tail", "-f", "/dev/null"]

# CMD [ "npm", "start", "--host", "0.0.0.0"]
# CMD [ "npm", "run", "dev", "--host", "0.0.0.0"]
CMD [ "npm", "run", "dev"]
# CMD [ "npm", "start", "--host", "0.0.0.0", "--port", "3000"]
