
# Hashnode CLI

Wanted to get [Hashnode](https://hashnode.com) on your terminal? Ahoy! now you can.

Here is a quick sample

```
hashnode --help
hashnode stories --hot
hashnode discussions --hot

(0) Download Monitoring: A Cross-Browser Story
    PreambleWhen you want to implement something in a cross-browser way, you are in for a ride down the bugtracker hole. After some ex
(1) How IDE is a blessing for beginners
    Quite recently I came across the post  IDE - The beginner's trap. It's a must a read, good article by YounesButI can't agree to it
(2) HTTP request with ES6 tagged templates
    ES6 template literals are one of my favorite features in es6. A few days ago I saw an open source project on github - htm, there a
(3) Benefits of Progressive Web Applications (PWAs) and How to Build One
    In this tutorial, we're going to build up the fundamentals on Progressive Web Applications (PWAs). I'll help you understand the pa
(4) The all-new Hashnode Chrome extension
    With the start of the new year, new team members, new features, and new product(s), we here at Hashnode are working hard to bring
(q) Quit
    Press to exit
```
## Installation
#### Linux
    curl -L https://github.com/Hashnode/hashnode-cli/releases/download/v0.1.7/hashnode-linux-amd64.tar.gz -o hashnode.tar.gz

```
tar xvf hashnode.tar.gz
chmod +x hashnode
sudo mv ./hashnode /usr/local/bin/hashnode
```
#### MacOS
`brew install hashnode/tap/hashnode`
    
#### Windows:

Download from [GitHub](https://github.com/Hashnode/hashnode-cli/releases) and add the binary to your PATH.

#### Go
Installing using go get pulls from the master branch with the latest development changes.

    go get -u github.com/hashnode/hashnode-cli
# Demo
[![asciicast](https://asciinema.org/a/221329.svg)](https://asciinema.org/a/221329)
