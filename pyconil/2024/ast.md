## Title

Hacking the AST for fun and durable execution

## Abstract (250 characters)

We'll see how you can hack the AST and intercept function calls.
I'll showcase how to convert function to activities in termporal for durable execution.

## Description

In the harsh world of production, things fail from time to time.
Dealing with failure and retries can be hard.
For example: If a data pipeline fails step 4 out of 7, you'd like to resume the pipeline only from step 4.
Temporal is a system that allow you to create durable workflows with retry policies and can restart workflows from the middle.

In this talk I'll show how we allow users to run their code almost unchanged,
while under the hood we converting it to a temporal workflow.

This talk will focus on changing the AST, but will include discussion on design decisions we made along the way.
