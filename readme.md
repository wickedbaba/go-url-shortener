# URL Shortener in Golang

I followed this [tutorial](https://www.eddywm.com/lets-build-a-url-shortener-in-go/) by Eddy WM.

To run this project, just clone this repo and 

1. Have redis [installed](https://redis.io/download/?ref=eddywm.com) and running in the background.
2. Run ```go run main.go``` in the terminal. This will start the server.
3. Then use any REST client (I used Postman) to hit the server up with a **POST request** with this JSON body - 

{   

    "long_url": "url-of-your-choice",

    "user_id" : "a-user-id"
}

e.g.- 

{

    "long_url": "https://stackoverflow.com/questions",
    
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"

}


4. Output - 

{
 
    "message": "Yay !!! short url created successfully :)",
 
    "short_url": "http://localhost:80805AXAg3iu"

}


