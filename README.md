# Memory Card Digital Landing Page

A beautiful, animated landing page built with Svelte 5, SvelteKit, and GSAP for scroll animations. This project is statically generated for optimal performance.

## Features

- 🚀 Built with Svelte 5 and SvelteKit
- 🎨 Smooth scroll animations using GSAP and ScrollTrigger
- 📱 Fully responsive design
- ⚡ Static site generation for blazing-fast performance
- 🎯 Modern, gradient-based UI design
- 🔧 Easy to customize

## Tech Stack

- **Framework**: SvelteKit
- **UI Library**: Svelte 5
- **Animations**: GSAP with ScrollTrigger
- **Adapter**: @sveltejs/adapter-static (for SSG)
- **Build Tool**: Vite

## Getting Started

### Prerequisites

- Node.js 18+ installed on your machine
- npm or pnpm package manager

### Installation

1. Clone the repository:
```bash
git clone https://github.com/gustavodslara/svelte-memorycarddigital-landingpage.git
cd svelte-memorycarddigital-landingpage
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm run dev
```

4. Open your browser and visit `http://localhost:5173`

## Building for Production

To create a production build:

```bash
npm run build
```

The static site will be generated in the `build` directory, ready to be deployed to any static hosting service.

## Preview Production Build

After building, you can preview the production build:

```bash
npm run preview
```

## Project Structure

```
├── src/
│   ├── lib/
│   │   └── components/
│   │       ├── Navbar.svelte
│   │       ├── Hero.svelte
│   │       ├── Features.svelte
│   │       ├── About.svelte
│   │       ├── CTA.svelte
│   │       └── Footer.svelte
│   ├── routes/
│   │   ├── +layout.svelte
│   │   ├── +layout.js
│   │   └── +page.svelte
│   ├── app.html
│   ├── app.css
│   └── app.d.ts
├── static/
├── svelte.config.js
├── vite.config.js
└── package.json
```

## Customization

### Colors

Edit the CSS custom properties in `src/app.css` to customize the color scheme:

```css
:root {
	--primary-color: #6366f1;
	--secondary-color: #8b5cf6;
	--accent-color: #ec4899;
	/* ... */
}
```

### Content

Each component in `src/lib/components/` can be easily customized with your own content, images, and styling.

## Deployment

This project can be deployed to any static hosting service:

- **Vercel**: `vercel deploy`
- **Netlify**: Drag and drop the `build` folder
- **GitHub Pages**: Push the `build` folder to gh-pages branch
- **Cloudflare Pages**: Connect your repository

## License

MIT License - feel free to use this project for your own purposes.

## Author

Gustavo Lara

## Acknowledgments

- Svelte team for the amazing framework
- GreenSock for GSAP animations
- The open-source community