package demo.lambda;

import com.amazonaws.services.lambda.runtime.Context; 
import com.amazonaws.services.lambda.runtime.RequestHandler;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class Hello implements RequestHandler<Request, Response> {
    private static final Logger log = LogManager.getLogger(Hello.class);

    public Response handleRequest(Request request, Context context) {
        log.info("Hello this is an info message");
        String greetingString = String.format("Hello %s %s.", request.firstName, request.lastName);
        return new Response(greetingString);
    }
}