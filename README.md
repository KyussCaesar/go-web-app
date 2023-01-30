# go-web-app

# Development

```bash
make setup
```

If your dev env breaks, try re-run ^that.
If it's still broken try `make clean` first then try `make setup` again.

# Structure/Conventions

I try to document my thinking about the structure and conventions here.

## `src/`: application source code

A lot of Go projects that I've seen don't have such a directory and just put
source code directly in the top-level.
Personally I prefer having a top-level `src/` directory and all code kept
underneath it.

The reason I prefer having the `src/` directory is because the top-level of
projects tends to get cluttered with all sort of other shit; e.g. `Dockerfile`,
`docker-compose.yml`, CI/CD configuration, `Makefile`, `README.md`... adding
source code in there is even more messy IMO.

## `src/main.go`: entrypoint

`main.go` is the entrypoint of the application.

Its only responsibility is to create the application object and run it.

## `github.com/google/wire`: Dependency Injection Framework

`wire` is a code-generation tool that helps us with dependency-injection.

`wire` defines some fundamental concepts which are described here at a high level:

- a "Provider" is a Go function that `wire` can use to fulfill a dependency.
- a "Set" is a group of Providers.
- an "Injector" resolves the dependencies an object using the Providers and Sets that you specify.

https://github.com/google/wire

### Random notes about how I use `wire`

*On the structure/conventions of source code files for Providers and Sets:* Each package defines its Providers and Sets in `*.wire-provider.go`.

This is just a convention I invented, it's not a requirement for `wire` to work.

But my thinking is it should make the repository easier to navigate because all
of the `wire` Provider stuff is in one place, just like how all the
routing/handling is in `*.controller.go`, all the business logic is in
`*.service.go`, all the persistence stuff is in `*.repository.go`, and so on.

Also, in theory it means the main code of the application doesn't know anything
about the DI framework being used, so we have the flexibility to change the DI
setup or even complete replace what framework we are using without having to
touch the other code.

Since Sets are just groups of Providers I think it makes sense just to put them
in the same file.

*On Providers:* By convention, Providers should be named like `Provide<TYPE-NAME>`, e.g.
`ProvideUserController` in `src/userapi/user.wire-provider.go`.

Most things probably only need 1 Provider, but in case you needed to use a
slightly different setup in a specific case then my proposal would be to write a
second provider with a more specific name like `Provide<TYPE-NAME>For<CONTEXT>`, e.g.
`ProvideUserServiceForMySpecialCase`.

*On Sets:* By convention, Sets should be named like `<PACKAGE-NAME>Set`, e.g.
`UserApiSet` in `src/userapi/user.wire-provider.go`.

Most packages should probably just have 1 Set that contains all the Providers
for that package.

*On the structure/conventions of source code files for Injectors:* The Injectors for a package are defined in `*.wire-injector.go`.

Injectors *must* be defined in a file with the following comments at the top:

```go
//go:build wireinject
//+build wireinject
```

Because of this "extra special" nature of Injectors I thought it makes sense to
put them in a separate file rather than just having one single `*.wire.go` with
everything in there.
I can imagine the case where:

- you have some package that did not previously have any Injectors in it.
- you go add an Injector to the `*.wire-provider.go` for the package.
- spend 6 hours trying to figure out why the bastard thing isn't working.
- realise you forgot to add 2 lines (that are _commented_, not even real code)
- spend another 2 hours reconsidering your life choices

At least if it's a separate file, you can just copy-paste some other
`*.wire-injector.go` and not have to worry about it.

*On Injectors:* By convention, Injectors should be named like
`Get<TYPE-NAME>`, e.g. `GetApplication` in `src/app/app.wire-injector.go`.

Note that the code you write in `*.wire-injector.go` is just a stub, basically
just a template for `wire` to use to generate the implementation of the
Injector.

For example the `GetApplication` Injector looks like this:

```go
func GetApplication() Application {
	wire.Build(userapi.UserApiSet, AppSet)
	return Application{}
}
```

When you run `wire` (or `go generate` after the first time), it will perform the
following steps:

1. Examine the return type of this Injector; in this case it is `Application`.
2. Search for an `Application` Provider in the Providers and Sets passed to
   `wire.Build`. In this case it should find `ProvideApplication` because
   it included in the `AppSet`.
3. Examine the types of the arguments for `ProvideApplication`, and search for
   a Provider for each of them. This process keeps going recursively until all
   of the dependencies have been fulfilled.
4. Generate an implementation of the Injector that calls all of the Providers as
   needed.
5. Write the implementation of the Injector into `**/wire_gen.go`.

Be advised that the `return Application{}` we wrote in our stub implementation
is just a placeholder; the generated implementation replaces it with the real
thing.

## `src/app`: app setup/creation

This package just sets up the `Application` object.

## `src/<API-NAME>`: Each different API/module for the application

This is kind of a "pivot" compared to how NestJS structures apps, but the reason
why is to make the package names better align with how Go wants you to use
packages.
Otherwise shit gets complicated and you need to make everything public because
stuff in `package controllers` is trying to call stuff in `package services`.

Personally I prefer it as well; with the NestJS approach it drives me crazy how
all the stuff for 1 API gets spread across so many different folders.

- `src/<API-NAME>/*.controller.go` routing and handlers for the respective API.
- `src/<API-NAME>/*.service.go` services for the respective API.
- `src/<API-NAME>/*.repository.go` repositories for the respective API.
- `src/<API-NAME>/*.model.go` models for the respective API.

