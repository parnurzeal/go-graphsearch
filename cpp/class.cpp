#include<iostream>
#include<vector>
#include<sstream>
#include<string>


using namespace std;

class Node{
  public:
  enum State{WALL='|', SPACE=' ', START='S',END='E', ANSWER='x'};
  State state;
  Node *up,*down, *left, *right;
  Node *pre;
  bool passed;
  Node(){
    passed = false;
  }
};


bool DFS(Node *node, Node *preNode){
 if(node == NULL){ return false; }
 if(node->passed){ return false; }
 node->passed=true;
 node->pre = preNode;
 switch(node->state){
  case Node::WALL: return false;
  case Node::END: return true;
  default:;break;
 }
 if( DFS(node->up,node) || DFS(node->down,node) || DFS(node->left,node) || DFS(node->right,node) ){
   return true;
 }
 return false;
}

int main(){
  // input
  vector<vector<char> > inputs;
  string line;
  while(getline(cin,line)){
    vector<char> line_tmp;
    for(int i = 0 ;i<line.length();i++){
      line_tmp.push_back(line[i]);
    }
    inputs.push_back(line_tmp);
  }

  // find matrix square size
  int maxR=inputs.size(), maxC=0;
  for(int i = 0;i<inputs.size();i++){
    if(inputs[i].size()>maxC){ maxC=inputs[i].size();}
  }
 
  // create nodes
  vector<vector<Node> > graph;
  for(int i =0;i<maxR;i++){
    vector<Node> lineNode;
    for(int j=0;j<maxC;j++){
      Node tmpNode;
      tmpNode.state = Node::State(inputs[i][j]);
      lineNode.push_back(tmpNode);
    }
    graph.push_back(lineNode);
  }
  // connect nodes
  for(int i = 0;i<maxR;i++){
    for(int j = 0;j<maxC;j++){
      // connect up
      if(i-1>=0){
        graph[i][j].up = &graph[i-1][j];
      }else{
        graph[i][j].up = NULL;
      }
      // connect down
      if(i+1<maxR){
        graph[i][j].down = &graph[i+1][j];
      }else{
        graph[i][j].down = NULL;
      }
      // connect left
      if(j-1>=0){
        graph[i][j].left = &graph[i][j-1];
      }else{
        graph[i][j].left = NULL;
      }
      // connect right
      if(j+1<maxC){
        graph[i][j].right = &graph[i][j+1];
      }else{
        graph[i][j].right = NULL;
      }
    }
  }
  // find start and end node
  Node *startNode, *endNode;
  for(int i = 0;i<maxR;i++){
    for(int j = 0;j<maxC;j++){
      if(graph[i][j].state==Node::START){
        startNode=&graph[i][j];
      }
      if(graph[i][j].state==Node::END){
        endNode=&graph[i][j];
      }
    }
  }
  // solution
  if(DFS(startNode,NULL)){
    while(endNode!=NULL){
      endNode->state = Node::ANSWER;
      endNode = endNode->pre;
    }
    // output
    for(int i = 0;i<maxR;i++){
      for(int j = 0;j<maxC;j++){
        cout<<char(graph[i][j].state);
      }
      cout<<endl;
    }
  }else{
    cout<<"No solution";
  }

}
