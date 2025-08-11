/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.{htmx,html,js}",
    "./internal/**/*.templ",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
