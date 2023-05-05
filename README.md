[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/fuzzy-repo-finder/workflows/Test/badge.svg?branch=master)](https://github.com/hoto/fuzzy-repo-finder/actions)
[![Release](https://img.shields.io/github/release/hoto/fuzzy-repo-finder.svg?style=flat-square)](https://github.com/hoto/fuzzy-repo-finder/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/fuzzy-repo-finder)](https://goreportcard.com/report/github.com/hoto/fuzzy-repo-finder)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/fuzzy-repo-finder/maintainability)
[![fuzzy-repo-finder](https://snapcraft.io//fuzzy-repo-finder/badge.svg)](https://snapcraft.io/fuzzy-repo-finder)
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


[![](https://mermaid.ink/img/pako:eNq1Vm1v4jgQ_iujfNqVgPK-XHRC4i1cVjRBkG5Pp0jIxA61GmzOcUrZqv_9xglhoVz3Kp2WDxBmnnl7PDPOixVJyizbStnfGRMRG3OyUWQbChJpqeBuOVkASeEuZSoUO6I0j_iOCA3DhX9_VDqZoCnMpdIkgd_Xqv9pqOQeLT6_MXEcAzc_M7JdU3KpHoxvXc8ACN1yUSU7fqlfTkZGm7IoU1wfELC7BLjet8ky8POcuHhiKVZwjRr5nuNODSaSIuabTBHNpbgGjocGNB5eSoN7dzZzfaMK9jxJuITB3A2F5jphMHnmqeZiA3HOCcn0A8SJ3IfCMFnt94-02XCTyA0XoTgKjMpxbJj7ywBusPab3NawCWtJDza8sC3hSQV2JE33UtEKbGMSyEcmXtGL46CHnMHShzaqjzmo5DAkIOFRToZbKJcsTfN_GCH3jTHMKfyiEJ7UDOQTU-aobWjUYMFSmTwxOLMDTkOBekxlPLQROZuMAnAW_i04k9VgPp-5o0Hg-l4oxsNqkbANV86bNRg9sOgReAy3zgB4CgpHgCtGAYMwBB4gP6JaUaAbQ1HuOXKdaRBSw06xlGFz7Bk-yjVJDojRmRLQrjc-lu6q6Mu7xQeSb5XJb8lzkSQQrdl2p9P3gs38qeutBndjN_i573YNvpGEU4LSDGf4dJbveZ76X92V6e6f--2c-c0Jj080lp4LBuyytpismCDrBHl2nFAU2rMAV-kkJNUrzJkWiWEULDj4Y-IF_94SJNF5KqfjJILClmNDik0oIP-YID8mC08TSk0xEOXcnmvMMJ6PutGZZPxgvrp1l0vXm4aCJSm7jo5ry5B0Hr3fP66ccuqQUh4fSshReSorb7hza0OQ62EiAf4E_mUnFMALVn5N3bibBzN3_F7d_7fqZr2OrsVFk75b9H-06ZQJpkybjpYL5zj0n7KM089X6C4OomIG-_U-ABJFuM0KizKR8k6yYToxVTRvuKD8idOMJKEotZd1XHKfi37EPUoHlMIDI5Sp1C4WlMkgf_izivRWlxNsNN-ruuNSasop93h5fLnzt-eGwsIGt9pOipThYg-tmKtUe2TLQgt5s0Krgt9m5C5lxjC08nvgHLhjKp8sdHcuNhtmjrfrwaWl-NWqWFtEE07xveTFnHho6QdWBqFEPSJOGBzekXJ5EJFlxwT7qmJlO7Ngji8xb6QTyvGF4CTE2_wvKRGkVVb8tewX69my27_VWp1er93p1nuNL51Wt2IdLLvXrDWa3Wa73e02u71Wp_lasb7n9o1ap1Vv9-rtbr3ZNA9fXv8BYBPlBg?type=png)](https://mermaid.live/edit#pako:eNq1Vm1v4jgQ_iujfNqVgPK-XHRC4i1cVjRBkG5Pp0jIxA61GmzOcUrZqv_9xglhoVz3Kp2WDxBmnnl7PDPOixVJyizbStnfGRMRG3OyUWQbChJpqeBuOVkASeEuZSoUO6I0j_iOCA3DhX9_VDqZoCnMpdIkgd_Xqv9pqOQeLT6_MXEcAzc_M7JdU3KpHoxvXc8ACN1yUSU7fqlfTkZGm7IoU1wfELC7BLjet8ky8POcuHhiKVZwjRr5nuNODSaSIuabTBHNpbgGjocGNB5eSoN7dzZzfaMK9jxJuITB3A2F5jphMHnmqeZiA3HOCcn0A8SJ3IfCMFnt94-02XCTyA0XoTgKjMpxbJj7ywBusPab3NawCWtJDza8sC3hSQV2JE33UtEKbGMSyEcmXtGL46CHnMHShzaqjzmo5DAkIOFRToZbKJcsTfN_GCH3jTHMKfyiEJ7UDOQTU-aobWjUYMFSmTwxOLMDTkOBekxlPLQROZuMAnAW_i04k9VgPp-5o0Hg-l4oxsNqkbANV86bNRg9sOgReAy3zgB4CgpHgCtGAYMwBB4gP6JaUaAbQ1HuOXKdaRBSw06xlGFz7Bk-yjVJDojRmRLQrjc-lu6q6Mu7xQeSb5XJb8lzkSQQrdl2p9P3gs38qeutBndjN_i573YNvpGEU4LSDGf4dJbveZ76X92V6e6f--2c-c0Jj080lp4LBuyytpismCDrBHl2nFAU2rMAV-kkJNUrzJkWiWEULDj4Y-IF_94SJNF5KqfjJILClmNDik0oIP-YID8mC08TSk0xEOXcnmvMMJ6PutGZZPxgvrp1l0vXm4aCJSm7jo5ry5B0Hr3fP66ccuqQUh4fSshReSorb7hza0OQ62EiAf4E_mUnFMALVn5N3bibBzN3_F7d_7fqZr2OrsVFk75b9H-06ZQJpkybjpYL5zj0n7KM089X6C4OomIG-_U-ABJFuM0KizKR8k6yYToxVTRvuKD8idOMJKEotZd1XHKfi37EPUoHlMIDI5Sp1C4WlMkgf_izivRWlxNsNN-ruuNSasop93h5fLnzt-eGwsIGt9pOipThYg-tmKtUe2TLQgt5s0Krgt9m5C5lxjC08nvgHLhjKp8sdHcuNhtmjrfrwaWl-NWqWFtEE07xveTFnHho6QdWBqFEPSJOGBzekXJ5EJFlxwT7qmJlO7Ngji8xb6QTyvGF4CTE2_wvKRGkVVb8tewX69my27_VWp1er93p1nuNL51Wt2IdLLvXrDWa3Wa73e02u71Wp_lasb7n9o1ap1Vv9-rtbr3ZNA9fXv8BYBPlBg)
