export const priceFormatter2dp = new Intl.NumberFormat("en-GB", {
  style: "currency",
  currency: "GBP",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

export function classNames(...classes: (string | undefined)[]): string {
  return classes.filter(Boolean).join(" ");
}
