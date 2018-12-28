[![CircleCI](https://circleci.com/gh/hoto/fuzzy-repo-finder/tree/master.svg?style=svg)](https://circleci.com/gh/hoto/fuzzy-repo-finder/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/fuzzy-repo-finder)](https://goreportcard.com/report/github.com/hoto/fuzzy-repo-finder)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/fuzzy-repo-finder/maintainability)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/hoto/fuzzy-repo-finder/releases/latest)
# Fuzzy Repo Finder

Command line tool for git projects navigation.

### Install

Download binary from [releases](<NEED_LINK>):

    curl <NEED_LINK>

Add to `~/.bashrc`, `~/.zshrc` or `~/.profile`:

    function go_to_project() {
      local pattern=$1
      fuzzy-repo-finder ${pattern}
      local selectedProjectPath=$(cat ~/.fuzzy-repo-finder/selected_project.txt)
      cd ${selectedProjectPath}
    }
    alias g='go_to_project'


### Usage

In terminal:

    $ g myprojectname

### Demo

![demo](https://github.com/hoto/fuzzy-repo-finder/wiki/images/fuzzy-repo-finder-demo-001.png)

### Development

Clean:

    make clean

Build:

    make build

Test:

    make test

Run: 

    make run
    make args="projectname" run

Install:

    make install

### TODO:
* Publish releases on github on every git tag
* Add installation instructions using github releases
* Setup brew,fedora,debian,ubuntu,arch,packman repo
* Stop walking directories tree when a match is find (optimization)
* Query by group
* Read config file from `~/.fuzzy-repo-finder/config.yml`
* Pass flags which can override `config.yml`
* Save found repositories in `~/.fuzzy-repo-finder/repositories_statistics.yml`
* Display cached repos from `repos.yml` before updating with real data
* Sort repos by usage or alphabetically
* Show dirty status of a repository using `*`
* How I want the presentation to look like:

```
Search: 
mango
    prices *                                 (116)
    purchase-notifications                   (29)
hoto
    home *                                   (150)
    home-private                             (81)
    jenkinsfile-examples *                   (49)
    jenkinsfile-loader                       (48)
    ansible-home-fedora                      (17)
    jenkins-shared-library *                 (17)
    ansible-sointeractive                    (16)
    deja-vu-sans-mono-font                   (15)
    demo-ansible-and-docker-swarm-comparison (15)
    demo-ansible-galaxy                      (12)
    demo-ansible-role-nginx                  (10)
    docker-presentation                      (9)
    flake8                                   (4)
    fuzzy-project-finder                     (1)
    git-branch-ps1                              
    hello-world                                 
    micro-twitter                               
    project-templates                           
    vagrant-ubuntu-base                         
    vagrant-ubuntu-workstation                  
github
    fedora-desktop-ansible-roles             (2)


=========


Search: pri
mango
    prices *               (116)
    purchase-notifications (29)

```

---
_Following_ [_Standard Go Project Layout_](https://github.com/golang-standards/project-layout)