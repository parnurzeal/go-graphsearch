#include<iostream>
#include<sstream>
#include<vector>


using namespace std;

/*class Node{
  public:
    Node *up, *down, *left, *right;
    int state;
    enum State { WALL, SPACE, START, END };
};*/

typedef struct {
  bool passed;
  char state;
  int preR, preC;
  enum State { WALL, SPACE, START, END };
} Node;

bool DFS(vector<vector<Node> > &all_nodes, int preR, int preC, int incR, int incC){
  int r = preR + incR, c = preC +incC;
  if(r<0 || r>=all_nodes.size() || c <0 || c>=all_nodes[0].size() || all_nodes[r][c].passed ){
    return false;
  }
  all_nodes[r][c].passed = true;
  if(all_nodes[r][c].state =='|' || all_nodes[r][c].state=='-'){
    return false;
  }
  all_nodes[r][c].preR = preR; all_nodes[r][c].preC = preC;
  if( all_nodes[r][c].state == 'E' ){
    return true;
  }
  if(DFS(all_nodes,r,c,-1,0) || DFS(all_nodes,r,c,0,-1) || DFS(all_nodes,r,c,1,0) || DFS(all_nodes,r,c,0,1)){
    return true;
  }
  return false;
}

int main(){
  vector<vector<Node> > all_nodes;
  string line;
  while(getline(cin, line)){
    istringstream iss(line);
    vector<Node> line_ins;
    for(int i= 0 ;i<line.length();i++){
      Node tmp; tmp.state = line[i];
      line_ins.push_back(tmp);
    }
    all_nodes.push_back(line_ins);
  }
  int rSize = all_nodes.size();
  int cSize = all_nodes[0].size();
  int rStart =0, cStart =0;
  int rEnd =0, cEnd =0;
  for(int i = 0 ;i<rSize;i++){
    for(int j = 0;j<cSize;j++){
      all_nodes[i][j].passed=false;
      if(all_nodes[i][j].state=='S'){
        rStart = i; cStart =j;
      }
      if(all_nodes[i][j].state=='E'){
        rEnd =i; cEnd=j;
      }
    }
  }

  if(DFS(all_nodes,rStart, cStart,0,0)){
    int rTmp=rEnd, cTmp=cEnd;
    while(rTmp!=rStart || cTmp!=cStart){
      if(all_nodes[rTmp][cTmp].state!='E' && all_nodes[rTmp][cTmp].state!='S'){
        all_nodes[rTmp][cTmp].state = 'a';
      }
      int r=all_nodes[rTmp][cTmp].preR,c=all_nodes[rTmp][cTmp].preC;
      rTmp=r;cTmp=c;
    }
  }else{
    cout << "No answer";
  }

  for(int i = 0 ;i<rSize;i++){
    for(int j = 0;j<cSize;j++){
      cout<<all_nodes[i][j].state;
    }
    cout<<endl;
  }
  return 0;
}
