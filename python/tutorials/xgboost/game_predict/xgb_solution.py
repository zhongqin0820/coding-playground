import xgboost as xgb
import os
import random
def mapfeat(dat_filename='game.txt'):
    if os.path.exists(dat_filename):
        print('data existed')
        return dat_filename
    print('start mapping feature')
    with open(dat_filename,'w+') as writer:
        with open('tap4fun_data/tap_fun_train.csv','r+') as reader:
            idx = 0
            for line in reader.readlines():
                idx += 1
                # feature loaded
                if idx == 1:
                    # line = line.strip('\n')
                    # data = line.split(',')[2:]
                    # print(data)
                    continue
                # data loaded label:prediction_pay_price
                line = line.strip('\n')
                data = line.split(',')
                label = data[-1]
                fea = data[2:-1]
                res = ''
                for i, v in enumerate(fea):
                    res += str(i) +':' + v + ' '
                res = label + ' ' + res + '\n'
                writer.write(res)
    return dat_filename

def mknfold(dat_filename='game.txt', k=1, n=5):
    tr_filename = dat_filename + '.train'
    te_filename = dat_filename + '.test'
    if (os.path.exists(tr_filename) or os.path.exists(te_filename)):
        print('data splitted')
        return tr_filename, te_filename
    print('splitting data into train and test')
    random.seed(10)
    with open(tr_filename, 'a+') as writer_tr:
        with open(te_filename, 'a+') as writer_te:
            with open(dat_filename,'r+') as reader:
                for line in reader.readlines():
                    if random.randint(1,n) == k:
                        writer_te.write(line)
                    else:
                        writer_tr.write(line)
                    # break
    return tr_filename, te_filename

def train(tr_filename='game.txt.train', te_filename='game.txt.test'):
    print('training...')
    dtrain = xgb.DMatrix(tr_filename)
    dtest = xgb.DMatrix(te_filename)
    params={
        'max_depth':10,
        'eta':1,
        'silent':1,
        'objective':'reg:linear',
        'eval_metric':'rsme'
    }
    num_rounds = 100
    watchlist = [(dtest,'eval'),(dtrain,'train')]
    bst = xgb.train(params, dtrain, num_rounds, watchlist, early_stopping_rounds=10)
    bst.save_model('0001.model')
    bst.dump_model('dump0001.raw.txt')
    # preds = bst.predict(dtest)
    # labels = dtest.get_label()
    return

if __name__ == '__main__':
    mapfeat()
    mknfold()
    train()