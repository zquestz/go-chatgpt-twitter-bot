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

## ENV Setup

There are quite a few API keys and ENV vars to setup. You can use a `.env` file if you want. We use the Twitter v2 API, and OpenAI.

```
TWITTER_BEARER_TOKEN=
TWITTER_CONSUMER_API_KEY=
TWITTER_CONSUMER_SECRET=
TWITTER_ACCESS_TOKEN=
TWITTER_ACCESS_SECRET=
OPENAI_API_KEY=
OPENAI_CHARACTER_BACKGROUND=
```

Make sure to setup a character background. This will be the personality that will post tweets!

## Usage

```
Usage:
  go-chatgpt-twitter-bot <handle> [flags]

Flags:
      --completion string   completion script for bash, zsh, fish or powershell
  -d, --dryrun              dryrun mode
      --handle string       twitter handle
  -h, --help                help for go-chatgpt-twitter-bot
  -t, --tweet               tweet to main account
      --userid string       twitter userid
  -v, --verbose             verbose mode
      --version             display version
```

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

* character_background (bot personality)
* dryrun (don't post to twitter)
* handle (twitter handle)
* tweet (tweet to main account)
* userid (twitter userid)
* verbose (verbose mode)

## Contributors

* [Josh Ellithorpe (zquestz)](https://github.com/zquestz/)

## License

go-chatgpt-twitter-bot is released under the MIT license.

[License-URL]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
