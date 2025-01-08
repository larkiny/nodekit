---
title: Getting Started
description: Troubleshooting nodekit installation
---

This section outlines **common errors encountered** when executing the nodekit installation command.

**If you are looking for the installation instructions instead, they are located [here](/guides/10-getting-started).**

### A nodekit file already exists in the current directory.

If you run the installer command more than once, you will see: 

> ERROR: An nodekit file already exists in the current directory. Delete or rename it before installing.

If you want to fetch the latest version of nodekit, you can delete the existing file:

```bash
rm nodekit
```

And then run the [Getting Started](/guides/10-getting-started) command again.
