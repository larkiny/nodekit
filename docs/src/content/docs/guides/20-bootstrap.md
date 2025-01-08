---
title: Installing and bootstrapping the Algorand node
description: Installing and bootstrapping the Algorand node
---

## How to start

The bootstrap process is automatically started after following the [nodekit installation instructions](/guides/10-getting-started), but it can also be trigerred manually by running this command:

```bash
./nodekit bootstrap
```

## What it does

Nodekit will check to see if there is a node already installed.

If there is none, it will ask if you want to start a node installation:

> Installing A Node
>
> It looks like you're running this for the first time. Would you like to install a node? (Y/n)

You can respond by pressing the `y` or `n` keys of your keyboard, for `yes` and `no` respectively.

---

It will then ask if you want to perform a "fast-catchup" with the network after installation is complete:

> Regular sync with the network usually takes multiple days to weeks. You can optionally perform fast-catchup to sync in 30-60 minutes instead.
>
> Would you like to preform a fast-catchup after installation? (Y/n)

Fast-catchup saves a lot of time, so we recommend responding `Yes`.
