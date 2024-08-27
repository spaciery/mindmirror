# Security

## Shared responsibility model

AWS is responsible for security *of* the cloud and customer is responsible for security *in* the cloud. 

![https://d1.awsstatic.com/security-center/Shared_Responsibility_Model_V2.59d1eccec334b366627e9295b304202faf7b899b.jpg](https://d1.awsstatic.com/security-center/Shared_Responsibility_Model_V2.59d1eccec334b366627e9295b304202faf7b899b.jpg)

## **AWS Identity and Access Management (IAM)**

AWS Identity and Access Management (IAM) enables you to manage access to AWS services and resources securely.   

IAM gives you the flexibility to configure access based on your company’s specific operational and security needs. You do this by using a combination of IAM features

- IAM users, groups, and roles
- IAM policies
- Multi-factor authentication

**AWS account root user**

The root user is accessed by signing in with the email address and password that you used to create your AWS account. It has complete access to all the AWS services and resources in the account.

Best practice : 

Do **not** use the root user for everyday tasks.

Instead, use the root user to create your first IAM user and assign it permissions to create other users.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/LdLxWyqSptNR3S6F_G6mCtOEHNBDNNKV4.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/LdLxWyqSptNR3S6F_G6mCtOEHNBDNNKV4.png)

**IAM users**

An **IAM user** is an identity that you create in AWS. It represents the person or application that interacts with AWS services and resources. It consists of a name and credentials.

By default, when you create a new IAM user in AWS, it has no permissions associated with it. To allow the IAM user to perform specific actions in AWS, such as launching an Amazon EC2 instance or creating an Amazon S3 bucket, you must grant the IAM user the necessary permissions.

**IAM policies**

An **IAM policy** is a document that allows or denies permissions to AWS services and resources.

IAM policies enable you to customize users’ levels of access to resources. For example, you can allow users to access all of the Amazon S3 buckets within your AWS account, or only a specific bucket.

**IAM groups**

An IAM group is a collection of IAM users. When you assign an IAM policy to a group, all users in the group are granted permissions specified by the policy.

**IAM roles**

An IAM role is an identity that you can assume to gain temporary access to permissions.  

Before an IAM user, application, or service can assume an IAM role, they must be granted permissions to switch to the role. When someone assumes an IAM role, they abandon all previous permissions that they had under a previous role and assume the permissions of the new role. 

## **Multi-factor authentication**

In IAM, multi-factor authentication (MFA) provides an extra layer of security for your AWS account.

First, a user enters their IAM user ID and password to sign in to an AWS website.

Next, the user is prompted for an authentication response from their AWS MFA device. This device could be a hardware security key, a hardware device, or an MFA application on a device such as a smartphone.

When the user has been successfully authenticated, they are able to access the requested AWS services or resources.

## **AWS Organizations**

