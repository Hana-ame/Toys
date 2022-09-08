import logging
import threading
import time
import traceback

import json

import pickle

import base64
import urllib
import urllib.request
import urllib.parse


class Query:
    # URI = scheme ":" ["//" authority] path ["?" query] ["#" fragment]
    # https://en.wikipedia.org/wiki/URL
    def __init__(self):
        self.m = {}
    def __setitem__(self, key, value):
        # 方括号运算符重载
        # https://blog.csdn.net/gavin_john/article/details/50717695
        self.m[key] = value
    def __iter__(self):
        # 迭代器重载
        # https://blog.csdn.net/qq_21594155/article/details/103210877
        return iter(self.m)
    def set(self, key, value):
        return self.__setitem__(key,value)
    def urlencode(self) -> str:
        return urllib.parse.urlencode(self.m)
    def __str__(self) -> str:
        return self.urlencode()

class Account:
    all_accts = []

    headers = {
        "accept"            :   "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
        "accept-language"   :   "zh-CN,zh;q=0.9",
        "cache-control"     :   "cache-control",
        "User-Agent"        :   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36 Edg/105.0.1343.27", 
    }
    proxy_support = urllib.request.ProxyHandler({'http' : 'http://localhost:10809/', 
                                             'https': 'http://localhost:10809/'})
    opener = urllib.request.build_opener(proxy_support)
    
    def __init__(self, host, acct, proxy=False):
        self.host = host
        self.acct = acct
        self.id = -1
        self.statuses = {}
        self.maxid=-1
        self.minid=-1
        self.proxy=proxy
        self.bio = None
        Account.all_accts.append(self)
    
    def get_id(self, id=-1):
        q = Query()
        if id == -1:
            q['acct'] = self.acct
        else:
            self.id = id
            return 


        req = urllib.request.Request(
            f"https://{self.host}/api/v1/accounts/lookup?{q}",
            headers = Account.headers,
        )
        r = b'{}'
        try:
            if self.proxy:
                with Account.opener.open(req, timeout=5) as resp:
                    r = resp.read()
            else:
                with urllib.request.urlopen(req,timeout=5) as resp:
                    r = resp.read()
        except Exception:
            print(traceback.format_exc())

        j = json.loads(r)
        self.bio = j
        print(r)
        try:
            # print(self.id)
            self.id = int(j['id'])
            # print(self.id)
        except Exception:
            self.id = -1

        if self.id == -1:
            print("请求失败。再来一次")
            print(f"https://{self.host}/api/v1/accounts/lookup?{q}")
            self.get_id()
    
    def request(self, hi=-1, lo=0):
        print(self)
        q = Query()
        if hi != -1:
            q['max_id'] = hi
        req = urllib.request.Request(
            f"https://{self.host}/api/v1/accounts/{self.id}/statuses?{q}",
            headers = Account.headers,
        )
        r = b'[]'
        try:
            if self.proxy:
                with Account.opener.open(req, timeout=15) as resp:
                    r = resp.read()
            else:
                with urllib.request.urlopen(req,timeout=15) as resp:
                    r = resp.read()
        except Exception:
            print(traceback.format_exc())
            self.request(hi,lo)
            return
        
        j = json.loads(r)
        if not j: # when empty
            return
        
        for item in j:
            self.statuses[int(item['id'])] = item

        self.minid = min(self.statuses)
        self.maxid = max(self.statuses)

        if self.minid < lo:
            return
                    
        self.request(hi=self.minid,lo=lo)
    def __str__(self) -> str:
        return f"""Account(
  https://{self.host}/web/accounts/{self.id}
  max_id : {self.maxid}
  min_id : {self.minid}
  length : {len(self.statuses)}
)"""
    def dump(self):
        with open(f"@{self.acct}@{self.host}.pickle", 'wb') as f:
            pickle.dump(self, f)
    @staticmethod
    def load(path):
        with open(path, 'rb') as f:
            return pickle.load(f)


class Pool:
    def __init__(self):
        self._pool = []
    def add(self, target, args=()):
        x = threading.Thread(target=target, args=args)
        self._pool.append(x)
    def start_all(self):
        for x in self._pool:
            x.start()
    def join_all(self):
        for x in self._pool:
            x.join()


class AcctThread(threading.Thread):
    def __init__(self, acct_para):
        threading.Thread.__init__(self)
        self.para = acct_para
    def run(self):
        a = Account(*self.para)
        a.get_id(-1)
        # a.get_id(60482)
        a.request()
        # a.dump()
        with open(f"@{a.acct}@{a.host}.pickle", 'wb') as f:
            pickle.dump(a, f)

acct_list = [
    # ('apihost':string,'id':string,isProxy:bool),
]

for acct in acct_list:
    a = AcctThread(acct)
    a.start()