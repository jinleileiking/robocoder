# 简介

利用Agently, 通过调用 chatgpt，根据需求生成代码

技术栈

- Python
- Go
  - gin
  - gorm
- Vue

# 步骤

# 安装依赖

- pip install -U Agently
- pip install -U openai httpx typer loguru

# 生成代码

- `python3 ./backend.py YOUR_PROJECT_NAME`

程序会读取prd.txt作为需求，在YOUR_PROJECT_NAME目录生成

> prd.md: 需求文档

> \*.go: 程序代码

# 微调

- db.go修改数据库连接信息
- main.go修改listen port
- 进入YOUR_PROJECT_NAME改改代码，因为LLM输出会有些问题

# 运行程序

进入 YOUR_PROJECT_NAME 目录，执行

> go mod tidy

> go build

> ./YOUR_PROJECT_NAME
