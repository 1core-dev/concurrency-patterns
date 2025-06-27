# Cupcake Factory Pipeline in Go

This project demonstrates a classic **pipeline pattern** implemented in Go, modeling
a cupcake factory assembly line. It illustrates how breaking down a problem into
dependent stages—where each task depends on the completion of the previous one—can
be efficiently managed using concurrent pipelines and Go’s goroutines and channels.

---

## Overview

In many real-world problems, tasks must be executed sequentially but can be decomposed
into multiple stages, where the output of one stage becomes the input of the next.
This often appears in manufacturing (e.g., car assembly lines) or digital signal
processing (e.g., audio filters). Here, we use a cupcake baking scenario as a metaphor
for such pipelines.

The pipeline stages include:

1. **PrepareTray** – Prepare an empty tray.  
2. **Mixture** – Pour cupcake mixture into the tray.  
3. **Bake** – Bake the mixture (slowest step).  
4. **AddToppings** – Add toppings to the baked cupcake.  
5. **Box** – Box the finished cupcake.

Each stage simulates work by sleeping for a certain number of seconds.

---

## Key Concepts

### Pipeline Pattern

The pipeline pattern is a powerful concurrency design where data flows through a series
of processing stages connected by channels. Each stage runs concurrently, processes
inputs, and sends results downstream. This allows overlapping of work, improving
overall throughput.

- **Throughput**: How many items the system produces per unit time.  
- **Latency**: How long it takes one item to pass through the entire pipeline.

### Bottlenecks and Performance

- The slowest stage (here, **baking**) determines the pipeline’s throughput.  
- Speeding up faster stages has little effect if the bottleneck remains unchanged.  
- Improving the bottleneck or parallelizing it improves total output.

### Timing observations
- With the original `everyThingElseTime = 1` and `ovenTime = 5`, producing `10`
cupcakes takes around `58 seconds`.
- Doubling speed of non-baking steps barely affects total time because baking
remains the bottleneck.
- Reducing `ovenTime` to `2` seconds and setting `everyThingElseTime` to 
`2` seconds reduces total time to about `30` seconds, increasing throughput.

### Summary
- Pipeline throughput depends on slowest stage.
- Improving faster stages alone has minimal effect.
- Focusing on bottlenecks is key to system performance improvements.
- This mirrors real-life scenarios like manufacturing lines, where optimizing or adding capacity to the slowest step is critical.