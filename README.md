# 🦂 scorpio: ___Client application to test lorax concurrency features___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![A B](https://img.shields.io/badge/branch%20history-linear-blue?style=flat)](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/defining-the-mergeability-of-pull-requests/managing-a-branch-protection-rule)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/scorpio.svg)](https://pkg.go.dev/github.com/snivilised/scorpio)
[![Go report](https://goreportcard.com/badge/github.com/snivilised/scorpio)](https://goreportcard.com/report/github.com/snivilised/scorpio)
[![Coverage Status](https://coveralls.io/repos/github/snivilised/scorpio/badge.svg?branch=master)](https://coveralls.io/github/snivilised/scorpio?branch=master&kill_cache=1)
[![Scorpio Continuous Integration](https://github.com/snivilised/scorpio/actions/workflows/ci-workflow.yml/badge.svg)](https://github.com/snivilised/scorpio/actions/workflows/ci-workflow.yml)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![A B](https://img.shields.io/badge/commit-conventional-commits?style=flat)](https://www.conventionalcommits.org/)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD014/commands-show-output: Dollar signs used before commands without showing output mark down lint -->
<!-- MarkDownLint-disable MD014 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<!-- MD028/no-blanks-blockquote: Blank line inside blockquote -->
<!-- MarkDownLint-disable MD028 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" /></a>
</p>

## 🔰 Introduction

This project will be used as an additional test aid of the currency features of [___lorax___](https://github.com/snivilised/lorax) outside of the scope the unit tests that it already contains.

## 📚 Usage

### Worker Pool Pipeline

From user provided flags, the __pool__ command submits a stream of jobs to the worker pool for concurrent execution. The job is defined as a function which takes a name and emits a greeting to this recipient. Since this function does not represent real work and therefore takes next to no time to run, it is infused with a synthetic delay to simulate real work. The delay is currently defined as a random interval.

- ▶️ __Producer(jobsChSize, stopAfter) => jobsCh:__ generates a workload

📌 Variables:

| Name       | Flag    | Unit    | Description |
|------------|---------|---------|-------------|
| jobsChSize | _jobq_  |         | Capacity of the jobs channel |
| stopAfter  | _after_ | seconds | Stops the producer after this time period |

- ▶️ __Pool(noOfWorkers, jobsCh) => resultsCh:__ handles the workload with multiple workers generating a result stream

📌 Variables:

| Name          | Flag    | Unit    | Description |
|---------------|---------|---------|-------------|
| noOfWorkers   | _now_   |         | No of workers in pool |

- ▶️ __Consumer(resultsChSize, resultsCh):__ consumes the result stream

📌 Variables:

| Name          | Flag    | Unit    | Description |
|---------------|---------|---------|-------------|
| resultsChSize | _resq_  |         | Capacity of the results channel |

___⚡ Invocation:___

```
scorpio pool --after 3 --now 5 --jobq 18 --resq 16
```

## 🎀 Features

<p align="left">
  <a href="https://onsi.github.io/ginkgo/"><img src="https://onsi.github.io/ginkgo/images/ginkgo.png" width="100" /></a>
  <a href="https://onsi.github.io/gomega/"><img src="https://onsi.github.io/gomega/images/gomega.png" width="100" /></a>
</p>

- unit testing with [Ginkgo](https://onsi.github.io/ginkgo/)/[Gomega](https://onsi.github.io/gomega/)
- implemented with [🐍 Cobra](https://cobra.dev/) cli framework, assisted by [🐲 Cobrass](https://github.com/snivilised/cobrass)
- i18n with [go-i18n](https://github.com/nicksnyder/go-i18n)
- linting configuration and pre-commit hooks, (see: [linting-golang](https://freshman.tech/linting-golang/)).

## 🔨 Developer Info

### ☑️ Github changes

Some general project settings are indicated as follows:

___General___

Under `Pull Requests`

- `Allow merge commits` 🔳 _DISABLE_
- `Allow squash merging` 🔳 _DISABLE_
- `Allow rebase merging` ✅ _ENABLE_

___Branch Protection Rules___

Under `Protect matching branches`

- `Require a pull request before merging` ✅ _ENABLE_
- `Require linear history` ✅ _ENABLE_
- `Do not allow bypassing the above settings` ✅ _ENABLE_

### ☑️ Code coverage

- `coveralls.io`: add scorpio project

### 🌐 l10n Translations

This template has been setup to support localisation. The default language is `en-GB` with support for `en-US`. There is a translation file for `en-US` defined as __i18n/deploy/scorpio.active.en-US.json__. This is the initial translation for `en-US` that should be deployed with the app.

Make sure that the go-i18n package has been installed so that it can be invoked as cli, see [go-i18n](https://github.com/nicksnyder/go-i18n) for installation instructions.

To maintain localisation of the application, the user must take care to implement all steps to ensure translate-ability of all user facing messages. Whenever there is a need to add/change user facing messages including error messages, to maintain this state, the user must:

- define template struct (__xxxTemplData__) in __i18n/messages.go__ and corresponding __Message()__ method. All messages are defined here in the same location, simplifying the message extraction process as all extractable strings occur at the same place. Please see [go-i18n](https://github.com/nicksnyder/go-i18n) for all translation/pluralisation options and other regional sensitive content.

For more detailed workflow instructions relating to i18n, please see [i18n README](./resources/doc/i18n-README.md)

### 🧪 Quick Test

To check the app is working (as opposed to running the unit tests), build and deploy:

> task tbd

(which performs a test, build then deploy)

NB: the `deploy` task has been set up for windows by default, but can be changed at will.

Check that the executable and the US language file __scorpio.active.en-US.json__ have both been deployed. Then invoke the __pool__ command with something like

> scorpio pool -now 5 --job 18 --resq 16
