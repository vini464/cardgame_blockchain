import json
import eth_keyfile
path = input("caminho: ")
kf = json.load(open(path))
print(eth_keyfile.decode_keyfile_json(kf, b"YOUR_PASSWORD").hex())
