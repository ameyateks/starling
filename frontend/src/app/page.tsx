"use client";
import {
  Chart,
  ArcElement,
  Tooltip,
  Legend,
  Title,
  LinearScale,
} from "chart.js";
import { Transactions } from "./_components/transactions";
import { Account } from "./_components/account";

export type StarlingAmount = { minorUnits: number; currency: string };

export default function Home() {
  Chart.register(ArcElement, Tooltip, Legend, Title, LinearScale);
  return (
    <main>
      <div className="flex flex-col items-center w-full h-full bg-backgroundBeige min-h-screen">
        <div className="flex text-darkGreen text-6xl pt-1 justify-center">
          Pocket Watcher
        </div>
        <div className="flex flex-col w-full h-96 md:px-24 lg:px-80">
          <div className="flex flex-row w-full h-full bg-beige rounded-md">
            {displayPanel("Account", <Account />)}
            {displayPanel("Transactions", <Transactions />)}
          </div>
        </div>
      </div>
    </main>
  );
}

function displayPanel(title: string, innerContent: JSX.Element): JSX.Element {
  return (
    <div className="flex flex-col space-y-2 p-5 w-1/2">
      <p className="font-heading font-bold">{title}</p>
      <div className="flex flex-col bg-darkGreen h-full rounded-md">
        {innerContent}
      </div>
    </div>
  );
}
