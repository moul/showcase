# bolosseum
:hocho: colosseum for bots

[![Build Status](https://travis-ci.org/moul/bolosseum.svg?branch=master)](https://travis-ci.org/moul/bolosseum)

## Usage

local battle

```console
$ bolosseum run tictactoe file://./test/tictactoe file://./test/tictactoe
[0000] Initializing game "tictactoe"
[0000] Game: "tictactoe": &{{<nil> []} map["0-0":"" "0-1":"" "0-2":"" "1-0":"" "1-1":"" "2-0":"" "2-2":"" "1-2":"" "2-1":""]}
[0000] Getting bot "file://./test/tictactoe"
[0000] Registering bot "./test/tictactoe"
[0000] Getting bot "file://./test/tictactoe"
[0000] Registering bot "./test/tictactoe"
[0000] ./test/tictactoe << {"game-id":"gameid","action":"init","game":"tictactoe","players":2,"board":null,"you":null,"player-index":0}
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"init","game":"tictactoe","players":2,"board":null,"you":null,"player-index":1}
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"","1-0":"","1-1":"","1-2":"","2-0":"","2-1":"","2-2":""},"you":"X","player-index":0}
[0000] body> {"name":"","play":"2-2","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"","1-0":"","1-1":"","1-2":"","2-0":"","2-1":"","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"1-1","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"","1-0":"","1-1":"O","1-2":"","2-0":"","2-1":"","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"0-2","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"X","1-0":"","1-1":"O","1-2":"","2-0":"","2-1":"","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"1-2","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"X","1-0":"","1-1":"O","1-2":"O","2-0":"","2-1":"","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"1-0","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"","2-1":"","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"0-1","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"O","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"","2-1":"","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"2-1","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"O","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"","2-1":"X","2-2":"X"},"you":"O","player-index":1}
[0000] body> {"name":"","play":"2-0","error":null}
[0000] ./test/tictactoe << {"game-id":"gameid","action":"play-turn","game":"tictactoe","players":0,"board":{"0-0":"","0-1":"O","0-2":"X","1-0":"X","1-1":"O","1-2":"O","2-0":"O","2-1":"X","2-2":"X"},"you":"X","player-index":0}
[0000] body> {"name":"","play":"0-0","error":null}
[0000] Draw
```

remote battle

```console
$ bolosseum run tictactoe https://showcase.m.42.am/bolosseum-tictactoe https://showcase.m.42.am/bolosseum-tictactoe
[0000] Initializing game "tictactoe"
[0000] Game: "tictactoe": &{{<nil> []} map["1-1":"" "1-2":"" "2-1":"" "2-2":"" "0-0":"" "0-1":"" "0-2":"" "1-0":"" "2-0":""]}
[0000] Getting bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] Registering bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] Getting bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] Registering bot "http://showcase.m.42.am/bolosseum-tictactoe"
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] body> {"name":"moul-tictactoe","play":null,"error":null}
[0000] body> {"name":"","play":"2-2","error":null}
[0001] body> {"name":"","play":"1-1","error":null}
[0001] body> {"name":"","play":"0-2","error":null}
[0001] body> {"name":"","play":"1-2","error":null}
[0001] body> {"name":"","play":"1-0","error":null}
[0002] body> {"name":"","play":"0-1","error":null}
[0002] body> {"name":"","play":"2-1","error":null}
[0002] body> {"name":"","play":"2-0","error":null}
[0002] body> {"name":"","play":"0-0","error":null}
[0002] Draw
```

stupid bots

```console
$ bolosseum run tictactoe stupid://tictactoe stupid://tictactoe
WARN[0000] Initializing game "tictactoe"
WARN[0000] Game: "tictactoe": &{{<nil> []} map["1-2":"" "2-1":"" "0-1":"" "1-0":"" "1-1":"" "2-2":"" "0-0":"" "0-2":"" "2-0":""]}
WARN[0000] Getting bot "stupid://tictactoe"
WARN[0000] Getting stupid IA "tictactoe"
WARN[0000] Registering bot "tictactoe"
WARN[0000] Getting bot "stupid://tictactoe"
WARN[0000] Getting stupid IA "tictactoe"
WARN[0000] Registering bot "tictactoe"
WARN[0000] SendMessage: {gameid init tictactoe 2 <nil> <nil> 0}
WARN[0000] SendMessage: {gameid init tictactoe 2 <nil> <nil> 1}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[0-1: 1-0: 1-1: 1-2: 2-1: 0-0: 0-2: 2-0: 2-2:] X 0}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[0-0: 0-2: 2-0: 2-2: 0-1: 1-0: 1-1: 1-2: 2-1:X] O 1}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[0-1: 1-0: 1-1: 1-2: 2-1:X 0-0: 0-2:O 2-0: 2-2:] X 0}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[0-2:O 2-0: 2-2: 0-0: 1-0: 1-1: 1-2:X 2-1:X 0-1:] O 1}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[2-1:X 0-1: 1-0: 1-1:O 1-2:X 0-0: 0-2:O 2-0: 2-2:] X 0}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[1-1:O 1-2:X 2-1:X 0-1:X 1-0: 2-0: 2-2: 0-0: 0-2:O] O 1}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[1-1:O 1-2:X 2-1:X 0-1:X 1-0:O 2-0: 2-2: 0-0: 0-2:O] X 0}
WARN[0000] SendMessage: {gameid play-turn tictactoe 0 map[1-1:O 1-2:X 2-1:X 0-1:X 1-0:O 2-0: 2-2: 0-0:X 0-2:O] O 1}
WARN[0000] Player 1 (Tictactoe StupidIA) won
```

### Supported schemes:

* `file://`: execute local script by passing the json as `argv[1]`
* `http://` or `http+post://`: execute a POST request on a remote http server with the json as body
* `https://` or `https+post://`: same as above with SSL
* `http+get://`: execute a GET request on a remote http server with the json as *message* query (`?message={"blah":...}`)
* `https+get://`: same as above with SSL

## Bots examples

* [moul/tictactoe](https://github.com/moul/tictactoe/blob/master/cmd/tictactoe-bolosseum/main.go) (Golang)
* [moul/showcase](https://github.com/moul/showcase/blob/master/bolosseum.go) (Golang)
* [gnieark/tictactoe](https://github.com/gnieark/tictactoeChallenge/blob/master/tictactoeJSON.php) (PHP)
* [gnieark/ias](https://github.com/gnieark/IAS) (PHP)
* [gnieark/botsarena](https://github.com/gnieark/botsArena) (PHP)

## Install

```console
$ go get github.com/moul/bolosseum/cmd/bolosseum
```

## License

MIT
