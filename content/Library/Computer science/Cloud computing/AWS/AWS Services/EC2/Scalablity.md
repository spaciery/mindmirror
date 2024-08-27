# Scalablity

**Scalability** involves beginning with only the resources you need and designing your architecture to automatically respond to changing demand by scaling out or in. As a result, you pay for only the resources you use. You don’t have to worry about a lack of computing capacity to meet your customers’ needs.

AWS service that provides this functionality for Amazon EC2 instances is **Amazon EC2 Auto Scaling**.

Within Amazon EC2 Auto Scaling, you can use two approaches: dynamic scaling and predictive scaling.

- *Dynamic scaling* responds to changing demand.
- *Predictive scaling* automatically schedules the right number of Amazon EC2 instances based on predicted demand.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717920000/PwqFIaIqtDDpydwytif-gw/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/ycQqvQ-dNtZo8148_f8VZ-ZFC2TOC7k5B.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717920000/PwqFIaIqtDDpydwytif-gw/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/ycQqvQ-dNtZo8148_f8VZ-ZFC2TOC7k5B.png)

The **minimum capacity** is the number of Amazon EC2 instances that launch immediately after you have created the Auto Scaling group.

Next, you can set the **desired capacity** at two Amazon EC2 instances even though your application needs a minimum of a single Amazon EC2 instance to run. by default this will be equal to minimum capacity.