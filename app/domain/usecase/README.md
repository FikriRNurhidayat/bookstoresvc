# Use Case Package

This package is being used to store the Use Case constructors.
Each use case should be defined by using struct with it's dependencies.

Also, use case function **MUST NOT** accept request payload and its context.
It should only accept defined interface so it can be less dependent.
