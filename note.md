```python
class Config:
    def __getitem__(self, ro):
        print(ro)
        return ro

c = Config()
# <__main__.Config object at 0x0000020B572CEE50>
c[1]
# 1
c[1:2]
# slice(1, 2, None)
c[1:]
# slice(1, None, None)
c[1:-1]
# slice(1, -1, None)
c[:-1]
# slice(None, -1, None)
c[1,2,3]
# (1, 2, 3)
c[1:2,3]
# (slice(1, 2, None), 3)
c[1:2:3]
# slice(1, 2, 3)
```