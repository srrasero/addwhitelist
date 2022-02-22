# Contributing to this project

:+1::tada: First off, thanks for taking the time to contribute! :tada::+1:

The following is a set of guidelines for contributing to this project and its
packages, which are hosted in GitHub. These are mostly guidelines, not rules.
Use your best judgment, and feel free to propose changes to this document in a
pull request.

## Table Of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

* [Code of Conduct](#code-of-conduct)
* [Prepare the environment](#prepare-the-environment)
* [What should I know before I get started?](#what-should-i-know-before-i-get-started)
* [How Can I Contribute?](#how-can-i-contribute)
  * [Reporting Bugs](#reporting-bugs)
    * [Before Submitting A Bug Report](#before-submitting-a-bug-report)
    * [How Do I Submit A (Good) Bug Report?](#how-do-i-submit-a-good-bug-report)
  * [Suggesting Enhancements](#suggesting-enhancements)
    * [Before Submitting An Enhancement Suggestion](#before-submitting-an-enhancement-suggestion)
    * [How Do I Submit A (Good) Enhancement Suggestion?](#how-do-i-submit-a-good-enhancement-suggestion)
  * [Your First Code Contribution](#your-first-code-contribution)
    * [Local development](#local-development)
  * [Pull Requests](#pull-requests)
* [Styleguides](#styleguides)
  * [Git Commit Messages](#git-commit-messages)
  * [Documentation Styleguide](#documentation-styleguide)
* [What do we expect from a contributor?](#what-do-we-expect-from-a-contributor)
* [Additional Notes](#additional-notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Code of Conduct

This project and everyone participating in it is governed by the
[Code Of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to
uphold this code. Please report unacceptable behavior by creating an issue or
reporting it directly to david.hernandezmartin@gruposantander.com.

## Prepare the environment

You'll get faster results by using the resources below instead of creating an
issue to aks a question.

* This project makes use of two main javascript dependencies. As mostly all the
  GitHub's Action projects, it relies on
  [actions/core](https://www.npmjs.com/package/@actions/core) package. And it
  also needs the [simple-git](https://www.npmjs.com/package/simple-git) package.

      npm init -y
      npm install @actions/core
      npm install @actions/github
      npm install simple-git

* Please,
  [sign your commits](https://docs.github.com/en/github/authenticating-to-github/managing-commit-signature-verification/signing-commits).

## What should I know before I get started?

This is a Javascript GitHub Action. To start your local development environment,
you need to install the dependencies listed before and, as usual for GitHub
Actions development based on Javascript, to initialize as a Node.js project.

The `node_modules` folder is included in the `.gitignore` file. Checking in your
`node_modules` directory can cause problems. As an alternative,
you should use a tool called [@vercel/ncc](https://github.com/vercel/ncc) to
compile your code and modules into one file used for distribution. To package
the action using `ncc` you can use the following command from the project's root
folder:

     ncc build index.js -o dist

## How Can I Contribute?

### Reporting Bugs

This section guides you through submitting a bug report for this project.
Following these guidelines helps maintainers and the community understand your
report :pencil:, reproduce the behavior :computer: :computer:, and find related
reports :mag_right:.

#### Before Submitting A Bug Report

#### How Do I Submit A (Good) Bug Report?

Bugs are tracked as [GitHub issues](https://guides.github.com/features/issues/).
After you've determined this is the repository your bug is related to, create an
issue and provide the information requested in
[the issue template](.github/ISSUE_TEMPLATE/bug_report.md).

### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion for this
repository, including completely new features and minor improvements to existing
functionality. Following these guidelines helps maintainers and the community
understand your suggestion :pencil: and find related suggestions :mag_right:.

#### Before Submitting An Enhancement Suggestion

Take a look at the
[discussions](https://github.com/santander-group/git-tag-counter-action/discussions)
and [issues](https://github.com/santander-group/git-tag-counter-action/issues)
sections. Maybe somebody is working on that.

#### How Do I Submit A (Good) Enhancement Suggestion?

Feature requests are tracked as
[GitHub issues](https://guides.github.com/features/issues/).
After you've determined this is the right repository, create an issue and
provide the information requested in
[new feature template](.github/ISSUE_TEMPLATE/feature_request.md).

### Pull Requests

The process described here has several goals:

* Maintain project's quality
* Fix problems that are important to users
* Engage the community in working toward their best
* Enable a sustainable system for this project's maintainers to review
  contributions

Please follow these steps to have your contribution considered by the maintainers:

1. Follow all instructions in
   [the pull request template](.github/PULL_REQUEST_TEMPLATE/pull_request_template.md)
2. Follow the [styleguides](#styleguides)
3. After you submit your pull request, verify that all
  [status checks](https://help.github.com/articles/about-status-checks/)
  are passing. If a status check is failing, and you believe that the failure
  is unrelated to your change, please leave a comment on the pull request
  explaining why you believe the failure is unrelated. A maintainer will re-run
  the status check for you. If we conclude that the failure was a false
  positive, then we will open an issue to track that problem with our status
  check suite.

While the prerequisites above must be satisfied prior to having your pull
request reviewed, the reviewer(s) may ask you to complete additional design
work, tests, or other changes before your pull request can be ultimately
accepted.

## Styleguides

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

### Documentation Styleguide

* Use [Markdown](https://daringfireball.net/projects/markdown).

## Getting help

If you need help when using this action please create take a look at the
[Discussions section](https://github.com/santander-group/git-tag-counter-action/discussions)
and search for somebody else topic. If your problem is something new, please
create a new entry.

## What do we expect from a contributor?

Only being nice and gentle. Be your best self.
