import cv2
import matplotlib.pyplot as plt
import numpy as np

import requests
import base64
import random
import json
from hashlib import md5
import time

class Recoder:
    def __init__(self, dest):
        self.f = open(dest, 'w+', encoding = 'utf-8')
    def 记录(self, ms, count):
        h,m,s,ms = self.时间点(ms)
        msg = "{}:{:0>2d}:{:0>2d}.{:0>3d},{}\n".format(h,m,s,ms,count)
        self.f.write(msg)
    def 时间点(self, milliseconds):
        ms = milliseconds % 1000
        ms = ms // 1
        s  = milliseconds // 1000
        m  = s // 60
        s  = s % 60
        h  = m // 60
        m  = m % 60    
        return int(h),int(m),int(s),int(ms)
    def close(self):
        self.f.close()

class BaiduTr:
    def __init__(self):
        self.appid = '20210419000788789'
        self.appkey = 'I8ScDR7t8OYU_dpumgyH'
        self.from_lang = 'jp'
        self.to_lang =  'zh'
        endpoint = 'https://api.fanyi.baidu.com'
        path = '/api/trans/vip/translate'
        self.url = endpoint + path
        self.salt = 0
        self.headers = {'Content-Type': 'application/x-www-form-urlencoded'}
    def make_md5(self, s, encoding='utf-8'):
        return md5(s.encode(encoding)).hexdigest()
    def tl(self, query):
        sign = self.make_md5(self.appid + query + str(self.salt) + self.appkey)
        payload = {'appid': self.appid, 'q': query, 'from': self.from_lang, 'to': self.to_lang, 'salt': self.salt, 'sign': sign}
        r = requests.post(self.url, params=payload, headers=self.headers)
        result = r.json()
        time.sleep(1)
        try:
            return result["trans_result"][0]["dst"]
        except:
            print(result)
            return "err"


baidu = BaiduTr()

# encoding:utf-8
import requests 

# client_id 为官网获取的AK， client_secret 为官网获取的SK
host = 'https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=w0GeSlSN6m6L8TV8MZG9jlW8&client_secret=Fanq3NtGNyB6ERTxRn1aESUrvvU1qVNc'
response = requests.get(host)
if response:
    # print(response.json())
    pass
else:
    print("https://cloud.baidu.com/doc/OCR/s/Ck3h7y2ia")
# 获取Access Token
access_token = response.json()['access_token']

def 机翻(图):
    # 保存
    cv2.imwrite('ocr.png',subFrame)
    request_url = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
    # request_url = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
    # 二进制方式打开图片文件
    f = open('ocr.png', 'rb')
    img = base64.b64encode(f.read())

    params = {"image":img, "language_type": "JAP", "probability": "true"}
    # access_token = access_token
    request_url = request_url + "?access_token=" + access_token
    headers = {'content-type': 'application/x-www-form-urlencoded'}
    response = requests.post(request_url, data=params, headers=headers)
    # if response:
        # return baidu.tl(response.json()["words_result"][0]['words'])
        # print (response.json())
    allresult = ''
    for i in response.json()["words_result"]:
        # print(i['words'])
        allresult += baidu.tl(i['words'])
    return allresult


k_h=-np.array([[0,0,-1,0,0],[0,-1,-2,-1,0],[0,0,0,0,0],[0,1,2,1,0],[0,0,1,0,0]])/5
k_v=-np.array([[0,0,0,0,0],[0,-1,0,1,0],[-1,-2,0,2,1],[0,-1,0,1,0],[0,0,0,0,0]])/5
# k_h=-np.array([[-1,-2,-1],[0,0,0],[1,2,1]])/4
# k_v=-np.array([[-1,0,1],[-2,0,2],[-1,0,1]])/4
def 卷积and计数(img, c1, c2 = np.array([255,255,255]), threshold = 128):
    dst   = cv2.filter2D(img, -1 ,  k_h)
    ndst  = c2-c1-dst
    mask  = np.sum(np.abs(ndst),axis=2) < threshold
    dst   = cv2.filter2D(img, -1 , -k_h)
    ndst  = c2-c1-dst
    mask |= np.sum(np.abs(ndst),axis=2) < threshold
    dst   = cv2.filter2D(img, -1 ,  k_v)
    ndst  = c2-c1-dst
    mask |= np.sum(np.abs(ndst),axis=2) < threshold
    dst   = cv2.filter2D(img, -1 , -k_v)
    ndst  = c2-c1-dst
    mask |= np.sum(np.abs(ndst),axis=2) < threshold
    return mask.sum()

# videoPath = "079.mp4"
videoPath = R"C:\Users\lumin\Downloads\[survive more] 家、行って寝取っていいですか？ ～上京したて田舎娘編～ The Motion Anime\AMCP-075_ieitte_MainMovie.mp4.mp4"
# colorName = "黑"
# 字色 = np.array([0,0,0])
colorName = "粉"
字色 = np.array([206,0,255])
# colorName = "青"
# 字色 = np.array([255,183,0])
# colorName = "红"
# 字色 = np.array([0,0,226])
# colorName = "绿"
# 字色 = np.array([25,182,0])
# colorName = "紫"
# 字色 = np.array([226,0,148])
# colorName = "黄"
# 字色 = np.array([0,130,255])
# colorName = "蓝"
# 字色 = np.array([255,61,0])
# colorName = "水"
# 字色 = np.array([192,92,33])
# colorName = "橙"
# 字色 = np.array([6,82,246])


cameraCapture = cv2.VideoCapture(videoPath)
videoPath = "075"
r1 =Recoder(videoPath + colorName + "1.csv")
r2 =Recoder(videoPath + colorName + "2.csv")

success = True
上次count = 0
# 抓轴过程
while success:
    # if cv2.waitKey(1) == 27:
        # break
    success, frame = cameraCapture.read()
    # cv2.imshow('Test camera', frame)
    milliseconds = cameraCapture.get(cv2.CAP_PROP_POS_MSEC)

    # 字幕区
    subFrame = frame[590:680,:]

    # 求差
    count = 卷积and计数(subFrame.copy(), 字色)

    r1.记录(milliseconds, count)
    if (np.abs(上次count - count) > 32):
        # r2.记录(milliseconds, '{},{}'.format(机翻(subFrame),count))
        cv2.imwrite('{:0>8d}.png'.format(int(milliseconds)),subFrame)
        r2.记录(milliseconds, count)
    上次count = count 
cv2.destroyAllWindows()
r1.close()
r2.close()

print("我准备好了")