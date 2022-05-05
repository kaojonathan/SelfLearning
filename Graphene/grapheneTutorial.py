# tutorial from graphene docs https://docs.graphene-python.org/en/latest/quickstart/
from graphene import ObjectType, String, Schema

class Query(ObjectType):
    # defines a field 'hello' with single argument 'name'
    hello = String(name=String(default_value="stranger"))
    goodbye = String()

    # Resolver method that takes the GraphQL context (root, info) and Argument (name)
    def resolve_hello(root, info, name):
        return f'Hello {name}!'

    def resolve_goodbye(root, info):
        return 'Bye bye!'

schema = Schema(query=Query)

# query with default argument
query_string = '{ hello }'
result = schema.execute(query_string)
print(result.data['hello'])
# "Hello stranger!"

# query with arg
query_with_argument = '{ hello(name: "graphQL") }'
result = schema.execute(query_with_argument)
print(result.data['hello'])
# "Hello GraphQL!"
