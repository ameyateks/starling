"use client";
import React, { useState, useEffect, useMemo } from "react";
import { pipe } from "fp-ts/lib/function";
import * as O from "fp-ts/lib/Option";
import { apply, array, option, string } from "fp-ts";
import * as util from "../../utils";
import { Doughnut } from "react-chartjs-2";
import {
  Chart,
  ArcElement,
  Tooltip,
  Legend,
  Title,
  LinearScale,
} from "chart.js";
import * as ord from "fp-ts/lib/Ord";
import Ord = ord.Ord;
import Spinner from "./components/Spinner";
export type StarlingAmount = { minorUnits: number; currency: string };

type BalanceAndSpaces = {
  balance: StarlingAmount;
  spaces: {
    savingsGoals: {
      savingsGoalUid: string;
      name: string;
      target: StarlingAmount;
      totalSaved: StarlingAmount;
      savedPercentage: number;
      sortOrder: number;
      state: string;
    }[];
    spendingSpaces: {
      spaceUid: string;
      name: string;
      sortOrder: number;
      state: string;
      spendingSpaceType: string;
      cardAssociationUid: string;
      balance: StarlingAmount;
    }[];
  };
  transactions: any[];
};

type NameAndBalanceSpace = { name: string; balance: number };

export function Account() {
  Chart.register(ArcElement, Tooltip, Legend, Title, LinearScale);
  const [posts, setPosts] = useState<O.Option<BalanceAndSpaces>>(O.none);

  useEffect(() => {
    fetch("http://localhost:8080/api/accounts")
      .then((response) => response.json())
      .then((data) => {
        setPosts(O.some(data));
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  const spacesAndBalancesData = useMemo(
    () => pipe(posts, O.map(data)),
    [posts]
  );

  const spacesAndBalancesDoughnutData = useMemo(
    () => pipe(posts, O.map(spacesDoughnutData)),
    [posts]
  );

  return pipe(
    posts,
    O.match(
      () => (
        <div className="flex justify-center items-center h-full">
          <Spinner />
        </div>
      ),
      (spacesAndBalance) => (
        <div className="flex flex-col space-y-1 p-2 justify-between h-full">
          <h5 className="heading text-white font-bold ">Available</h5>
          <h5 className="heading text-[4rem] leading-none text-glass mt-1 ">
            {util.priceFormatter2dp.format(
              (spacesAndBalance?.balance.minorUnits ?? 0) / 100
            )}
          </h5>
          <div className="flex w-full justify-center pb-5">
            {pipe(
              {
                data: spacesAndBalancesData,
                doughnutData: spacesAndBalancesDoughnutData,
              },
              apply.sequenceS(option.Apply),
              O.match(
                () => <>Data unretrievable</>,
                (spacesAndBalancesData) => (
                  <Doughnut
                    data={spacesAndBalancesData.data}
                    options={{
                      plugins: {
                        legend: {
                          display: true,
                        },
                        tooltip: {
                          enabled: true,
                          callbacks: {
                            label: (context) => {
                              const calcLabel = () =>
                                `${
                                  pipe(
                                    spacesAndBalancesData.doughnutData.spaces,
                                    array.map((s) => s.name),
                                    array.prepend("Available")
                                  )[context.dataIndex]
                                }: ${context.dataset.data[context.dataIndex]}`;

                              return calcLabel();
                            },
                          },
                        },
                        title: {
                          display: false,
                          text: "Spaces Breakdown",
                          font: {
                            size: 24,
                          },
                        },
                      },
                      rotation: -90,
                      circumference: 180,
                      cutout: "40%",
                      maintainAspectRatio: false,
                      responsive: true,
                    }}
                  />
                )
              )
            )}
          </div>
        </div>
      )
    )
  );
}

const data = (spacesAndBalance: BalanceAndSpaces) => ({
  datasets: [
    {
      data: pipe(
        spacesDoughnutData(spacesAndBalance).spaces,
        array.map((s) => s.balance),
        array.prepend(spacesDoughnutData(spacesAndBalance).availableBalance)
      ),
      backgroundColor: pipe(
        generateDarkerHexCodes(
          spacesDoughnutData(spacesAndBalance).spaces,
          "#e9a3de"
        ),
        array.prepend("#CFC0BD")
      ),
      borderColor: "#FFFFFF",
      borderWidth: 1,
    },
  ],
});

function spacesDoughnutData(balanceAndSpaces: BalanceAndSpaces): {
  availableBalance: number;
  spaces: { name: string; balance: number }[];
} {
  const availableBalance = balanceAndSpaces.balance.minorUnits / 100;

  const spaces: { name: string; balance: number }[] = pipe(
    pipe(
      balanceAndSpaces?.spaces?.savingsGoals ?? [],
      array.map((s) => ({
        name: s.name,
        balance: s.totalSaved.minorUnits / 100,
      })),
      (t) => t
    ).concat(
      pipe(
        balanceAndSpaces?.spaces?.spendingSpaces ?? [],
        array.map((s) => ({
          name: s.name,
          balance: s.balance.minorUnits / 100,
        }))
      )
    ),
    array.sort(ordInstance)
  );

  return { availableBalance, spaces };
}

export const ordInstance: Ord<NameAndBalanceSpace> = ord.contramap(
  (nameAndBalanceSpace: NameAndBalanceSpace) => nameAndBalanceSpace.name
)(string.Ord);

function generateDarkerHexCodes(
  objects: NameAndBalanceSpace[],
  startColor: string
) {
  const hexToRgb = (hex: string) => {
    const bigint = parseInt(hex.slice(1), 16);
    return {
      r: (bigint >> 16) & 255,
      g: (bigint >> 8) & 255,
      b: bigint & 255,
    };
  };

  const rgbToHex = (r: number, g: number, b: number) =>
    `#${((1 << 24) + (r << 16) + (g << 8) + b).toString(16).slice(1)}`;

  const startRgb = hexToRgb(startColor);

  return objects.map((obj, index) => {
    const factor = (index + 1) * 0.1;
    const darkerRgb = {
      r: Math.max(0, Math.round(startRgb.r - startRgb.r * factor)),
      g: Math.max(0, Math.round(startRgb.g - startRgb.g * factor)),
      b: Math.max(0, Math.round(startRgb.b - startRgb.b * factor)),
    };

    const darkerHex = rgbToHex(darkerRgb.r, darkerRgb.g, darkerRgb.b);

    return darkerHex;
  });
}
