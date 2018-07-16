import xgboost as xgb
import numpy as np
import random

def mknfold(fname='../binary_classification/agaricus.txt', k=1, nfold=5):
    random.seed(10)
    fi = open(fname, 'r')
    ftr = open('agaricus.txt.train', 'w')
    fte = open('agaricus.txt.test', 'w')
    for l in fi:
        if random.randint(1, nfold) == k:
            fte.write(l)
        else:
            ftr.write(l)
    fi.close()
    ftr.close()
    fte.close()
    return

def train_and_eval(tr_filename='agaricus.txt.train',te_filename='agaricus.txt.test'):
    dtrain = xgb.DMatrix(tr_filename)
    dtest = xgb.DMatrix(te_filename)
    params = {
        'max_depth':2,
        'eta':1,
        'silent':1,
        'objective':'reg:linear',
        'eval_metric':'rmse'
    }
    watchlist = [(dtest,'eval'), (dtrain,'train')]
    num_round = 2
    bst = xgb.train(params, dtrain, num_round, watchlist)
    # testing
    # preds = bst.predict(dtest)
    # labels = dtest.get_label()
    bst.dump_model('dump.raw.txt')
    return

def tune_model():
    pass

if __name__ == '__main__':
    mknfold()
    train_and_eval()
    pass