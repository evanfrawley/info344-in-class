from flask import Flask, request
from datetime import datetime

app = Flask(__name__)

@app.route('/time')
def time():
    return "The time serveice {} says the current time is: {}".format(request.host, str(datetime.now().time()))
