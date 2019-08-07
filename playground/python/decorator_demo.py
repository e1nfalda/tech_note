# encoding: utf8

"""
    1. decorator修饰的方法会在import时就执行decorator(func),class 也会实例化。
"""

import functools


def decorator(func):
    print "init decorator"

    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print "wrapper start"
        func(*args, **kwargs)
        print "wrapper end"
    return wrapper


@decorator
def dec_func(string):
    print "run_func: ", string


if __name__ == "__main__":
    dec_func("hello")
    print dec_func.__name__
