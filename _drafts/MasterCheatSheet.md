---
title: `vinckr Cheatsheet`
---

good vid
https://www.youtube.com/watch?v=2sjqTHE0zok

## git town

git config --global alias.append "town append"

git config --global alias.hack "town hack"

git config --global alias.kill "town kill"

git config --global alias.new-pull-request "town new-pull-request"

git config --global alias.prepend "town prepend"

git config --global alias.prune-branches "town prune-branches"

git config --global alias.rename-branch "town rename-branch"

git config --global alias.repo "town repo"

git config --global alias.ship "town ship"

git config --global alias.sync "town sync"

### UNIX

`ps uax` all network shizzle

`CTRL+z` -> `bg` (programm in background) -> `fg` (programm visible again)

`Time` CRON: https://crontab.guru/

Make bash file executable: chmod +x name_of_file.sh Execute bash file: sh
./name_of_file.sh

---

### Github CLI

gh pr create --title "docs: os motivation fix"

format: npm i npm run format git commit -a -m "styles: format" git push

#clone kratos git clone https://github.com/ory/kratos.git #clone hydra git clone
https://github.com/ory/hydra.git #clone oathkeeper git clone
https://github.com/ory/oathkeeper.git #clone keto git clone
https://github.com/ory/keto.git

echo Setting up username

Set your user name globally for commits

git config --global user.name "vinckr"

Set your email globally for commits

git config --global user.email "vincent@ory.sh"

Some helpful colours

git config --global color.ui auto

---

## Github API

https://api.github.com/repos/ory/kratos > kratos.json

curl -i -H "Accept: application/json" -H "Content-Type: application/json"
https://api.github.com/repos/ory/kratos -o kratos.json

GET /repos/:owner/:repo/contributors

https://api.github.com/repos/ory/kratos/contributors

https://api.github.com/repos/ory/kratos/contributors?anon=1

GET /repos/:owner/:repo/stats/contributors

https://api.github.com/repos/ory/kratos/stats/contributors?anon=1

This search provides all discussions from ory with less than 1 comment.
https://github.com/search?q=user%3Aory+comments%3A%3C1&type=Discussions user:ory
comments:<1 type=Discussions

`https://github.com/search?q=created%3A%3E%3D2021-01-25+author%3Avinckr` pull
requests in a given timeframe

---

## After Effects

Control Shift D = Cut

Cmd+K = Composition Settings

## Other

`curl ipinfo.io`

If you want to know your external ip address using ipinfo.io you can simple run this command or you even add this command as your bash alias.

history | awk '{print $2}' | sort | uniq -c | sort -rn | head.

If you want to see which commands you run often you can run the following command and you can even add in your .bashrc as an alias.

with grip

> render markdown in browser (for exporting as pdf) `grip your_markdown.md`

> spellcheck `spellchecker -f file`

## Useful links

docs.new, sheets.new, slides.new