Suppose that your company has multiple AWS accounts. You can use [**AWS Organizations](https://aws.amazon.com/organizations)** to consolidate and manage multiple AWS accounts within a central location.

When you create an organization, AWS Organizations automatically creates a **root**, which is the parent container for all the accounts in your organization.

In AWS Organizations, you can centrally control permissions for the accounts in your organization by using [**service control policies (SCPs)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps.html)** . SCPs enable you to place restrictions on the AWS services, resources, and individual API actions that users and roles in each account can 
access.

**Organizational units**

In AWS Organizations, you can group accounts into organizational units (OUs) to make it easier to manage accounts with similar business or security requirements. When you apply a policy to an OU, all the accounts in the OU automatically inherit the permissions specified in the policy.

By organizing separate accounts into OUs, you can more easily isolate workloads or applications that have specific security requirements. For instance, if your company has accounts that can access only the AWS services that meet certain regulatory requirements, you can put these accounts into one OU. Then, you can attach a policy to the OU that blocks access to all other AWS services that do not meet the regulatory requirements.

## **AWS Artifact**

Depending on your company’s industry, you may need to uphold specific standards. An audit or inspection will ensure that the company has met those standards.

[**AWS Artifact](https://aws.amazon.com/artifact)**  is a service that provides on-demand access to AWS security and compliance reports and select online agreements. AWS Artifact consists of two main sections: AWS Artifact Agreements and AWS Artifact Reports.

**AWS Artifact Agreements**

Suppose that your company needs to sign an agreement with AWS regarding your use of certain types of information throughout AWS services. You can do this through **AWS Artifact Agreements**. 

**AWS Artifact Reports**.

suppose that a member of your company’s development team is building an application and needs more information about their responsibility for complying with certain regulatory standards. You can advise them to access this information in **AWS Artifact Reports**.

## **Customer Compliance Center**

The [**Customer Compliance Center**(opens in a new tab)](https://aws.amazon.com/compliance/customer-center/) contains resources to help you learn more about AWS compliance.

In the Customer Compliance Center, you can read customer compliance stories to discover how companies in regulated industries have solved various compliance, governance, and audit challenges.

You can also access compliance whitepapers and documentation on topics such as:

- AWS answers to key compliance questions
- An overview of AWS risk and compliance
- An auditing security checklist

## **AWS Shield**

**AWS Shield Standard** automatically protects all AWS customers at no cost. It protects your AWS resources from the most common, frequently occurring types of DDoS attacks.

As network traffic comes into your applications, AWS Shield Standard uses a variety of analysis techniques to detect malicious traffic in real time and automatically mitigates it.

**AWS Shield Advanced** is a paid service that provides detailed attack diagnostics and the ability to detect and mitigate sophisticated DDoS attacks.

It also integrates with other services such as Amazon CloudFront, Amazon Route 53, and Elastic Load Balancing. Additionally, you can integrate AWS Shield with AWS WAF by writing custom rules to mitigate complex DDoS attacks.

## **AWS Key Management Service (AWS KMS)**

You must ensure that your applications’ data is secure while in storage **(encryption at rest)** and while it is transmitted, known as **encryption in transit**.

[**AWS Key Management Service (AWS KMS)**(opens in a new tab)](https://aws.amazon.com/kms) enables you to perform encryption operations through the use of **cryptographic keys**. A cryptographic key is a random string of digits used for locking (encrypting) and unlocking (decrypting) data. You can use AWS KMS to create, manage, and use cryptographic keys. You can also control the use of keys across a wide range of services and in your applications.With AWS KMS, you can choose the specific levels of access control that you need for your keys. For example, you can specify which IAM users and roles are able to manage keys. Alternatively, you can temporarily disable keys so that they are no longer in use by anyone. Your keys never leave AWS KMS, and you are always in control of them

## **AWS WAF**

[**AWS WAF**(opens in a new tab)](https://aws.amazon.com/waf) is a web application firewall that lets you monitor network requests that come into your web applications.

AWS WAF works together with Amazon CloudFront and an Application Load Balancer. Recall the network access control lists that you learned about in an earlier module. AWS WAF works in a similar way to block or allow traffic. However, it does this by using a [**web access control list (ACL)](https://docs.aws.amazon.com/waf/latest/developerguide/web-acl.html)** to protect your AWS resources.

## **Amazon Inspector**

Amazon Inspector helps to improve the security and compliance of applications by running automated security assessments. It checks applications for security vulnerabilities and deviations from security best practices, such as open access to Amazon EC2 instances and installations of vulnerable software versions. 

## **Amazon GuardDuty**

[**Amazon GuardDuty**](https://aws.amazon.com/guardduty) is a service that provides intelligent threat detection for your AWS infrastructure and resources. It identifies threats by continuously monitoring the network activity and account behavior within your AWS environment.

After you have enabled GuardDuty for your AWS account, GuardDuty begins monitoring your network and account activity. You do not have to deploy or manage any additional security software. GuardDuty then continuously analyzes data from multiple AWS sources, including VPC Flow Logs and DNS logs. 

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/lFD5roorPinvKt-p_phaFiuF9QHWbEoD0.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1718002800/yDo4O5XSKlpJKXyrF7kAnQ/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/lFD5roorPinvKt-p_phaFiuF9QHWbEoD0.png)