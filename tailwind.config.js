/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    '/public/**/*.{html,js}',
    './public/views/*.{html,js}'
  ],
  theme: {
    extend: {
        fontFamily: {
          'vt323': ['"VT323"', 'monospace'],
          'pixelify': ['"Pixelify Sans"', 'sans-serif'],
          'jersey10': ['"Jersey 10"', 'sans-serif'],
          'modern-dos16': ['"ModernDOS8x16"', 'monospace'],
          'modern-dos14': ['"ModernDOS8x14"', 'monospace'],
        },
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



