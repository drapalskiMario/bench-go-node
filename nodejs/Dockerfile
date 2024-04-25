FROM node:20.12.2
WORKDIR /app
COPY . .
RUN npm install
COPY wait-for-it.sh /usr/wait-for-it/wait-for-it.sh
RUN chmod +x /usr/wait-for-it/wait-for-it.sh

CMD /usr/wait-for-it/wait-for-it.sh mysql:3306 -- npm start
