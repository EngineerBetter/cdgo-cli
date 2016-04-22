source /usr/local/opt/chruby/share/chruby/chruby.sh
export MAVEN_HOME=/usr/local/Cellar/maven/3.3.9/
export ANDROID_HOME=~/devtools/android-sdk-macosx
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin:~/devtools:~/scripts
eval "$(direnv hook bash)"

# Git Radar
export PS1="\W\$(git-radar --bash --fetch) \$ "
