---
slug: "/git/"
date: "2022-03-14"
title: "My git cheatsheet"
description: "git commands I often forget"
published: false
---

## Basics

1. git init: start a new repository
1. git status: check current status
1. git commit -am: Commit all changes in the working directory with a message added on

5. git switch -c [branch name]: Create a branch and switch to it. (git checkout -b will do the same thing)

6. git switch [branch name]: switch between named branches

7. git merge[branch]: merge named branch into current branch

8. git branch [branch name] -d : delete branch if not tracked

9. git log --pretty=oneline: show a graph of commit history.

11. git push

12. git clone [repo]: Copy of a project/file onto your local computer.

git diff

git rebase -i

git reset

`git commit -a -m`

`git switch -c some-name` Take your changes to a new branch, creates new branch

`git switch -t some-name` Take your changes to a new branch, creates new branch
and sets origin as upstream

`git fetch origin && git checkout -b vince-patch-name origin/master`

`git add .` will add all files of the current directory to track (including
untracked)

`git remote set-url origin git@github.com:someuser/somerepo.git` set origin

`git remote add upstream` add upstream

`git remote -v` check origins

`git stash branch branchname` move changes to new branch

`git reset hard HEAD~2` reset to last commit WARNING: deletes your uncommited
changes

`git branch | grep -v "master" | xargs git branch -D ` delete all local branches
except master

`git rev-list --count --since="Dec 3 2015" --before="Jan 3 2016" --all` number
of commits over timeframe

git log --graph --oneline
git diff --word-diff-regex=<regex>

Merge conflict:

```
    press "i" (i for insert)
    write your merge message
    press "esc" (escape)
    write ":wq" (write & quit)
    then press enter
```

_Use with caution!_

```
git remote add upstream /url/to/original/repo
git fetch upstream
git checkout master
git reset --hard upstream/master
git push origin master --force
```

git rm --cached -r . (remove everything from git cache)
add to bootcamp



    Delete your local branch: 
    git branch -d local_branch

    Fetch the latest remote branch: 
    git fetch origin remote_branch

    Rebuild the local branch based on the remote one:
    git checkout -b local_branch origin/remote_branch





https://dev.to/guivern/git-how-to-fix-a-bad-commit-4068
