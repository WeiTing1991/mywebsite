/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
  theme: {
    container: {
      center: true,
    },
    extend: {
     fontFamily: {
        roboto: ["Roboto", "sans-serif"],
        robotoThin: ["Roboto Thin", "sans-serif"],
        robotoLight: ["Roboto Light", "sans-serif"],
        robotoMedium: ["Roboto Medium", "sans-serif"],
        robotoBlack: ["Roboto Black", "sans-serif"],
        ibmPlexSans: ["IBM Plex Sans", "sans-serif"],
      },
    },
  },
  plugins: [],
}

