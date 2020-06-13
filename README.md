[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/fuzzy-repo-finder/workflows/Build%20and%20test/badge.svg?branch=master)](https://github.com/hoto/fuzzy-repo-finder/actions)
[![Release](https://img.shields.io/github/release/hoto/fuzzy-repo-finder.svg?style=flat-square)](https://github.com/hoto/fuzzy-repo-finder/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/fuzzy-repo-finder)](https://goreportcard.com/report/github.com/hoto/fuzzy-repo-finder)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/fuzzy-repo-finder/maintainability)
[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/fuzzy-repo-finder)
# Fuzzy Repo Finder

Command line tool for navigating git repositories.

### Installation

Mac:

    brew install hoto/repo/fuzzy-repo-finder

Mac or Linux:

    sudo curl -L \
      "https://github.com/hoto/fuzzy-repo-finder/releases/download/2.2.1/fuzzy-repo-finder_2.2.1_$(uname -s)_$(uname -m)" \
       -o /usr/local/bin/fuzzy-repo-finder

    sudo chmod +x /usr/local/bin/fuzzy-repo-finder
    
Snap:

    sudo snap install fuzzy-repo-finder
    
Or manually download binary from [releases](https://github.com/hoto/fuzzy-repo-finder/releases).

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
    group_B/group_C
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

    make run arg="--projectRoots=${HOME}/projects project_to_look_for"

Install to global golang bin directory:

    make install

### TODO:
* Fix order when scrolling through projects
* Fix versioning in snapcraft builds (they work only with goreleaser ATM)

---
_Following_ [_Standard Go Project Layout_](https://github.com/golang-standards/project-layout)
