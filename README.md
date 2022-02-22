---
permalink: actions/add-to-whitelist-action.html
tags: [github, action]
summary: Action to add to an organization's whitelist
---

# Add to Whitelist

This action adds a github action to an organization's whitelist. First it reads
the active patterns included in the whitelist using the github API and checks
whether the new pattern is already included. If it is not, it adds it to the
list, calling again the github.com API to update the whitelist.
This action receives a token -which will have to have admin rights-, a pattern
and the organization where the new action is to be allowed.

## Inputs

## `token`

**Required**: The token to be used to call the github.com API. It has to have
admin rights.

E.g.: `token: "xxxxxxxx"`

## `pattern`

**Required**: The pattern to be added to the organization's list of allowed
patterns.

E.g.: `pattern: "santander-group/sonar-issues-to-annotations"`

## `org`

**Required**: The organization in which the pattern is to be allowed.

E.g.: `org: "the-iron-bank-of-volantis"`

## Outputs

## `allowedPatterns`

Updated list of allowd patterns.

E.g.: given a token, a pattern `test4/*` and an organization, the output could be like this:

    "docker/*,santander-group/*,test3/*,test4/*"

## Example usage

      - name: Add to whitelist
        uses: srrasero/addwhitelist@v0.1
        id: goaction
        with: 
          token: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
          pattern: "test4/*"
          org: the-iron-bank-of-volantis

## Getting involved

This project is far from being perfect, so any help is always welcome. Please,
review our [CONTRIBUTING](CONTRIBUTING.md) file to know how to get involved.

---

## Licensing info

* [LICENSE](LICENSE)

---

## Credits and references

1. [Creating github actions with go](https://www.youtube.com/watch?v=8IgNY8QT3vk)
2. [Creating actions official documentation](https://docs.github.com/en/actions/creating-actions)
