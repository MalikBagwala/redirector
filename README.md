# Redirector

## Motivation

So you want to share your resume at 10 different places only to have all of them break when you make a change to your resume?
Or you just changed your instagram handle and now all of your previous links are broken?

Enter redirector

the point of a service like this is to maintain a single record of all your links and the web app simply redirects it to the link, not only that but there is also support for closest word matching so accessing all the below paths will fetch the same result!, this is really helpful when you are trying to programatically access the links and your key does not match the route exactly.

`/insta`
`/inta`
`/instagram`
`/INSTAGRAM`
`/Ins_Tagram`

## Setup

1. populate `.env` file to load all the env variables (this is optional and environment variables can be loaded anyway you like)
2. Install all the packages
3. `go run main.go`

## Deployment

The application is deployed on [https://railway.app](railway)