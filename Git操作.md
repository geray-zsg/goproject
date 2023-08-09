# Git管理Github

## 1. Git查看状态
> Git提交文件到远程仓库流程
- `git status` （红色表示还未提交到暂存区，需要执行git add 添加修改的文件到暂存区）

- `git add` 文件名 （添加文件到git暂存区）

- `git status` （绿色表示文件为提交到Git本地仓库，需要使用`git commit -m '提交文件的描述'`命令提交到本地仓库）

- `git push master` 提交到本地仓库后便可以使用该命令提交到远程仓库

## VSCode提交文件到Github显示连接Github超时
> 尚不清楚VSCode哪里配置不合适，可以使用上面Git命令的流程排查并执行命令操作