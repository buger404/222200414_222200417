FROM docker.m.daocloud.io/library/node:20-alpine AS build

# 复制 package.json 和 package-lock.json 到容器中
COPY package*.json ./

# 安装项目依赖
RUN npm install

# 复制所有项目文件到容器中
COPY . .

# 构建项目
RUN npm run build

# 使用 Nginx 作为 Web 服务器
FROM nginx:alpine

# 将 Vue 项目的构建文件复制到 Nginx 的默认静态文件目录
COPY --from=0 /dist /usr/share/nginx/html

# 暴露端口
EXPOSE 80

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]
