[![License][License-Image]][License-URL] 

# go-chatgpt-twitter-bot
Twitter bot backed by ChatGPT

## Install

```
go get -v github.com/zquestz/go-chatgpt-twitter-bot
cd $GOPATH/src/github.com/zquestz/go-chatgpt-twitter-bot
make
make install
```

## Twitter Setup

Make sure the env var `TWITTER_BEARER_TOKEN` is set. The Twitter v2 API is used. You can generate a Bearer Token within the developer dashboard. There is also support for `.env` files.

## Autocompletion

Autocompletion is supported. To set up autocompletion:

### Bash Linux

```
go-chatgpt-twitter-bot --completion bash > /etc/bash_completion.d/go-chatgpt-twitter-bot
```

### Bash MacOS

```
go-chatgpt-twitter-bot --completion bash > /usr/local/etc/bash_completion.d/go-chatgpt-twitter-bot
```


### Zsh

Generate a `_go-chatgpt-twitter-bot` completion script and put it somewhere in your `$fpath`:

```
go-chatgpt-twitter-bot --completion zsh > /usr/local/share/zsh/site-functions/_go-chatgpt-twitter-bot
```

### Fish

```
go-chatgpt-twitter-bot --completion fish > ~/.config/fish/completions/go-chatgpt-twitter-bot
```

### Powershell

```
(& go-chatgpt-twitter-bot --completion powershell) | Out-String | Invoke-Expression
```

## Configuration

To setup your own default configuration just create `~/.config/go-chatgpt-twitter-bot/config`. The configuration file is in UCL format. JSON is also fully supported as UCL can parse JSON files.

For more information about UCL visit:
[https://github.com/vstakhov/libucl](https://github.com/vstakhov/libucl)

The following keys are supported:

* handle (twitter handle)
* userid (twitter userid)
* verbose (verbose mode)

## Contributors

* [Josh Ellithorpe (zquestz)](https://github.com/zquestz/)

## License

go-chatgpt-twitter-bot is released under the MIT license.

[License-URL]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
