exports.handler = async (event) => {
    const n=30
    const response = {
        statusCode: 200,
        body: 'Fibonacci(' + n + ') = ' + Fibo(n),
    };
    return response;
};

function Fibo(n) {
    if (n <= 2)
        return n - 1;
    else
        return Fibo(n - 1) + Fibo(n - 2);
}
