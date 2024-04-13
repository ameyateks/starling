export const priceFormatter2dp = new Intl.NumberFormat("en-GB", {
  style: "currency",
  currency: "GBP",
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

export function classNames(...classes: (string | undefined)[]): string {
  return classes.filter(Boolean).join(" ");
}

export function categoryToFormattedString(category: string): string {
  const categoryLowerCaseSpaceOut = category.toLowerCase().replace("_", " ");
  return categoryLowerCaseSpaceOut.replace(/\w\S*/g, function (txt) {
    return txt.charAt(0).toUpperCase() + txt.substring(1).toLowerCase();
  });
}
