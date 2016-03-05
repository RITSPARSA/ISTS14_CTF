from flask import Flask, render_template, redirect, request, url_for
import random

app = Flask(__name__)

@app.route("/", methods=["GET", "POST"])
def home():
	if request.method == "GET":
		return render_template("index.html")
	else:
		low = int(request.form['lower'])
		high = int(request.form['upper'])

		num = random.randint(low, high)
		return render_template("post.html", number=num)

if __name__ == "__main__":
	#application.run()
	app.run(host="0.0.0.0", port=8080, debug=True)

