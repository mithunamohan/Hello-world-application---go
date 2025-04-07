FROM node:lts
WORKDIR /usr/src/app
COPY ["package.json","package-lock.json","./"]
RUN npm install --silent
COPY . .
EXPOSE 5173
CMD ["npm","run","dev","--","--host"]