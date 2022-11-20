---
pagetitle: "Misc cheatsheet"
---

# Misc cheatsheet

_Last updated on 17.11.2022_

A mixed collection of commands I had to look up at some point and collect here for future reference. 
Probably not useful for anyone else but me ;-).

## Mixed/UNIX

- `history | awk '{print $2}' | sort | uniq -c | sort -rn | head`: see which commands you run often.
- `curl ipinfo.io`: external ip address using ipinfo.io
- `lsof -iTCP -sTCP:LISTEN -n -P`: show all active ports.
- `kill -9 <PID>`: kill process for a port.
- `dig WWW.VINCKR.COM +nostats +nocomments +nocmd`: DNS lookup. 
- `ps uax`: networks.
- `chmod +x name_of_file.sh`: make bash file executable.

## CSS

- `div > p`: child selector selects all elements that are the children of a specified element ([See also](https://www.w3schools.com/css/css_combinators.asp))

## Docker

- `docker kill $(docker ps -q)`: stop all Docker containers.
- `docker rm $(docker ps -a -q)`: remove all Docker containers.
- `docker rmi $(docker images -q)`: remove all Docker images. 

## git town

Aliases for git-town, add to git config:

```
git config --global alias.append "town append"
git config --global alias.hack "town hack"
git config --global alias.kill "town kill"
git config --global alias.repo "town repo"
git config --global alias.ship "town ship"
git config --global alias.sync "town sync"
git config --global alias.new-pr "town new-pull-request"
git config --global alias.prepend "town prepend"
git config --global alias.prune-branches "town prune-branches"
git config --global alias.rename-branch "town rename-branch"
```

<img class="center" src="./img/hero-git.png" alt="Content ends here" />
