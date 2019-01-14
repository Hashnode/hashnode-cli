
# Hashnode CLI

Wanted to get [hashnode](https://hashnode.com) on your terminal, Ahoy! now you can.

Here is a quick sample

```
hashnode posts --trending

(0) [Look Inside] There are two kinds of developers in the world. Which one are you?
    A fun question - Which one are you?
(1) What do you do when you have no work at office?
    There's always a phase where we have no work assigned to us at the office. How do you best utilise the time?
(2) What should 'good communication skills' mean for a developer?
    We often hear that apart from technical skills companies also look for 'good communication skills' in a software developer. In your opinion, what are the attributes of someone h
(3) What are some good podcasts for developers?
    I personally listen to the followingCommand Line Heroes (by RedHat)Software Engineering DailyIRL (by Mozilla)Hanselminutes (by Scott Hanselman)Masters of scale with Reid Hoffman
(4) Need advice: How do you write tests?
    We've our Rest API server built on NodeJS and we've started to implement "test".Our app consists of a lot of DB calls (mostly using ORM), API calls to FB and other 3rd party ser
(q) Quit
    Press to exit
```
## Installation
#### Linux
    curl -L https://github.com/Hashnode/hashnode-cli/releases/download/v0.1.5/hashnode-linux-amd64.tar.gz -o hashnode.tar.gz

```
tar xvf hashnode.tar.gz
chmod +x hashnode
sudo mv ./hashnode /usr/local/bin/hashnode
```
#### MacOS
`brew install hashnode/tap/hashnode`
    
### Windows:

Download from [GitHub](https://github.com/Hashnode/hashnode-cli/releases) and add the binary to your PATH.

### Go
Installing using go get pulls from the master branch with the latest development changes.

    go get -u github.com/hashnode/hashnode-cli
