# !/usr/bin/python

def loadfmap(fname):
    fmap = {}
    nmap = {}
    for l in open(fname):
        arr = l.split()
        if arr[0].find('.')!=-1:
            idx = int( arr[0].strip('.') )
            assert idx not in fmap
            fmap[idx]={}
            ftype=arr[1].strip(':')
            content = arr[2]
        else:
            content = arr[0]
        for it in content.split(','):
            if it.strip() == '':
                continue
            k, v = it.split('=')
            fmap[idx][v] = len(nmap) + 1
            nmap[len(nmap)] = ftype+'='+k
    return fmap, nmap

def write_nmap(fo, nmap):
    for i in range( len(nmap) ):
        fo.write( '%d\t%s\ti\n' % (i, nmap[i]) )

def mapfeat():
    # load feature
    fmap, nmap = loadfmap('agaricus-lepiota.fmap')
    fo = open('featmap.txt', 'w')
    write_nmap(fo, nmap)
    fo.close()
    # write data
    fo = open('agaricus.txt', 'w')
    for l in open('agaricus-lepiota.data'):
        arr = l.split(',')
        if arr[0] == 'p':
            fo.write('1')
        else:
            assert arr[0] == 'e'
            fo.write('0')
        for i in range(1,len(arr)):
            fo.write(' %d:1' % fmap[i][arr[i].strip()])
        fo.write('\n')
    fo.close()
    return 'agaricus.txt'

def mknfold(fname='agaricus.txt', k=1, nfold=5):
    '''
    1st arg: dataset filename [agaricus.txt]
    2nd arg: k fold num used in tset [k = 1]
    3rd arg: nfold num [nfold = 5]
    '''
    import random
    random.seed(10)
    fi = open(fname, 'r')
    ftr = open(fname+'.train', 'w')
    fte = open(fname+'.test', 'w')
    for l in fi:
        if random.randint(1, nfold) == k:
            fte.write(l)
        else:
            ftr.write(l)
    fi.close()
    ftr.close()
    fte.close()

if __name__ == '__main__':
    # mapfeat()
    mknfold()
