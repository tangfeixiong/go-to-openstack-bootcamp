import eventlet
from eventlet.green import urllib2
import re, json


urls = ["http://127.0.0.1:5000/v2.0/tokens", "http://127.0.0.1:5000/v3/auth/tokens", "http://python.org/images/python-logo.gif", "http://www.baidu.com"]

data2 = {"auth": {"tenantName": "admin", "passwordCredentials": {"username": "admin", "password": "changeme1122"}}}
data3 = {"auth":{"identity":{"methods":["password"],"password":{"user":{"name":"admin","domain":{"id":"default"},"password":"changeme1122"}}}}}
headers1 = {'User-Agent': 'Mozilla/5.0 (Windows NT 5.1; rv:35.0) Gecko/20100101 Firefox/35.0', 'Content-Type': 'application/json'}
headers2 = {'User-Agent': 'python-keystoneclient', 'Content-Type': 'application/json'}

def fetch(url):
    print("opening", url)
    if (re.search(r'\:5000/v2\.0', url) != None) :
        d = json.dumps(data2)
        req = urllib2.Request(url, d, headers1)
        f = urllib2.urlopen(req)
        body = f.read()
        print  '\n%(response)s\n' % {"response": body}
    elif (re.search(r'\:5000/v3/', url) != None) :
        d = json.dumps(data3)
        req = urllib2.Request(url, d, headers2)
        f = urllib2.urlopen(req)
        body = f.read()
        print  '\n%(response)s\n' % {"response": body}    
    else : 
        body = urllib2.urlopen(url).read()
    print("done with", url)
    return url, body

pool = eventlet.GreenPool(200)
for url, body in pool.imap(fetch, urls):
    print("got body from", url, "of length", len(body))