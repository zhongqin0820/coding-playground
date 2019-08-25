# initialization, mainly two categories: adjacency matrix and adjacency list
# here, we use adjacency list
graph = {
    'A': set(['B','C']),
    'B': set(['A', 'D', 'E']),
    'C': set(['A', 'F']),
    'D': set(['B']),
    'E': set(['B', 'F']),
    'F': set(['C', 'E'])
}

# graph search: depth-fisrt search and breath-first search
# DFS
def dfs1(graph, start):
    """
    :type graph: set dict
    :type start: char
    :rtype: list
    """
    visited, stack = set(), [start]
    while stack:
        vertex = stack.pop()
        if vertex not in visited:
            visited.add(vertex)
            stack.extend(graph[vertex]-visited)
    return visited

def dfs2(graph, start, visited=None):
    """
    :type grpah:
    :type start:
    :type visited:
    :rtype: 
    """
    if visited is None:
        visited=set()
    visited.add(start)
    for next in graph[start]-visited:
        dfs2(graph, next, visited)
    return visited

# return all possible paths between a start and goal vertex.
def dfs1_path(graph, start, goal):
    stack = [(start, [start])]
    while stack:
        (vertex, path) = stack.pop()
        for next in graph[vertex]-set(path):
            if next == goal:
                yield path + [next]
            else:
                stack.append((next, path+[next]))

# def dfs2_path(graph, start, goal, path=None):
#     if path is None:
#         path = [start]
#     if start == goal:
#         yield path
#     for next in graph(start) - set(path):
#         dfs2_path(graph, next, goal, path+[next])

# BFS
from collections import deque
def bfs(graph, start):
    visited, deq = set(), deque(start)
    while deq:
        vertex=deq.popleft()
        if vertex not in visited:
            visited.add(vertex)
            deq.extend(graph[vertex]-visited)
    return visited

def bfs_path(graph, start, goal):
    deq = [(start, [start])]
    while deq:
        (vertex, path) = deq.pop(0)
        for next in graph[vertex] - set(path):
            if next == goal:
                yield path + [next]
            else:
                deq.append((next, path+[next]))

def shortest_path(graph, start, goal):
    # the shortest path will be returned first from the BFS path generator 
    try:
        return next(bfs_path(graph, start, goal))
    except StopIteration:
        return None

if __name__ == '__main__':
    print('dfs1: ', dfs1(graph,'A'))
    print('dfs2: ', dfs2(graph,'A'))
    print('dfs1 path', list(dfs1_path(graph, 'A', 'F'))) # [['A', 'C', 'F'], ['A', 'B', 'E', 'F']])
    # print('dfs2 path', list(dfs2_path(graph, 'A', 'F'))) # [['A', 'C', 'F'], ['A', 'B', 'E', 'F']])
    print('bfs1: ', bfs(graph,'A'))
    print(list(bfs_path(graph, 'A', 'F'))) # [['A', 'C', 'F'], ['A', 'B', 'E', 'F']]
    print(shortest_path(graph, 'A', 'F'))
