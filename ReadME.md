With this API you can consult all about restaurants, delete, insert and update.
This Code generate a log and consult a MongoDB.

- This API is used to test a MongoDB connection with GO.
- You have 2 EndPoint:
	- /Restaurants, With GET method you obtain all JSON restaurants, with POST method you insert a new Restaurant.
	- /Restaurants/{id} With GET method you obtain one JSON restaurant, with PUT method you update the existed Restaurant, with DELETE method you delete one Restaurant.
- The code is based on SOLID Go Structured.
- In this case i don't develop Unit testing, but i usually do that.


*"myAPI" run in Linux