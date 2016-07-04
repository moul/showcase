#!/usr/bin/env python

import json
import random
import sys


class BolosseumBot():
    def handle(self, args):
        method = args['action'].replace("-", "_")
        return getattr(self, method)(args)


class CoinflipBot(BolosseumBot):
    def init(self, args):
        return {"name": "coinflip-bot"}

    def play_turn(self, args):
        if random.randint(0, 1):
            return {"play": "head"}
        else:
            return {"play": "ship"}


if __name__ == "__main__":
    bot = CoinflipBot()
    args = json.loads(sys.argv[1])
    print(bot.handle(args))
