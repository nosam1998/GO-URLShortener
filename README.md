# GO-URLShortener

A URL Shortener API written in Go! Keep in mind that this is ONLY an API. 
You still need a frontend for user input (Unless you want to use postman).

Utilizes [gorilla/mux](https://www.github.com/gorilla/mux) and [mattn/go-sqlite3](https://www.github.com/mattn/go-sqlite3)

Main endpoints:
 - (`METHODS ALLOWED`) : `Endpoint`
 - (`GET`) : `/ShortenedURLName`
   - Redirects the user to the original URL retrieved from the DB
 - (`POST`) : `/api/shorten`
   - JSON POST data format
     - Example:
     ```json
       {
           "ShortUrlSlug": "myGithub",
           "OriginalUrl": "https://www.github.com/nosam1998"
       }
       ```
     - Format
     ```json
     {
         "ShortUrlSlug": "YourShortenedUrlName",
         "OriginalUrl": "TheUrlToRedirectTo"
     }
     ```