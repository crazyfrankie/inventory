# Inventory Management System
## 简介
Inventory Management System 是一个使用 Go 语言编写的简单清单管理系统。它提供了基本的清单管理功能，包括添加、更新、删除和查询库存项。

## 特性
- 添加新库存项
- 更新现有库存项
- 删除库存项
- 查询库存项

## 安装和运行
### 先决条件
- Go 1.18+
- MySQL 数据库

### 克隆项目

    git clone https://github.com/crazyfrankie/inventory.git
    cd inventory
    
### 配置数据库

创建一个名为 bubble 的 MySQL数据库，并将相关的用户权限分配给它。可以通过以下命令创建数据：
    
```sql
CREATE DATABASE bubble;
CREATE USER 'debian-sys-main'@'localhost' IDENTIFIED BY 'SFpNZbKpkNOE94AA';
GRANT ALL PRIVILEGES ON bubble.* TO 'debian-sys-main'@'localhost';
FLUSH PRIVILEGES;
```

### 配置文件
在 config 文件夹下创建一个 config.ini 文件，并按以下内容进行配置：

```ini
port = 8080
release = false
    
[database]
user = debian-sys-main
password = SFpNZbKpkNOE94AA
host = localhost
port = 3306
name = bubble
```
### 安装依赖
确保在项目根目录下运行以下命令来安装所需的 Go 依赖包：

```go
go mod tidy
```

### 运行项目

编译并运行项目：

```sh
go build -o inventory
./inventory
```

## 使用方法
启动服务器后，可以通过以下端点进行操作：

- 添加新库存项：POST /api/v1/item
- 更新库存项：PUT /api/v1/item/:id
- 删除库存项：DELETE /api/v1/item/:id
- 查询库存项：GET /api/v1/items

具体的 API 使用方式可以查看项目中的相关代码和注释。

## 贡献指南

欢迎贡献代码！请确保您的贡献符合以下准则：

1. Fork 此项目
2. 创建您的特性分支 (git checkout -b feature/AmazingFeature)
3. 提交您的更改 (git commit -m 'Add some AmazingFeature')
4. 推送到分支 (git push origin feature/AmazingFeature)
5. 打开一个 Pull Request

## 许可证
该项目使用 MIT 许可证。
