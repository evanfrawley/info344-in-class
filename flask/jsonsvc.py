from flask import Flask, request, Response, json

app = Flask(__name__)

@app.route("/json")
def json_handler():
    data = {
        "foo": "bar",
        "Hello": "World!"
    }

    js = json.dumps(data)
    resp = Response(js, content_type="application/json")
    resp.headers["Access-Control-Allow-Origin"] = "*"

    return resp
