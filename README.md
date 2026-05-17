# payment-system

Read Development instructions in [CONTRIBUTING.md](CONTRIBUTING.md)

## What has been done

The following decisions were taken
- Monorepo / Mono repository - a single repository to place the two services
- Single `go.mod` file
- Single `Dockerfile` file
- `Dockerfile.debug` file
- `COPY . .` container build instruction
- Include `go mod download` instruction

## Why

### Why Monorepo?

For simplicity. If it becomes complex to handle a monorepo, we can split into separate repositories

Monorepo can also help with sharing common functionalities / features in the form of libraries (modules) without having to have multiple repositories and versioning them and releasing them for every change in the modules

### Why single `go.mod` file?

For simplicity. If it becomes complex to handle a single `go.mod` file to manage the dependencies for both the services, then we can use separate `go.mod` files to manage the dependencies for both the services

### Why single `Dockerfile` file?

For simplicity. As of now, it works for both the services as a common `Dockerfile` as the container build instructions are the same for both the services. When the container build instructions differ in such a way that it requires two `Dockerfile`s, then we can create them. Or if there's any other need that's not satisfied by the current single `Dockerfile`, then we can create separate `Dockerfile`s

### Why `Dockerfile.debug` file?

To be able to debug anything related to the `Dockerfile` or the container build. For example - if we want to understand what was sent as part of build context, then we can use it, so that we can debug if `.dockerignore` is also working correctly. Also, we use `COPY . .` as a container build instruction - so, it's important to be able to find out what's sent as part of build context

## Why `COPY . .` container build instruction?

For simplicity. It's known that it's generally a bad idea to use this instruction, yes. The other way would be to explicitly mention the files and directories to be copied from the build context using `COPY`. But it can get pretty cumbersome to include any new files or directories every time there is a change - and hence sometimes developer might miss including the necessary files in the `COPY` instruction, causing an error. `COPY . .` has been done so as to make it easier for the developer. The downside is - the build context size has to be taken care of and for this we have `.dockerignore`. So, the developer still has to be mindful of `.dockerignore` and ignore any unnecessary files that may increase the size of the container image or pollute the container image unnecessarily. So, either the developer has to be mindful to add files in `COPY` instruction if `COPY . .` is not used. Or the developer has to be mindful of ignoring unnecessary files in the `.dockerignore` file. In both cases, a miss from the developer's side can possibly lead to issues - for example - not ignoring unnecessary files in `.dockerignore` if using `COPY . .`, or not including necessary files in `COPY` if not using `COPY . .`. And both can lead to errors / issues - but yes, these issues can be caught easily in pre-prod environments. I feel there will be more possibility problems when necessary files are not included as part of `COPY` when not using `COPY . .`, and there's less possibility of problems when unnecessary files are not ignored as part of `.dockerignore`. Hence this decision

### Why include `go mod download` instruction?

Though `go build` downloads dependencies, it's better to do `go mod download` before hand separately - for example in container build instructions so that we can leverage and take advantage of the container image layer caching - which won't change if there are no changes in the dependencies

But yes, also note that - `go mod download` downloads all modules defined in the `go.mod` while `go build` only downloads specific dependencies needed to compile the given package. But we use `go mod download` for efficiency and speed as part of the container build process which can leverage caching and speed up the build when there are no changes in the list of dependencies

## How to run it

Check [CONTRIBUTING.md](CONTRIBUTING.md)
