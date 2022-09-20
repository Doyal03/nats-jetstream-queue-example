# nats-jetstream-queue-example
Another Nats Jetstream example, with queueing system.

Making two queues listen to the same stream will help distribute the incoming messages between them in a load-balanced way.
In the "QueueSubscribe" func, we pass the stream string as well as the "queue name" string. It is important that the two (or more than two, depending upon how many queues are required) queues have same queue name.

Here, in the sample example streaming below, the publisher is sending 100 messages at once on two subscribers with same queue name, and we can see the number of messages read by each are quite load balanced, i.e., one read 51 while other read 49 data.

![image](https://user-images.githubusercontent.com/91259249/191339139-ecb3da04-87f2-43b8-a043-a8f87d7c401d.png)

Note here, that timeout has occurred even if all the messages have been read by the two queues. This timeout is because all of the messages have not been ack'd in given alloted time - it does not comment upon the delivery of all the messages to the subscriber(s). 
