import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    colors: {
      darkGreen: "#586F6B",
      lightGreen: "#7F9183",
      beige: "#B8B8AA",
      lightBeige: "#CFC0BD",
      backgroundBeige: "#DDD5D0",
    },
  },
  plugins: [],
};
export default config;
