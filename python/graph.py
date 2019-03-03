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

    def dfs(self, vertex, beggin=True):
        if beggin:
            self.paths = []
        self.paths.append(vertex.value)
        vertex.visited = True
        for v in  vertex.edges:
            if not v.visited:
                self.dfs(v, False)

    def bfs(self, vertex):
        pass

    def bds(self, v_start, v_end):
        pass

    def clear(self):
        for v in self.vertexs:
            v.visited = False




class Vertex(object):
    def __init__(self, value, weight):
        self.value = value
        self.weight = weight
        self.visited = False
        self.edges = []

    def add(self, obj):
        self.edges.append(obj)


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

    graph.dfs(v2)
    print(graph.paths)
