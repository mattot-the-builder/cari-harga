/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./internal/app/views/*.html", "./node_modules/flowbite/**/*.js"],
	theme: {
		extend: {},
	},
	plugins: [
		require('flowbite/plugin')
	],
}

