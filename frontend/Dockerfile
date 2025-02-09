FROM node:alpine AS node_builder

COPY package.json .
COPY package-lock.json .

RUN npm ci

COPY . .

RUN npm install
RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]