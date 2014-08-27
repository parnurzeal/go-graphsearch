#include<iostream>
#include<sstream>
#include<vector>


using namespace std;

class Node{
  public:
    Node *up, *down, *left, *right;
    int state;
    enum State { WALL, SPACE, START, END };
};


int main(){
  vector<vector<char> > inputs;
  string line;
  while(getline(cin, line)){
    istringstream iss(line);
    for(int i= 0 ;i<line.length();i++){
      cout<<line[i];
    }
    cout<<endl;
  }
   
  cout<<"Hello world!";
  return 0;
}
