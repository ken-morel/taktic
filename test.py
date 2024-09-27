def do_things():
    def foo():
        nonlocal bar
        bar += 1
        print(bar)
        return locals()
    bar = 0
    return locals()

l = do_things()
print(l, flush=True)
n = l['foo']()
print(l, n, flush=True)
