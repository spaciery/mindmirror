# mindmirror

Mindmirror is a minimal static site generator intented to use as a note taking/sharing application.It is customizable and can support standard markdown
thanks to goldmark.

mindmirror can
- Show the markdown as a webpage
- provide a homepage with a filetree
- inline and block math
- code
and much more

## Installation

1. Clone this repo and build it using go
```sh
git clone https://github.com/samintejas/mindmirror.git
cd mindmirror
go build mindmirror.go
```

2. replace the content folder with your notes

3. run the server with following flags

- g - generate the pages
- c - clean previously generated pages if any
- s - start the server

```sh
./mindmirror -c -g -s
```
by default the application is accesible on localhost:8080/
