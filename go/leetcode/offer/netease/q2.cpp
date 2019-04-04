#include <iostream>
int MAX = 100000;
using namespace std;
int main() {
	int n;
	cin>>n;
	int a[MAX];
	for(int i=0;i<n;i++){
		cin>>a[i];
	}
	for(int i=0;i<n;i++) {
		for(int j=i+1;j<n;j++) {
			for(int k=j+1;k<n;k++) {
				
			}
		}
	}


	for(int i=0;i<n;i++){
		cout<<a[i];
	}
	cout<<"\n";
	return 0;
}

