# Deployment Instructions

## GitHub Pages Deployment

Your site is now built and ready to deploy! The static files are in the `docs` folder.

### Steps to Deploy:

1. **Commit the docs folder**:
   ```bash
   git add docs
   git commit -m "Add static build for GitHub Pages"
   git push origin main
   ```

2. **Enable GitHub Pages**:
   - Go to your repository on GitHub: https://github.com/gustavodslara/svelte-memorycarddigital-landingpage
   - Click on **Settings**
   - Scroll down to **Pages** section (in the left sidebar)
   - Under **Source**, select **Deploy from a branch**
   - Under **Branch**, select:
     - Branch: `main`
     - Folder: `/docs`
   - Click **Save**

3. **Wait for deployment**:
   - GitHub will automatically build and deploy your site
   - After a few minutes, your site will be live at:
     - **https://gustavodslara.github.io/svelte-memorycarddigital-landingpage/**

### Testing Locally

The site is currently running on the preview server at:
- **http://localhost:4173/**

Navigation between pages (Home, Features, About, Contact) works with SvelteKit's client-side routing, so page transitions are smooth without full reloads.

### Rebuilding

Whenever you make changes to the source code, rebuild the site:

```bash
npm run build
```

This will regenerate the `docs` folder with your latest changes.

### Features

✅ Static site generation (SSG)
✅ Client-side navigation with SvelteKit
✅ PS2 Memory Card animations with GSAP
✅ Scroll-triggered animations on all pages
✅ Responsive design
✅ GitHub Pages ready
