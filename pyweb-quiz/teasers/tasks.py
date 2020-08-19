# A Task to Do

from heapq import heappush, heappop

# Tasks priority queue
# A task is a tuple of (priority, payload)
tasks = []
heappush(tasks, (30, 'work out'))
heappush(tasks, (10, 'wake up'))
heappush(tasks, (20, 0xCAFFE))
heappush(tasks, (20, 'feed cat'))
heappush(tasks, (40, 'write book'))

while tasks:
    _, payload = heappop(tasks)
    print(payload)
