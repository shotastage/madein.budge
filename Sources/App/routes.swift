import Fluent
import Vapor

func routes(_ app: Application) throws {
    app.get { req in
        return req.view.render("index", ["title": "Hello Vapor!"])
    }

    app.get("hello") { req -> String in
        return "Hello, world!"
    }

    app.get("defualt-budge") { req -> String in
        return "Default, Budge!"
    }

    try app.register(collection: TodoController())
}
