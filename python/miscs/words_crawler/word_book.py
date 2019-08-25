from bs4 import BeautifulSoup
import urllib3
import pandas as pd


def get_trans(word):

    url = 'http://dict.youdao.com/search?q=%s'%word
    response = urllib3.PoolManager().request('GET', url)
    context = response.data
    soup = BeautifulSoup(context, 'html.parser', from_encoding='utf-8')

    try:
        baav = soup.find_all(name='div',attrs={"class":"baav"})
        trans = soup.find_all(name='div',attrs={"class":"trans-container"})
        phonetic = baav[0].text.replace(' ','').replace('\n','').encode('utf-8')
        meaning = trans[0].text.replace(' ','').replace('\n','').encode('utf-8')
    except IndexError as e:
        print word, e
        return 'None','None'
    return phonetic, meaning


def text_to_excel():
    pass
    '''save to excel'''
    data = pd.read_csv(r'./mean.txt',delim_whitespace = True)
    df = pd.DataFrame(data)
    df.to_csv(r'./output.csv')


def main():
    pass
    with open(r'./words.txt','r') as f:
        a = [w.rstrip('\n') for w in f.readlines() if w[-1]=='\n']
        p = [get_trans(w) for w in a ]
    with open(r'./mean.txt','wb+') as writer:
        line = 'words phonetic trans\n'
        writer.write(line)
        for idx in range(p.__len__()):
            line = '%s %s %s\n'%(a[idx],p[idx][0],p[idx][1])
            writer.write(line)


if __name__ == '__main__':
    pass
    # text_to_excel()
    # main()
