# Assignment 1: Mini Project: Exploring Modern Programming Languages and GO Fundamentals

This repository contains the solution for the "New age programming languages" (ENSP415) assignment.

## 1. Overview of Modern Languages

This project explores four modern programming languages: GO, F#, Clojure, and Kotlin.

* **GO (Golang):** A statically typed, compiled language from Google, known for its simplicity, performance, and excellent concurrency support via goroutines and channels.
* **F#:** A functional-first, statically typed language running on the .NET platform, emphasizing immutability and data-oriented programming.
* **Clojure:** A dynamic, functional Lisp dialect that runs on the JVM (and other platforms), known for its simplicity, powerful macro system, and focus on immutable data structures.
* **Kotlin:** A statically typed language from JetBrains that runs on the JVM. It's fully interoperable with Java, concise, and features null safety, making it a popular choice for Android development.

## 2. Comparison with Traditional Languages

| Feature | Traditional (C, C++, Java) | Modern (GO, F#, Clojure, Kotlin) |
| :--- | :--- | :--- |
| **Concurrency** | Difficult, based on manual thread management and locks. | Built-in, simpler models (e.g., GO's goroutines). |
| **Syntax** | Often verbose (e.g., Java, C++). | Generally more concise and expressive. |
| **Memory Mgmt.** | Manual (C/C++) or Garbage Collection (Java). | Automatic Garbage Collection (All four). |
| **Paradigm** | Primarily Object-Oriented or Procedural. | Multi-paradigm, with strong functional support. |

## 3. Environment Setup and Outputs

This section documents the installation of **GO** and **F#**.

### GO Installation
1. Download GO from the official website: https://golang.org/dl/
2. Run the installer and follow the setup wizard
3. Add GO to system PATH:
    - Windows: Automatically added during installation
    - Linux/Mac: Add `export PATH=$PATH:/usr/local/go/bin` to ~/.bashrc
4. Verify installation by opening a terminal and running:
    ```bash
    go version
    ```
5. Create a workspace directory:
    ```bash
    mkdir go-workspace
    cd go-workspace
    ```
6. Initialize a new module:
    ```bash
    go mod init hello
    ```
7. Create and run first program:
    ```bash
    echo 'package main\n\nfunc main() {\n\tprintln("Hello, World!")\n}' > hello.go
    go run hello.go
    ```

**GO "Hello World" Output:**

![GO Hello World Output](./images/Screenshot%20(17).png)

The above screenshot shows the successful execution of the GO "Hello World" program displaying "Hello, World!" in the terminal.

### F# Installation
1. Download F# from the official website: https://dotnet.microsoft.com/en-us/
2. Install the .NET SDK which includes F# support
3. Verify installation by opening a terminal and running:
    ```bash
    dotnet --version
    ```
4. Create a new F# console project:
    ```bash
    dotnet new console -lang F# -n fsharp-hello
    cd fsharp-hello
    ```
5. Replace the Program.fs with the Hello World example:
    ```fsharp
    printfn "Hello, World!"
    ```
6. Run the program:
    ```bash
    dotnet run
    ```

**F# "Hello World" Output:**

![F# Hello World Output](./images/fsharp_hello_world.png)

The above screenshot shows the successful execution of the F# "Hello World" program displaying "Hello, World!" in the terminal.

## 4. Summary of GO Basics Implemented

All source code is located in the `Source Code Folder` (or root).

* **Task 3 (Syntax & Data Types):** `task3_syntax.go` demonstrates variable declaration (`var`, `:=`), constants (`const`), basic types (int, float64, string, bool), and explicit type conversions.
* **Task 4 (Control Structures):** `task4_control.go` implements `if-else` (including with a short statement), `switch` (with and without expressions), `for` loops (C-style, while-style), and `range` to iterate over a slice.
* **Task 5 (Functions & Packages):** `task5_project/` directory shows how to create a custom function and organize code into a separate package (`custommath`) which is then imported and used by the main program.
* **Task 6 (Arrays, Slices, Maps):** `task6_datastructures.go` shows the creation and manipulation of fixed-size arrays, flexible slices (including `append`, `copy`, and slicing), and key-value maps.
* **Task 7 (Structs):** `task7_structs.go` defines a custom `Employee` struct, attaches methods to it (`Display`, `Deactivate`), and demonstrates how methods can be used to control data modification.
* **Task 8 (Pointers & Memory):** `task8_pointers.go` explains Go's automatic garbage collection and demonstrates pointer usage (`&`, `*`). It includes a function (`updateValue`) that takes a pointer to modify the original variable's value.

## 5. Program Execution Outputs

### Task 3 Output: Syntax & Data Types

![Task 3 Output](./images/Screenshot%20(19).png)

The output demonstrates variable declarations using both `var` and `:=` operators, constant definitions, display of different data types (int, float64, string, bool), and type conversion examples.

### Task 4 Output: Control Structures

![Task 4 Output](./images/task4_control_output.png)

The output shows the execution of if-else statements, switch cases with and without expressions, various for loop patterns (C-style, while-style, and range-based), and iteration over slices.

### Task 5 Output: Functions & Packages

![Task 5 Output](./images/task5_functions_output.png)

The output demonstrates the successful import and use of the custom `custommath` package, showing function calls from the imported package working correctly with the main program.

### Task 6 Output: Arrays, Slices, Maps

![Task 6 Output](./images/task6_datastructures_output.png)

The output displays array creation and access, slice operations (append, copy, slicing), and map usage with key-value pairs and iteration.

### Task 7 Output: Structs

![Task 7 Output](./images/task7_structs_output.png)

The output shows struct instantiation, method calls on struct instances (Display, Deactivate), and the controlled modification of struct fields through methods.

### Task 8 Output: Pointers & Memory

![Task 8 Output](./images/task8_pointers_output.png)

The output demonstrates pointer declaration and dereferencing, passing pointers to functions for value modification, and showcasing Go's automatic garbage collection in action.

## 6. Conclusion

This assignment successfully explores GO fundamentals and compares modern programming languages with traditional ones. The hands-on implementation of syntax, data types, control structures, functions, data structures, structs, and pointers provides a comprehensive understanding of GO's core concepts and its advantages in modern software development.