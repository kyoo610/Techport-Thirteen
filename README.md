To run:
"go build"
"go run main.go"

Endpoints:
- GET /api/listCookies
    - Retrieves a list of all cookies stored in the database
- POST /api/insertCookie
    - Creates a new Cookie based on the provided POST data
    - Returns an error if the cookie's Name is not given
- GET /api/deleteCookie/:id
    - Removes a cookie from the database based on its Mongo id
    Returns an error if no id is provided or the id is incorrect

The structure of a cookie 
- A cookie contains the following information:
- Name - The name of the cookie
- Quantity - The number of cookies of this type available 
- Last Baked - A timestamp containing when the last cookie batch was baked. Format: 1970-01-31 00:00:00, stored as a date object in Mongo
- Expiry - A timestamp of when the cookies will go bad. Format: 1970-01-31 00:00:00, stroed as a date object in Mongo 
- Price per cookie - The dollar amount of each cookie
- Description - A description of all the gooey deliciousness 