# Selecting a Region

When determining the right Region for your services, data, and applications, consider the following four business factors. 

## Compliance

Depending on your company and location, you might need to run your data out of specific areas. For example, if your company requires all of its data to reside within the boundaries of the UK, you would choose the London Region. 

## Proximity

Selecting a Region that is close to your customers will help you to get content to them faster. For example, your company is based in Washington, DC, and many of your customers live in Singapore. You might consider running your infrastructure in the Northern Virginia Region to be close to company headquarters, and run your applications from the Singapore Region.

## Features

Sometimes, the closest Region might not have all the features that you want to offer to customers. AWS is frequently innovating by creating new services and expanding on features within existing services. However, making new services available around the world sometimes requires AWS to build out physical hardware one Region at a time.

Suppose that your developers want to build an application that uses Amazon Braket (AWS quantum computing platform). As of this course, Amazon Braket is not yet available in every AWS Region around the world, so your developers would have to run it in one of the Regions that already offers it.

## Pricing

Suppose that you are considering running applications in both the United States and Brazil. The way Brazil’s tax structure is set up, it might cost 50% more to run the same workload out of the São Paulo Region compared to the Oregon Region. You will learn in more detail that several factors determine pricing, but for now know that the cost of services can vary from Region to Region.

## **Availability Zones**

![https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717920000/PwqFIaIqtDDpydwytif-gw/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/MXfZDFcx-tg6KT4q_zBL1ijNGsMMu3-4j.png](https://explore.skillbuilder.aws/files/a/w/aws_prod1_docebosaas_com/1717920000/PwqFIaIqtDDpydwytif-gw/tincan/c8957d640c3a0f679bf6e6bb7f125de41ac6ad45/assets/MXfZDFcx-tg6KT4q_zBL1ijNGsMMu3-4j.png)

An **Availability Zone** is a single data center or a group of data centers within a Region. Availability Zones are located tens of miles apart from each other. This is close enough to have low latency 
(the time between when content requested and received) between Availability Zones. However, if a disaster occurs in one part of the Region, they are distant enough to reduce the chance that multiple Availability Zones are affected.