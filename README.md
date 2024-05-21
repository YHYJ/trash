<h1 align="center">Trash</h1>

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-11-26 11:40:09 -->

---

<p align="center">
  <a href="https://github.com/YHYJ/trash/actions/workflows/release.yml"><img src="https://github.com/YHYJ/trash/actions/workflows/release.yml/badge.svg" alt="Go build and release by GoReleaser"></a>
</p>

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Install](#install)
  * [一键安装](#一键安装)
* [Usage](#usage)
* [Compile](#compile)
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)
    * [macOS](#macos)

<!-- vim-markdown-toc -->

---

<!------------------------------->
<!--  _                 _      -->
<!-- | |_ _ __ __ _ ___| |__   -->
<!-- | __| '__/ _` / __| '_ \  -->
<!-- | |_| | | (_| \__ \ | | | -->
<!--  \__|_|  \__,_|___/_| |_| -->
<!------------------------------->

---

CLI 版文件回收站，防止 rm 命令直接彻底删除文件。支持 Linux 和 macOS 系统

## Install

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/trash/main/install.sh | sudo bash -s
```

## Usage

- `put`子命令

  该子命令将指定文件或目录放入回收站

- `list`子命令

  该子命令会列出回收站中的文件，格式是：文件删除日期 文件删除时间 文件原路径

- `restore`子命令

  该子命令交互式恢复回收站中的文件

- `empty`子命令

  该子命令会清空回收站

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/trash/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/trash/general.BuildTime=`date +%s` -X github.com/yhyj/trash/general.BuildBy=$USER" -o build/trash main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/trash/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/trash/general.BuildTime=`date +%s` -X github.com/yhyj/trash/general.BuildBy=$USER" -o build/trash main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/trash/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/trash/general.BuildTime=`date +%s` -X github.com/yhyj/trash/general.BuildBy=$USER" -o build/trash main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64
