# %% import 
import cv2
import numpy as np
from numpy.lib.shape_base import split
from scipy.special import comb

shallcache = True
weights = {}

def genweights(n,split=100):
    w = np.ones((split+1,n+1))
    t = np.linspace(0, 1, split+1)
    tprod = np.ones(t.shape)
    for i in range(n+1):
        w[:,i] *= tprod
        tprod *= t
    tprod = np.ones(t.shape)
    for i in range(n+1):
        w[:,n-i] *= tprod
        tprod *= 1-t
    b = comb(n, np.arange(0,n+1,1,dtype=np.uint32))
    for i in range(n+1):
        w[:,i] *= b[i]
    return w

def bezier(pts,split=100):
    global weights
    n = len(pts) - 1
    if weights.get((n,split)) is None and shallcache:
        weights[(n,split)] = genweights(n,split)
    return np.dot(weights[(n,split)], pts).astype(np.int32)
    
# def bezier(pts,split=100):
#     global weights    
#     n = len(pts) - 1
#     weights[1] = genWeights(n,split)
#     return np.dot(weights[1], pts).astype(np.int32)

def polysmooth(img, k, n, pts, thickness=0):
    _pts = pts.reshape(-1,2)
    pts = np.zeros( (_pts.shape[0]+2, _pts.shape[1]) , dtype=np.int32 )
    pts[1:-1] = _pts
    pts[0] = pts[1]
    pts[-1] = pts[-2]

    线段length = np.linalg.norm(pts[:-1] - pts[1:],axis = 1)
    中点list = (pts[:-1] + pts[1:]) / 2
    for i in range(n):
        if i == 0: continue
        pctrl=np.zeros((4,2))
        pctrl[0] = pts[i]
        pctrl[3] = pts[i+1]
        pctrl[1] = pctrl[0] + (中点list[i] - 中点list[i-1]) * 线段length[i]/(线段length[i-1]+线段length[i]) * k
        pctrl[2] = pctrl[3] + (中点list[i] - 中点list[i+1]) * 线段length[i]/(线段length[i+1]+线段length[i]) * k
        # img = cv2.polylines(img, [bezier(pctrl)], False, (255,255,255), thickness=thickness)
        cv2.polylines(img, [bezier(pctrl)], False, (255,255,255), thickness=thickness)
    

