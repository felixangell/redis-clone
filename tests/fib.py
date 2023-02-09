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

    result = fibonacci(n - 1) + fibonacci(n - 2)
    r.set(key, result)

    return result

n = 300
result = fibonacci(n)

print(f"The {n}th Fibonacci number is: {result}")
