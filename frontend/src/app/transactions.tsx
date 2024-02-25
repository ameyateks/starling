"use client";
import Loading from "./loading";
import React, { useState, useEffect, Fragment, useMemo } from "react";
import { pipe } from "fp-ts/lib/function";
import * as O from "fp-ts/lib/Option";
import * as E from "fp-ts/lib/Either";
import * as A from "fp-ts/Array";
import { array, string } from "fp-ts";
import { classNames, priceFormatter2dp } from "../../utils";
export type StarlingAmount = { minorUnits: number; currency: string };
import { Dialog, Transition } from "@headlessui/react";
import {
  ChevronDoubleDownIcon,
  ChevronDoubleUpIcon,
} from "@heroicons/react/24/solid";
import * as t from "io-ts";

type Transaction = {
  feedItemUid: string;
  categoryUid: string;
  amount: StarlingAmount;
  sourceAmount: StarlingAmount;
  direction: "IN" | "OUT";
  updatedAt: string; //TODO: change to date
  transactionTime: string; //TODO: change to date
  source: string;
  sourceSubType: string;
  status: string;
  transactingApplicationUserUid: string;
  counterPartyType: string;
  counterPartyUid: string;
  counterPartyName: string;
  counterPartySubEntityUid: string;
  reference: string;
  country: string;
  spendingCategory: string;
  hasAttachment: boolean;
  hasReceipt: boolean;
};

type ClassifiedTransaction = {
  category: string;
};

