import time
from functools import wraps

def timethis(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        startprocess = time.process_time()
        r = func(*args, **kwargs)
        end = time.perf_counter()
        endprocess = time.process_time()
        print('{}.{} : {}'.format(func.__module__, func.__name__, end - start))
        print('{}.{} : {}'.format(func.__module__, func.__name__, endprocess - startprocess))
        return r
    return wrapper