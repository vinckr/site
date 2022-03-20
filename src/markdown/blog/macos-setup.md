---
slug: "macos-setup"
date: "2022-03-14"
title: "My MacOS Setup"
description: "My MacOS setup."
---

Start with a clean install.
Open the [terminal](https://en.wikipedia.org/wiki/Terminal_(macOS)) app and copy and paste the commands.

>While nothing here can harm your hardware, only run these commands if you know what you are doing.

#### Widen the space between apps in Dock

```bash
defaults write com.apple.dock persistent-apps -array-add '{"tile-type"="spacer-tile";}' && \
killall Dock
```

#### Enable autohide dock

```bash
defaults write com.apple.dock autohide -boolean true && \
killall Dock
```

#### Show hidden files in finder

```bash
defaults write com.apple.finder AppleShowAllFiles -boolean true && \
killall Finder
```

#### Create projects folder

(here I store my coding related files & repositories)

```bash
mkdir  -p ~/projects
```

#### Install [Xcode](https://en.wikipedia.org/wiki/Xcode)

```bash
xcode-select --install
```

#### Install [brew](https://brew.sh/)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)
```

#### Install shells

brew install zsh
brew install zsh-completions
brew install --cask iterm2

#### Customize shell

https://github.com/ohmyzsh/ohmyzsh

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

https://github.com/romkatv/powerlevel10k

brew install romkatv/powerlevel10k/powerlevel10k && \
echo '# Theme configuration: PowerLevel10K' >>! ~/.zshrc && \
echo 'source /usr/local/opt/powerlevel10k/powerlevel10k.zsh-theme' >>! ~/.zshrc && \
echo '# To customize prompt, run `p10k configure` or edit ~/.p10k.zsh.' >>! ~/.zshrc && \
echo '[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh' >>! ~/.zshrc
omz update

#### Install & customize git

brew install git

```bash
cd ~
touch .gitconfig
open .
```

Add to .gitconfig:

```bash
[alias]
  # Show verbose output about tags, branches or remotes
  tags = tag -l
  branches = branch -a
  remotes = remote -v
  # Pretty log output
  hist = log --graph --pretty=format:'%Cred%h%Creset %s%C(yellow)%d%Creset %Cgreen(%cr)%Creset [%an]' --abbrev-commit --date=relative

[color]
  # Use colors in Git commands that are capable of colored output when
  # outputting to the terminal. (This is the default setting in Git ≥ 1.8.4.)
  ui = auto
[color "branch"]
  current = yellow reverse
  local = yellow
  remote = green
[color "diff"]
  meta = yellow bold
  frag = magenta bold
  old = red bold
  new = green bold
[color "status"]
  added = yellow
  changed = green
  untracked = cyan
```

#### Install browsers

```bash
brew install --cask firefox
brew install --cask chromium
brew install --cask google-chrome
brew install --cask brave-browser
brew install --cask opera
```

#### Install development tools

```bash
brew install --cask sublime-text
brew install --cask sublime-merge
brew install --cask visual-studio-code
brew install go
brew install node
brew install yarn
brew install git-town
brew install --cask docker
brew install gh
```

#### Install social tools

```bash
brew install --cask discord
brew install --cask slack
brew install --cask signal
```

#### Install design & content tools

```bash
brew install --cask figma
brew install --cask obs
brew install --cask adobe-creative-cloud
brew install vcl
````

#### update npm

`npm i -g npm@latest`

#### Install misc. tools

```bash
brew install alt-tab
brew install --cask rectangle
````

#### Reset Launchpad & Dock

```bash

defaults write com.apple.dock ResetLaunchPad -boolean true; killall Dock
brew update && brew upgrade && brew doctor && brew cleanup
```


gh repo clone ory/kratos; gh repo clone ory/hydra; gh repo clone ory/keto; gh repo clone ory/oathkeeper; gh repo clone ory/web; gh repo clone ory/docs; gh repo clone ory/meta

Browser settings:

- Turn off _save passwords_ on all browsers
extensions:
- [uBlock Origin](https://github.com/gorhill/uBlock)
- [HTTPS Everywhere](https://www.eff.org/https-everywhere)
- [Privacy Badger](https://privacybadger.org/)
- [Bitwarden](https://addons.mozilla.org/en-US/firefox/addon/bitwarden-password-manager/)
