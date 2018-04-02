# ghcli - Extremely simple Github CLI test program

This project is a very simple command-line tool to get some basic data about 
the repos in a given organization. 

## Building

### With Docker

To build `ghcli` with Docker, you must first build the builder. Do this by running

```
    ./scripts/bootstrap
```

You will only need to do this once.

After the bootstrap script finishes, you can build the tool by running

```
    ./scripts/build-with-docker
```

### Without Docker

If you have the following installed on your system

* Go v1.9 or higher
* Dep v0.4.1 or higher

you can run `./scripts/test && ./scripts/build` to build.

## Running

After building, there will be binaries in the `./target` directory, named after the arch/OS they were built for. You should see binaries for macOS and Linux. 

```
Display repos from a Github organization, sorted by different criteria.

Usage:
  ghcli [flags]
  ghcli [command]

Available Commands:
  contributions sort by number of pull requests/fork
  forks         sort by number of forks
  help          Help about any command
  login         Set your Github OAuth token
  pull-requests sort by number of pull requests
  stars         sort by the number of stars

Flags:
  -h, --help           help for ghcli
      --org string     organization to use
      --token string   Github OAuth token
      --top int        number of results to return (default 10)
  -v, --verbose        verbose output

Use "ghcli [command] --help" for more information about a command.
```

There are four commands that display information about a given Github organization: `contributions`, `forks`, `pull-requests`, and `stars`. The fifth command, `login`, is for setting your Github OAuth token into the `ghcli` config file.

### Parameters

You can use `ghcli` anonymously, but some repos will be hidden from you, and there is a rate-limit of 60 requests/hour, so you will probably run into that quickly.

If you have a Github account, you can generate an OAuth token that can be used for authorization. Go to [Settings/Developer settings/Personal access tokens](https://github.com/settings/tokens) to generate an OAuth token. You can pass this token on the command line, by adding `--token <oauth token>` to each request, but it's much easier to add it to your config file. To do that, run

```
ghcli login <oauth token>
```

this will add the token to your settings file, `~/.ghcli.yaml`. 

Once you have done this, all future operations will automatically use this token.

All of the commands require the name of an organization to operate on, which you can specify each time by adding `--org <org name>` to the command line. If you usually work with the same organization, you can add it to your config file, `~/.ghcli.yaml`, as

```yaml
org: my-organization
```

The default for all the commands is to return the top 10 repos. You can change this by either passing `--top n` with each request. You can also set it in your config file with

```yaml
top: 23
```

For all three of these parameters, if you have added them to your settings file, you can override them for a one-off by adding the command line option.

### Commands

* `contributions` will sort based on the ratio of pull requests to forks. 
* `forks` will sort based on the number of times the repo has been forked
* `pull-requests` will sort based on the number of pull requests created against the repo
* `stars` will sort based on the number of stars a repo has been given

## Examples

Running anonymously, on a Mac, fetching repos sorted by stars:

```
./target/ghcli-darwin-amd64 --org my-org stars
```

Running anonymously, on a Mac, sorted by forks, but limiting the number returned to 5:

```
./target/ghcli-darwin-amd64 --org my-org forks --top 5
```

If you had a settings file that looked like this:

```yaml
top: 5
org: my-org
```

then the previous command would just be:

```
./target/ghcli-darwin-amd64
```

If you need to specify an OAuth token, it would look like this:

```
./target/ghcli-darwin-amd64 --org my-org forks --top 5 --token 12345
```
