feature: friday

  Background:
    Given that the algorithm specifies 21 days between special fridays

  Scenario: Next Friday

    Given that I have a friday
      |Year|Month|Day|
      |2023|03|31|
    Given that this friday is for: "Dad"
    When I ask for the next friday
    Then the date is
      |Year|Month|Day|
      |2023|04|21|