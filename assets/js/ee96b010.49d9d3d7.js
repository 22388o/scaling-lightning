"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[7459],{3905:(n,e,t)=>{t.d(e,{Zo:()=>l,kt:()=>h});var r=t(7294);function i(n,e,t){return e in n?Object.defineProperty(n,e,{value:t,enumerable:!0,configurable:!0,writable:!0}):n[e]=t,n}function o(n,e){var t=Object.keys(n);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(n);e&&(r=r.filter((function(e){return Object.getOwnPropertyDescriptor(n,e).enumerable}))),t.push.apply(t,r)}return t}function a(n){for(var e=1;e<arguments.length;e++){var t=null!=arguments[e]?arguments[e]:{};e%2?o(Object(t),!0).forEach((function(e){i(n,e,t[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(n,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(e){Object.defineProperty(n,e,Object.getOwnPropertyDescriptor(t,e))}))}return n}function s(n,e){if(null==n)return{};var t,r,i=function(n,e){if(null==n)return{};var t,r,i={},o=Object.keys(n);for(r=0;r<o.length;r++)t=o[r],e.indexOf(t)>=0||(i[t]=n[t]);return i}(n,e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(n);for(r=0;r<o.length;r++)t=o[r],e.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(n,t)&&(i[t]=n[t])}return i}var u=r.createContext({}),c=function(n){var e=r.useContext(u),t=e;return n&&(t="function"==typeof n?n(e):a(a({},e),n)),t},l=function(n){var e=c(n.components);return r.createElement(u.Provider,{value:e},n.children)},p="mdxType",g={inlineCode:"code",wrapper:function(n){var e=n.children;return r.createElement(r.Fragment,{},e)}},m=r.forwardRef((function(n,e){var t=n.components,i=n.mdxType,o=n.originalType,u=n.parentName,l=s(n,["components","mdxType","originalType","parentName"]),p=c(t),m=i,h=p["".concat(u,".").concat(m)]||p[m]||g[m]||o;return t?r.createElement(h,a(a({ref:e},l),{},{components:t})):r.createElement(h,a({ref:e},l))}));function h(n,e){var t=arguments,i=e&&e.mdxType;if("string"==typeof n||i){var o=t.length,a=new Array(o);a[0]=m;var s={};for(var u in e)hasOwnProperty.call(e,u)&&(s[u]=e[u]);s.originalType=n,s[p]="string"==typeof n?n:i,a[1]=s;for(var c=2;c<o;c++)a[c]=t[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,t)}m.displayName="MDXCreateElement"},5795:(n,e,t)=>{t.r(e),t.d(e,{assets:()=>u,contentTitle:()=>a,default:()=>g,frontMatter:()=>o,metadata:()=>s,toc:()=>c});var r=t(7462),i=(t(7294),t(3905));const o={sidebar_position:5},a="Running on github actions",s={unversionedId:"how-to/run-on-github-actions",id:"how-to/run-on-github-actions",title:"Running on github actions",description:"As a scaling lightning network can be setup and interacted with from code it lends itself to being usable for end to end or functional tests running on Github actions (or CI/CD tool of choice).",source:"@site/docs/how-to/run-on-github-actions.md",sourceDirName:"how-to",slug:"/how-to/run-on-github-actions",permalink:"/docs/how-to/run-on-github-actions",draft:!1,editUrl:"https://github.com/scaling-lightning/scaling-lightning/tree/main/docs/docs/how-to/run-on-github-actions.md",tags:[],version:"current",sidebarPosition:5,frontMatter:{sidebar_position:5},sidebar:"docsSidebar",previous:{title:"Running in the cloud",permalink:"/docs/how-to/run-in-cloud"}},u={},c=[],l={toc:c},p="wrapper";function g(n){let{components:e,...t}=n;return(0,i.kt)(p,(0,r.Z)({},l,t,{components:e,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"running-on-github-actions"},"Running on github actions"),(0,i.kt)("p",null,"As a scaling lightning network can be setup and interacted with from code it lends itself to being usable for end to end or functional tests running on Github actions (or CI/CD tool of choice)."),(0,i.kt)("p",null,"The following is an example github actions file that utilises minikube to run the cluster:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},'name: Scaling Lightning On MiniKube using GitHub Actions\n\n"on":\n  push:\n    branches:\n      - main\n\njobs:\n  run-sl:\n    runs-on: ubuntu-latest\n    steps:\n      - name: start minikube\n        id: minikube\n        uses: medyagh/setup-minikube@master\n      - name: Setup helmfile\n        uses: mamezou-tech/setup-helmfile@v1.2.0\n      - name: Install traefik CRDs\n        run: |\n          helm repo add traefik https://traefik.github.io/charts\n          helm repo update\n          helm install traefik traefik/traefik -n sl-traefik --create-namespace -f https://raw.githubusercontent.com/scaling-lightning/scaling-lightning/v0.0.33/charts/traefik-values.yml\n      - name: Start minikube\'s loadbalancer tunnel\n        run: minikube tunnel &> /dev/null &\n      - uses: actions/checkout@v3\n      - uses: actions/setup-go@v4\n        with:\n          go-version: ">=1.21.0"\n      - name: Run example test\n        run: go test -run ^TestMain$ github.com/scaling-lightning/scaling-lightning/examples/go -count=1 -v -timeout=15m\n')),(0,i.kt)("p",null,"The ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/scaling-lightning/scaling-lightning/blob/main/examples/go/example_test.go"},"example test")," can be found in our repo under examples/go."))}g.isMDXComponent=!0}}]);