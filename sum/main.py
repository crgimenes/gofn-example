import sys
import json

d = json.load(sys.stdin)
ret = {"sum": d["factor1"] + d["factor2"]}
json.dump(ret, sys.stdout, ensure_ascii=False)
