# webkit

C bindings to webkit for Go (Golang).

[![](https://img.shields.io/circleci/token/abf9e47762afcbbd936490819683ad44594f67b5/project/abcum/webkit/master.svg?style=flat-square)](https://circleci.com/gh/abcum/webkit) [![](https://img.shields.io/badge/status-beta-ff00bb.svg?style=flat-square)](https://github.com/abcum/webkit) [![](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abcum/webkit) [![](https://goreportcard.com/badge/github.com/abcum/webkit?style=flat-square)](https://goreportcard.com/report/github.com/abcum/webkit) [![](https://img.shields.io/badge/license-Apache_License_2.0-00bfff.svg?style=flat-square)](https://github.com/abcum/webkit) 

#### Installation

```bash
go get github.com/abcum/webkit
```

#### Building

[![](https://img.shields.io/badge/api-docs-ffbb00.svg?style=flat-square)](https://webkitgtk.org/reference/webkit2gtk/stable/index.html) [![](https://img.shields.io/badge/version-webkit2gtk+2.14-0ffbbb.svg?style=flat-square)](https://webkitgtk.org/reference/webkit2gtk/stable/api-index-2-14.html)

```bash
go install -tags 'webkit_2_14' 
```

[![](https://img.shields.io/badge/api-docs-ffbb00.svg?style=flat-square)](https://webkitgtk.org/reference/webkit2gtk/stable/index.html) [![](https://img.shields.io/badge/version-webkit2gtk+2.16-0ffbbb.svg?style=flat-square)](https://webkitgtk.org/reference/webkit2gtk/stable/api-index-2-16.html)

```bash
go install -tags 'webkit_2_16' 
```

#### Requirements

- Golang >= 1.2
- GTK+ >= 3.10
- WebKit2GTK+ >= 2.0.0

_Alpine Linux_
```bash
apk --update install gcc musl-dev gtk+3.0-dev webkit2gtk-dev
```

_Ubuntu 13.10_
```bash
sudo add-apt-repository ppa:gnome3-team/gnome3-staging
sudo apt-get update
sudo apt-get install libwebkit2gtk-3.0-dev
```

_Ubuntu 13.04_
```bash
sudo add-apt-repository ppa:gnome3-team/gnome3
sudo apt-get update
sudo apt-get install libwebkit2gtk-3.0-dev
```

_Arch Linux_
```bash
sudo pacman -S webkitgtk
```

#### Copyright

For WebKit licensing and copyright information, [click here](https://webkit.org/licensing-webkit/).
