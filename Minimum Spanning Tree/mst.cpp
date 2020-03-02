int spanningTree(int V,int E,vector<vector<int> > graph)
{
    // code here
    int i,j;
    bool mstSet[V] = {false};
    int keys[V] = {INT_MAX};
    
    int leftVerticeNum = V;
    int minIndex = 0, minIndex2=0;
    int minKey = INT_MAX;
    unsigned result=0;
    
    for(i=0; i<V ; ++i)
    
    
    while(leftVerticeNum){
        
        for(i=0; i<V ; ++i){
            if(graph[minIndex][i] < keys[i]){
                keys[i] = graph[minIndex][i];
            }
        }
        prevMinIndex=minIndex;
        minKey = INT_MAX;
        for(i=0; i<V ; ++i){
            if(!mstSet[i] && keys[i]<minKey){
                minIndex = i;
                minKey = keys[i];
            } 
        }
        
        mstSet[minIndex] = true;
        leftVerticeNum --;
        result += graph[prevMinIndex][minIndex];
    }
    
    return result;
}