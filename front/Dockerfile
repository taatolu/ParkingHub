FROM node:20

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

COPY . .

EXPOSE 3000

# Dockerコンテナ内でファイル監視のポーリングを有効化（ホットリロードに便利）
ENV CHOKIDAR_USEPOLLING=true

CMD ["yarn", "start"]