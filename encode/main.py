import sys
import json
import base64

d = json.load(sys.stdin)
d["encoded"] = base64.b64encode(d["phrase"].encode('utf-8')).decode('utf-8')
json.dump(d, sys.stdout, ensure_ascii=False)
