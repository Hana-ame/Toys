import base64
import webbrowser

arr = []


host = 'http://127.111.0.1:8080'
pref = f'{host}/exec/'

while True:
    x = input()
    if x:        
        encoded = base64.urlsafe_b64encode(x.encode())
        arr.append(encoded.decode())
        print(arr)
        print(' '.join(arr))
    else:
        print(' '.join(arr))        
        for i in arr:
            print(base64.urlsafe_b64decode(i),end='  ')
        print(f'{pref}{"%20".join(arr)}')
        webbrowser.open(f'{pref}{"%20".join(arr)}')
        print('clear!')
        arr = []



# bmV0c3RhdA== LWFubw==