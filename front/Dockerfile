FROM node:20

WORKDIR /app

COPY . .

EXPOSE 3000

# Dockerコンテナ内でファイル監視のポーリングを有効化（ホットリロードに便利）
ENV CHOKIDAR_USEPOLLING=true

CMD ["sh"]