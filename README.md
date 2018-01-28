# Text To Img Api

WebAPI convert text to image with customization

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Pre-request

- ImageMagick 7.0 or latest (<https://www.imagemagick.org>)
- Go 1.9.3 (<https://golang.org/>)
- Go dep (Dependency management) <https://github.com/golang/dep>

## Start

1. Clone this project
2. Change directory to project directory
    ```bash
    cd /project/root/dir
    ```
3. Download dependencies
    ```bash
    dep ensure
    ```
4. Finally, start server =D
    ```bash
    go run main.go
    ```

## Config

You can change the configuration by config.yml

If you not familiar with yaml, you can visit official YAML website <http://www.yaml.org/spec/1.2/spec.html> to learn the syntax.

## API

### Convert text to image

`GET /-/{text}.png`

`GET /api/text/{text}.png`

#### Accept query

| Name     | Key         | Type                | Example |
|----------|-------------|---------------------|---------|
|Font size |`fsize`      |Integer              |20       |
|Font color|`fcolor`     |Hex Color (RRGGBBAA) |FFAA00FF |
|Font      |`f`          |String               |Helvetica|

#### Examples

##### Simple text

`http://<url:port>/-/Hello, World!.png`

##### With green color

`http://<url:port>/-/Hello, World!.png?fcolor=00AA00FF`

##### With Courier font and blue color

`http://<url:port>/-/Hello, World!.png?f=Courier&fcolor=0055CCFF`

##### With Courier font, red color and 36px size

`http://<url:port>/-/Hello, World!.png?f=Helvetica&fcolor=FF3333FF&fsize=36`

### Get fonts

`GET /api/fonts`

#### Accept query

| Name     | Key         | Type                | Example |
|----------|-------------|---------------------|---------|
|Status    |`status`     |String               |disabled |

#### Examples

##### List available fonts

`http://<url:port>/api/fonts`

###### Response

```json
[
  "Courier",
  "Helvetica"
]
```

##### List disabled fonts

`http://<url:port>/api/fonts?status=disabled`

###### Response

```json
[
  "Some-Font",
  "Another-Font-Name"
]
```

## License

Licensed under MIT

Copyright (c) 2018 [OnikurYH](https://github.com/OnikurYH)