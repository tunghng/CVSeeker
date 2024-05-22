/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                'primary': '#0D6EFD',
                'primary-hover': '#045adb',
                'primary-subtle': '#CFE2FF',
                'secondary-subtle': '#DEE2E6',
                'success-subtle': '#C7EDD9',
                'success-primary' : '#41AF79',
                'danger-subtle': '#FBBFBF',
                'danger-primary': '#DC3545',

                'disable-light': '#E9ECEF',
                'disable-dark': '#C2C7CD',

                'border': '#DEE2E6',

                'title': '#222222',
                'subtitle': '#878787',
                'text': '#444444',

                'background': '#F9F9F9',
            },
        },
    },
    plugins: [],
}

