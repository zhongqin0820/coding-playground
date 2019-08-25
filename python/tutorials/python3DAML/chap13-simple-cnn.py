from keras.models import Sequential
from keras.layers import Dense, Dropout, Flatten
from keras.layers.convolutional import Convolution2D, MaxPooling2D
from keras.utils import np_utils
from keras.datasets import mnist
import numpy as np

# /Users/Zoking/.keras/datasets/
(x_train, y_train), (x_test, y_test) = mnist.load_data()
num_pixes = x_train.shape[1]*x_train.shape[2]
n_channels = 1
def preprocess(matrix):
    return matrix.reshape(matrix.shape[0], n_channels, matrix.shape[1], matrix.shape[2]).astype('float32')/255.
x_train, x_test = preprocess(x_train), preprocess(x_test)
y_train = np_utils.to_categorical(y_train)
y_test = np_utils.to_categorical(y_test)
num_classes = y_train.shape[1]


def baseline_model():
    model = Sequential()
    model.add(Flatten(input_shape=(1, 28, 28)))
    model.add(Dense(num_pixes, init='normal', activation='relu'))
    model.add(Dense(num_classes, init='normal', activation='softmax'))
    model.compile(loss='categorical_crossentropy', optimizer='adam', metrics=['accuracy'])
    return model


def simple_cnn():
    model = Sequential()
    model.add(Convolution2D(32, 5, 5, border_mode='valid',input_shape=(1,28,28),activation='relu'))
    model.add(MaxPooling2D(pool_size(2,2)))
    model.add(Dropout(0.2))
    model.add(Flatten())
    model.add(Dense(128, activation='relu'))
    model.add(Dense(num_classes, activation='softmax'))
    model.compile(loss='categorical_crossentropy', optimizer='adam', metrics=['accuracy'])
    return model


def complex_cnn():
    model = Sequential()
    model.add(Convolution2D(30, 5, 5, border_mode='valid',input_shape=(1,28,28),activation='relu'))
    model.add(MaxPooling2D(pool_size(2,2)))
    model.add(Convolution2D(15, 3, 3, activation='relu'))
    model.add(MaxPooling2D(pool_size(2,2)))
    model.add(Dropout(0.2))
    model.add(Flatten())
    model.add(Dense(128, activation='relu'))
    model.add(Dense(50, activation='relu'))
    model.add(Dense(num_classes, activation='softmax'))
    model.compile(loss='categorical_crossentropy', optimizer='adam', metrics=['accuracy'])
    return model

np.random.seed(2019)
models=[('baseline', baseline_model()),('simple', simple_cnn()), ('complex',complex_cnn())]

for name, model in models:
    print(name, ' ')
    model.fit(x_train, y_train, validation_data=(x_test, y_test), nb_epoch=10, batch_size=100, verbose=2)
    scores = model.evaluate(x_test, y_test, verbose=0)
    print('error: %.2f%%'%(100-scores[1]*100))
