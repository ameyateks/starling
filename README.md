# Starling Spend Tracker

This project was an exploration into Golang and using classification algorithms in a real world application.

I used the Starling API to query my spends monthly to create a dashboard highlighting monthly transactions as well as available money.

The more fascinating aspect of this project is being able to use the existing data from Starling and the KNN algorithm to assign a spending category under the assumption that my spending is habitual based on both timing and magnitude of purchase. We can then compare the knn assigned spending category with what Starling assigned and have the chance to overwrite or keep the category the same.

An end goal will be to save the transactions to a relational database and have a look at how my KNN classifier performs across all historic transactions.

Screenshots:
![The Main View](/frontend/MainView.png)
![The Transactions View](/frontend/TransactionView.png)
