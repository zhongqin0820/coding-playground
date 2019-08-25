import sys

def data_fmap():
    '''
    mapping feature into xgboost
    '''
    pass

def train():
    '''
    setup parameter
    training with xgboost
    '''
    pass

def predict():
    '''
    using xgboost to predict 
    '''
    ind = 0
    with open('sub_prediction.csv','a+') as writer:
        writer.write('user_id,prediction_pay_price\n')
        for line in sys.stdin:
            ind += 1
            line = line.strip('\n')
            if ind == 1:
                continue
            data = line.split(',')
            user_id = int(data[0])
            pay_predict = float(data[-2])#day7
            if pay_predict == 0.0:
                final_pay = 0
            else:
                final_pay = pay_predict
            writer.write('{},{}\n'.format(user_id, final_pay))
            # if ind == 4:
            #     break

if __name__ == '__main__':
    data_fmap()
    train()
    predict()