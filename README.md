## About
This is a template for a web app using the following technologies:
- [Go](https://golang.org/) (backend)
- [Templ](https://templ.guide/) (templating engine)
- [Tailwind](https://tailwindcss.com/) (CSS framework)
- [HTMX](https://htmx.org/) (for making the frontend dynamic)
- [Typescript](https://www.typescriptlang.org/) (for having a better dev experience)
- [Parcel](https://parceljs.org/) (for bundling the frontend code)

I made this template because I wanted to have a web app that is easy to develop and that is fast. I also wanted to have a web app that is easy to deploy. This template is meant as a starting point to use these technologies together.

## Features
- Hot reloading for both the backend and the frontend
- Automatic bundling of the frontend code
- Automatic compilation of the templ files
- Automatic compilation of the typescript files

## Dependencies
- [WGO (for having a stable dev server)](https://github.com/bokwoon95/wgo)
- [Templ cli](github.com/a-h/templ/cmd/templ@latest)
- NodeJS (for tailwind and typescript)
- NPM (for tailwind and typescript)

To add WGO and templ to your PATH, run the following command:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
(I recommend adding this to your .bashrc or .zshrc)

## How to run
1. Install dependencies
```bash
make install_deps
```
2. Run the dev server
```bash
make dev
```

## Structure
The Go code is in the `app` folder. The frontend typescript is in the `app/ts` folder. The templ files are in the `app/templ` folder. The compiled typescript file is in the `public/js` folder.

You can just install packages from NPM and use them in your typescript code (`app/ts`). Parcel will bundle them automatically.

## Important for VSCode
For a better dev experience, install the following extensions:
- [templ-vscode](https://marketplace.visualstudio.com/items?itemName=a-h.templ)
- [Tailwind CSS IntelliSense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
- [HTMX Attributes (for autocomplete in templ files)](https://marketplace.visualstudio.com/items?itemName=CraigRBroughton.htmx-attributes)

And then add the following to your settings.json:
```json
"tailwindCSS.includeLanguages": {
    "templ": "html"
},
"emmet.includeLanguages": {
    "templ": "html"
}
```

## Caching in development
During development, open your network tab in your browser and disable caching. This will make sure that you always get the latest version of your files.