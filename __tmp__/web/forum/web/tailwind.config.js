import daisyui from 'daisyui'

/** @type {import('tailwindcss').Config & {daisyui: import("daisyui").Config}} */
export default {
  darkMode: 'class',
  content: ['./index.html', './src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      container: {
        center: true
      }
    }
  },
  plugins: [daisyui],
  daisyui: {
    themes: ['business']
  }
}
