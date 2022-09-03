
import pandas as pd 

df = pd.read_stata('test02.dta')

dic = {}
for index, row in df.iterrows():
    if dic.get(row['nkcode']) is None:
        dic[row['nkcode']] = []
    if not pd.isna(row['d01_it103']):
        dic[row['nkcode']].append(row['d01_it103'])
    else:
        # print('提示：无效的数据于{index}行:{data}'.format(index=index,data=row['d01_it103']))
        pass

# 此处，dic中存以nkcode为index，所有d01_it103（str）的列表为value

def 求平均(l):
    n = len(l)
    if n == 0:
        return "无有效数据"
    sum = 0
    for i in l:
        sum += float(i)
    return sum/n
    
for k in dic:
    # print(k, 求平均(dic[k])) 
    # print(int(k), 求平均(dic[k]))
    pass

# df['average'] = 0
# for k in dic:
#     df['average'][df['nkcode']==k] = 求平均(dic[k])
# df[['nkcode','d01_it103','average']]
# print(df[['nkcode','d01_it103','average']])

df['业界'] = df['nkil']//1000%100
print( df[ ['nkil','业界'] ].drop_duplicates(subset=['nkil']) )

