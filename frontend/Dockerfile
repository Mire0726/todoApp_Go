# Node.jsの最新安定版イメージを使用
FROM node:20.10.0

# アプリケーションのワーキングディレクトリを設定
WORKDIR /app

# アプリケーションの依存関係ファイルをコピー
COPY package.json package-lock.json ./

# 依存関係をインストール
RUN npm install

# ソースコードをコピー
COPY . .

# アプリケーションがリッスンするポートを指定
EXPOSE 3000

# アプリケーションの起動コマンドを指定
CMD ["npm", "run", "dev"]

