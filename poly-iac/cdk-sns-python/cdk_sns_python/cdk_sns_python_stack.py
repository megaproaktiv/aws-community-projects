from aws_cdk import Stack
from constructs import Construct
from aws_cdk import aws_sns as sns

class CdkSnsPythonStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # The code that defines your stack goes here
        sns.Topic(self, "pythontopic")
