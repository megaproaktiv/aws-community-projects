


def handler(event, context):
    n = 30
    return 'Fibonacci({}) = {}'.format(n, Fibo(int(n)))




# Function for nth Fibonacci number
def Fibo(n):
    if n <= 2:
        return n - 1
    else:
        return Fibo(n - 1) + Fibo(n - 2)
