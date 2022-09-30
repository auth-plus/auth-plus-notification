# Core

Here is the main application. The construction is:

A Usecase has dependencies like a repository or provider-manager and must contain the business rule.

A Repository can access a database to look for a user or a template.

A Manager returns a provider based on a technical rule.

A Provider has the implementation.
