# Webhook

Webhook hosted by golang

## Download

```bash
git clone https://github.com/LucienShui/Webhook.git -b build --depth=1
```

## Usage

```plaintext
Usage of webhook:
  -c string
        -c <config file> (default "./config.json")
  -version
        Print version information
```

## Config

For example, `config.json` looks like below:

```json
{
  "address": "0.0.0.0",
  "port": 10086,
  "webhooks": [
    {
      "name": "beat",
      "script": "echo 'It works!'",
      "password": ""
    },
    {
      "name": "hello",
      "script": "set -x && echo 'Hello World!'",
      "password": "world"
    }
  ]
}
```

One can access the API `/:name?password=` to execute the script

## API

### Execute

`GET` `/:name?password=[password]`

#### Params

| Name | Type | Description |
| :---: | :---: | --- |
| name | string | record's name |
| password | string | record's password |

### Log

`GET` `/:name/log?raw=`

#### Params

| Name | Type | Description |
| :---: | :---: | --- |
| name | string | record's name |
| raw | bool | show execute's output in raw or not, false for default |
