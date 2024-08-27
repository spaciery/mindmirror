# Networking

A networking service that you can use to establish boundaries around your AWS resources is [**Amazon Virtual Private Cloud (Amazon VPC)**](https://aws.amazon.com/vpc/)

Amazon VPC enables you to provision an isolated section of the AWS Cloud. In this isolated section, you can launch resources in a virtual network that you define. Within a virtual private cloud (VPC), you can organize your resources into subnets. A **subnet** is a section of a VPC that can contain resources such as Amazon EC2 instances.

## **Internet gateway**

To allow public traffic from the internet to access your VPC, you attach an **internet gateway** to the VPC. An internet gateway is a connection between a VPC and the internet. You can think of an internet gateway as being similar to a doorway that customers use to enter the coffee shop. Without an internet gateway, no one can access the resources within your VPC.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/zUWapGL6CscxDBO-_NEblbQjD0vn0-pPU.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/zUWapGL6CscxDBO-_NEblbQjD0vn0-pPU.png)

## **Virtual private gateway**

To access private resources in a VPC, you can use a **virtual private gateway**.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/Qnh2PeTY4KJ1Bg4E_s8U3lQzEONXm1FMX.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/Qnh2PeTY4KJ1Bg4E_s8U3lQzEONXm1FMX.png)

## **AWS Direct Connect**

[**AWS Direct Connect](https://aws.amazon.com/directconnect/)** is a service that lets you to establish a dedicated private connection between your data center and a VPC.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/ICJ98A_JYucjg9bl_YdzRvczPABE_j-yV.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/ICJ98A_JYucjg9bl_YdzRvczPABE_j-yV.png)

## **Subnets**

In a VPC, **subnets** are separate areas that are used to group together resources.

**Public subnets** contain resources that need to be accessible by the public, such as an online store’s website.

**Private subnets** contain resources that should be accessible only through your private network, such as a database that contains customers’ personal information and order histories.

In a VPC, subnets can communicate with each other. For example, you might have an application  that involves Amazon EC2 instances in a public subnet communicating with databases that are located in a private subnet.

## **Network ACLs**

When a customer requests data from an application hosted in the AWS Cloud, this request is sent as a packet. A **packet** is a unit of data sent over the internet or a network.

It enters into a VPC through an internet gateway. Before a packet can enter into a subnet or exit from a subnet, it checks for permissions. These permissions indicate who sent the packet and how the packet is trying to communicate with the resources in a subnet.

The VPC component that checks packet permissions for subnets is a [**network access control list (ACL)**](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html)

A network ACL is a virtual firewall that controls inbound and outbound traffic at the subnet level.
Network ACLs perform **stateless** packet filtering. They remember nothing and check packets that cross the subnet border each way: inbound and outbound. 

Each AWS account includes a default network ACL. When configuring your VPC, you can use your account’s default network ACL or create custom network ACLs.

By default, your account’s default network ACL allows all inbound and outbound traffic, but you can modify it by adding your own rules. For custom network ACLs, all inbound and outbound traffic is denied until you add rules to specify which traffic to allow. Additionally, all network ACLs have an explicit deny rule. This rule ensures that if a packet doesn’t match any of the other rules on the list, the packet is denied.

## **Security groups**

A security group is a virtual firewall that controls inbound and outbound traffic for an Amazon EC2 instance.

By default, a security group denies all inbound traffic and allows all outbound traffic. You can add custom rules to configure which traffic should be allowed; any other traffic would then be denied.If you have multiple Amazon EC2 instances within the same VPC, you can associate them with the same security group or use different security groups for each instance. 
Security groups perform **stateful** packet filtering.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/TYGTiKs6Rex-XWXm_ha8um-1InZb0jryB.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/TYGTiKs6Rex-XWXm_ha8um-1InZb0jryB.png)

## **Domain Name System (DNS)**

Suppose that StreetFuse has a website hosted in the AWS Cloud. Customers enter the web address into their browser, and they are able to access the website. This happens because of **Domain Name System (DNS)** resolution. DNS resolution involves a customer DNS resolver communicating with a company DNS server.

You can think of DNS as being the phone book of the internet. DNS resolution is the process of translating a domain name to an IP address.

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/fuKyjNqTn5bpoPAr_AFBIDcTkYnYj1OOb.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717938000/5fUinWtdNCCY9VRVnNP54w/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/fuKyjNqTn5bpoPAr_AFBIDcTkYnYj1OOb.png)

## **Amazon Route 53**

[**Amazon Route 53](https://aws.amazon.com/route53)** is a DNS web service. It gives developers and businesses a reliable way to route end users to internet applications hosted in AWS.