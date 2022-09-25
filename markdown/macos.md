---
pagetitle: "My MacOS Setup"
tags: ["go", "yaml", "json", "toml"]
---

### My MacOS Setup

_Last updated on 24.03.2022_

Start with a clean install.
Open the [terminal](<https://en.wikipedia.org/wiki/Terminal_(macOS)>) app and copy and paste the commands.  
This is probably outdated as I only update it when setting up MacOs.

> ! Only run these commands if you know what you are doing. !

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

```bash
brew install zsh
brew install zsh-completions
brew install --cask iterm2
```

#### Customize shell

[ohmyzsh](https://github.com/ohmyzsh/ohmyzsh)

```bash
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

[powerlevel10k](https://github.com/romkatv/powerlevel10k)

```bash
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >>~/.zshrc
```

##### Autosuggestions

```bash
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
```

Add inside ~/.zshrc:

```bash
plugins=(
    # other plugins...
    zsh-autosuggestions
)
```

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
  sync = town sync
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

#### Install design, content, and other tools

```bash
brew install --cask figma
brew install --cask obs
brew install --cask adobe-creative-cloud
brew install --cask vlc
brew install --cask foobar2000
brew install --cask figma
brew install youtube-dl
```

#### Install misc. tools

```bash
brew install alt-tab
brew install --cask rectangle
```

#### Reset Launchpad & Dock

```bash
defaults write com.apple.dock ResetLaunchPad -boolean true; killall Dock
brew update && brew upgrade && brew doctor && brew cleanup
```

Browser settings:

- Turn off _save passwords_ on all browsers
- [uBlock Origin](https://github.com/gorhill/uBlock)
- [HTTPS Everywhere](https://www.eff.org/https-everywhere)
- [Privacy Badger](https://privacybadger.org/)

<img class="center" src="./img/hero-macos.png" alt="A decorative image" />
