# slack-bot-file-monitor

[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/slack-bot-file-monitor.svg)](https://github.com/suzuki-shunsuke/slack-bot-file-monitor/releases)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/slack-bot-file-monitor.svg)](https://github.com/suzuki-shunsuke/slack-bot-file-monitor)

slack bot to monitor file uploads

## Note

This bot can't post a message to private channels because the [files.info](https://api.slack.com/methods/files.info) API does not returns private channel ids.

## Required Slack token's scope

* [chat.postMessage](https://api.slack.com/methods/chat.postMessage)
* [files.info](https://api.slack.com/methods/files.info)

## Run

```
$ cp env.sh.tmpl env.sh
$ vi env.sh
$ dep ensure
$ make run
```

## License

[MIT](LICENSE)
