lower_bound:
	python ./dma/contests/run_tests.py ./solutions/1-binary-search/lower_bound/lower_bound.go ./dma/contests/1-binary-search/lower_bound_tests.zip

nearest_element:
	python ./dma/contests/run_tests.py ./solutions/1-binary-search/nearest_element/nearest_element.go ./dma/contests/1-binary-search/nearest_element_tests.zip

peak:
	python ./dma/contests/run_tests.py ./solutions/1-binary-search/peak/peak.go ./dma/contests/1-binary-search/peak_tests.zip

any_sort:
	python ./dma/contests/run_tests.py ./solutions/2-quicksort-mergesort/any_sort/any_sort.go ./dma/contests/2-quicksort-mergesort/any_sort_tests.zip

merge_sort:
	python ./dma/contests/run_tests.py ./solutions/2-quicksort-mergesort/merge_sort/merge_sort.go ./dma/contests/2-quicksort-mergesort/merge_sort_tests.zip

quicksort:
	python ./dma/contests/run_tests.py ./solutions/2-quicksort-mergesort/quicksort/quicksort.go ./dma/contests/2-quicksort-mergesort/quicksort_tests.zip

dfs:
	python ./dma/contests/run_tests.py ./solutions/3-dfs-bfs/dfs/dfs.go ./dma/contests/3-dfs-bfs/dfs_tests.zip

bfs:
	python ./dma/contests/run_tests.py ./solutions/3-dfs-bfs/bfs/bfs.go ./dma/contests/3-dfs-bfs/bfs_tests.zip

components:
	python ./dma/contests/run_tests.py ./solutions/3-dfs-bfs/components/components.go ./dma/contests/3-dfs-bfs/components_tests.zip

bfs_path:
	python ./dma/contests/run_tests.py ./solutions/4-topology-dijkstra/bfs_path/bfs_path.go ./dma/contests/4-topology-dijkstra/bfs_path_tests.zip

dijkstra:
	python ./dma/contests/run_tests.py ./solutions/4-topology-dijkstra/dijkstra/dijkstra.go ./dma/contests/4-topology-dijkstra/dijkstra_tests.zip

kan:
	python ./dma/contests/run_tests.py ./solutions/4-topology-dijkstra/kan/kan.go ./dma/contests/4-topology-dijkstra/kan_tests.zip

prim:
	python ./dma/contests/run_tests.py ./solutions/5-ford-bellman-prim/prim/prim.go ./dma/contests/5-ford-bellman-prim/prim_tests.zip	

ford_bellman:
	python ./dma/contests/run_tests.py ./solutions/5-ford-bellman-prim/ford_bellman/ford_bellman.go ./dma/contests/5-ford-bellman-prim/ford_bellman_tests.zip

all_tests:
	python ./dma/contests/run_tests.py ./solutions/1-binary-search/lower_bound/lower_bound.go ./dma/contests/1-binary-search/lower_bound_tests.zip
	python ./dma/contests/run_tests.py ./solutions/1-binary-search/nearest_element/nearest_element.go ./dma/contests/1-binary-search/nearest_element_tests.zip
	python ./dma/contests/run_tests.py ./solutions/1-binary-search/peak/peak.go ./dma/contests/1-binary-search/peak_tests.zip
	python ./dma/contests/run_tests.py ./solutions/2-quicksort-mergesort/any_sort/any_sort.go ./dma/contests/2-quicksort-mergesort/any_sort_tests.zip
	python ./dma/contests/run_tests.py ./solutions/2-quicksort-mergesort/merge_sort/merge_sort.go ./dma/contests/2-quicksort-mergesort/merge_sort_tests.zip
	python ./dma/contests/run_tests.py ./solutions/2-quicksort-mergesort/quicksort/quicksort.go ./dma/contests/2-quicksort-mergesort/quicksort_tests.zip
	python ./dma/contests/run_tests.py ./solutions/3-dfs-bfs/dfs/dfs.go ./dma/contests/3-dfs-bfs/dfs_tests.zip
	python ./dma/contests/run_tests.py ./solutions/3-dfs-bfs/bfs/bfs.go ./dma/contests/3-dfs-bfs/bfs_tests.zip
	python ./dma/contests/run_tests.py ./solutions/3-dfs-bfs/components/components.go ./dma/contests/3-dfs-bfs/components_tests.zip
	python ./dma/contests/run_tests.py ./solutions/4-topology-dijkstra/bfs_path/bfs_path.go ./dma/contests/4-topology-dijkstra/bfs_path_tests.zip
	python ./dma/contests/run_tests.py ./solutions/4-topology-dijkstra/dijkstra/dijkstra.go ./dma/contests/4-topology-dijkstra/dijkstra_tests.zip
	python ./dma/contests/run_tests.py ./solutions/4-topology-dijkstra/kan/kan.go ./dma/contests/4-topology-dijkstra/kan_tests.zip
	python ./dma/contests/run_tests.py ./solutions/5-ford-bellman-prim/prim/prim.go ./dma/contests/5-ford-bellman-prim/prim_tests.zip	
	python ./dma/contests/run_tests.py ./solutions/5-ford-bellman-prim/ford_bellman/ford_bellman.go ./dma/contests/5-ford-bellman-prim/ford_bellman_tests.zip
