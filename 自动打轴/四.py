import cv2
import matplotlib.pyplot as plt
import numpy as np


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

k_h=-np.array([[-1,-2,-1],[0,0,0],[1,2,1]])/4
k_v=-np.array([[-1,0,1],[-2,0,2],[-1,0,1]])/4
def 卷积and计数(img, c1, c2 = np.array([255,255,255]), threshold = 72):
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

cameraCapture = cv2.VideoCapture('./079.mp4')
r1 =Recoder("r41.csv")
r2 =Recoder("r42.csv")

success = True
上次count = 0
字色 = np.array([0,0,229])
# 抓轴过程
while success:
    success, frame = cameraCapture.read()
    milliseconds = cameraCapture.get(cv2.CAP_PROP_POS_MSEC)

    # 字幕区
    subFrame = frame[600:720,:]

    # 求差
    count = 卷积and计数(subFrame.copy(), 字色)

    r1.记录(milliseconds, count)
    if (np.abs(上次count - count) > 192):
        r2.记录(milliseconds, count)
    # cv2.imshow('Test camera', canvas)
    上次count = count 
cv2.destroyAllWindows()
r1.close()
r2.close()