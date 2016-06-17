Feature: Basic Features for lodash-match-pattern

  Scenario: Check shell steps
    When this shell script runs
      """
      ls
      """
    Then STDOUT matched
      """
      features
      node_modules
      package.json
      """
    And the shell output is reset
    Then STDOUT matched
      """
      """
    And this shell script runs
      """
      pwd
      """
    Then STDOUT matched
      """
      /apitest/
      """

  Scenario: Check request steps
    Given the client gets "/"
    Then the response had status code "404"

    Given the database is reset
    Given the client gets "/users"
    And the response matched the pattern
      """
      []
      """
