FROM node:10.16-alpine

COPY . .

ENTRYPOINT ["node","index.js"]
