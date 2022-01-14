FROM public.ecr.aws/lambda/java:11

COPY target/classes ${LAMBDA_TASK_ROOT}
COPY target/dependency/* ${LAMBDA_TASK_ROOT}/lib/

CMD [ "demo.lambda::handleRequest" ]