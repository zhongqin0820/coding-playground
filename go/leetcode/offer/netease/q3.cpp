#include <iostream>
#include <algorithm>
#include <vector>

int MAX = 100000;
int MAXW = 1000000000;
using namespace std;
int main() {
	int n,m;
	cin>>n>>m;
	int w[MAX];
	for(int i=0;i<n;i++){
		cin>>w[i];
	}
	sort(w,w+n);
	int i=0;
	int j=n-1;
	int res = 0;
	while (i<j){
		if (w[i]+w[j]>m) {
			res+=2;
			i++;
			j--;
		} else {
			res+=1;
			i++;
			j--;
		}
	}

	// for(int i=0;i<n;i++){
		// cout<<w[i];
	// }
	cout<<res<<endl;
	// cout<<endl;
	return 0;
}
