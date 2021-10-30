# 字节后端训练营-红包雨项目

### 功能实现

* [ ] 流水线部署（火山引擎，github，docker）
* [ ] 反爬机制 
* [ ] 压力测试（webbench）
* [ ] 缓存（redis） 

### 环境配置

- 安装vue
    - sudo npm install -g vue
    - sudo npm install -g @vue/cli
    - sudo npm install @vue/cli-service -g
- 前端安装
    - cd frontend && sudo npm install
    - 降低eslint版本：sudo npm install --save-dev eslint@5
- 安装docker
    - sudo apt -y install docker.io
    - systemctl start docker
- 从docker构建
    - cd backend && sudo docker build -t my_app .
- 运行
    - sudo docker run -p 8080:8080 --name my_app --rm my_app

### 开发日志

2021-10-30：火山引擎创建集群（单实例），创建镜像仓库，上传镜像（nginx），部署资源，创建流水线（绑定github仓库），运行流水线推送至nginx。