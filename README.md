

![image](https://github.com/user-attachments/assets/da802864-6a8a-43de-9ed0-be80223ddee0)



# Go-HTMX Template Project

This is a comprehensive template project that combines Go, HTMX, Templ, Tailwind CSS, and Echo to create modern, efficient web applications. It's designed to provide a solid starting point for developers looking to leverage these technologies in their projects.

## Tech Stack

- **Go**: The backend language, providing robust and efficient server-side logic.
- **HTMX**: A lightweight JavaScript library that allows for dynamic HTML updates without writing JavaScript.
- **Templ**: A typed templating language for Go that compiles to native Go code.
- **Tailwind CSS**: A utility-first CSS framework for rapidly building custom user interfaces.
- **Echo**: A high-performance, minimalist Go web framework.

## Project Structure
```
go-htmx/ 
  ├── controllers 
    ├── api  
        ├── auth.go 
        ├── dashboard.go 
        └── profile.go
    ├── web 
        ├── auth.go  
        ├── dashboard.go 
        └── profile.go
  ├── css/ 
    ├── input.css  
    └── output.css 
  ├── routes
    ├── api
        ├── auth.go 
        ├── dashboard.go    
        └── profile.go 
    ├── web
        ├── auth.go 
        ├── dashboard.go
        └── profile.go
    ├── routes.go 
  ├── static
    ├── go-htmx-logo.png 
    └── templ.svg 
  ├── templates
    ├── common 
      ├── auth.templ 
      ├── dashboard.templ
      ├── login.templ
      ├── signup.templ 
      └── index.templ 
  ├── .air.toml 
  ├── go.mod 
  ├── go.sum 
  ├── main.go 
  ├── Makefile 
  ├── tailwind.config.js
```

### Directory Explanations

- **controllers/**: Contains the logic for handling requests, separated into API and web controllers.
- **css/**: Holds the Tailwind CSS input and compiled output files.
- **routes/**: Defines the routing structure for both API and web endpoints.
- **static/**: Stores static assets like images and SVGs.
- **templates/**: Houses all the Templ template files for rendering views.
- **.air.toml**: Configuration for Air, a live-reloading tool for Go applications.
- **go.mod & go.sum**: Go module files for dependency management.
- **main.go**: The entry point of the application.
- **Makefile**: Contains useful commands for building and running the project.
- **tailwind.config.js**: Configuration file for Tailwind CSS.

## Key Features

1. **Modular Structure**: The project is organized into logical components, making it easy to extend and maintain.
2. **HTMX Integration**: Enables dynamic content updates without full page reloads.
3. **Templ Templates**: Provides type-safe, efficient templating for Go.
4. **Tailwind CSS**: Allows for rapid UI development with utility classes.
5. **Echo Framework**: Offers a simple and performant web server implementation.
6. **API and Web Separation**: Clear separation between API and web routes for better organization.

## Getting Started

1. Clone this repository.
2. Run `go mod tidy` to install dependencies.
3. Install Air for live reloading: go install github.com/cosmtrek/air@latest
4. Use the following commands to build and run the project:
   - `make css`: Builds the Tailwind CSS file.
   -  `air` : Starts the application with live reloading using the configuration in .air.toml.

## Customization

This template is designed to be a starting point. Feel free to modify the structure, add new routes, controllers, and templates as needed for your specific project requirements.

## Contributing

Contributions to improve this template are welcome. Please feel free to submit issues or pull requests.

## License

[MIT License](LICENSE)
