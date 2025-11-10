# Assignment 1: Mini Project: Exploring Modern Programming Languages and GO Fundamentals

[cite_start]This repository contains the solution for the "New age programming languages" (ENSP415) assignment[cite: 1, 3].

## 1. Overview of Modern Languages

[cite_start]This project explores four modern programming languages: GO, F#, Clojure, and Kotlin[cite: 26].

* **GO (Golang):** A statically typed, compiled language from Google, known for its simplicity, performance, and excellent concurrency support via goroutines and channels.
* **F#:** A functional-first, statically typed language running on the .NET platform, emphasizing immutability and data-oriented programming.
* **Clojure:** A dynamic, functional Lisp dialect that runs on the JVM (and other platforms), known for its simplicity, powerful macro system, and focus on immutable data structures.
* **Kotlin:** A statically typed language from JetBrains that runs on the JVM. It's fully interoperable with Java, concise, and features null safety, making it a popular choice for Android development.

## 2. Comparison with Traditional Languages

| Feature | Traditional (C, C++, Java) | Modern (GO, F#, Clojure, Kotlin) |
| :--- | :--- | :--- |
| **Concurrency** | Difficult, based on manual thread management and locks. | [cite_start]Built-in, simpler models (e.g., GO's goroutines)[cite: 8]. |
| **Syntax** | Often verbose (e.g., Java, C++). | [cite_start]Generally more concise and expressive[cite: 9]. |
| **Memory Mgmt.** | Manual (C/C++) or Garbage Collection (Java). | Automatic Garbage Collection (All four). |
| **Paradigm** | [cite_start]Primarily Object-Oriented or Procedural[cite: 13]. | [cite_start]Multi-paradigm, with strong functional support[cite: 8]. |

## 3. Environment Setup and Outputs

[cite_start]This section documents the installation of **GO** and **[Your Chosen Language]**.

### GO Installation
*(TODO: Add your documented steps for installing GO here)*
1.  ...
2.  ...

**GO "Hello World" Output:**
[cite_start]*(TODO: Add your screenshot for the GO 'Hello World' program here)* 
``

### [Your Chosen Language] Installation
*(TODO: Add your documented steps for installing F# / Clojure / Kotlin here)*
1.  ...
2.  ...

**[Language] "Hello World" Output:**
[cite_start]*(TODO: Add your screenshot for the other language's 'Hello World' program here)* 
` Hello World]`

## 4. Summary of GO Basics Implemented

[cite_start]All source code is located in the `Source Code Folder` (or root)[cite: 30].

* [cite_start]**Task 3 (Syntax & Data Types):** `task3_syntax.go` demonstrates variable declaration (`var`, `:=`), constants (`const`), basic types (int, float64, string, bool), and explicit type conversions[cite: 21, 29].
* [cite_start]**Task 4 (Control Structures):** `task4_control.go` implements `if-else` (including with a short statement), `switch` (with and without expressions), `for` loops (C-style, while-style), and `range` to iterate over a slice[cite: 21, 29].
* [cite_start]**Task 5 (Functions & Packages):** `task5_project/` directory shows how to create a custom function and organize code into a separate package (`custommath`) which is then imported and used by the main program[cite: 21, 29].
* [cite_start]**Task 6 (Arrays, Slices, Maps):** `task6_datastructures.go` shows the creation and manipulation of fixed-size arrays, flexible slices (including `append`, `copy`, and slicing), and key-value maps[cite: 21, 29].
* [cite_start]**Task 7 (Structs):** `task7_structs.go` defines a custom `Employee` struct, attaches methods to it (`Display`, `Deactivate`), and demonstrates how methods can be used to control data modification[cite: 21, 29].
* **Task 8 (Pointers & Memory):** `task8_pointers.go` explains Go's automatic garbage collection and demonstrates pointer usage (`&`, `*`). [cite_start]It includes a function (`updateValue`) that takes a pointer to modify the original variable's value[cite: 21, 29].

## 5. Program Execution Outputs

[cite_start]*(TODO: Add screenshots of your GO programs running, as required by the submission instructions)* 

**Task 3 Output:**
``

**Task 4 Output:**
``

**(Continue for Tasks 5-8)**