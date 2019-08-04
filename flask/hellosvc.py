from flask import Flask, request

app = Flask(__name__)

@app.route('/hello')
def hello():
    if 'name' in request.args:
        name = request.args['name']
        return 'hello {} {} {}'.format(name, name, name)
    return "RIP NO NAME", 400
