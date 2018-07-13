import sys
import time

def data_rate():
    '''
    train data:
        online_minute = 0 but paied samples' index:
            133665
            274623
            394692
            448108
            661424
            790134
            850466
            885132
            913089
            1105041
            1180427
            1221005
            1386851
            1493818
            1564177
            1735839
            1825074
            1861181
            1958652
            2074448
            2188535
        samplenum:2288007:41439:7457.95:0.99
        samplenum:2288007:45988:32977.81:0.99
        zero:4549,not zero:11309
        zero:1098.833333,0.166667
    test data:
        samplenum:828934:19549:14495.22:0.99
        zero:1081.5,0.333333
    '''
    ind = 0
    day7_pay_num = 0
    day45_pay_num = 0
    day_pay_thres = 0.0
    day7_pay_max = 0.0
    day7_pay_min = 666666.0
    day45_pay_max = 0.0
    day45_pay_min = 666666.0
    day45_more_day7_zero = 0
    day45_more_day7_notzero = 0
    per_online_minute_max = 0.0
    per_online_minute_min = 666666.0
    for line in sys.stdin:
        ind += 1
        line = line.strip('\n')
        if ind == 1:
            continue
        data = line.split(',')
        per_online_minute = float(data[-4])#avg_online_minutes
        day7_pay_price = float(data[-3])#day7_pay_price
        # online_count = float(data[-2])#
        day45_pay_price = float(data[-1])#day45_pay_price
        if day7_pay_price > day_pay_thres:
            day7_pay_num += 1
            #     # figure out online minute
            # if per_online_minute > per_online_minute_max:
            #     per_online_minute_max = per_online_minute
            # if per_online_minute < per_online_minute_min:
            #     if per_online_minute == day_pay_thres:
            #         print(ind)
            #     else:
            #         per_online_minute_min = per_online_minute
            if day45_pay_price > day7_pay_price:
                day45_more_day7_notzero += 1
            if day7_pay_price < day7_pay_min:
                day7_pay_min = day7_pay_price
            if day7_pay_price > day7_pay_max:
                day7_pay_max = day7_pay_price
        elif day45_pay_price > day7_pay_price:
            # if day7_pay_price == day_pay_thres:
            #     print(data)
            #     time.sleep(5)
            day45_more_day7_zero += 1
            # figure out online minute
            if per_online_minute > per_online_minute_max:
                per_online_minute_max = per_online_minute
            if per_online_minute < per_online_minute_min:
                if per_online_minute == day_pay_thres:
                    print(ind)
                else:
                    per_online_minute_min = per_online_minute
        
        if day45_pay_price > day_pay_thres:
            day45_pay_num += 1
            if day45_pay_price < day45_pay_min:
                day45_pay_min = day45_pay_price
            if day45_pay_price > day45_pay_max:
                day45_pay_max = day45_pay_price
    print('samplenum:{}:{}:{}:{}'.format(ind-1, day7_pay_num, day7_pay_max, day7_pay_min))
    print('samplenum:{}:{}:{}:{}'.format(ind-1, day45_pay_num, day45_pay_max, day45_pay_min))
    print('zero:{},not zero:{}'.format(day45_more_day7_zero,day45_more_day7_notzero))
    print('zero:{},{}'.format(per_online_minute_max, per_online_minute_min))
    return

def data_clean():
    '''
    828934:17413:791972:19549
    0.166667:1605.833333
    '''
    i = 0
    filter_num = 0 # sample num filtered
    pay_num = 0 # sample pay in 7 days
    payzero_num = 0
    per_time_min = 66666.0
    per_time_max = 0.0
    count_online_zero = 0
    count_online_notzero = 0
    for line in sys.stdin:
        i += 1
        line = line.strip('\r\n')
        if i == 1:
            continue
        data = line.split(',')
        data_4 = float(data[-3])#avg_online_minutes
        data_3 = float(data[-2])#day7
        data_2 = float(data[-1])#online count
        if data_4 == 0.0:
            filter_num += 1
        else:
            if data_3 == 0.0:
                payzero_num += 1
                if data_4 > per_time_max:
                    per_time_max = data_4
                if data_4 < per_time_min:
                    per_time_min = data_4
                if data_2 > count_online_zero:
                    count_online_zero += 1
            else:
                # print(data)
                # time.sleep(5)
                pay_num += 1
                if data_2 > count_online_notzero:
                    count_online_notzero += 1
    print('{}:{}:{}:{}'.format(i-1, filter_num, payzero_num, pay_num))
    print('{}:{}'.format(per_time_min, per_time_max))
    print('{}:{}'.format(count_online_zero, count_online_notzero))
    return

if __name__ == '__main__':
    # data_rate()
    data_clean()