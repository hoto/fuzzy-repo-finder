[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/fuzzy-repo-finder)
[![CircleCI](https://circleci.com/gh/hoto/fuzzy-repo-finder/tree/master.svg?style=svg)](https://circleci.com/gh/hoto/fuzzy-repo-finder/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/fuzzy-repo-finder)](https://goreportcard.com/report/github.com/hoto/fuzzy-repo-finder)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/fuzzy-repo-finder/maintainability)
[![Release](https://img.shields.io/github/release/hoto/fuzzy-repo-finder.svg?style=flat-square)](https://github.com/hoto/fuzzy-repo-finder/releases/latest)
# Fuzzy Repo Finder

Command line tool for navigating git repositories.

### Installation

Use snap:

    sudo snap install fuzzy-repo-finder
    
Or download binary from [releases](https://github.com/hoto/fuzzy-repo-finder/releases):

Linux and mac:

    sudo curl -L \
      "https://github.com/hoto/fuzzy-repo-finder/releases/download/2.1.0/fuzzy-repo-finder_2.1.0_$(uname -s)_$(uname -m)" \
       -o /usr/local/bin/fuzzy-repo-finder

    sudo chmod +x /usr/local/bin/fuzzy-repo-finder

### Configuration and running

Add to your `~/.bashrc` or `~/.zshrc` or `~/.profile`:  

    function go_to_project() {
      cd $(fuzzy-repo-finder --projectRoots "${HOME}/projects,${HOME}/go/src" $@)
    }
    alias g='go_to_project'

In terminal:

    $ g
    
Find projects by partial name:

    $ g myprojectname
    
Debug:
  
    $ fuzzy-repo-finder --projectRoots "${HOME}/projects,${HOME}/go/src" --debug myprojectname
    
Help:
  
    $ fuzzy-repo-finder --help

### Demo

From directory structure:

    ~/projects
      ├── group_A
      │   ├── project_1
      │   ├── project_2
      │   └── project_3
      └── group_B
          ├── project_1
          ├── project_2
          └── group_C
              └── project_1

Unfiltered:

    Search: 
    group_A
        project_1
        project_2
        project_3
    group_B
        project_1
        project_2
        project_3
    group_B/group_C
        project_1

Filtered:

    Search: pr1
    group_A
        project_1
    group_B
        project_1
    groupB/group_C
        project_1

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

### TODO:
* Fix order when scrolling through projects
* Fix versioning in snapcraft builds (they work only with goreleaser ATM)

---
_Following_ [_Standard Go Project Layout_](https://github.com/golang-standards/project-layout)
