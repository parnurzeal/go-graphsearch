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

void DFS(vector<vector<Node> > &all_nodes, int preR, int preC, int incR, int incC){
  int r = preR + incR, c = preC +incC;
  if(r<0 || r>=all_nodes.size() || c <0 || c>=all_nodes[0].size() | all_nodes[r][c].passed ){
    return;
  }
  all_nodes[r][c].passed = true;
  all_nodes[r][c].preR = preR; all_nodes[r][c].preC = preC;
  if( all_nodes[r][c].state == 'E' ){
    cout<< "foundE";
    return;
  }
  DFS(all_nodes,r,c,-1,0);
  DFS(all_nodes,r,c,0,-1);
  DFS(all_nodes,r,c,1,0);
  DFS(all_nodes,r,c,0,1);
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
  for(int i = 0 ;i<rSize;i++){
    for(int j = 0;j<cSize;j++){
      all_nodes[i][j].passed=false;
      if(all_nodes[i][j].state=='S'){
        rStart = i; cStart =j;
      }
    }
  }

  DFS(all_nodes,rStart, cStart,0,0);

  return 0;
}
