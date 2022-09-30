# Usecases

The Usecase is a struct that should implement methods constrained by its context. All business rules should be here

## Example

Let's say that you can accept two kinds of contract when sending an email:

1. payload has the user id and the template id
2. payload has the email, template_id, and all other variables

These two case means two methods and each one has a differently flow.
