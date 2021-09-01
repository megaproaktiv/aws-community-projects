#!/usr/bin/env python
from constructs import Construct
from cdktf import App, TerraformStack
from imports.aws import SnsTopic


class MyStack(TerraformStack):
    def __init__(self, scope: Construct, ns: str):
        super().__init__(scope, ns)

        # define resources here
        SnsTopic(self, 'cdktf-phython-topic', display_name='sns-cdktf-python')



app = App()
MyStack(app, "cdktf-sns-python")

app.synth()
