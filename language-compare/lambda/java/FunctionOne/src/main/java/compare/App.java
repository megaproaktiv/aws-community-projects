package compare;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.URL;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;

/**
 * Handler for requests to Lambda function.
 */
public class App implements RequestHandler<APIGatewayProxyRequestEvent, APIGatewayProxyResponseEvent> {

    public APIGatewayProxyResponseEvent handleRequest(final APIGatewayProxyRequestEvent input, final Context context) {
        Map<String, String> headers = new HashMap<>();
        headers.put("Content-Type", "application/json");
        headers.put("X-Custom-Header", "application/json");

        APIGatewayProxyResponseEvent response = new APIGatewayProxyResponseEvent()
                .withHeaders(headers);
        
        // long leftLimit = 28L;
        // long rightLimit = 32L;
        // long n = leftLimit + (long) (Math.random() * (rightLimit - leftLimit));
        long n = 30L;
        String output = String.format("{ \"message\": \"hello world\", \"fibo\": \"%s\" : \"%s\" }",n, fibo(n));
        return response
                .withStatusCode(200)
                .withBody(output);
        
    }

    // Function for nth Fibonacci number
    private long fibo(long n) {
        if (n <= 2)
            return n - 1;
        else
        return fibo(n - 1) + fibo(n - 2);
    }
}
