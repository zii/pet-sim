#coding: utf-8
import requests
import os

def loadimgnos():
    nos = []
    lines = open("../bin/data/enemybase2.txt", "rb").readlines()
    for line in lines:
        line = line.strip()
        line = line.decode("utf-8")
        rows = line.split(",")
        if len(rows) < 3:
            continue

        print("rows:", len(rows), rows[35])
        no = int(rows[36])
        nos.append(no)
    return nos

def saveimg(no):
    url = "http://t.shiqi.me/cw/{}.gif".format(no)
    localpath = "../bin/static/pet/{}.gif".format(no)
    if os.path.exists(localpath):
        return 1
    r = requests.get(url, stream=True)
    if r.status_code != 200:
        return 2
    with open(localpath, "wb") as f:
        for chunk in r:
            f.write(chunk)
    return 0

if __name__ == "__main__":
    nos = loadimgnos()
    notfound = []
    for no in nos:
        r = saveimg(no)
        if r == 2:
            notfound.append(no)
    notfound = list(set(notfound))
    print("notfound:", notfound)