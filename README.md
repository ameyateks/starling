# Starling Spend Tracker

This project was an exploration into Golang and using classification algorithms in a real world application.

I used the Starling API to query my spends monthly to create a dashboard highlighting monthly transactions as well as available money.

The more fascinating aspect of this project is being able to use the existing data from Starling and the KNN algorithm to assign a spending category under the assumption that my spending is habitual based on both timing and magnitude of purchase. We can then compare the knn assigned spending category with what Starling assigned and have the chance to overwrite or keep the category the same.

You can set this up too!

- Obtain your personal token [here](https://developer.starlingbank.com/login?next=/personal/token)
- Set it in the the backend .env file under ACCESS_TOKEN
- Set IS_DEMO to false
- Then run the backend using `go run .` and the frontend using `npm run build && npm run start` (alternatively you can use `npm run dev` for hot module replacement if you are making changes).

Further changes to expect:
Saving transactions to the database via a cron job. This will allow for more heavy duty analysis to be carried out on the data without constantly having to call the Starling API.
With transactions stored in the database an interesting project I inted to embark on will be running time series modelling, specifically Autoregressive Integrated Moving Average to forecast upcoming spending given historic monthly spending.

Screenshots:
![The Main View](/frontend/MainView.png)
![The Transactions View](/frontend/TransactionView.png)
