# 明日方舟斗蛐蛐活动

这是一个模拟明日方舟斗蛐蛐活动的Web应用程序。

## 功能特性

- 活动信息展示
- 用户注册和登录
- 个人中心管理
- 活动报名功能

## 技术栈

- 后端：Go
- 前端：HTML, CSS, JavaScript, Bootstrap
- 认证：QQ登录集成

## 本地开发

1. 克隆仓库
   ```bash
   git clone https://github.com/sunist-c/arknights-cricket-fighting.git
   cd arknights-cricket-fighting
   ```

2. 安装依赖
   ```bash
   go mod download
   ```

3. 运行应用
   ```bash
   go run main.go
   ```

4. 访问应用
   打开浏览器访问 `http://localhost:18080`

## Docker部署

### 使用预构建镜像

```bash
docker pull ghcr.io/sunist-c/arknights-cricket-fighting:latest
docker run -d -p 18080:18080 ghcr.io/sunist-c/arknights-cricket-fighting:latest
```

### 本地构建镜像

```bash
docker build -t arknights-cricket-fighting .
docker run -d -p 18080:18080 arknights-cricket-fighting
```

## 环境变量

应用支持以下环境变量配置：

- `PORT`: 应用监听端口，默认为18080

## 部署到生产环境

项目使用GitHub Actions自动构建Docker镜像：

- `main`分支的提交会构建`nightly`标签镜像
- 发布标签（如`v1.0.0`）会构建对应版本标签和`latest`标签镜像

## 贡献指南

1. Fork本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

[MIT](LICENSE) 