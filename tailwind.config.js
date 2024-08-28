/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    '/public/**/*.{html,js}',
    './public/views/*.{html,js}'
  ],
  theme: {
    extend: {
        animation: {
        border: 'background ease infinite',
        },
        keyframes: {
        background: {
            '0%, 100%': { backgroundPosition: '0% 50%' },
            '50%': { backgroundPosition: '100% 50%' },
        },
        },
    },
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
}



