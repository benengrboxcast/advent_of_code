import heapq
from typing import Generic, TypeVar, List

T = TypeVar('T')

class MinHeap(Generic[T]):
    """
    Min heap implementation using heapq library
    """

    def __init__(self):
        self.heap: List[T] = []

    def insert(self, element: T):
        heapq.heappush(self.heap, element)

    def pop(self) -> T:
        return heapq.heappop(self.heap)

    def size(self) -> int:
        return len(self.heap)