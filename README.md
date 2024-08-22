## Windows Setup

Download apps:

```
scoop install scoop
scoop install ojdkbuild-jre
scoop install gnuplot
```

Download and extract [Maelstrom](https://github.com/jepsen-io/maelstrom/releases/tag/v0.2.3). There would a `lib` folder inside it. Rename it to `maelstrom-lib` and copy and paste it into the project root.

## Run

To run & verify, first build the project using `go build .` and then run it like:

```pwsh
cd maelstrom-lib
sudo java -jar .\maelstrom.jar test -w echo --bin ..\gossip-glomers.exe --node-count 1 --time-limit 10
# or
sudo java -jar .\maelstrom.jar test -w unique-ids --bin ..\gossip-glomers.exe --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
# etc.
```
