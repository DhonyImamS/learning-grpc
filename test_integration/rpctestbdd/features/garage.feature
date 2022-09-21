Feature: GRPC Integration Garage RPC using BDD

    In Order to fetch List Garage RPC we need to register New Garage First

    Scenario: Fetch List Garage RPC should be succeed return registered data
        Given Client set data for new Garage
        When Client registered new Garage
        Then Client will be able to retrieve new Garage data