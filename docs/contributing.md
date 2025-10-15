---
layout: default
title: Contributing
description: Contributing guidelines for GoliteFlow
---

# Contributing to GoliteFlow

Thank you for your interest in contributing to GoliteFlow! This document provides guidelines and information for contributors.

## 📋 Table of Contents

- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Contributing Process](#contributing-process)
- [Code Style Guidelines](#code-style-guidelines)
- [Testing Guidelines](#testing-guidelines)
- [Documentation Guidelines](#documentation-guidelines)
- [Release Process](#release-process)

## Getting Started

### Prerequisites

- Go 1.19 or later
- Git
- Basic understanding of Go development
- Familiarity with YAML and workflow orchestration

### Development Setup

1. **Fork the Repository**

   ```bash
   # Fork on GitHub, then clone your fork
   git clone https://github.com/YOUR_USERNAME/goliteflow.git
   cd goliteflow
   ```

2. **Add Upstream Remote**

   ```bash
   git remote add upstream https://github.com/sintakaridina/goliteflow.git
   ```

3. **Install Dependencies**

   ```bash
   go mod tidy
   ```

4. **Verify Setup**
   ```bash
   go test ./...
   go build ./cmd/goliteflow
   ```

## Contributing Process

### 1. Create a Feature Branch

```bash
# Make sure you're on main and up to date
git checkout main
git pull upstream main

# Create your feature branch
git checkout -b feature/your-feature-name
```

### 2. Make Your Changes

- Write clean, idiomatic Go code
- Follow existing code style and patterns
- Add comprehensive tests
- Update documentation as needed

### 3. Test Your Changes

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/parser

# Format your code
go fmt ./...

# Lint your code
go vet ./...

# Build to ensure everything compiles
go build ./cmd/goliteflow
```

### 4. Commit Your Changes

```bash
# Stage your changes
git add .

# Commit with a descriptive message
git commit -m "feat: add new workflow validation feature

- Add validation for workflow dependencies
- Include comprehensive test coverage
- Update documentation with examples"
```

### 5. Push and Create Pull Request

```bash
# Push your branch
git push origin feature/your-feature-name

# Create a Pull Request on GitHub
```

## 📝 Commit Message Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

### Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Examples

```
feat(parser): add support for custom task timeouts

fix(executor): resolve memory leak in task runner

docs: update README with new examples

test(scheduler): add tests for cron validation
```

## Testing Guidelines

### Test Requirements

- **Unit Tests**: All new functionality must have unit tests
- **Integration Tests**: Test component interactions
- **Edge Cases**: Test error conditions and edge cases
- **Coverage**: Maintain or improve test coverage

### Test Structure

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected ExpectedType
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    validInput,
            expected: expectedOutput,
            wantErr:  false,
        },
        {
            name:     "invalid input",
            input:    invalidInput,
            expected: nil,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := FunctionName(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("FunctionName() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

### Test Data

- Use `testdata/` directory for test configuration files
- Keep test data minimal and focused
- Include both valid and invalid test cases

## Code Style Guidelines

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use `gofmt` for formatting
- Use `golint` for style checking
- Keep functions small and focused
- Use meaningful variable and function names

### Documentation

- Add comments for all exported functions, types, and variables
- Use complete sentences in comments
- Include examples for complex functions
- Update README.md for user-facing changes

### Error Handling

```go
// Good
if err != nil {
    return fmt.Errorf("failed to parse config: %w", err)
}

// Avoid
if err != nil {
    return err
}
```

### Logging

```go
// Use structured logging
logger.Info("workflow started", "workflow", workflowName, "tasks", len(tasks))

// Avoid
logger.Info("workflow started")
```

## 📖 Documentation Guidelines

### README Updates

When adding new features:

- Update the feature list
- Add usage examples
- Update installation instructions if needed
- Add new configuration options

### API Documentation

- Document all exported functions and types
- Include parameter descriptions
- Provide usage examples
- Update API reference when adding new functions

### Examples

- Add examples for new features
- Update existing examples if behavior changes
- Include both simple and complex examples
- Test all examples before submitting

## Bug Reports

### Before Submitting

1. Check existing issues to avoid duplicates
2. Try to reproduce the issue
3. Gather relevant information (OS, Go version, etc.)

### Bug Report Template

```markdown
**Bug Description**
A clear description of the bug.

**Steps to Reproduce**

1. Go to '...'
2. Run command '...'
3. See error

**Expected Behavior**
What you expected to happen.

**Actual Behavior**
What actually happened.

**Environment**

- OS: [e.g., Windows 10, macOS 12, Ubuntu 20.04]
- Go Version: [e.g., 1.19.3]
- GoliteFlow Version: [e.g., v1.0.0]

**Additional Context**
Any other relevant information.
```

## Feature Requests

### Before Submitting

1. Check existing feature requests
2. Consider if the feature aligns with project goals
3. Think about implementation complexity

### Feature Request Template

```markdown
**Feature Description**
A clear description of the feature.

**Use Case**
Why is this feature needed? What problem does it solve?

**Proposed Solution**
How would you like this feature to work?

**Alternatives Considered**
Other solutions you've considered.

**Additional Context**
Any other relevant information.
```

## Code Review Process

### For Contributors

1. **Self-Review**: Review your own code before submitting
2. **Address Feedback**: Respond to reviewer comments promptly
3. **Update Tests**: Add tests for any new functionality
4. **Update Documentation**: Keep docs in sync with code changes

### For Reviewers

1. **Be Constructive**: Provide helpful, specific feedback
2. **Test Changes**: Run tests and verify functionality
3. **Check Style**: Ensure code follows project guidelines
4. **Approve Promptly**: Don't let PRs sit without review

## Pull Request Guidelines

### PR Template

```markdown
## Description

Brief description of changes.

## Type of Change

- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing

- [ ] Tests pass locally
- [ ] New tests added for new functionality
- [ ] Manual testing completed

## Checklist

- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No breaking changes (or clearly documented)
```

### PR Requirements

- **Title**: Clear, descriptive title
- **Description**: Detailed description of changes
- **Tests**: All tests must pass
- **Documentation**: Update docs for user-facing changes
- **Breaking Changes**: Clearly document any breaking changes

## Release Process

### Version Numbering

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Release Checklist

- [ ] All tests pass
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
- [ ] Version bumped in go.mod
- [ ] Release notes prepared
- [ ] GitHub release created

## Community Guidelines

### Be Respectful

- Use welcoming and inclusive language
- Be respectful of differing viewpoints
- Accept constructive criticism gracefully
- Focus on what's best for the community

### Be Collaborative

- Help others when you can
- Share knowledge and experience
- Be patient with newcomers
- Work together toward common goals

## Getting Help

- **GitHub Issues**: For bugs and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Code Review**: Ask questions in PR comments

## Recognition

Contributors will be recognized in:

- CONTRIBUTORS.md file
- Release notes
- Project documentation

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Conventional Commits](https://www.conventionalcommits.org/)

---

<div class="next-steps">
  <h3>Ready to contribute?</h3>
  <p>Start by forking the repository and creating your first pull request.</p>
  <a href="https://github.com/sintakaridina/goliteflow" class="btn btn-primary">Fork on GitHub</a>
  <a href="/getting-started" class="btn btn-secondary">Getting Started</a>
</div>

<style>
.next-steps {
  text-align: center;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 1rem;
  margin: 2rem 0;
}

.btn {
  display: inline-block;
  padding: 0.75rem 1.5rem;
  margin: 0.5rem;
  border-radius: 0.5rem;
  text-decoration: none;
  font-weight: 600;
  transition: all 0.3s ease;
}

.btn-primary {
  background: #007bff;
  color: white;
}

.btn-primary:hover {
  background: #0056b3;
  transform: translateY(-2px);
}

.btn-secondary {
  background: transparent;
  color: #007bff;
  border: 2px solid #007bff;
}

.btn-secondary:hover {
  background: #007bff;
  color: white;
}
</style>
