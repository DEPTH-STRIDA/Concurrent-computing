# Concurrent Computing

This repository contains a coursework project that simulates concurrent computing. The program is implemented in Java, and you can find the coursework details here.

## Introduction

In computer programming, concurrency refers to the ability of a computer to handle multiple tasks simultaneously. For example, when using a web browser, multiple things happen at once: downloading a file, listening to music, and scrolling through a web page in another tab. This means the computer performs several tasks at the same time. Without concurrency, we would have to wait for the file to download completely before continuing to listen to music or scroll through the website.

The architecture of a CPU is such that a single core can perform only one action at a time. Today, there are still many single-core processors on the market. But even with a single-core processor, we can download a file, listen to music, and browse a news feed simultaneously. This is achieved through concurrency. Let's look at a diagram demonstrating how a single-core CPU handles the browser example.

![image](https://github.com/DEPTH-STRIDA/Concurrent-computing/assets/92984389/08efeef9-7008-4ee6-8d79-048b788bdfc6)

As we can see, all tasks are broken down into smaller pieces, which are prioritized, and the CPU constantly switches between them, creating the illusion of simultaneous execution.

## Program Description

We will create a program in Go that receives pre-divided tasks and executes them. The program will be implemented using an event-driven programming model. The main loop in the central class will select and process events, with events being generated in the loop.

### Main Loop

The main loop processes a list of events, to which new events are dynamically added.

### Event Types

There will be three types of events:
1. **Request a new task and grant access to the processor:** Initially, we have one event of this type. When processed, it adds a task (increases L), adds the next event of the first type (removing the processed one), and if the processor is free, it adds the second event (processing the newly added task). The time associated with the event indicates when it occurred, meaning it executes immediately.
2. **Task completion:** This event signifies the end of task processing and checks if there are any remaining tasks. If there are (L>0), it adds the next event of the second type. The time associated with the event indicates when the task ended.
3. **Program termination:** This event occurs when a pre-planned task is completed or the maximum number of tasks is exceeded.

We will implement this program to demonstrate how concurrency can be managed and simulated on a single-core processor.
