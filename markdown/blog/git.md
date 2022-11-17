---
pagetitle: "git Cheatsheet"
---

# git cheatsheet

_Last updated on 17.11.2022_

> git gets easier once you get the basic idea that branches are homeomorphic endofunctors mapping submanifolds of a Hilbert space.

## Basics

- `git init`: start a new repository.
- `git status`: check current status.
- `git commit -am "<commit-message>`: commit all changes in the working directory with a message added on.
- `git switch -c <branch-name>`: create a branch and switch to it (`git checkout -b` will do the same, omit `-c` to just switch branches, use `-t` to set origin as upstream).
- `git merge [branch]`: merge named branch into current branch.
- `git branch <branch-name> -d` : delete branch if not tracked (`-D` to force, caution!).
- `git log --pretty=oneline`: show a graph of commit history.
- `git remote -v` : check origins.
- `git branch | grep -v "master" | xargs git branch -D` : delete all local branches except master.
- `git add .` : add all files of the current directory to track (including untracked).
- `git rev-list --count --since="Jan 1 2020" --before="Jan 1 2021" --all` number of commits over timeframe.
- `git log --format='%aN' | sort -u` : list all contributors by name.
- `git log --all --full-history -- "**/some-file.*"`: search git history for a specific file. Useful when something got deleted.
- `git fetch origin <your-branch>:<your-branch>`: fetch a branch.

### config 

- `git config --global color.ui auto`: helpful colors.
- `git config --global user.name "vinckr"`: set user name globally for commits.
- `git config --global user.email "vinckr@example.com"`: set email globally for commits.

## Troubleshooting

**Use the following with caution!**

- `git revert <commit-id>` : revert a commit that is already pushed.
- `git commit --amend -m "<new-commit-message>"` : fix commit message (only use if you haven't pushed already!).
- `git rm --cached -r .` : remove everything from git cache.
- `git reset --soft HEAD~N` : undo last N commits, but keep changes locally.
- `git reset --hard HEAD~N` : undo and delete last N commits.
- `git reset --hard origin/<branch-name>` : undo all local commits to a branch.
- add files to last commit (only use if you haven't pushed already!).

```bash
git add <file_name>
git commit --amend HEAD~1
```

- delete local copy of branch and fetch upstream:

```bash 
git switch master
git branch -d <branch-name>
git fetch origin <branch-name>
git checkout -b <branch-name> origin/<branch-name>
```

## Further reference

- [Oh Shit, Git!?!](https://ohshitgit.com/)
- [A Visual Git Reference](http://marklodato.github.io/visual-git-guide/index-en.html)
- [Git for the lazy](https://wiki.spheredev.org/index.php/Git_for_the_lazy)
- [Git Immersion: A guided tour](https://gitimmersion.com/)
- [git-fu](https://gitfu.wordpress.com/)

<img class="center" src="./img/hero-git.png" alt="Content ends here" />
