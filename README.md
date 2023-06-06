# 智能交互

## 环境搭建

### go相关配置
go下载：https://cn.vuejs.org/<br>
go依赖的安装：go mod tidy<br>

### vue相关配置
vue下载：https://cn.vuejs.org/<br>
vue依赖的安装：yarn 或者 npm install

## 运行项目

### 后端
在config.ini和utils目录中setting.go中有项目相关配置，例如数据库，以及相应的端口号等，可以按照自己的意愿修改。我的电脑是mac，在一些文件中，使用的是路径符为/，而win中为\，会出现路径错误，需要修改路径。.idea是jetbrains公司下相关ide的配置文件，所以建议在goland中运行本项目。
### 前端
web中是前端部分，将web部分在另外一个软件中打开（如vscode），使用npm run serve，或者在package.json中的命令自行选择如何启动项目，启动后可以在自己的网页输入url来进行访问。
