import redis

# Connect to Redis instance
r = redis.Redis(host='localhost', port=9093)

def fibonacci(n):
    if n < 0:
        return None
    if n <= 1:
        return n

    key = f"fibonacci:{n}"
    if r.exists(key):
        return int(r.get(key))

    if not r.exists("fibonacci:prev"):
        r.set("fibonacci:prev", 0)
    if not r.exists("fibonacci:curr"):
        r.set("fibonacci:curr", 1)
    prev = int(r.get("fibonacci:prev"))
    curr = int(r.get("fibonacci:curr"))

    for i in range(2, n + 1):
        prev, curr = curr, prev + curr
        r.set("fibonacci:prev", prev)
        r.set("fibonacci:curr", curr)
        r.set(f"fibonacci:{i}", curr)

    return curr

# Calculate the 100th Fibonacci number
n = 1000
result = fibonacci(n)

print(f"The {n}th Fibonacci number is: {result}")