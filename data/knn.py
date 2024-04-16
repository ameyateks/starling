from dataclasses import dataclass, asdict
from datetime import datetime, timedelta
import numpy as np
import json
import sys


@dataclass
class Transaction:
    spendingCategory: str
    amount: float
    transactionTime: datetime


def countUniqueSpendCats(input_list):
    categoryCounts = {}

    for tpl in input_list:
        string = tpl[0]

        categoryCounts[string] = categoryCounts.get(string, 0) + 1

    res = [(string, count) for string, count in categoryCounts.items()]

    return res


def knnAlgoToReturnClassifiedSpendingCategory(
    transaction: Transaction, allTransactions: list[Transaction], k: int
):
    def euclidianDistance(indexedTransaction: Transaction) -> tuple[str, float]:
        return (
            indexedTransaction.spendingCategory,
            np.sqrt(
                (transaction.amount - indexedTransaction.amount) ** 2
                + (
                    (
                        (
                            transaction.transactionTime.time().hour * 60
                            + transaction.transactionTime.time().minute
                        )
                        - indexedTransaction.transactionTime.time().hour * 60
                        + indexedTransaction.transactionTime.time().minute
                    )
                )
                ** 2
            ),
        )

    catsAndDists = list(map(euclidianDistance, allTransactions))

    catsAndDists.sort(key=lambda a: a[1])
    res = countUniqueSpendCats(catsAndDists[0:k])

    res.sort(key=lambda a: a[1], reverse=True)

    return res[0][0]


currentTransaction = json.loads(sys.argv[1])
file = open("/tmp/transactions.json", "r")

allTransactions = json.loads(file.read())

file.close()

parseTransaction = lambda x: Transaction(
    x["spendingCategory"],
    x["amount"],
    datetime.fromisoformat(x["transactionTime"]),
)
currentTransaction = parseTransaction(currentTransaction)
allTransactions = list(
    map(
        parseTransaction,
        allTransactions,
    )
)

print(
    json.dumps(
        {
            "category": knnAlgoToReturnClassifiedSpendingCategory(
                currentTransaction,
                allTransactions,
                np.round(np.sqrt(len(allTransactions))).astype(int),
            )
        }
    )
)
