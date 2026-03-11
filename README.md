# Dev Setup

Scripts and dotfiles to set up my local environment.

OS: Ubuntu 24.04

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/rafaelim/dev-setup/main/install.sh | bash
```

Clones the repo to `~/.rig`, installs all tools, and deploys dotfiles.

## CLI

```
rig sync       deploy dotfiles from repo to ~/
rig push       copy dotfile changes from ~/ back to repo and push
rig upgrade    download latest rig binary from GitHub releases
```
