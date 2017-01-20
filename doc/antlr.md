# ANTLR

## Install

### Download Jar

The instruction assume you have `~/app/bin` in your `PATH`

- from http://www.antlr.org/download or use the jar in the [third_party](../third_party) directory
- put the jar in i.e. `~/app/`, you have `~/app/antlr-4.6-complete.jar`
- in your `~/.bashrc` add the following

````
export CLASSPATH=".:$HOME/app/antlr-4.6-complete.jar:$CLASSPATH"
alias antlr4='java -Xmx500M org.antlr.v4.Tool'
alias grun='java org.antlr.v4.gui.TestRig'
````

- run `antlr4` you should see something like the following

````
$ antlr4                                    
ANTLR Parser Generator  Version 4.6
 -o ___              specify output directory where all output is generated
 -lib ___            specify location of grammars, tokens files
...
````

### Build ANTLR from source

- `git clone git@github.com:antlr/antlr4.git`
- install [maven](https://maven.apache.org/download.cgi)
- `mvn -DskipTests install`
- jars are in `tool/target`
