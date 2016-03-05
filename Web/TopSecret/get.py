#!/usr/bin/env python
import requests, sys, time

headers = {
	"Host" : "52.91.204.109",
	"Cache-Control" : "max-age=0",
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36",
	"Accept-Language" : "en-US,en;q=0.8",
	"Cookie" : "PHPSESSID=9dac3b45763171dc81fa3a8d14334a45d7fa63744625af07fdb458c50cae818b71f179123cdf57093cbd86daec5e2bd75829a01565d44adec81e26dd4a9b133f; User=admin"
}

def main():
	while True:
		r = requests.get("https://52.91.204.109", headers=headers, verify=False)
		print "made request with status '%d', sleeping 10" % r.status_code
		time.sleep(10);

if __name__ == "__main__":
	sys.exit(main())

