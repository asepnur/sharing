Feature: Account Registration and Authentication

  As an account holder,
  I want to be able to register and log in using different authentication methods,
  So that I can choose the method that is most convenient for me.

  Background:
    Given the application is running
    And the database is initialized

  # Scenario for Email/Password Registration
  Scenario: Successful registration with Email/Password
    Given an account with username "testaccount" and password "password"
    When the account registers with email and password
    Then the registration should be successful

    # Acceptance Criteria:
    # 1. The account is created in the database.
    # 2. The account receives a confirmation email.
    # 3. The account is marked as active in the system.

  # Scenario for Google Registration
  Scenario: Successful registration with Google
    Given an account with Google ID "google123"
    When the account registers with Google
    Then the registration should be successful

    # Acceptance Criteria:
    # 1. The account is created in the database.
    # 2. The account receives a confirmation message.
    # 3. The account is marked as active in the system.

  # Scenario for Facebook Registration
  Scenario: Successful registration with Facebook
    Given an account with Facebook ID "fb123"
    When the account registers with Facebook
    Then the registration should be successful

    # Acceptance Criteria:
    # 1. The account is created in the database.
    # 2. The account receives a confirmation message.
    # 3. The account is marked as active in the system.

  # Scenario for Duplicate Email Registration
  Scenario: Registration with existing username
    Given an account with username "testaccount" and password "password"
    And the account is already registered with email and password
    When the account tries to register again with email and password
    Then the registration should fail with an error "account already exists"

    # Acceptance Criteria:
    # 1. The system checks for existing usernames.
    # 2. If a duplicate username is found, an error message is displayed.
    # 3. The duplicate registration attempt is logged for auditing.

  # Scenario for Successful Email/Password Login
  Scenario: Successful login with Email/Password
    Given an account with username "testaccount" and password "password"
    And the account is registered with email and password
    When the account logs in with email and password
    Then the login should be successful

    # Acceptance Criteria:
    # 1. The system verifies the account's credentials.
    # 2. The account is granted access if the credentials are correct.
    # 3. A login success message is displayed to the account holder.

  # Scenario for Unsuccessful Email/Password Login
  Scenario: Unsuccessful login with incorrect password
    Given an account with username "testaccount" and password "wrongpassword"
    And the account is registered with email and password
    When the account tries to log in with the incorrect password
    Then the login should fail with an error "invalid credentials"

    # Acceptance Criteria:
    # 1. The system verifies the account's credentials.
    # 2. If the credentials are incorrect, an error message is displayed.
    # 3. The unsuccessful login attempt is logged for security monitoring.
