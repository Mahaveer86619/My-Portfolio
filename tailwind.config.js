/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.tmpl", "./static/css/input.css"],
  theme: {
    extend: {
      keyframes: {
        'slide-in': {
          '0%': { transform: 'translateX(100%)', opacity: '0' },
          '100%': { transform: 'translateX(0)', opacity: '1' },
        },
      },
      animation: {
        'slide-in': 'slide-in 0.8s ease-out forwards',
      },
      scale: {
        '80': '.80',
      },
    },
  },
  plugins: [],
}