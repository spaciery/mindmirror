# Java

Resource : https://dev.java/learn/

## OOP Principles

1. Encapsulation
    
    It is the technique of bundling the data (variables) and the methods (functions) that operate on the data into a single unit called a class. Key principles of encapsulation are data hiding, access control
    
2. Inheritance
3. Polymorphism
It refers to the ability of a single interface to support multiple underlying forms 
Compile-time Polymorphism - Achieved through method overloading.
Runtime Polymorphism - Achieved through method overriding.
4. Abstraction 
Abstraction involves hiding the complex implementation details and showing only the essential features of an object. It focuses on the "what" an object does rather than "how" it does it.

## SOLID Principles

### Single responsibility principle

every class should have only a single responsibility , or a single purpose

example : lets say we have a class for `bankingServices` , the class contains `deposit`, `withdraw`, `sentOTP`, `addLoan` methods . this class wont satisfy SRP , so we need to create separate classes for `sentOTP` and `sentEmail` say `notificationService` and separate class for adding loan say `LoanService`

### Open and close principle

According to new requirements , classes should be open for extension but closed for modification.

example 1 :

lets say we have a class that can calculate area for shapes and the class have calc methods that take in circle and square, this doesn't satisfy OCP since lets say for ex we tried to calculate the area of a triangle we will have to modify the code.

example 2 : 
lets say we have a class for sending notification , this code will violate the OCP since if we need to add a new requirement to sent WhatsApp notification , we need to modify the code.

```java
public class NotificationService {

	public void setNotification(String medium) {
	
			if(medium.equals("email")) {
				// sent email
			}
			
			if(medium.equals("sms")) {
				// sent sms
			}
	}

}
```

So we need to design our code in such a way that we can extend it based on new requirement. Here’s one of making it 

create an interface for notification

```java
public interface NotificationService {

	public void setNotification();

}
```

create Separate implementation for each kind of notification

```java
public class EmailNotificationService implements NotificationService  {

	public void setNotification() {
				// sent email
	}

}
```

### Liskov’s Substitution principle

Derived classes must be completely substitutable for their base class, in other words lets say we have class A which is a subtype of class B , then if we replace class A with class B program behavior should not be affected.

### Interface segregation principle

Larger interface should be broken down into smaller ones so that implementation classes can only use what it need, we should not force clients to implement the methods they don't want.

### Dependency inversion principle

## Design patterns

A [design patter](https://www.geeksforgeeks.org/software-design-patterns/)n is a generic repeatable solution to a frequently occurring problem in 
software design that is used in software engineering. It isn’t a complete design that can be written in code right away.

Creational design patterns

- singleton
- Factory
- Abstract Factory
- Prototype
- Builder

Factory : 

Factory Method Design Pattern define an interface for creating an object, but let subclass decide which class to instantiate. Factory Method lets a class defer instantiation to subclass.

example : 

```java
public interface MotorVehicle {
    void build();
}

public class Motorcycle implements MotorVehicle {
    @Override
    public void build() {
        System.out.println("Build Motorcycle");
    }
}

public class Car implements MotorVehicle {
    @Override
    public void build() {
        System.out.println("Build Car");
    }
}

public abstract class MotorVehicleFactory {
    public MotorVehicle create() {
        MotorVehicle vehicle = createMotorVehicle();
        vehicle.build();
        return vehicle;
    }
    protected abstract MotorVehicle createMotorVehicle();
}

public class MotorcycleFactory extends MotorVehicleFactory {
    @Override
    protected MotorVehicle createMotorVehicle() {
        return new Motorcycle();
    }
}

public class CarFactory extends MotorVehicleFactory {
    @Override
    protected MotorVehicle createMotorVehicle() {
        return new Car();
    }
}
```