# awsss - AWS Security Scanner

**awsss** is a lightweight and efficient AWS IAM role analyzer and privilege escalation risk scanner, built in Go with multithreading support for better performance. It is inspired by the excellent [PMapper](https://github.com/nccgroup/PMapper) project but addresses several limitations in the original Python implementation.

## Why Rewrite?

The original PMapper project was a great tool but hasn’t been updated in over **three years** and is stuck on **Python 3.8** 🤮. Additionally, PMapper only uses a **single thread** for its operations, making it slow for large environments. It is also **memory inefficient**, leading to performance issues during long-running scans in complex AWS environments.

To solve these issues, **awsss** was rewritten in Golang.

## Features

- **Multithreaded AWS Role and Privilege Escalation Analysis**
- **Efficient IAM Role and Trust Policy Parsing**
- **Support for Common AWS Services (EC2, IAM, Lambda, etc.)**
- **Graph Export in PNG or SVG Format**
- **Color-Coded Nodes**:
  - **Red** for administrator roles
  - **Blue** for IAM users
  - **White** for regular roles

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ProDefense/awsss
   ```
2. Build the project:
   ```bash
   cd awsss
   go build
   ```

## Usage

To generate a graph that analyzes privilege escalation paths in your AWS environment:

```bash
./awsss graph output -t png
```

You can specify whether the output should be in PNG or SVG format using the `-t` flag.

### Commands

- **`graph`**: Generate a graph of AWS IAM roles and their relationships.
- **`privesc`**: Analyze for potential privilege escalation paths.
- **`scan`**: Scan AWS accounts and roles for security issues.

## Acknowledgments

This tool is inspired by the [PMapper](https://github.com/nccgroup/PMapper) project created by **nccgroup**, which laid the foundation for AWS IAM role mapping and privilege escalation analysis.
