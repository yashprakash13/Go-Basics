# Python code

```python
class Graph:
    def __init__(self, vertices):
        self.graph = defaultdict(list)
        self.V = vertices
    
    def insert_edge(self, from_node, to_node):
        self.graph[from_node].append(to_node)
    
    def if_all_courses_complete(self):
        in_degree = [0] * self.V
        for node in self.graph:
            for prereq in self.graph[node]:
                in_degree[prereq] += 1
        
        queue = deque()
        for i, degree in enumerate(in_degree):
            if degree == 0 :
                queue.append(i)
        
        count_vis = 0
        
        while queue:
            node = queue.popleft()
            count_vis += 1
            
            for prereq in self.graph[node]:
                in_degree[prereq] -= 1
                if in_degree[prereq] == 0:
                    queue.append(prereq)
        
        if count_vis == self.V:
            return True
        else:
            return False



class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        if numCourses == 1:
            return True
        if prerequisites == []:
            return True
        
        graph = Graph(numCourses)
        
        for pre in prerequisites:
            graph.insert_edge(pre[1], pre[0])
        
        return graph.if_all_courses_complete()
        
                    
        
```