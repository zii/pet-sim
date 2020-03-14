#coding: utf-8
import requests

def loadimgnos():
    nos = []
    lines = open("../bin/data/enemybase.txt", "rb").readlines()
    for line in lines:
        line = line.strip()
        line = line.decode("utf-8")
        #print(line)
        rows = line.split(",")
        no = int(rows[36])
        nos.append(no)
    return nos

def saveimg(no):
    url = "http://t.shiqi.me/cw/{}.gif".format(no)
    r = requests.get(url, stream=True)
    if r.status_code != 200:
        print("not found:", no, r.status_code, r.content)
        return
    with open("../bin/static/pet/{}.gif".format(no), "wb") as f:
        for chunk in r:
            f.write(chunk)

if __name__ == "__main__":
    nos = loadimgnos()
    for no in nos:
        saveimg(no)