import json

# while True:
command = input(">")
data = []
with open(command, "r") as f:
    data = json.load(f)


def rec(data):
    for id, i in enumerate(data):
        i["o"] = False
        i["f"] = []

        if isinstance(i["n"], list):
            data[id]["n"] = rec(i["n"])
        data[id]["awcf"] = rec(i["awcf"])
    return data


data = rec(data)
with open(command, "w") as f:
    f.write(json.dumps(data))
