---
title: Generating consensus participation keys
description: Generating consensusparticipation keys with nodekit
---

If it is not running already, start nodekit with this command:

```bash
./nodekit
```

After your node has fully synced with the network, you will see a green `RUNNING` label at the top:

![](/assets/nodekit-state-running.png)

You will only be able to generate participation keys after your node is in a `RUNNING` state

## Generate participation keys

Press the `G` key to start generating participation keys.

Nodekit will ask the account address that you will be participating in consensus:

![](/assets/nodekit-partkey-gen-1.png)

Enter your account address and press ENTER

## Select participation key duration

Nodekit will ask the number of days that the participation keys will be valid for:

![](/assets/nodekit-partkey-gen-2.png)

You can press the S key to cycle through duration modes in days / months / rounds.

The longer your duration, the longer the participation key generation step will take to complete. 

:::note
**Did you encounted any errors?**
Check out the [Troubleshooting: Generating consensus participation keys](/troubleshooting/30-generate-participation-keys.md) section.
:::
