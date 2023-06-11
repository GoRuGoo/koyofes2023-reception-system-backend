# koyofes2023-reception-system-backend

## Overview

2023 年度こうよう祭の一般来場者受付システムのバックエンド API

## Requirement

### OS

- Mac OS Ventura 13.0(動作確認済み)

### Library

- Go
  - Gin
- Docker
- docker-compose

## Installation(local)

1. Clone this repository

```
git clone https://github.com/GoRuGoo/koyofes2023-reception-system-backend.git
```

2. Change directory

```
cd koyofes2023-reception-system-backend
```

3. Build docker image

```
docker-compose up -d
```

4. Create database

ボリュームマウントでは何故か docker-compose.yml ファイルで DB 構築出来なかった為仕方なく...

```
docker-compose exec -it mysql bash
```

```
mysql -u root -p
```

```
gorupass
```

```
CREATE DATABASE reception;
```

## Usage(local)

1. Build & start container

```
docker-compose up
```

2.

```
docker-compose exec go run main.go
```

## Author

- [Yuta Ito](https://github.com/GoRuGoo)
