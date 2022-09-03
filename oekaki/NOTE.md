# todo
- ~~平滑~~

- 对曲线拟合
- 曲线归一化
- 


# a smooth line is drawed by
  
    import mybezier

    mybezier.polysmooth(
        img,            # 图片
        k,              # 平滑系数
        length,         # 控制点个数
        pts,            # 控制点
        thickness=0     # 粗细
    )

----
trembling

    __pts = np.array([0,0,2,-.75,2,.75,0,0,-.75,-2,.75,-2,0,0,-2,.75,-2,-.75,0,0,.75,2,-.75,2,0,0])*scale
    __pts +=  np.random.rand(*__pts.shape) * (scale*0.1)

----
测试文件执行时间

    from timethis import *

    @timethis
    def func():
        ...
参考：[给你的程序做性能测试](https://python3-cookbook.readthedocs.io/zh_CN/latest/c14/p13_profiling_and_timing_your_program.html)


----
这个是生成可执行文件

    pyinstaller main.py --onefile

参考：[官网](https://pyinstaller.readthedocs.io/en/stable/genindex.html)

Linux下未测试
 

