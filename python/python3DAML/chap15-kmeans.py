import numpy as np
import matplotlib.pyplot as plt
# %matplotlib inline
import pandas as pd
from sklearn.datasets.samples_generator import make_blobs
from sklearn.cluster import KMeans
from sklearn.preprocessing import StandardScaler
from sklearn.pipeline import make_pipeline
from sklearn import metrics

def toy_data():
    # generate sample data
    X, y = make_blobs(n_samples=1000, n_features=2, centers=[[-1,-1],[0,0],[1,1],[2,2]],cluster_std=[0.4,0.2,0.2,0.2],random_state=2019)

    # plotting data
    # plt.scatter(X[:,0], X[:,1], marker='o')
    # plt.show()

    y_pred = KMeans(n_clusters=4, random_state=2019).fit_predict(X)
    # plt.scatter(X[:,0],X[:,1],c=y_pred)
    # plt.show()

    print(metrics.calinski_harabaz_score(X,y_pred))


def toy_fish():
    df = pd.read_csv('fish.csv')
    species = list(df['species'])
    del df['species']
    df.head()
    # build pipeline
    scalar = StandardScaler()
    samples = df.values
    kmeans = KMeans(n_clusters=4)
    pipeline = make_pipeline(scalar, kmeans)
    # training and testing
    pipeline.fit(samples)
    labels = pipeline.predict(samples)
    # analysis output
    df = pd.DataFrame({'labels':labels, 'species':species})
    ct = pd.crosstab(df['labels'], df['species'])
    print(ct)


def main():
    toy_fish()


if __name__ == '__main__':
    main()