export function Transactions() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  useEffect(() => {
    fetch("http://localhost:8080/api/transactions")
      .then((response) => response.json())
      .then((data) => {
        console.log("trans", data);
        setTransactions(data.transactions);
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  return !transactions.length ? (
    <div className="flex justify-center items-center h-full">
      <Loading />
    </div>
  ) : (
    <div className="flex flex-col space-y-1 p-2 h-full">
      <h5 className="heading text-white font-bold ">
        Transactions (Last 30 days)
      </h5>
      {transactions.length > 0 ? transactions.length : 0} transactions
      <div className="bg-beige rounded-md h-full p-2">
        <List transactions={transactions} />
      </div>
    </div>
  );
}

function List(props: { transactions: Transaction[] }): JSX.Element {
  let [isOpen, setIsOpen] = useState(O.none as O.Option<Transaction>);

  function closeModal() {
    setIsOpen(O.none);
  }

  function openModal(transaction: Transaction) {
    setIsOpen(O.some(transaction));
  }

  const [isLoading, setIsLoading] = useState(false);
  const [responseData, setResponseData] = useState<
    O.Option<ClassifiedTransaction>
  >(O.none);

  const handleClick = async (transaction: Transaction) => {
    setIsLoading(true);
    try {
      const response = await fetch("http://localhost:8080/api/knn", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(transaction),
      });

      const responseData = await response.json();
      pipe(
        responseData,
        t.type({ category: t.string }).decode,
        E.match(
          (err) => console.error(JSON.stringify(err)),
          (data) => setResponseData(O.some(data))
        )
      );
    } catch (error) {
      console.error("Error:", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="parent flex flex-col h-full max-h-[228px] bg-white rounded-md">
      <div className="child flex-1 border-b border-gray-300 p-1 overflow-y-auto space-y-1">
        {pipe(
          props.transactions,
          A.map((transaction) => (
            <>
              <button
                className="h-1/3 w-full flex flex-row text-red text-sm bg-backgroundBeige rounded-md"
                onClick={() => {
                  openModal(transaction);
                  handleClick(transaction);
                }}
              >
                <div className="flex flex-col justify-between p-2 w-2/3">
                  <div
                    className={classNames(
                      transaction.direction === "IN"
                        ? "text-darkGreen"
                        : "text-red",
                      "flex justify-items-start text-sm font-bold"
                    )}
                  >
                    {transaction.direction}
                  </div>
                  <div className="flex flex-col">
                    <div className="flex justify-items-start text-sm text-black truncate">
                      {transaction.counterPartyName}
                    </div>
                    <div className="flex justify-items-start text-sm text-black">
                      {pipe(
                        transaction.spendingCategory,
                        string.replace("_", " ")
                      )}
                    </div>
                  </div>
                </div>
                <div className="w-1/3 relative">
                  <div className="absolute top-2 right-2 text-black font-bold">
                    {priceFormatter2dp.format(
                      transaction.amount.minorUnits / 100
                    )}
                  </div>
                </div>
              </button>

              <Transition appear show={pipe(isOpen, O.isSome)} as={Fragment}>
                <Dialog as="div" className="relative z-10" onClose={closeModal}>
                  <Transition.Child
                    as={Fragment}
                    enter="ease-out duration-500"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in duration-200"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                  >
                    <div className="fixed inset-0 bg-transparent" />
                  </Transition.Child>

                  <div className="fixed inset-0 overflow-y-auto">
                    <div className="flex items-center pt-64 justify-center p-4 text-center relative">
                      <Transition.Child
                        as={Fragment}
                        enter="ease-out duration-300"
                        enterFrom="opacity-0 scale-95"
                        enterTo="opacity-100 scale-100"
                        leave="ease-in duration-200"
                        leaveFrom="opacity-100 scale-100"
                        leaveTo="opacity-0 scale-95"
                      >
                        {pipe(
                          isOpen,
                          O.match(
                            () => <></>,
                            (transaction) => (
                              <Dialog.Panel className="w-full max-w-sm transform overflow-hidden rounded-2xl bg-white text-left align-middle shadow-md transition-all">
                                <div className="flex flex-row w-full h-full p-6 space-x-4">
                                  <div className="flex flex-col w-3/5 h-full">
                                    <h3 className="text-lg font-medium leading-6 text-black">
                                      {transaction.counterPartyName}
                                    </h3>
                                    <div className="flex flex-col w-full">
                                      <div className="flex flex-row w-full">
                                        <div className="w-1/2 text-black">
                                          Direction:
                                        </div>
                                        <div className="flex w-1/2 text-black justify-end">
                                          {transaction.direction}
                                        </div>
                                      </div>
                                      <div className="flex flex-row w-full">
                                        <div className="w-1/2 text-black">
                                          Category:
                                        </div>
                                        <div className="flex w-1/2 text-black justify-end">
                                          {pipe(
                                            transaction.spendingCategory,
                                            string.replace("_", " ")
                                          )}
                                        </div>
                                      </div>
                                      <div className="flex flex-row w-full">
                                        <div className="w-1/2 text-black">
                                          Date of:
                                        </div>
                                        <div className="flex w-1/2 text-black justify-end">
                                          {transaction.transactionTime.substring(
                                            0,
                                            10
                                          )}
                                        </div>
                                      </div>
                                    </div>
                                  </div>
                                  <div className="w-2/5 text-black bg-lightGray rounded-md">
                                    <div className="flex flex-col justify-between" />
                                    <div className="h-1/3"></div>
                                    <div
                                      className={classNames(
                                        transaction.direction === "OUT"
                                          ? "text-red"
                                          : "text-darkGreen",
                                        "flex flex-row space-x-1 font-bold place-content-center place-items-center text-lg h-1/3"
                                      )}
                                    >
                                      {priceFormatter2dp.format(
                                        transaction.amount.minorUnits / 100
                                      )}
                                      {transaction.direction === "OUT" ? (
                                        <ChevronDoubleDownIcon className="h-5 w-5" />
                                      ) : (
                                        <ChevronDoubleUpIcon className="h-5 w-5" />
                                      )}
                                    </div>
                                    <div className="h-1/3"></div>
                                  </div>
                                </div>
                              </Dialog.Panel>
                            )
                          )
                        )}
                      </Transition.Child>
                    </div>
                  </div>
                </Dialog>
              </Transition>
            </>
          ))
        )}
      </div>
    </div>
  );
}
