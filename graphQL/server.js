// from https://www.youtube.com/watch?v=ZQL7tL2S0oQ&list=WL&index=2&t=345s
const express = require('express')
const app = express()
const expressGraphQL = require('express-graphql').graphqlHTTP
const {
    GraphQLSchema, // schema for query
    GraphQLObjectType, // for object typing
    // some graphQL types
    GraphQLString,
    GraphQLList,
    GraphQLInt,
    GraphQLNonNull // ensures the object returned is not null
} = require('graphql')

// dummy data
const authors = [
	{ id: 1, name: 'J. K. Rowling' },
	{ id: 2, name: 'J. R. R. Tolkien' },
	{ id: 3, name: 'Brent Weeks' }
]

const books = [
	{ id: 1, name: 'Harry Potter and the Chamber of Secrets', authorId: 1 },
	{ id: 2, name: 'Harry Potter and the Prisoner of Azkaban', authorId: 1 },
	{ id: 3, name: 'Harry Potter and the Goblet of Fire', authorId: 1 },
	{ id: 4, name: 'The Fellowship of the Ring', authorId: 2 },
	{ id: 5, name: 'The Two Towers', authorId: 2 },
	{ id: 6, name: 'The Return of the King', authorId: 2 },
	{ id: 7, name: 'The Way of Shadows', authorId: 3 },
	{ id: 8, name: 'Beyond the Shadows', authorId: 3 }
]

/* dummy query schema
const schema = new GraphQLSchema({
    // define the schema here
    // define query section ("the getting of data")
    query: new GraphQLObjectType({
        name: 'HelloWorld', // name of query
        fields: () => ({
            message: { 
                type: GraphQLString,
                resolve: () => 'Hello World' // tells graphQL where to get the information from
                } // an object that defines the type of message
        }) // the fields to be returned
    })
})
*/

// define our own custom object "BookType"
const BookType = new GraphQLObjectType({
    name: 'Book',
    description: 'This represents a book written by an author',
    fields: () => ({
        id: { type: GraphQLNonNull(GraphQLInt) },
        // do not need a resolve: as in dummy schema since the id will be automatically pulled from object
        name: { type: GraphQLNonNull(GraphQLString) },
        authorId: { type: GraphQLNonNull(GraphQLInt) },
        // what if we want to access the author object from the book query?
        author: {
            type: AuthorType,
            resolve: (book) => { // the author is inside the book type so we take in a parent parameter of book
                return authors.find(author => author.id === book.authorId) // match the author with the book author
            }
        }
    })
})

// define our own custom object "BookType"
const AuthorType = new GraphQLObjectType({
    name: 'Author',
    description: 'The author of a book',
    fields: () => ({
        id: { type: GraphQLNonNull(GraphQLInt) },
        name: { type: GraphQLNonNull(GraphQLString) },
        // what if we want to return the books of the author!
        books: { type : new GraphQLList(BookType),
            resolve: (author) => {
                return books.filter(book => book.authorId === author.id)
            }
        }
    })
})

// create a rootQuery type that can query all the data
const RootQueryType = new GraphQLObjectType({
    name: 'Query',
    description: 'Root Query', // shows up in docs
    fields: () => ({
        books: {
            type: new GraphQLList(BookType),
            description: 'List of all Books',
            resolve: () => books // normally would query a database
        }, // return a list of books
        // add a query for authors
        authors: {
            type: new GraphQLList(AuthorType),
            description: 'List of Authors',
            resolve: () => authors
        },
        // what if we just want to query a SINGLE book
        book: {
          type: BookType,
          description: "A single book",
          args: {
              id: { type : GraphQLInt}
          }, // graphiQL takes in a single parameter of the book id to return the book
          resolve: (parent, args) => books.find(book => book.id === args.id)
          // takes in arguments from graphiQL to return a single book
          // note the parent parameter doesnt matter here
        },
        author: {
          type: AuthorType,
          description: "A single author",
          args: {
              id: { type : GraphQLInt}
          }, // graphiQL takes in a single parameter of the book id to return the book
          resolve: (parent, args) => authors.find(author => author.id === args.id)
          // takes in arguments from graphiQL to return a single book
          // note the parent parameter doesnt matter here
        }
    })
})

// used for mutating the database data
const RootMutationType = new GraphQLObjectType({
    name: 'Mutation',
    description: 'Root Mutation',
    fields: () => ({
        // different mutation operations
        addBook: {
            type: BookType,
            description: 'Add a book',
            args: { // graphiQL args to pass as new book
                name: { type: GraphQLNonNull(GraphQLString)},
                authorId: { type: GraphQLNonNull(GraphQLInt)}
            },
            resolve: (parent, args) => {
                const book = { 
                    id: books.length + 1, // would be auto gen by db
                    name: args.name,
                    authorId: args.authorId
                }
                // add to our array (do the corresponding database action)
                books.push(book)
                return book
            }
        },
        addAuthor: {
            type: AuthorType,
            description: 'Add a author',
            args: { // graphiQL args to pass as new book
                name: { type: GraphQLNonNull(GraphQLString)}
            },
            resolve: (parent, args) => {
                const author = { 
                    id: authors.length + 1, 
                    name: args.name
                }
                authors.push(author)
                return author
            }
        }
    })
})

const schema = new GraphQLSchema({
    query: RootQueryType,
    mutation: RootMutationType
})

app.use('/graphql', expressGraphQL({ // check out the graphIQL interface at localhost:5000/graphql
    schema: schema, // gives schema
    graphiql: true // give user interface to access graphql without using Postman
}))
app.listen(5000, () => console.log('Server Running'))

// Notes about graphQL interface
/*
    can open docs to see all the possible queries
    -- Dummy Query --
    make a query with:
        query {
            message
        }
    returns:
    {
        "data": {
            "message": "Hello World"
        }
    }

    -- Query Books (list of books) --
    query {
        books {
            //fields...
            name,
            authorId
        }
    }

    -- Query Books and query their author --
    query {
        books {
            author {
                name
                // more fields of author...
            }
        }
    }

    -- Query Authors (list of authors)  and their books --
    query {
        authors {
            name 
            books {
                name,
                id
            }
        }
    }

    -- Query a single book based on id argument --
    query {
        book (id: 1) {
            name
            author {
                name
            }
        }
    }

    fields in types returns a function since it needs to let Book and Author Type find each other

    -- addBook Mutation --
    mutation {
        addBook(name: "New Book", authorId: 1) {
            // the fields we want to be returned of the book we added
            id
            name
        }
    }
*/