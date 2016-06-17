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
    And the client gets "/users"
    Then the response had status code "200"
    And the response matched the pattern
      """
      []
      """
    Then the client posts to "/users" with query string
      """
      name=alec&age=19
      """
    Then the response had status code "200"
    Then the response matched the pattern
      """
      {
        Id: 1,
        Name: "alec",
        Age: 19
      }
      """
