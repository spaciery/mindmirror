# Monitoring

## **Amazon CloudWatch**

[**Amazon CloudWatch**](https://aws.amazon.com/cloudwatch/) is a web service that enables you to monitor and manage various metrics
 and configure alarm actions based on data from those metrics.

CloudWatch uses [**metrics**](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/working_with_metrics.html) to represent the data points for your resources. AWS services send 
metrics to CloudWatch. CloudWatch then uses these metrics to create graphs automatically that show how performance has changed over time.

The CloudWatch [**dashboard**](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Dashboards.html) feature enables you to access all the metrics for your resources from a
 single location. For example, you can use a CloudWatch dashboard to monitor the CPU utilization of an Amazon EC2 instance, the total number of requests made to an Amazon S3 bucket, and more. You can even customize separate dashboards for different business purposes, applications, or resources.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/waRne92bTSuHCL68_rLfuExSsSpuXOVqf.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/waRne92bTSuHCL68_rLfuExSsSpuXOVqf.png)

**CloudWatch alarms**

With CloudWatch, you can create [**alarms](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html)** that automatically perform actions if the value of your metric has gone above or below a predefined threshold. 

## **AWS CloudTrail**

[**AWS CloudTrail**](https://aws.amazon.com/cloudtrail/) records API calls for your account. The recorded information includes the identity of the API caller, the time of the API call, the source IP address of the API caller, and more. You can think of CloudTrail as a “trail” of breadcrumbs (or a log of actions) that someone has left 
behind them.

Recall that you can use API calls to provision, manage, and configure your AWS resources. With CloudTrail, you can view a complete history of user activity and API calls for your applications 
and resources.

Events are typically updated in CloudTrail within 15 minutes after an API call. You can filter events by specifying the time and date that an API call occurred, the user who requested the action, the type of resource that was involved in the API call, and more.

Within CloudTrail, you can also enable [**CloudTrail Insights](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html).** This optional feature allows CloudTrail to automatically detect unusual API activities in your AWS account. 

## **AWS Trusted Advisor**

[**AWS Trusted Advisor**](https://aws.amazon.com/premiumsupport/technology/trusted-advisor/) is a web service that inspects your AWS environment and provides real-time recommendations in accordance with AWS best practices.

Trusted Advisor compares its findings to AWS best practices in five categories: cost optimization, performance, security, fault tolerance, and service limits. For the checks in each category, Trusted Advisor offers a list of recommended actions and additional resources to learn more about AWS 
best practices.

The guidance provided by AWS Trusted Advisor can benefit your company at all stages of deployment. For example, you can use AWS Trusted Advisor to assist you while you are creating new workflows and developing new applications. You can also use it while you are making ongoing improvements to existing applications and resources.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/z2nlPeKUxP4VmCbC_2gGWhu3Np9FWdzHm.jpg](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/z2nlPeKUxP4VmCbC_2gGWhu3Np9FWdzHm.jpg)

For each category:

- The green check indicates the number of items for which it detected **no problems**.
- The orange triangle represents the number of recommended **investigations**.
- The red circle represents the number of recommended **actions**.