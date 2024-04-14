import { array, option } from "fp-ts";
import { pipe } from "fp-ts/lib/function";
import { Errors } from "io-ts";
import { formatValidationErrors } from "io-ts-reporters";

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

function getPaths(validationErrors: Errors): Array<string> {
  return pipe(
    validationErrors,
    array.map((error) =>
      pipe(
        error.context.map(({ key }) => key),
        array.last,
        option.getOrElse(() => "")
      )
    )
  );
}

export const validationErrorsToString: (validationErrors: Errors) => string = (
  validationErrors
) => formatValidationErrors(validationErrors).join(", ");

export function getPathsErrorMessage(validationErrors: Errors): string {
  return `Fields missing or invalid: ${getPaths(validationErrors).join(", ")}`;
}
