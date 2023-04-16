FROM node:19-bullseye
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm prune --production
RUN npm run build --production
RUN npm install -g serve
EXPOSE 3000

# Run application
#CMD [ "npm", "start" ]
CMD serve -s build
