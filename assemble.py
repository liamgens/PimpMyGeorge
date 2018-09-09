#!/usr/bin/env python3
import base64
import json


def encode_the_whore(path):
    with open(path, "rb") as image_file:
        encoded_string = base64.b64encode(image_file.read())
    return str(encoded_string)[2:-1]


def write_to_json(dictionary):
    filename = './pmg/georges.json' if gb == 'g' else './pmg/blang.json'
    with open(filename, 'r') as f:
        json_in = json.load(f)

    if gb == 'g':
        json_in["georges"].append(dictionary)
    else:
        json_in["blang"].append(dictionary)
    with open(filename, 'w') as f:
        json.dump(json_in, f, indent=2)


def g_json(name, path):
    jsonn = {
        "name": name,
        "image": encode_the_whore(path),
        "pieces": {
            "head": [200, -3, 315, 65],
            "face": [210, 73, 290, 90],
            "neck": [217, 164, 300, 180],
            "torso": [165, 165, 375, 325],
            "legs": [0, 0, 0, 0],
            "hands": [178, 352, 270, 430],
            "feet": [0, 0, 0, 0]
        }
    }

    write_to_json(jsonn)


def b_json(noun, adj, location, path):
    jsonn = {
        "noun": noun,
        "adj": adj,
        "location": location,
        "image": encode_the_whore(path)
    }

    write_to_json(jsonn)


gb = input("g for george or b for blang: ")
filename = input("filename: ")

if gb == 'g':
    name = input("name: ")
    g_json(name, filename)

elif gb == 'b':
    noun = input("noun: ")
    adj = input("adj: ")
    location = input("location: ")
    b_json(noun, adj, location, filename)

else:
    print("Sugondese NUTZ")
