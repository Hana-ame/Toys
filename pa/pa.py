from bs4 import BeautifulSoup
import urllib3
import certifi
import os
import time
import json
import sys

http = urllib3.PoolManager(
        cert_reqs='CERT_REQUIRED',
         ca_certs=certifi.where()
        )

已读id列表 = []
# http = urllib3.ProxyManager("http://localhost:10809")



def act():
        爬目录(int(sys.argv[1]))

def 读json():
        global 已读id列表
        f = open('/home/luminovoez/已读id列表.json','r')
        data = f.read()
        f.close()
        已读id列表 += json.loads(data)

def 写json():
        data = json.dumps(已读id列表)
        f = open('/home/luminovoez/已读id列表.json','w')
        f.write(data)
        f.close()


def 爬目录(Page):
        r = http.request( 'GET','https://asiansister.com/_page{}'.format(Page) )
        soup = BeautifulSoup(r.data)
        for link in soup.find_all('a'):
                href = link.get('href')
                if(href[0:5] == 'view_'):
                        爬画廊(href)

def 爬画廊(href):
        id = href.split('_')[1] # 得到id
        r = http.request('GET','https://asiansister.com/{0}'.format(href))
        soup = BeautifulSoup(r.data)
        # 提取标题
        title = soup.find('h1')
        title = title.string
        # 提取图数
        pics = 0 # 初始化
        for img in soup.find_all('img'):
                if( img.get('onclick') ):  # 有这个属性
                   pics += 1
        pics -= 4 # 修正
        爬图整页(id,pics,title)


def 爬图整页(id, pics, title):
        global 已读id列表
        dest = "{0}-{1}".format(id, title)
        os.system("mkdir \"{0}\"".format(dest))
        i = 0
        while i <= pics:
                try:
                        pics += 爬(id, i, dest)
                        i += 1
                except:
                        i = pics + 1
        print(id,pics,title)
        file_object = open('/home/luminovoez/asiansister.txt', 'a')
        file_object.write('{0}\t{1}\t{2}\n'.format(id,pics,title))
        file_object.close()
        已读id列表.append({"id":id,"pages":pics})
        写json()

def 爬(id, i, dest):
        补正 = 1
        r = http.request('GET','https://asiansister.com/viewImg.php?code={1}&id={0}'.format(i,id))
        print('https://asiansister.com/viewImg.php?code={1}&id={0}'.format(i,id))
        soup = BeautifulSoup(r.data)
        try:
         img = soup.find("meta",property="og:image")
         if len(img['content'])<5:
             time.sleep(2)
             r = http.request('GET','https://asiansister.com/viewImg.php?code={1}&id={0}'.format(i,id))
             print('https://asiansister.com/viewImg.php?code={1}&id={0}'.format(i,id))
             soup = BeautifulSoup(r.data)
             img = soup.find("meta",property="og:image")
         if img['content'][-4:] == ".jpg":
                 补正 = 0
         elif img['content'][-4:] == "jpeg":
                 补正 = 0
         elif img['content'][-4:] == ".png":
                 补正 = 0
        except:
         print("https://asiansister.com/{0}".format(img['content']))
        # os.system("wget \"https://asiansister.com/{0}\" ;".format(img['content']))
        r = http.request( 'GET',"https://asiansister.com/{0}".format(img['content']) )
        f = open(dest + '/' + img['content'].split('/')[-1],'wb')
        f.write(r.data)
        f.close()
        return 补正




if __name__ == "__main__":
        # 爬目录(int(sys.argv[1]))
        f = open("../lixt.txt")
        s = f.read()
        s = s.replace("\n","")
        s = s.split(" ")
        for i in s:
            print(i)
            try:
                爬画廊(i)
            except:
                print(i,"!!")
        pass
#

