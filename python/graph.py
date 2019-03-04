class Graph(object):

    def __init__(self):
        self.vertexs = []
        self.paths = []

    def add(self, vertex):
        self.vertexs.append(vertex)

    def visual(self):

        for v in self.vertexs:
            print("Vertex {} {} ->".format(v.value, v.weight))
            for e in v.edges:
                print("Edges {} {}".format(e.value, e.weight))

    def dfs(self, vertex):
        self.paths.append(vertex.value)
        vertex.visited = True
        for v in  vertex.edges:
            if not v.visited:
                self.dfs(v)

    def bfs(self, vertex, queue=None, beggin=True):
        if beggin or queue == None:
            queue = []

        vertex.visited = True
        queue.append(vertex)

        while queue:
            v = queue.pop(0)
            self.paths.append(v.value)
            for e in v.edges:
                if not e.visited:
                    e.visited = True
                    queue.append(e)

    def is_visited(self, s_visited, e_visited):
        for v in self.vertexs:
            if (v.value in s_visited) and (v.value in e_visited):
                return v

        return None

    def bds_bfs(self, vertex, queue, visited, parent):
        v = queue.pop(0)
        for e in v.edges:
            if e.value not in visited:
                visited.append(e.value)
                parent[e.value] = v
                queue.append(e)

    def print_Path(self, s_parent, e_parent, v_start, v_end, match):
        self.paths.append(match)
        tmp = match

        while tmp != v_start:
            self.paths.append(s_parent[tmp.value])
            tmp = s_parent[tmp.value]

        self.paths.reverse()
        tmp = match

        while tmp != v_end:
            self.paths.append(e_parent[tmp.value])
            tmp = e_parent[tmp.value]

        print(" -> ".join([str(v.value) for v in self.paths]))

    def bds(self, v_start, v_end):
        s_queue = []
        e_queue = []
        s_visited = []
        e_visited = []

        s_parent = {}
        e_parent = {}

        s_queue.append(v_start)
        s_visited.append(v_start.value)

        e_queue.append(v_end)
        e_visited.append(v_end.value)

        while s_queue and e_queue:
            self.bds_bfs(v_start, s_queue, s_visited, s_parent)
            self.bds_bfs(v_end, e_queue, e_visited, e_parent)
            # print(s_visited, e_visited)
            result = self.is_visited(s_visited, e_visited)
            if result:
                self.print_Path(s_parent, e_parent, v_start, v_end, result)
                break

    def clear(self):
        for v in self.vertexs:
            v.visited = False
        self.paths = []




class Vertex(object):
    def __init__(self, value, weight):
        self.value = value
        self.weight = weight
        self.visited = False
        self.edges = []

    def add(self, obj):
        self.edges.append(obj)
        obj.edges.append(self)


if __name__ == "__main__":
    graph = Graph()
    v1 = Vertex(1, "Quang")
    v2 = Vertex(2, "Tuan")
    v3 = Vertex(3, "Duc")
    v4 = Vertex(4, "Hiep")
    v5 = Vertex(5, "Tung")

    v1.add(v2)
    v2.add(v3)
    v3.add(v1)
    v3.add(v4)
    v1.add(v4)
    v4.add(v5)
    v2.add(v5)

    graph.add(v1)
    graph.add(v2)
    graph.add(v3)
    graph.add(v4)
    graph.add(v5)

    graph.visual()

    graph.clear()
    graph.dfs(v2)
    print(graph.paths)

    graph.clear()
    graph.bfs(v3)
    print(graph.paths)

    graph.clear()
    graph.bds(v1, v5)
