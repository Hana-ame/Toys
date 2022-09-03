import cv2
import numpy as np
import mybezier
import time

from timethis import *

scale = 120
img = np.zeros((scale*5,scale*5,3),dtype=np.uint8)

# __pts = np.array([[2,2,4,2,4,4,3,3,2,4]])
# __pts *= 200
# mybezier.polysmooth(img, 0.75, 5, __pts)

# __pts = np.array([[2,2,4,2,4,4,3,3,2,4,0,0,2,2]])
# __pts *= 200
# mybezier.polysmooth(img, 0.75, 7, __pts+100)


# __pts = np.array([0,0,2,1,1,2,0,0]) * scale
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,8)
# __pts = -__pts
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,8)
# __pts = np.array([0,0,2,-1,1,-2,0,0])*scale
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,8)
# __pts = -__pts
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,8)

# __pts = np.array([0,0,2,-0.75,2,0.75,0,0]) * scale 
# __pts +=  np.random.rand(*__pts.shape) * scale*0.2
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,18)

# __pts = np.array([0,0,2,-0.75,2,0.75,0,0]) * scale 
# __pts = -__pts
# __pts +=  np.random.rand(*__pts.shape) * scale*0.2
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,18)

# __pts = np.array([0,0,-0.75,2,0.75,2,0,0]) * scale
# __pts +=  np.random.rand(*__pts.shape) * scale*0.2
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,18)

# __pts = np.array([0,0,-0.75,2,0.75,2,0,0]) * scale
# __pts = -__pts
# __pts +=  np.random.rand(*__pts.shape) * scale*0.2
# mybezier.polysmooth(img, 0.75, 4, __pts+scale*2.75,18)

print(time.time_ns()%(2**32-1))
np.random.seed(time.time_ns()%(2**32-1))
__pts = np.array([0,0,2,-.75,2,.75,0,0,-.75,-2,.75,-2,0,0,-2,.75,-2,-.75,0,0,.75,2,-.75,2,0,0])*scale
__pts +=  np.random.rand(*__pts.shape) * scale*0.1
# __pts *= np.array([0,0,1,1,1,1,0,0,1,1,1,1,1,1,1,1,1,1,1,1,0,0])
print(__pts.reshape(-1,2))
mybezier.polysmooth(img, 0.8, len(__pts)//2, __pts+scale*2.5,18)

def test():
    a = 0
    for i in range(99999999):
        a += i
    return a


def main():
    for i in range(20000):
        time.sleep(0.001)
        img = np.zeros((scale*5,scale*5,3),dtype=np.uint8)
        __pts = np.array([0,0,0,0,2,-.75,2,.75,0,0,-.75,-2,.75,-2,0,0,-2,.75,-2,-.75,0,0,.75,2,-.75,2,0,0,0,0])*scale
        __pts +=  np.random.rand(*__pts.shape) * scale*0.1
        # print(__pts.reshape(-1,2))
        mybezier.polysmooth(img, 0.8, len(__pts)//2, __pts+scale*2.5,18)
        cv2.imshow("",255-img)
        if cv2.waitKey(1) == 27:
            cv2.destroyAllWindows()
            exit(0)

if __name__ == '__main__':
    main()