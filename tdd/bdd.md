# What is BDD?
Behavior-Driven Development (BDD) is a software development process that encourages collaboration among developers, QA, and non-technical stakeholders.
- Enhances communication and understanding
- Focuses on the expected behavior of the software


# Key Concepts of BDD
- **Collaboration**: Involves all stakeholders
- **Common Language**: Uses Gherkin for writing tests
- **Executable Specifications**: Tests are automated and executable
- **Documentation**: Tests serve as living documentation

# The Gherkin Language
Gherkin is a business-readable, domain-specific language.
- Used for writing user stories and scenarios
- Follows a simple format

# Gherkin Syntax
- **Feature**: Describes the feature being tested
- **Scenario**: Describes a specific test case
- **Given**: Sets up the initial context
- **When**: Describes an action or event
- **Then**: Describes the expected outcome

Example:
```gherkin
Feature: User Registration

Scenario: Successful registration with Email/Password
  Given a user with username "testuser" and password "password"
  When the user registers with email and password
  Then the registration should be successful

