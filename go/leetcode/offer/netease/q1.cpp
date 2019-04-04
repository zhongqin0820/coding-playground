#include <iostream>
#include <stack>
#include <string>
static const int npos = -1;
using namespace std;
int main() {
	string s;
	string sc (".,;:[{]}/?\'\"");
	string emp (" ");
	getline(cin,s);
	stack<string> q;
	int n = s.length();
	int i=1;
	int j=0;
	if( sc.find_first_of(s.substr(j,i-j))!=npos) {
		q.push(s.substr(j,i-j));
	}
	while (j<i && i<n) {
		if(emp.find_first_of(s.substr(i,1))!=npos ){
			if (sc.find_first_of(s.substr(j,1))==npos) {
				q.push(s.substr(j,i-j));
			}
			i=i+1;
			j=i;
			i=i+1;
		}else if (sc.find_first_of(s.substr(i,1))!=npos) {
			q.push(s.substr(i,1).append(s.substr(j,i-j)));
			j=i;
			i=i+1;
		} else {
			i++;
		}
	}
	while(q.size()!=0){
		s = q.top();
		q.pop();
		cout<<s<<" ";
	}
	cout<<endl;
}
