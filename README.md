# mu

A Micro app platform

## Overview

Mu is a Micro app platform that provides a simple set of building blocks for life. It was born out of frustration with existing services using Ads or creating addictive behaviours. The goal is to create a separate system that addresses daily needs and nothing else. Part of that 
is focusing on Muslim related requirements like prayer times, the quran in english, etc. Other parts are related to dev interests like 
hacker news, open source, crypto.

## Apps

The current list of apps:

- **Chat** - Channel based AI chat
- **News** - Topic based news feed
- **Pray** - Islamic prayer times
- **Reminder** - The Quran in English

## Dependencies

- Go toolchain

## Usage

Download source

```bash
go install mu.dev/cmd/mu@latest
```

Run it

```
mu
```

Goto `localhost:8080`
## APIs

Set `OPENAI_API_KEY` from `openai.com` for ability to chat with AI

```
export OPENAI_API_KEY=xxx
```

Set `SUNNAH_API_KEY` from `sunnah.com` for daily hadith in news app

```
export SUNNAH_API_KEY=xxx
```

Set `CRYPTO_API_KEY` from `cryptocompare.com` for crypto market tickers

```
export CRYPTO_API_KEY=xxx
```

## PWA

Mu operates as a progressive web app. The main app can be installed just like a native app. 

Caching and offline mode is still a WIP.

## Development

Currently hacking on this for personal reasons. Join [Discord](https://mu.xyz/discord) if interested.
