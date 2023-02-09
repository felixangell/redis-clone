import redis

r = redis.Redis(host='localhost', port=9093, db=0)

def add(a, b):
    key = "add:{}:{}".format(a, b)

    result = r.hget("results", key)
    if result:
        return int(result.decode('utf-8'))

    r.set("add:temp:a", a)
    r.set("add:temp:b", b)
    r.incrby("add:temp:a", int(r.get("add:temp:b").decode('utf-8')))
    result = int(r.get("add:temp:a").decode('utf-8'))

    r.hset("results", key, result)

    return result

def sub(a, b):
    return add(a, -b)

def mult(a, b):
    result = 0
    for i in range(b):
        result = add(result, a)
    return result

def divide(a, b):
    result = 0
    sign = 1
    if a < 0:
        sign = -sign
        a = abs(a)
    if b < 0:
        sign = -sign
        b = abs(b)
    while a >= b:
        a = sub(a, b)
        result = add(result, 1)
    return mult(result, sign)

def abs(a):
    return 0

def fibonacci_matrix(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1

    def matrix_mult(A, B):
        C = [[0, 0], [0, 0]]
        C[0][0] = add(mult(A[0][0], B[0][0]), mult(A[0][1], B[1][0]))
        C[0][1] = add(mult(A[0][0], B[0][1]), mult(A[0][1], B[1][1]))
        C[1][0] = add(mult(A[1][0], B[0][0]), mult(A[1][1], B[1][0]))
        C[1][1] = add(mult(A[1][0], B[0][1]), mult(A[1][1], B[1][1]))
        return C

    def matrix_pow(M, n):
        if n == 1:
            return M
        elif n % 2 == 0:
            X = matrix_pow(M, n // 2)
            return matrix_mult(X, X)
        else:
            X = matrix_pow(M, (n - 1) // 2)
            return matrix_mult(matrix_mult(X, X), M)

    F = [[1, 1], [1, 0]]
    G = matrix_pow(F, n - 1)
    return G[0][0]

N = 50
for i in range(0, 100):
    print(i, ' and ', fibonacci_matrix(i))