[![CircleCI](https://circleci.com/gh/hoto/fuzzy-repo-finder/tree/master.svg?style=svg)](https://circleci.com/gh/hoto/fuzzy-repo-finder/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/fuzzy-repo-finder)](https://goreportcard.com/report/github.com/hoto/fuzzy-repo-finder)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/fuzzy-repo-finder/maintainability)
[![Release](https://img.shields.io/github/release/hoto/fuzzy-repo-finder.svg?style=flat-square)](https://github.com/hoto/fuzzy-repo-finder/releases/latest)
# Fuzzy Repo Finder

Command line tool for navigating git repositories.

WIP: This app is still in prototype phase.

### Installation

Download binary from [releases](https://github.com/hoto/fuzzy-repo-finder/releases):

Linux:

    mkdir ~/bin
    
    curl -L \
      "https://github.com/hoto/fuzzy-repo-finder/releases/download/1.0.0/fuzzy-repo-finder_1.0.0_$(uname -s)_$(uname -m)" \
       -o ~/bin/fuzzy-repo-finder

    chmod +x ~/bin/fuzzy-repo-finder

Create a config file:

    $ vim ~/.fuzzy-repo-finder/config.yml 
    
    ---
    project_roots:
      - '${HOME}/projects'
      - '${HOME}/go/src'


Add alias to your `~/.bashrc` or `~/.zshrc`:  

    function go_to_project() {
      local pattern=$1
      ~/bin/fuzzy-repo-finder ${pattern}
      local selectedProjectPath="$(cat ~/.fuzzy-repo-finder/selected_project.txt)"
      cd "${selectedProjectPath}"
    }
    alias g='go_to_project'

You can chose any alias name you want.   
In my case I'm using `g`.  

Use in terminal:

    $ g myprojectname

Or without arguments:

    $ g

### Demo

From directory structure:

```
~/projects
  ├── group_A
  │   ├── project_1
  │   ├── project_2
  │   └── project_3
  └── group_B
      ├── project_1
      ├── project_2
      └── project_3
```

Unfiltered:

```
Search: 
group_A
    project_1
    project_2
    project_3
group_B
    project_1
    project_2
    project_3
```

Filtered:

```
Search: t_1
group_A
    project_1
group_B
    project_1
```

![demo](https://github.com/hoto/fuzzy-repo-finder/wiki/images/001.png)  

![demo](https://github.com/hoto/fuzzy-repo-finder/wiki/images/002.gif)  

![demo](https://github.com/hoto/fuzzy-repo-finder/wiki/images/005.gif)  

---

### Development

Get:

    go get github.com/hoto/fuzzy-repo-finder/cmd/fuzzy-repo-finder/

Download dependencies:

    make dependencies

Build, test and run:

    make clean
    make build
    make test
    make run

Run with arguments:

    make args="myprojectname" run

Install to global golang bin directory:

    make install
    
Dry run gorelease (auto releasing to github release page):

    make release_dry_run

### TODO:
* Fix order when scrolling through projects
* Pass flags which can override `config.yml`
* Make $HOME a default project root if not provided
* Add --version info

---
_Following_ [_Standard Go Project Layout_](https://github.com/golang-standards/project-layout)
