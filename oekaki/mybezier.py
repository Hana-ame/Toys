# %% import 
import cv2
import numpy as np
# from numpy.lib.shape_base import split
from scipy.special import comb

weights = {}

def genWeights(n,split=100):
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
    if weights.get((n,split)) is None:
        weights[(n,split)] = genWeights(n,split)
    return np.dot(weights[(n,split)], pts).astype(np.int32)
    
# def bezier(pts,split=100):
#     global weights    
#     n = len(pts) - 1
#     weights[1] = genWeights(n,split)
#     return np.dot(weights[1], pts).astype(np.int32)

def polysmooth(img, k, n, pts, thickness=0):
    # _pts = pts.reshape(-1,2)
    # pts = np.zeros( (_pts.shape[0]+2, _pts.shape[1]) , dtype=np.int32 )
    # pts[1:-1] = _pts
    # pts[0] = pts[1]
    # pts[-1] = pts[-2]
    pts = pts.reshape(-1,2) # 修改为值

    线段length = np.linalg.norm(pts[:-1] - pts[1:],axis = 1)
    中点list = (pts[:-1] + pts[1:]) / 2
    for i in range(n-2):
        if i == 0: continue
        pctrl=np.zeros((4,2))
        pctrl[0] = pts[i]
        pctrl[3] = pts[i+1]
        pctrl[1] = pctrl[0] + (中点list[i] - 中点list[i-1]) * 线段length[i]/(线段length[i-1]+线段length[i]) * k
        pctrl[2] = pctrl[3] + (中点list[i] - 中点list[i+1]) * 线段length[i]/(线段length[i+1]+线段length[i]) * k
        # img = cv2.polylines(img, [bezier(pctrl)], False, (255,255,255), thickness=thickness)
        cv2.polylines(img, [bezier(pctrl)], False, (255,255,255), thickness=thickness)
    