if __name__ == '__main__' :
    img = np.zeros((240*5,240*5,3),dtype=np.uint8)
    __pts = np.array([[2,2,4,2,4,4,3,3,2,4]])
    __pts *= 200
    img = cv2.polylines(img, [__pts.reshape(-1,2).astype(np.int32)], False, (255,255,0), 0)
    k = 0.5

    polysmooth(img, k, 5, __pts)

    # %% 新建画板
    img = np.zeros((240*5,240*5,3),dtype=np.uint8)
    # 设置控制点
    pts = np.array([[0,0],[240,0],[240,240],[0,240]])
    # 插值
    line = bezier(pts)
    # 绘图
    arr = cv2.polylines(img, [line],False, (255,255,255), 0)

    # cv2.line(arr, (0,0), (240,2400), (255,255,255))

    #   显示图片
    cv2.imshow("",img)
    if cv2.waitKey(0):
        cv2.destroyAllWindows()
    # %% 测试

    pts = np.array([[0,0],[240,0],[240,240]])

    n = len(pts) - 1
    split = 300 
    w = np.ones((split+1,n+1))
    t = np.linspace(0, 1, split+1)

    tprod = t
    for i in range(n+1):
        if i == 0: continue
        w[:,i] *= tprod
        tprod *= t
    # print(w)

    talt = 1-t
    tprod = talt
    for i in range(n+1):
        if i == 0: continue
        w[:,n-i] *= tprod
        tprod *= talt
    # print(w)

    # b = comb(n, np.arange(0,n+1,1,dtype=np.uint32))
    # for i in range(n+1):
    #     w[:,i] *= b[i]
    # print(w)

    line = np.dot(w, pts)
    print(line.astype(np.uint32))
    arr = cv2.polylines(arr, [line.astype(np.int32)],False, (255,255,255), 0)
    # print(arr)
    cv2.imshow("",arr)
    if cv2.waitKey(0):
        cv2.destroyAllWindows()
    
    def getWeight(n,splits=100):
        t = np.linspace(0, 1, splits+1)
        w = np.zeros(n,2)
        w[:,0] = np.power()
        w[:,1] = w[:,0]
        return w

    def genPoint(n, t, pts):
        powerarr = np.arange(0,n,1)
        temparr = np.zeros(pts.shape, dtype=np.uint32)
        temparr[:,0] = np.power(1-t,n-powerarr) * np.power(t,powerarr) * pts[:,0] + 0.5
        temparr[:,1] = np.power(1-t,n-powerarr) * np.power(t,powerarr) * pts[:,1] + 0.5
        np.sum(temparr, axis=0)
    def genLine(n, pts, split=100):
        pass
    ### 参考资料 ###
    # https://zhuanlan.zhihu.com/p/409585038 这是个坏东西
    # https://juejin.cn/post/6952831906060697636
    # https://qiita.com/mo256man/items/122228f3ab7101f24059#opencv%E3%81%A7%E8%87%AA%E7%94%B1%E3%81%AB%E6%8F%8F%E3%81%8F
    exit(0)



    img = np.zeros((240*5,240*5,3),dtype=np.uint8)
    __pts = np.array([[2,2,4,2,4,4,3,3,2,4]])
    __pts *= 200
    img = cv2.polylines(img, [__pts.reshape(-1,2).astype(np.int32)], False, (255,255,0), 0)
    k = 0.5

    _pts = __pts.reshape(-1,2)
    thickness = 0
    length = 4
    pts = np.zeros( (_pts.shape[0]+2, _pts.shape[1]) , dtype=np.int32 )
    pts[1:-1] = _pts
    pts[0] = pts[1]
    pts[-1] = pts[-2]

    线段length = np.linalg.norm(pts[:-1] - pts[1:],axis = 1)
    中点list = (pts[:-1] + pts[1:]) / 2

    print(pts,pts[:-1] - pts[1:],中点list)



    for i in range(len(pts)-2):
        if i == 0: continue
        pctrl=np.zeros((4,2))
        pctrl[0] = pts[i]
        pctrl[3] = pts[i+1]
        pctrl[1] = pctrl[0] + (中点list[i] - 中点list[i-1]) * 线段length[i]/(线段length[i-1]+线段length[i]) * k
        pctrl[2] = pctrl[3] + (中点list[i] - 中点list[i+1]) * 线段length[i]/(线段length[i+1]+线段length[i]) * k
        # print(pctrl, (中点list[i] - 中点list[i+1]), 线段length[i],(线段length[i+1]+线段length[i]), k)
        img = cv2.polylines(img, [bezier(pctrl)], False, (255,255,255), 0)
        



    pts = np.array([[25, 70], [25, 160], 
                    [110, 200], [200, 160], 
                    [200, 70], [110, 20]],
                np.int32)
    print(pts)
    print([pts])
    pts = pts.reshape((-1, 1, 2))
    print(pts)
    image = cv2.polylines(image, [pts], 
                        isClosed, color, thickness)

    print(np.array([1,2,3,4,5]) * np.array([1,2,3,4,5]))

    sep = 0.01
    t = np.arange(0,1+sep,sep)
    print(t.shape)
    t = t.reshape(-2,1)
    print(t.shape)

    # print(comb(n, i) * t**i * (1 - t)**(n-i))
    print('----')
    n = 10
    k = np.arange(0,n+1,1,dtype=np.uint32)
    print(k)
    b = comb(n,k)
    print(b)
    # def genTriangle(pointPair):
    #     n = pointPair.shape[0]-1
    #     k = np.arange(0,n+1,1)
    #     b = comb(n,k)
    #     return b
    # def genWeight(pointPair,split=100):
    #     n = pointPair.shape[0]-1
    #     sep = 1/split
    #     t = np.arange(0,1+sep,sep)
    #     t = t.reshape(-1,1)
    #     oneMinusT = 1-t
    #     b = genTriangle(pointPair)
    #     weight = np.power(t,k)*np.power(oneMinusT,n-k)*b
    #     return weight
    # def getBesLine(pointPair,split=100):
    #     weight = genWeight(pointPair,split)
    #     x = np.sum(weight*pointPair[:,0],axis=1)
    #     y = np.sum(weight*pointPair[:,1],axis=1)
    #     line = np.concatenate([x.reshape(-1,1),y.reshape(-1,1)],axis=1)
    #     return line
    # getBesLine(np.array([[1000,535],[780,855],[460,215],[600,535],[1000,535]]))
    # n = 10
    # i = 1