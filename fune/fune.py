vis = {}

state = [0,0,0,0,0,0]
fune = 0
res = None

def dfs(state, fune, history):
    print(state,fune,history)
    global res
    # 失败条件
    if state[0] != state[1]:
        if state[0] == state[1] or state[2] == state[1] or state[4] == state[1]:
            print("fail")
            return
    if state[2] != state[3]:
        if state[0] == state[3] or state[2] == state[3] or state[4] == state[3]:
            print("fail")
            return
    if state[4] != state[5]:
        if state[0] == state[5] or state[2] == state[5] or state[4] == state[5]:
            print("fail")
            return
    # 成功条件
    if state == [1,1,1,1,1,1]:
        print("SUCCESS!")
        res = history
        return     
    if vis.get((tuple(state),fune)) is not None:
        return
    vis[(tuple(state),fune)] = True
    # 转移
    if fune == 0 or fune == 1:
        for i in range(6):
            if state[i] == fune:
                if i not in [0,1,2,4]:
                    continue
                nstate = state.copy()
                nstate[i] = 1-fune
                nhistory = history.copy()
                nhistory.append(nstate)
                dfs(nstate, 1-fune, nhistory)
        for i in range(6):
            for j in range(i):
                if state[i] == fune and state[j] == fune:
                    if i not in [0,1,2,4] and j not in [0,1,2,4]:
                        continue
                    nstate = state.copy()
                    nstate[i] = 1-fune
                    nstate[j] = 1-fune
                    nhistory = history.copy()
                    nhistory.append(nstate)
                    dfs(nstate, 1-fune, nhistory)

dfs(state,fune,[])

print(res)

