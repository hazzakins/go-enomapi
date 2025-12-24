API Tokens
==========

What are API Tokens
-------------------

API Tokens are a replacement for using your account password as part of your API credentials.
- You can have more than one API Token for your account allowing you to give multiple users API access without sharing your account password.
- You can instantly delete \(revoke\) an API Token if there's ever an issue.
- Using API Tokens gives you the ability to update your account password frequently without having to worry about your API connections suddenly breaking because they are now using wrong credentials.

How do I get one?
-----------------

You can manage API tokens in your Enom account, or you can manage API tokens from the API itself.

**API Token Management in your Enom account:**
1. Login to your Enom account.
2. Then go to the following URL to manage your API tokens: [https://cp.enom.com/apitokens/](https://cp.enom.com/apitokens/)
3. Click **Generate an API token**

**API Token Management in the API for new users:**
1. ```
Create a new user with either <a href="http://www.enom.com/api/API%20topics/api_CreateAccount.htm" target="_blank">CREATEACCOUNT</a> or <a href="http://www.enom.com/api/API%20topics/api_CreateSubAccount.htm" target="_blank">CREATESUBACCOUNT</a> commands.
```
2. ```
In the response you will find the APIToken and APITokenName, which you can then store on your side to make any future calls for this user in place of their password.
```
**API Token Management in the API for existing users:**
1. ```
If the user does not already have an API Token use the [ApiToken_Create](doc:apitoken-create) command to create one for them. The response will contain the APIToken and APITokenName.
```
2. ```
If the user already has a token(s) and you need to retrieve them you can use [ApiToken_Details](doc:apitoken-details) command. This returns a list of API tokens for the user.
```
3. ```
If, for some reason, you need to delete/revoke an API token, you can quickly do this with the [ApiToken_Delete](doc:apitoken-delete) command.
```
4. ```
If you need to change the name of a token you can use the [ApiToken_Update](doc:apitoken-update) command.
```
How do I use an API Token?
--------------------------

Simply replace the password with the API token and all you API calls will go through as expected.

Command Reference
-----------------

| Description                              | Command |
| --------------------------------------------------------------------- | -------------------------------------------------- |
| - \*Create\*\* a new API token.                   | [ApiToken\_Create](../docs/apitoken-create.md) |
| - \*Retrieve\*\* one or all API tokens associated with your account. | [ApiToken\_Details](../docs/apitoken-details.md) |
| - \*Update**the**name\*\* of an API token.              | [ApiToken\_Update](../docs/apitoken-update.md) |
| - \*Delete\*\* an API token. | [ApiToken\_Delete](../docs/apitoken-delete.md)  |

> ### Tip
>
>
>
> When creating a new account or sub-account, a **default** API token is created and returned in the output.