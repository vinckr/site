/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./content/**/*.{md,njk,html}", "./**/*.{njk,html}"],
  theme: {
    extend: {
      colors: {
        bg: "var(--bg)",
        "bg-2": "var(--secondary-bg)",
        accent: "var(--accent)",
        "accent-2": "var(--secondary-accent)",
        text: "var(--text)",
        "text-2": "var(--secondary-text)",
        "text-3": "var(--text-3)",
        "border-light": "var(--border-light)",
        "border-dark": "var(--border-dark)",
      },
      fontFamily: {
        pixel: ["ArialPixel", "system-ui", "sans-serif"],
      },
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
